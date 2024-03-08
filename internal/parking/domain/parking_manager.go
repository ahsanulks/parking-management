package domain

type ParkingManager struct {
	ParkingLots []*ParkingLot
}

func (pm *ParkingManager) CreateParkingLot(numSlots int) {
	parkingLot := &ParkingLot{}
	for i := 1; i <= numSlots; i++ {
		parkingLot.slots = append(parkingLot.slots, &ParkingSlot{number: i, available: true})
	}
	pm.ParkingLots = append(pm.ParkingLots, parkingLot)
}
