package server

import (
	"supertaltest/infrastructure/postgresql"
	"supertaltest/internal/parking/usecase"
	"supertaltest/server/handler/api"
)

type ApiHandler struct {
	ParkingHandler api.ApiParkingHandler
}

func NewApiHandler() *ApiHandler {
	db := postgresql.MustNewConnection()
	parkingRepo := postgresql.NewParkingPostgresqlRepository(db)
	guestParkingUsecase := usecase.NewGuestParkerUsecase(parkingRepo)
	return &ApiHandler{
		ParkingHandler: *api.NewApiParkingHandler(guestParkingUsecase),
	}
}
