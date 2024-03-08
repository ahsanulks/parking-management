package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"supertaltest/internal/parking/domain"
	"time"

	"github.com/jmoiron/sqlx"
)

type ParkingPostgresqlRepository struct {
	db *sqlx.DB
}

func NewParkingPostgresqlRepository(db *sqlx.DB) *ParkingPostgresqlRepository {
	return &ParkingPostgresqlRepository{
		db: db,
	}
}

type ParkingLotModel struct {
	Id        int    `db:"id"`
	Name      string `db:"name"`
	Capacity  int    `db:"capacity"`
	SlotLeft  int    `db:"slot_left"`
	ManagerId int    `db:"manager_id"`
}

type ParkingSlotModel struct {
	Id           int  `db:"id"`
	ParkingLotId int  `db:"parking_lot_id"`
	Number       int  `db:"number"`
	Available    bool `db:"available"`
	Maintenance  bool `db:"in_maintenance"`
}

type TicketModel struct {
	Id            int        `db:"id"`
	ParkingLotId  int        `db:"parking_lot_id"`
	ParkingSlotId int        `db:"parking_slot_id"`
	Code          string     `db:"code"`
	EntryTime     time.Time  `db:"entry_time"`
	ExitTime      *time.Time `db:"exit_time"`
	Fee           *int       `db:"fee"`
}

func (repo *ParkingPostgresqlRepository) ParkVehicle(
	ctx context.Context,
	lotId int,
	updateFunc func(parkingLot *domain.ParkingLot) (*domain.Ticket, error),
) (ticket *domain.Ticket, err error) {
	tx := repo.db.MustBeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
		ReadOnly:  false,
	})
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	var parkingLotModel ParkingLotModel
	err = tx.Get(&parkingLotModel, "SELECT * FROM parking_lots WHERE id = $1 LIMIT 1", lotId)
	if err != nil {
		return nil, err
	}

	var availableSlotModel []*ParkingSlotModel
	err = tx.Select(&availableSlotModel, "SELECT * FROM parking_slots WHERE parking_lot_id = $1 AND available = $2", lotId, true)
	if err != nil {
		return nil, err
	}

	parkingLot := unmarshallParkingLot(&parkingLotModel, availableSlotModel)

	ticket, err = updateFunc(parkingLot)
	if err != nil {
		return nil, err
	}

	_, err = execAndCheckTx(tx, "UPDATE parking_lots SET slot_left = $1 where id = $2", parkingLot.SlotLeft(), lotId)
	if err != nil {
		return nil, err
	}

	_, err = execAndCheckTx(tx, "UPDATE parking_slots SET available = $1 where id = $2", false, ticket.SlotID())
	if err != nil {
		return nil, err
	}

	_, err = execAndCheckTx(tx, "INSERT INTO tickets (code, parking_lot_id, parking_slot_id, entry_time) VALUES ($1, $2, $3, $4)", ticket.Code(), ticket.LotId(), ticket.SlotID(), ticket.EntiryTime())
	if err != nil {
		return nil, err
	}

	return ticket, nil
}

