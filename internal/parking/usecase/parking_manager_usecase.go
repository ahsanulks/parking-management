package usecase

import (
	"context"
	"supertaltest/internal/parking/domain"
)

type ParkingManagerUsecase struct {
	parkingRepo domain.Repository
}

func NewParkingManagerUsecase(repo domain.Repository) *ParkingManagerUsecase {
	return &ParkingManagerUsecase{
		parkingRepo: repo,
	}
}

type RequestCreateParkingLot struct {
	NumSlot        int
	ParkingLotName string
}

func (pu *ParkingManagerUsecase) CreateParkingLot(ctx context.Context, parkingManager *domain.ParkingManager, request *RequestCreateParkingLot) (id int, err error) {
	parkingLot := parkingManager.CreateParkingLot(request.ParkingLotName, request.NumSlot)

	err = pu.parkingRepo.CreateParkingLot(ctx, parkingLot)
	return parkingLot.Id(), err
}
