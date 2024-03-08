package usecase

import (
	"context"
	"supertaltest/internal/parking/domain"
)

type GuestParkerUsecase struct {
	parkRepository domain.Repository
}

func NewGuestParkerUsecase(repo domain.Repository) *GuestParkerUsecase {
	return &GuestParkerUsecase{
		parkRepository: repo,
	}
}

func (parker *GuestParkerUsecase) ParkUserVehicle(ctx context.Context, lotId int) (ticketId string, err error) {
	ticket, err := parker.parkRepository.ParkVehicle(
		ctx,
		lotId,
		func(parkingLot *domain.ParkingLot) (*domain.Ticket, error) {
			return parkingLot.ParkVehicle()
		})
	if err != nil {
		return "", err
	}
	return ticket.Code(), err
}

func (parker *GuestParkerUsecase) ExitUserPark(ctx context.Context, ticketCode string) (*domain.Ticket, error) {
	return parker.parkRepository.ExitParking(
		ctx,
		ticketCode,
		func(ticket *domain.Ticket, lot *domain.ParkingLot) (*domain.Ticket, error) {
			err := ticket.Exit()
			if err != nil {
				return nil, err
			}
			lot.MakeSlotAvailable()
			return ticket, err
		})
}
