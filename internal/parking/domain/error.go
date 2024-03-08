package domain

import "net/http"

type NoParkingSlotLeftError struct {
	lotName string
	code    int
}

func NewNoParkingSlotLeftError(parkingLot *ParkingLot) *NoParkingSlotLeftError {
	return &NoParkingSlotLeftError{
		lotName: parkingLot.name,
		code:    http.StatusBadRequest,
	}
}

func (err NoParkingSlotLeftError) Error() string {
	return "No parking slot left on " + err.lotName
}

func (err NoParkingSlotLeftError) Code() int {
	return err.code
}
