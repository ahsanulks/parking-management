package main

import (
	"context"
	"fmt"
	"supertaltest/infrastructure/postgresql"
	"supertaltest/internal/parking/usecase"

	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()

	db := postgresql.MustNewConnection()
	parkingRepo := postgresql.NewParkingPostgresqlRepository(db)
	uc := usecase.NewGuestParkerUsecase(parkingRepo)
	a, err := uc.ParkUserVehicle(context.Background(), 1)
	fmt.Println(a, err)
	// a2, err := uc.ParkUserVehicle(context.Background(), 1)
	// fmt.Println(a2, err)
	// b, err := uc.ExitUserPark(context.Background(), a)
	// fmt.Println(b, err)
	// b2, err := uc.ExitUserPark(context.Background(), a2)
	// fmt.Println(b2, err)
}
