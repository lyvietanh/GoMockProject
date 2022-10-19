package booking_api_handlers

import (
	"anhlv11/go-mock-project/pb"
	"anhlv11/go-mock-project/util"
	"net/http"

	booking_api_payloads "anhlv11/go-mock-project/api/booking-api/payloads"

	"github.com/gin-gonic/gin"
)

func (h *BookingApiHandler) ViewBooking(ctx *gin.Context) {
	request := booking_api_payloads.ViewBookingRequest{}

	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusOK, util.CreateApiErrorResponse(BAD_REQUEST_ERROR, err.Error()))
		return
	}
	if err := ctx.ShouldBindQuery(&request); err != nil {
		ctx.JSON(http.StatusOK, util.CreateApiErrorResponse(BAD_REQUEST_ERROR, err.Error()))
		return
	}

	grpcViewBookingRequest := &pb.GrpcViewBookingRequest{
		ReservationCode: request.ReservationCode,
		LastName:        request.LastName,
		TravelDocNumber: request.TravelDocNumber,
	}

	grpcBookFlightResponse, err := h.GrpcClient.ViewBooking(ctx.Request.Context(), grpcViewBookingRequest)
	if err != nil {
		ctx.JSON(http.StatusOK, util.CreateCustomApiErrorResponse(BAD_REQUEST_ERROR, err.Error(), nil))
		return
	}

	bfResPassengers := []*booking_api_payloads.BfResPassenger{}
	for _, grpcPassenger := range grpcBookFlightResponse.Passengers {
		bfResPassengers = append(bfResPassengers, &booking_api_payloads.BfResPassenger{
			Title:           grpcPassenger.Title,
			LastName:        grpcPassenger.LastName,
			FirstName:       grpcPassenger.FirstName,
			DateOfBirth:     util.JsonDate(grpcPassenger.DateOfBirth.AsTime()),
			TravelDocType:   grpcPassenger.TravelDocType,
			TravelDocNumber: grpcPassenger.TravelDocNumber,
			FfpNumber:       grpcPassenger.FfpNumber,
		})
	}

	bfResFlightOptions := []*booking_api_payloads.BfResFlightOption{}
	for _, grpcFlightOption := range grpcBookFlightResponse.FlightOptions {
		bfResFlightOptions = append(bfResFlightOptions, &booking_api_payloads.BfResFlightOption{
			Id:                     grpcFlightOption.Id,
			OriginAirportCode:      grpcFlightOption.OriginAirportCode,
			DestinationAirportCode: grpcFlightOption.DestinationAirportCode,
			DepartureDateTime:      util.JsonDateTime(grpcFlightOption.DepartureDateTime.AsTime()),
			ArrivalDateTime:        util.JsonDateTime(grpcFlightOption.ArrivalDateTime.AsTime()),
			BookingClass:           grpcFlightOption.BookingClass,
			FlightNumber:           grpcFlightOption.FlightNumber,
			FlightDuration:         grpcFlightOption.FlightDuration,
			CurrencyCode:           grpcFlightOption.CurrencyCode,
			Price:                  grpcFlightOption.Price,
		})
	}

	response := booking_api_payloads.BookFlightResponse{
		RequestData:     request,
		ReservationCode: grpcBookFlightResponse.ReservationCode,
		ContactEmail:    grpcBookFlightResponse.ContactEmail,
		ContactPhone:    grpcBookFlightResponse.ContactPhone,
		CurrencyCode:    grpcBookFlightResponse.CurrencyCode,
		TotalPrice:      grpcBookFlightResponse.TotalPrice,
		Passengers:      bfResPassengers,
		FlightOptions:   bfResFlightOptions,
	}
	ctx.JSON(http.StatusOK, util.CreateSuccessResponse(response))
}
