package server

import (
	"supertaltest/infrastructure/postgresql"
	"supertaltest/internal/parking/usecase"
	"supertaltest/server/handler/api"
)

type ApiHandler struct {
	ParkingHandler        *api.ApiParkingHandler
	ParkingManagerHandler *api.ApiParkingManagerHandler
}

func NewApiHandler() *ApiHandler {
	db := postgresql.MustNewConnection()
	parkingRepo := postgresql.NewParkingPostgresqlRepository(db)

	guestParkingUsecase := usecase.NewGuestParkerUsecase(parkingRepo)
	parkingManagerUsecase := usecase.NewParkingManagerUsecase(parkingRepo)

	return &ApiHandler{
		ParkingHandler:        api.NewApiParkingHandler(guestParkingUsecase),
		ParkingManagerHandler: api.NewApiParkingManagerHandler(parkingManagerUsecase),
	}
}
