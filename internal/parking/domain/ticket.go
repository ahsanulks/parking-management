package domain

import (
	"errors"
	"fmt"
	"math"
	"time"
)

const (
	FEE_PER_HOUR int = 10
)

type Ticket struct {
	id        int
	lotID     int
	slotID    int
	code      string
	entryTime time.Time
	exitTime  *time.Time
	fee       *int
}

func NewTicket(lotID, slotID int) *Ticket {
	ticket := &Ticket{
		lotID:     lotID,
		slotID:    slotID,
		entryTime: time.Now().UTC(),
	}
	ticket.generateCode()
	return ticket
}

func (t *Ticket) UnmarshallFromDatabase(
	id, lotID, slotID int,
	code string,
	entryTime time.Time,
	exitTime *time.Time,
	fee *int,
) {
	t.id = id
	t.lotID = lotID
	t.slotID = slotID
	t.code = code
	t.entryTime = entryTime.UTC()
	t.exitTime = exitTime
	t.fee = fee
}

func (t *Ticket) generateCode() {
	t.code = fmt.Sprintf("%d-%d-%d", t.lotID, t.slotID, t.entryTime.Unix())
}

func (t *Ticket) Exit() error {
	if t.exitTime != nil {
		return errors.New("ticket already exit")
	}
	now := time.Now().UTC()
	t.exitTime = &now
	duration := time.Since(t.entryTime)
	hours := int(math.Ceil(duration.Hours()))
	totalFee := hours * FEE_PER_HOUR
	t.fee = &totalFee
	return nil
}

func (t Ticket) LotId() int {
	return t.lotID
}

func (t Ticket) Code() string {
	return t.code
}

func (t Ticket) SlotID() int {
	return t.slotID
}

func (t Ticket) EntiryTime() time.Time {
	return t.entryTime
}

func (t Ticket) Fee() *int {
	return t.fee
}

func (t Ticket) ExitTime() *time.Time {
	return t.exitTime
}
