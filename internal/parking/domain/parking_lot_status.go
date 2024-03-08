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

// func (parkingLotStatus *ParkingLotStatus) UnmarshallFromDatabase(
// 	id int,
// 	name string,
// 	slotLeft int,
// 	capcity int,
// 	slotsStatus []*ParkingSlotStatus,
// ) {
// 	parkingLotStatus.id = id
// 	parkingLotStatus.name = name
// 	parkingLotStatus.capcity = capcity
// 	parkingLotStatus.slotLeft = slotLeft
// 	parkingLotStatus.slotsStatus = slotsStatus
// }
