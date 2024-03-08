package domain

import (
	"fmt"
)

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

func (pm *ParkingManager) UnparkVehicle(registrationNumber string) int {
	// for _, lot := range pm.ParkingLots {
	// 	for _, slot := range lot.Slots {
	// 		if slot.Vehicle != nil && slot.Vehicle.RegistrationNumber == registrationNumber {
	// 			duration := time.Since(slot.Vehicle.EntryTime)
	// 			hours := int(duration.Hours())
	// 			fee := 10 * hours
	// 			fmt.Printf("Vehicle with registration number %s unparked successfully. Parking fee: Rs %d.\n", registrationNumber, fee)
	// 			slot.Vehicle = nil
	// 			slot.Available = true
	// 			return fee
	// 		}
	// 	}
	// }
	fmt.Println("Vehicle not found in any parking slot.")
	return 0
}
