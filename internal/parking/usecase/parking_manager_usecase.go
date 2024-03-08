package usecase

import (
	"context"
	"supertaltest/internal/parking/domain"
	"time"
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

func (pu *ParkingManagerUsecase) GetParkingSummary(ctx context.Context, startDate, endDate time.Time) (*domain.ParkingSummary, error) {
	startDate = startDate.UTC().Truncate(24 * time.Hour)
	endDate = endDate.UTC().Truncate(24 * time.Hour)
	return pu.parkingRepo.GetParkingSummary(ctx, startDate, endDate)
}
