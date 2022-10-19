package flight_api_handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/timestamppb"

	flight_api_payloads "anhlv11/go-mock-project/api/flight-api/payloads"
	"anhlv11/go-mock-project/constant"
	"anhlv11/go-mock-project/pb"
	"anhlv11/go-mock-project/util"
)

func (h *FlightApiHandler) UpdateFlight(ctx *gin.Context) {
	// need ADMINISTRATOR permission to use this feature
	if _, err := h.VerifyRolePermission(ctx, &[]string{constant.ROLE_ID_ADMINISTRATOR}); err != nil {
		ctx.JSON(http.StatusOK, util.CreateApiErrorResponse(BAD_REQUEST_ERROR, err.Error()))
		return
	}

	// binding request
	request := flight_api_payloads.UpdateFlightRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusOK, util.CreateApiErrorResponse(BAD_REQUEST_ERROR, err.Error()))
		return
	}

	grpcUpdateFlightRequest := &pb.GrpcUpdateFlightRequest{
		Id: request.Identifier.Id,
		// OriginAirportCode:      request.OriginAirportCode,
		// DestinationAirportCode: request.DestinationAirportCode,
		DepartureDateTime: timestamppb.New(time.Time(request.DepartureDateTime)),
		BookingClass:      request.BookingClass,
		FlightNumber:      request.FlightNumber,
		FlightDuration:    request.FlightDuration,
		SeatRemaining:     request.SeatRemaining,
		CurrencyCode:      request.CurrencyCode,
		Price:             request.Price,
		Enabled:           request.Enabled,
	}
	grpcUpdateFlightResponse, err := h.FlightGrpcClient.UpdateFlight(ctx.Request.Context(), grpcUpdateFlightRequest)
	if err != nil {
		ctx.JSON(http.StatusOK, util.CreateCustomApiErrorResponse(BAD_REQUEST_ERROR, err.Error(), nil))
		return
	}

	response := flight_api_payloads.UpdateFlightResponse{
		RequestData: request,
		UpdateBy:    grpcUpdateFlightResponse.UpdateBy,
		UpdateAt:    grpcUpdateFlightResponse.UpdateAt.AsTime(),
	}
	ctx.JSON(http.StatusOK, util.CreateSuccessResponse(response))
}
