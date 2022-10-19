package flight_api_handlers

import (
	flight_api_payloads "anhlv11/go-mock-project/api/flight-api/payloads"
	"anhlv11/go-mock-project/constant"
	"anhlv11/go-mock-project/pb"
	"anhlv11/go-mock-project/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *FlightApiHandler) CreateFlight(ctx *gin.Context) {
	// need ADMINISTRATOR permission to use this feature
	if _, err := h.VerifyRolePermission(ctx, &[]string{constant.ROLE_ID_ADMINISTRATOR}); err != nil {
		ctx.JSON(http.StatusOK, util.CreateApiErrorResponse(BAD_REQUEST_ERROR, err.Error()))
		return
	}

	// binding request
	request := flight_api_payloads.CreateFlightRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusOK, util.CreateApiErrorResponse(BAD_REQUEST_ERROR, err.Error()))
		return
	}

	grpcCreateFlightRequest := &pb.GrpcCreateFlightRequest{
		OriginAirportCode:      request.OriginAirportCode,
		DestinationAirportCode: request.DestinationAirportCode,
		DepartureDateTime:      timestamppb.New(time.Time(request.DepartureDateTime)),
		BookingClass:           request.BookingClass,
		FlightNumber:           request.FlightNumber,
		FlightDuration:         request.FlightDuration,
		SeatRemaining:          request.SeatRemaining,
		CurrencyCode:           request.CurrencyCode,
		Price:                  request.Price,
	}
	grpcCreateFlightResponse, err := h.FlightGrpcClient.CreateFlight(ctx.Request.Context(), grpcCreateFlightRequest)
	if err != nil {
		ctx.JSON(http.StatusOK, util.CreateCustomApiErrorResponse(BAD_REQUEST_ERROR, err.Error(), nil))
		return
	}

	response := flight_api_payloads.CreateFlightResponse{
		RequestData: request,
		Id:          grpcCreateFlightResponse.Id,
		CreateBy:    grpcCreateFlightResponse.CreateBy,
		CreateAt:    grpcCreateFlightResponse.CreateAt.AsTime(),
	}
	ctx.JSON(http.StatusOK, util.CreateSuccessResponse(response))
}
