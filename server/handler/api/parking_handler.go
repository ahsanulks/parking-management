package api

import (
	"context"
	"net/http"
	"strconv"
	"supertaltest/internal/parking/domain"
	"time"

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

type ParkingVehicleResponse struct {
	TicketCode string
}

// @BasePath /api/v1

// ParkUserVehicle godoc
// @Summary park user vehicle
// @Schemes
// @Description parking user vehicle to obtain ticket code
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "1"
// @Success 200 object ParkingVehicleResponse
// @Router /parking-lots/{id}/park [post]
func (handler *ApiParkingHandler) ParkUserVehicle(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "invalid lots id")
		return
	}
	ticketCode, err := handler.parkingUsecase.ParkUserVehicle(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, ParkingVehicleResponse{
		TicketCode: ticketCode,
	})
}

type ExitVehicleResponse struct {
	TicketCode    string
	TotalFee      int
	ParkingSlotId int
	EntryTime     time.Time
	ExitTime      time.Time
}

// ExitUserVehicle godoc
// @Summary unpark user vehicle
// @Schemes
// @Description unpark user vehicle to obtain fee
// @Tags user
// @Accept json
// @Produce json
// @Param code path string true "00001-00001-1709942742"
// @Success 200 object ExitVehicleResponse
// @Router /tickets/{code}/exit [post]
func (handler *ApiParkingHandler) ExitUserVehicle(ctx *gin.Context) {
	ticket, err := handler.parkingUsecase.ExitUserPark(ctx.Request.Context(), ctx.Param("code"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, ExitVehicleResponse{
		TicketCode:    ticket.Code(),
		TotalFee:      *ticket.Fee(),
		ParkingSlotId: ticket.SlotID(),
		EntryTime:     ticket.EntiryTime(),
		ExitTime:      *ticket.ExitTime(),
	})
}
