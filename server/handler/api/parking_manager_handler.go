package api

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"supertaltest/internal/parking/domain"
	"supertaltest/internal/parking/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

type ParkingManagerUsecase interface {
	CreateParkingLot(ctx context.Context, parkingManager *domain.ParkingManager, request *usecase.RequestCreateParkingLot) (id int, err error)
	GetLotStatus(ctx context.Context, lotId int) (*domain.ParkingLotStatus, error)
	ToggleParkingSlotToMaintenance(ctx context.Context, parkingManager *domain.ParkingManager, slotId int) error
	GetParkingSummary(ctx context.Context, startDate, endDate time.Time) (*domain.ParkingSummary, error)
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

func (handler *ApiParkingManagerHandler) GetParkingSlotStatus(ctx *gin.Context) {
	lotId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "invalid lots id")
		return
	}
	parkingLotStatus, err := handler.usecase.GetLotStatus(ctx, lotId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, parkingLotStatus)
}

func (handler *ApiParkingManagerHandler) ToggleParkingSlotMaintenance(ctx *gin.Context) {
	slotId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "invalid lots id")
		return
	}

	err = handler.usecase.ToggleParkingSlotToMaintenance(ctx, domain.NewParkingManager(1), slotId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, "ok")
}

func (handler *ApiParkingManagerHandler) ParkingSummary(ctx *gin.Context) {
	startDateStr := ctx.Query("startDate")
	endDateStr := ctx.Query("endDate")

	var startDate, endDate time.Time
	var err error

	if startDateStr != "" {
		startDate, err = time.Parse("2006-01-02", startDateStr)
		fmt.Println("111", startDate)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid start date format"})
			return
		}
	}

	if endDateStr != "" {
		endDate, err = time.Parse("2006-01-02", endDateStr)
		fmt.Println("1112222", endDate)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid end date format"})
			return
		}
	}
	fmt.Println(startDate, endDate)

	if startDate.IsZero() {
		startDate = time.Now().AddDate(0, 0, -7)
	}
	if endDate.IsZero() {
		endDate = time.Now()
	}
	summary, err := handler.usecase.GetParkingSummary(ctx.Request.Context(), startDate, endDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, summary)
}
