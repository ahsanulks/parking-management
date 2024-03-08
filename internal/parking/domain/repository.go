package domain

import (
	"context"
	"time"
)

type Repository interface {
	ParkVehicle(
		ctx context.Context,
		lotId int,
		updateFunc func(parkingLot *ParkingLot) (*Ticket, error),
	) (*Ticket, error)

	ExitParking(
		ctx context.Context,
		ticketCode string,
		updateFunc func(ticket *Ticket, lot *ParkingLot) (*Ticket, error),
	) (*Ticket, error)

	CreateParkingLot(ctx context.Context, lot *ParkingLot) error
	GetParkingLotStatus(ctx context.Context, lotId int) (*ParkingLotStatus, error)
	ToggleParkingSlotMaintenance(ctx context.Context, managerId, slotId int) error
	GetParkingSummary(ctx context.Context, startDate, endDate time.Time) (*ParkingSummary, error)
}
