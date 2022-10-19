package flight_api_handlers

import (
	flight_api_payloads "anhlv11/go-mock-project/api/flight-api/payloads"
	"anhlv11/go-mock-project/pb"
	"anhlv11/go-mock-project/util"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *FlightApiHandler) SearchOneWayFlight(ctx *gin.Context) {
	request := flight_api_payloads.SearchOneWayFlightRequest{}
	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusOK, util.CreateApiErrorResponse(BAD_REQUEST_ERROR, err.Error()))
		return
	}

	originDestinationCols := strings.Split(request.OriginDestination, "-")
	if len(originDestinationCols) != 2 {
		ctx.JSON(http.StatusOK, util.CreateCustomApiErrorResponse(BAD_REQUEST_ERROR, "Origin/Destination is not valid", nil))
		return
	}

	originAirportCode := originDestinationCols[0]
	destinationAirportCode := originDestinationCols[1]
	if len(originAirportCode) != 3 {
		ctx.JSON(http.StatusOK, util.CreateCustomApiErrorResponse(BAD_REQUEST_ERROR, "Origin airport code is not valid", nil))
		return
	}
	if len(destinationAirportCode) != 3 {
		ctx.JSON(http.StatusOK, util.CreateCustomApiErrorResponse(BAD_REQUEST_ERROR, "Destination airport code is not valid", nil))
		return
	}

	departureDate, err := util.ConvertToIsoDate(request.DepartureDate)
	if err != nil {
		ctx.JSON(http.StatusOK, util.CreateCustomApiErrorResponse(BAD_REQUEST_ERROR, "Departure date is not valid", err.Error()))
		return
	}

	if err := ctx.ShouldBindQuery(&request); err != nil {
		ctx.JSON(http.StatusOK, util.CreateApiErrorResponse(BAD_REQUEST_ERROR, err.Error()))
		return
	}

	grpcSearchOneWayFlightRequest := &pb.GrpcSearchOneWayFlightRequest{
		OriginAirportCode:      originAirportCode,
		DestinationAirportCode: destinationAirportCode,
		DepartureDate:          timestamppb.New(departureDate),
		BookingClasses:         request.BookingClasses,
		Seats:                  int32(request.Seats),
	}
	grpcSearchOneWayFlightResponse, err := h.FlightGrpcClient.SearchOneWayFlight(ctx.Request.Context(), grpcSearchOneWayFlightRequest)
	if err != nil {
		ctx.JSON(http.StatusOK, util.CreateCustomApiErrorResponse(BAD_REQUEST_ERROR, err.Error(), nil))
		return
	}

	sfrFlightOptions := []flight_api_payloads.SfrFlightOption{}
	for _, grpcSfrFlightOption := range grpcSearchOneWayFlightResponse.DepartureFlightOptions {
		sfrFlightOptions = append(sfrFlightOptions, flight_api_payloads.SfrFlightOption{
			Id:                     grpcSfrFlightOption.Id,
			OriginAirportCode:      grpcSfrFlightOption.OriginAirportCode,
			DestinationAirportCode: grpcSfrFlightOption.DestinationAirportCode,
			DepartureDateTime:      util.JsonDateTime(grpcSfrFlightOption.DepartureDateTime.AsTime()),
			ArrivalDateTime:        util.JsonDateTime(grpcSfrFlightOption.ArrivalDateTime.AsTime()),
			BookingClass:           grpcSfrFlightOption.BookingClass,
			FlightNumber:           grpcSfrFlightOption.FlightNumber,
			FlightDuration:         grpcSfrFlightOption.FlightDuration,
			SeatRemaining:          grpcSfrFlightOption.SeatRemaining,
			CurrencyCode:           grpcSfrFlightOption.CurrencyCode,
			Price:                  grpcSfrFlightOption.Price,
		})
	}

	response := flight_api_payloads.SearchOneWayFlightResponse{
		RequestData:            request,
		DepartureFlightOptions: sfrFlightOptions,
	}
	ctx.JSON(http.StatusOK, util.CreateSuccessResponse(response))
}

