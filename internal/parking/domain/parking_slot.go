package domain

type ParkingSlot struct {
	id          int
	number      int
	available   bool
	maintenance bool
}

func (parkingSlot ParkingSlot) Number() int {
	return parkingSlot.number
}

func (parkingSlot ParkingSlot) Available() bool {
	return parkingSlot.available
}

func (parkingSlot ParkingSlot) InMaintenance() bool {
	return parkingSlot.maintenance
}

func (parkingSlot *ParkingSlot) makeAvailable() {
	parkingSlot.available = true
}

func (parkingSlot *ParkingSlot) UnmarshallFromDatabase(id, number int, availale, maintenance bool) {
	parkingSlot.id = id
	parkingSlot.number = number
	parkingSlot.available = availale
	parkingSlot.maintenance = maintenance
}
