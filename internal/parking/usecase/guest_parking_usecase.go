package usecase

import (
	"context"
	"supertaltest/internal/parking/domain"
)

type GuestParkingUsecase struct {
	parkRepository domain.Repository
}

func NewGuestParkerUsecase(repo domain.Repository) *GuestParkingUsecase {
	return &GuestParkingUsecase{
		parkRepository: repo,
	}
}

func (parker *GuestParkingUsecase) ParkUserVehicle(ctx context.Context, lotId int) (ticketId string, err error) {
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

func (parker *GuestParkingUsecase) ExitUserPark(ctx context.Context, ticketCode string) (*domain.Ticket, error) {
	return parker.parkRepository.ExitParking(
		ctx,
		ticketCode,
		func(ticket *domain.Ticket, lot *domain.ParkingLot) (*domain.Ticket, error) {
			err := ticket.Exit()
			if err != nil {
				return nil, err
			}
			err = lot.MakeSlotAvailable()
			return ticket, err
		})
}
