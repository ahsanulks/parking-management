package usecase

import (
	"context"
	entity "supertaltest/internal/parking/domain"
)

func (pu *ParkUsecase) CreateParkingLot(ctx context.Context, numSlot int) (id int, err error) {
	manager := entity.ParkingManager{}
	manager.CreateParkingLot(numSlot)

	err = pu.parkRepository.CreateParkingLot(ctx, manager.ParkingLots[0])
	return 1231, err
}
