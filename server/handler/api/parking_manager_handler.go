package api

import (
	"context"
	"net/http"
	"supertaltest/internal/parking/domain"
	"supertaltest/internal/parking/usecase"

	"github.com/gin-gonic/gin"
)

type ParkingManagerUsecase interface {
	CreateParkingLot(ctx context.Context, parkingManager *domain.ParkingManager, request *usecase.RequestCreateParkingLot) (id int, err error)
}

type ApiParkingManagerHandler struct {
	usecase ParkingManagerUsecase
}

func NewApiParkingManagerHandler(usecase ParkingManagerUsecase) *ApiParkingManagerHandler {
	return &ApiParkingManagerHandler{
		usecase: usecase,
	}
}

type JsonRequestCreateParkingLot struct {
	NumSlot        int    `json:"numSlot"`
	ParkingLotName string `json:"name"`
}

func (handler *ApiParkingManagerHandler) CreateParkingLot(ctx *gin.Context) {
	params := &JsonRequestCreateParkingLot{}
	if err := ctx.BindJSON(params); err != nil {
		ctx.JSON(http.StatusBadRequest, "invalid body request")
		return
	}
	lotId, err := handler.usecase.CreateParkingLot(ctx.Request.Context(), domain.NewParkingManager(1), &usecase.RequestCreateParkingLot{
		NumSlot:        params.NumSlot,
		ParkingLotName: params.ParkingLotName,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]int{
		"parkingLotId": lotId,
	})
}
