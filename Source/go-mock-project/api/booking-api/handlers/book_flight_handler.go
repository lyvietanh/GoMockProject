package booking_api_handlers

import (
	"anhlv11/go-mock-project/pb"
	"anhlv11/go-mock-project/util"
	"net/http"
	"time"

	booking_api_payloads "anhlv11/go-mock-project/api/booking-api/payloads"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *BookingApiHandler) BookFlight(ctx *gin.Context) {
	request := booking_api_payloads.BookFlightRequest{}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusOK, util.CreateApiErrorResponse(BAD_REQUEST_ERROR, err.Error()))
		return
	}

	grpcBfReqPassengers := []*pb.GrpcBfReqPassenger{}
	for _, requestPassenger := range request.Passengers {
		grpcBfReqPassengers = append(grpcBfReqPassengers, &pb.GrpcBfReqPassenger{
			Title:           requestPassenger.Title,
			LastName:        requestPassenger.LastName,
			FirstName:       requestPassenger.FirstName,
			DateOfBirth:     timestamppb.New(time.Time(requestPassenger.DateOfBirth)),
			TravelDocType:   requestPassenger.TravelDocType,
			TravelDocNumber: requestPassenger.TravelDocNumber,
			FfpNumber:       requestPassenger.FfpNumber,
		})
	}
	grpcBookFlightRequest := &pb.GrpcBookFlightRequest{
		CurrencyCode:    request.CurrencyCode,
		ContactEmail:    request.ContactEmail,
		ContactPhone:    request.ContactPhone,
		Passengers:      grpcBfReqPassengers,
		FlightOptionIds: request.FlightOptionIds,
	}

	grpcBookFlightResponse, err := h.GrpcClient.BookFlight(ctx.Request.Context(), grpcBookFlightRequest)
	if err != nil {
		ctx.JSON(http.StatusOK, util.CreateCustomApiErrorResponse(BAD_REQUEST_ERROR, err.Error(), nil))
		return
	}

	responsePassengers := []*booking_api_payloads.BfResPassenger{}
	for _, grpcPassenger := range grpcBookFlightResponse.Passengers {
		responsePassengers = append(responsePassengers, &booking_api_payloads.BfResPassenger{
			Title:           grpcPassenger.Title,
			LastName:        grpcPassenger.LastName,
			FirstName:       grpcPassenger.FirstName,
			DateOfBirth:     util.JsonDate(grpcPassenger.DateOfBirth.AsTime()),
			TravelDocType:   grpcPassenger.TravelDocType,
			TravelDocNumber: grpcPassenger.TravelDocNumber,
			FfpNumber:       grpcPassenger.FfpNumber,
		})
	}

	responseFlightOptions := []*booking_api_payloads.BfResFlightOption{}
	for _, grpcFlightOption := range grpcBookFlightResponse.FlightOptions {
		responseFlightOptions = append(responseFlightOptions, &booking_api_payloads.BfResFlightOption{
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
		CurrencyCode:    grpcBookFlightResponse.CurrencyCode,
		TotalPrice:      grpcBookFlightResponse.TotalPrice,
		Passengers:      responsePassengers,
		FlightOptions:   responseFlightOptions,
	}
	ctx.JSON(http.StatusOK, util.CreateSuccessResponse(response))
}
