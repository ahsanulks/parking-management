package domain

import "errors"

type ParkingSlot struct {
	id          int
	number      int
	available   bool
	maintenance bool
}

type ParkingLot struct {
	id       int
	slots    []*ParkingSlot
	name     string
	capacity int
	slotLeft int
}

func (parkingLot ParkingLot) hasSlotLeft() bool {
	return parkingLot.slotLeft > 0
}

func (parkingLot ParkingLot) SlotLeft() int {
	return parkingLot.slotLeft
}

func (parkingSlot *ParkingSlot) makeAvailable() {
	parkingSlot.available = true
}

func (parkingLot *ParkingLot) MakeSlotAvailable() error {
	parkingLot.slotLeft++
	if len(parkingLot.slots) == 0 {
		return errors.New("failed when make slots available")
	}
	parkingLot.slots[0].makeAvailable()
	return nil
}

func (parkingLot *ParkingLot) ParkVehicle() (*Ticket, error) {
	if !parkingLot.hasSlotLeft() {
		return nil, NewNoParkingSlotLeftError(parkingLot)
	}

	parkingLot.slotLeft -= 1
	nearestSlot, err := parkingLot.findNearestSlot()
	if err != nil {
		return nil, err
	}
	nearestSlot.available = false

	return NewTicket(parkingLot.id, nearestSlot.id), nil
}

func (parkingLot ParkingLot) findNearestSlot() (*ParkingSlot, error) {
	var nearestSlot *ParkingSlot
	slots := parkingLot.slots
	for i := 0; i < len(parkingLot.slots); i++ {
		if slots[i].available && !slots[i].maintenance {
			nearestSlot = parkingLot.slots[i]
			break
		}
	}

	if nearestSlot == nil {
		return nil, errors.New("unable find parking slot")
	}

	return nearestSlot, nil
}

func (parkingLot *ParkingLot) UnmarshallFromDatabase(id, slotLeft int, name string) {
	parkingLot.id = id
	parkingLot.slotLeft = slotLeft
	parkingLot.name = name
}

func (parkingLot *ParkingLot) Slot(slots []*ParkingSlot) {
	parkingLot.slots = slots
}

func (parkingSlot *ParkingSlot) UnmarshallFromDatabase(id, number int, availale, maintenance bool) {
	parkingSlot.id = id
	parkingSlot.number = number
	parkingSlot.available = availale
	parkingSlot.maintenance = maintenance
}