func (repo *ParkingPostgresqlRepository) ExitParking(
	ctx context.Context,
	ticketCode string,
	updateFunc func(ticket *domain.Ticket, lot *domain.ParkingLot) (*domain.Ticket, error),
) (ticket *domain.Ticket, err error) {
	tx := repo.db.MustBeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
		ReadOnly:  false,
	})

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var parkingLotModel ParkingLotModel
	var parkingSlotModel ParkingSlotModel
	var ticketModel TicketModel

	query := `
		SELECT
			pl.id AS parking_lot_id, pl.name AS parking_lot_name, pl.capacity, pl.slot_left,
			ps.id AS parking_slot_id, ps.number, ps.available, ps.in_maintenance,
			t.id AS ticket_id, t.code, t.entry_time, t.exit_time, t.fee
		FROM tickets t
		INNER JOIN parking_slots ps ON t.parking_slot_id = ps.id
		INNER JOIN parking_lots pl ON ps.parking_lot_id = pl.id
		WHERE t.code = $1
		AND t.fee is null
		LIMIT 1
	`

	err = tx.QueryRow(query, ticketCode).Scan(
		&parkingLotModel.Id, &parkingLotModel.Name, &parkingLotModel.Capacity, &parkingLotModel.SlotLeft,
		&parkingSlotModel.Id, &parkingSlotModel.Number, &parkingSlotModel.Available, &parkingSlotModel.Maintenance,
		&ticketModel.Id, &ticketModel.Code, &ticketModel.EntryTime, &ticketModel.ExitTime, &ticketModel.Fee,
	)
	if err != nil {
		return nil, err
	}
	ticketModel.ParkingLotId = parkingLotModel.Id
	ticketModel.ParkingSlotId = parkingSlotModel.Id

	parkingLot := unmarshallParkingLot(&parkingLotModel, []*ParkingSlotModel{&parkingSlotModel})

	ticket = &domain.Ticket{}
	unmarshallTicket(&ticketModel, ticket)

	ticket, err = updateFunc(ticket, parkingLot)
	if err != nil {
		return nil, err
	}

	_, err = execAndCheckTx(tx, "UPDATE tickets SET fee = $1, exit_time = $2 where code = $3", ticket.Fee(), ticket.ExitTime(), ticketCode)
	if err != nil {
		return nil, err
	}

	_, err = execAndCheckTx(tx, "UPDATE parking_slots SET available = $1 where id = $2", true, ticketModel.ParkingSlotId)
	if err != nil {
		return nil, err
	}

	_, err = execAndCheckTx(tx, "UPDATE parking_lots SET slot_left = $1 where id = $2", parkingLot.SlotLeft(), ticketModel.ParkingLotId)
	if err != nil {
		return nil, err
	}

	return ticket, nil
}

func (repo ParkingPostgresqlRepository) CreateParkingLot(ctx context.Context, lot *domain.ParkingLot) (err error) {
	tx := repo.db.MustBeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
		ReadOnly:  false,
	})
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var lotId int
	err = tx.QueryRow("INSERT INTO parking_lots (name, capacity, slot_left, manager_id) VALUES($1, $2, $3, $4) returning id", lot.Name(), lot.Capacity(), lot.SlotLeft(), lot.ManagerId()).Scan(&lotId)
	if err != nil {
		return err
	}

	var params []any
	var query []string
	for _, parkingSlot := range lot.Slots() {
		query = append(query, fmt.Sprintf("($%d, $%d, $%d, $%d)", len(params)+1, len(params)+2, len(params)+3, len(params)+4))
		params = append(params, lotId, parkingSlot.Number(), parkingSlot.Available(), parkingSlot.InMaintenance())
	}

	stmt := "INSERT INTO parking_slots (parking_lot_id, number, available, in_maintenance) VALUES" + strings.Join(query, ", ")
	_, err = execAndCheckTx(tx, stmt, params...)
	if err != nil {
		return err
	}

	lot.UnmarshallFromDatabase(int(lotId), lot.SlotLeft(), lot.Capacity(), lot.ManagerId(), lot.Name(), lot.Slots())
	return
}

func unmarshallParkingLot(
	parkingLotModel *ParkingLotModel,
	slotModel []*ParkingSlotModel,
) *domain.ParkingLot {
	availableSlots := make([]*domain.ParkingSlot, len(slotModel))
	for i, slot := range slotModel {
		parkingSlot := new(domain.ParkingSlot)
		parkingSlot.UnmarshallFromDatabase(
			slot.Id,
			slot.Number,
			slot.Available,
			slot.Maintenance,
		)
		availableSlots[i] = parkingSlot
	}

	parkingLot := new(domain.ParkingLot)
	parkingLot.UnmarshallFromDatabase(
		parkingLotModel.Id,
		parkingLotModel.SlotLeft,
		parkingLotModel.Capacity,
		parkingLotModel.ManagerId,
		parkingLotModel.Name,
		availableSlots,
	)

	return parkingLot
}

func unmarshallTicket(ticketModel *TicketModel, ticket *domain.Ticket) {
	ticket.UnmarshallFromDatabase(
		ticketModel.Id,
		ticketModel.ParkingLotId,
		ticketModel.ParkingSlotId,
		ticketModel.Code,
		ticketModel.EntryTime,
		ticketModel.ExitTime,
		ticketModel.Fee,
	)
}

func execAndCheckTx(tx *sqlx.Tx, query string, args ...any) (sql.Result, error) {
	result, err := tx.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	return result, nil
}
