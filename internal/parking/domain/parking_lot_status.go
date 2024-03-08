package domain

import "time"

type ParkingLotStatus struct {
	Id          int
	Name        string
	SlotLeft    int
	Capcity     int
	SlotsStatus []*ParkingSlotStatus
}

type ParkingSlotStatus struct {
	Number     int
	TicketCode *string
	EntryTime  *time.Time
}