func (h *FlightApiHandler) SearchRoundTripFlight(ctx *gin.Context) {
	// binding request
	request := flight_api_payloads.SearchRoundTripFlightRequest{}
	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusOK, util.CreateApiErrorResponse(BAD_REQUEST_ERROR, err.Error()))
		return
	}

	originDestinationCols := strings.Split(request.OriginDestination, "-")
	if len(originDestinationCols) != 2 {
		ctx.JSON(http.StatusOK, util.CreateCustomApiErrorResponse(BAD_REQUEST_ERROR, "Origin/Destination is not valid", nil))
		return
	}

	originAirportCode := originDestinationCols[0]
	destinationAirportCode := originDestinationCols[1]
	if len(originAirportCode) != 3 {
		ctx.JSON(http.StatusOK, util.CreateCustomApiErrorResponse(BAD_REQUEST_ERROR, "Origin airport code is not valid", nil))
		return
	}
	if len(destinationAirportCode) != 3 {
		ctx.JSON(http.StatusOK, util.CreateCustomApiErrorResponse(BAD_REQUEST_ERROR, "Destination airport code is not valid", nil))
		return
	}

	departureDate, err := util.ConvertToIsoDate(request.DepartureDate)
	if err != nil {
		ctx.JSON(http.StatusOK, util.CreateCustomApiErrorResponse(BAD_REQUEST_ERROR, "Departure date is not valid", err.Error()))
		return
	}
	returnDate, err := util.ConvertToIsoDate(request.ReturnDate)
	if err != nil {
		ctx.JSON(http.StatusOK, util.CreateCustomApiErrorResponse(BAD_REQUEST_ERROR, "Return date is not valid", err.Error()))
		return
	}

	if err := ctx.ShouldBindQuery(&request); err != nil {
		ctx.JSON(http.StatusOK, util.CreateApiErrorResponse(BAD_REQUEST_ERROR, err.Error()))
		return
	}

	grpcSearchRoundTripFlightRequest := &pb.GrpcSearchRoundTripFlightRequest{
		OriginAirportCode:      originAirportCode,
		DestinationAirportCode: destinationAirportCode,
		DepartureDate:          timestamppb.New(departureDate),
		ReturnDate:             timestamppb.New(returnDate),
		BookingClasses:         request.BookingClasses,
		Seats:                  int32(request.Seats),
	}
	grpcSearchRoundTripFlightResponse, err := h.FlightGrpcClient.SearchRoundTripFlight(ctx.Request.Context(), grpcSearchRoundTripFlightRequest)
	if err != nil {
		ctx.JSON(http.StatusOK, util.CreateCustomApiErrorResponse(BAD_REQUEST_ERROR, err.Error(), nil))
		return
	}

	responseDepartureFlightOptions := []flight_api_payloads.SfrFlightOption{}
	for _, grpcSfrFlightOption := range grpcSearchRoundTripFlightResponse.DepartureFlightOptions {
		responseDepartureFlightOptions = append(responseDepartureFlightOptions, flight_api_payloads.SfrFlightOption{
			Id:                     grpcSfrFlightOption.Id,
			OriginAirportCode:      grpcSfrFlightOption.OriginAirportCode,
			DestinationAirportCode: grpcSfrFlightOption.DestinationAirportCode,
			DepartureDateTime:      util.JsonDateTime(grpcSfrFlightOption.DepartureDateTime.AsTime()),
			ArrivalDateTime:        util.JsonDateTime(grpcSfrFlightOption.ArrivalDateTime.AsTime()),
			BookingClass:           grpcSfrFlightOption.BookingClass,
			FlightNumber:           grpcSfrFlightOption.FlightNumber,
			FlightDuration:         grpcSfrFlightOption.FlightDuration,
			SeatRemaining:          grpcSfrFlightOption.SeatRemaining,
			CurrencyCode:           grpcSfrFlightOption.CurrencyCode,
			Price:                  grpcSfrFlightOption.Price,
		})
	}
	responseReturnFlightOptions := []flight_api_payloads.SfrFlightOption{}
	for _, grpcSfrFlightOption := range grpcSearchRoundTripFlightResponse.ReturnFlightOptions {
		responseReturnFlightOptions = append(responseReturnFlightOptions, flight_api_payloads.SfrFlightOption{
			Id:                     grpcSfrFlightOption.Id,
			OriginAirportCode:      grpcSfrFlightOption.OriginAirportCode,
			DestinationAirportCode: grpcSfrFlightOption.DestinationAirportCode,
			DepartureDateTime:      util.JsonDateTime(grpcSfrFlightOption.DepartureDateTime.AsTime()),
			ArrivalDateTime:        util.JsonDateTime(grpcSfrFlightOption.ArrivalDateTime.AsTime()),
			BookingClass:           grpcSfrFlightOption.BookingClass,
			FlightNumber:           grpcSfrFlightOption.FlightNumber,
			FlightDuration:         grpcSfrFlightOption.FlightDuration,
			SeatRemaining:          grpcSfrFlightOption.SeatRemaining,
			CurrencyCode:           grpcSfrFlightOption.CurrencyCode,
			Price:                  grpcSfrFlightOption.Price,
		})
	}
	response := flight_api_payloads.SearchRoundTripFlightResponse{
		RequestData:            request,
		DepartureFlightOptions: responseDepartureFlightOptions,
		ReturnFlightOptions:    responseReturnFlightOptions,
	}
	ctx.JSON(http.StatusOK, util.CreateSuccessResponse(response))
}
