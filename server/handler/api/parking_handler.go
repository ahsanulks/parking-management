package api

import (
	"context"
	"net/http"
	"strconv"
	"supertaltest/internal/parking/domain"

	"github.com/gin-gonic/gin"
)

type ParkingUsecase interface {
	ParkUserVehicle(ctx context.Context, lotId int) (ticketId string, err error)
	ExitUserPark(ctx context.Context, ticketCode string) (*domain.Ticket, error)
}

type ApiParkingHandler struct {
	parkingUsecase ParkingUsecase
}

func NewApiParkingHandler(parkingUsecase ParkingUsecase) *ApiParkingHandler {
	return &ApiParkingHandler{
		parkingUsecase: parkingUsecase,
	}
}

func (handler *ApiParkingHandler) ParkUserVehicle(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "invalid lots id")
		return
	}
	ticketCode, err := handler.parkingUsecase.ParkUserVehicle(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, map[string]string{
		"ticketCode": ticketCode,
	})
}

func (handler *ApiParkingHandler) ExitUserVehicle(ctx *gin.Context) {
	ticket, err := handler.parkingUsecase.ExitUserPark(ctx.Request.Context(), ctx.Param("code"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]any{
		"code":          ticket.Code(),
		"totalFee":      ticket.Fee(),
		"parkingSlotId": ticket.SlotID(),
		"entryTime":     ticket.EntiryTime(),
		"exitTime":      ticket.ExitTime(),
	})
}
