package usecase

import (
	"context"
	entity "supertaltest/internal/parking/domain"
)

type ParkRepository interface {
	CreateParkingLot(ctx context.Context, lot *entity.ParkingLot) error
}

type ParkUsecase struct {
	parkRepository ParkRepository
}
