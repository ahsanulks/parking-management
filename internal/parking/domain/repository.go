package domain

import (
	"context"
)

type Repository interface {
	ParkVehicle(
		ctx context.Context,
		lotId int,
		updateFunc func(parkingLot *ParkingLot) (*Ticket, error),
	) (*Ticket, error)

	ExitParking(
		ctx context.Context,
		ticketCode string,
		updateFunc func(ticket *Ticket, lot *ParkingLot) (*Ticket, error),
	) (*Ticket, error)
}
