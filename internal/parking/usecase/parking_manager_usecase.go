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

func (pu *ParkingManagerUsecase) GetLotStatus(ctx context.Context, lotId int) (*domain.ParkingLotStatus, error) {
	return pu.parkingRepo.GetParkingLotStatus(ctx, lotId)
}

func (pu *ParkingManagerUsecase) ToggleParkingSlotToMaintenance(ctx context.Context, parkingManager *domain.ParkingManager, slotId int) error {
	return pu.parkingRepo.ToggleParkingSlotMaintenance(ctx, parkingManager.Id(), slotId)
}
