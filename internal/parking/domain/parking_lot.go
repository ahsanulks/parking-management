package domain

import "errors"

type ParkingLot struct {
	id        int
	managerId int
	slots     []*ParkingSlot
	name      string
	capacity  int
	slotLeft  int
}

func (parkingLot ParkingLot) hasSlotLeft() bool {
	return parkingLot.slotLeft > 0
}

func (parkingLot ParkingLot) SlotLeft() int {
	return parkingLot.slotLeft
}

func (parkingLot ParkingLot) Id() int {
	return parkingLot.id
}

func (parkingLot ParkingLot) Name() string {
	return parkingLot.name
}

func (parkingLot ParkingLot) Capacity() int {
	return parkingLot.capacity
}

func (parkingLot ParkingLot) ManagerId() int {
	return parkingLot.managerId
}

func (parkingLot ParkingLot) Slots() []*ParkingSlot {
	return parkingLot.slots
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

func (parkingLot *ParkingLot) UnmarshallFromDatabase(id, slotLeft, capacity, managerId int, name string, slots []*ParkingSlot) {
	parkingLot.id = id
	parkingLot.slotLeft = slotLeft
	parkingLot.capacity = capacity
	parkingLot.managerId = managerId
	parkingLot.name = name
	parkingLot.slots = slots
}
