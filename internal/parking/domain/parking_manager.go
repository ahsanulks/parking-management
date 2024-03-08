package domain

type ParkingManager struct {
	id int
}

func NewParkingManager(id int) *ParkingManager {
	return &ParkingManager{
		id: id,
	}
}

func (pm *ParkingManager) Id() int {
	return pm.id
}

func (pm *ParkingManager) CreateParkingLot(lotName string, numSlots int) *ParkingLot {
	parkingLot := &ParkingLot{
		managerId: pm.id,
		name:      lotName,
		capacity:  numSlots,
		slotLeft:  numSlots,
	}
	for i := 1; i <= numSlots; i++ {
		parkingLot.slots = append(parkingLot.slots, &ParkingSlot{number: i, available: true})
	}
	return parkingLot
}
