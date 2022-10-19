package booking_api_handlers

import (
	"anhlv11/go-mock-project/pb"
	"anhlv11/go-mock-project/util"
	"net/http"

	booking_api_payloads "anhlv11/go-mock-project/api/booking-api/payloads"

	"github.com/gin-gonic/gin"
)

func (h *BookingApiHandler) CancelFlight(ctx *gin.Context) {
	request := booking_api_payloads.CancelFlightRequest{}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusOK, util.CreateApiErrorResponse(BAD_REQUEST_ERROR, err.Error()))
		return
	}

	grpcCancelBookingRequest := &pb.GrpcCancelBookingRequest{
		ReservationCode: request.ReservationCode,
		LastName:        request.LastName,
		TravelDocNumber: request.TravelDocNumber,
	}

	grpcCancelBookingResponse, err := h.GrpcClient.CancelBooking(ctx.Request.Context(), grpcCancelBookingRequest)
	if err != nil {
		ctx.JSON(http.StatusOK, util.CreateCustomApiErrorResponse(BAD_REQUEST_ERROR, err.Error(), nil))
		return
	}

	response := booking_api_payloads.CancelFlightResponse{
		RequestData:     request,
		ReservationCode: grpcCancelBookingResponse.ReservationCode,
		CancelDate:      util.JsonDateTime(grpcCancelBookingResponse.CancelDate.AsTime()),
	}
	ctx.JSON(http.StatusOK, util.CreateSuccessResponse(response))
}
