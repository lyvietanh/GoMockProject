package customer_api_handlers

import (
	customer_api_payloads "anhlv11/go-mock-project/api/customer-api/payloads"
	"anhlv11/go-mock-project/pb"
	"anhlv11/go-mock-project/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *CustomerApiHandler) ViewBookingHistory(ctx *gin.Context) {
	// need valid token to use this feature
	grpcVerifyAuthResponse, err := h.VerifyRolePermission(ctx, nil)
	if err != nil {
		ctx.JSON(http.StatusOK, util.CreateApiErrorResponse(BAD_REQUEST_ERROR, err.Error()))
		return
	}

	email := grpcVerifyAuthResponse.Email

	grpcViewBookingHistoryRequest := &pb.GrpcViewBookingHistoryRequest{
		Email: email,
	}

	grpcViewBookingHistoryResponse, err := h.UserGrpcClient.ViewBookingHistory(ctx.Request.Context(), grpcViewBookingHistoryRequest)
	if err != nil {
		ctx.JSON(http.StatusOK, util.CreateCustomApiErrorResponse(BAD_REQUEST_ERROR, err.Error(), nil))
		return
	}

	vbhResItems := []*customer_api_payloads.VbhResItem{}
	for _, grpcItem := range grpcViewBookingHistoryResponse.Items {
		vbhResItemFlightOptions := []*customer_api_payloads.VbhResItemFlightOption{}
		for _, grpcItemFlightOption := range grpcItem.FlightOptions {
			vbhResItemFlightOptions = append(vbhResItemFlightOptions, &customer_api_payloads.VbhResItemFlightOption{
				Id:                     grpcItemFlightOption.Id,
				OriginAirportCode:      grpcItemFlightOption.OriginAirportCode,
				DestinationAirportCode: grpcItemFlightOption.DestinationAirportCode,
				DepartureDateTime:      util.JsonDateTime(grpcItemFlightOption.DepartureDateTime.AsTime()),
				ArrivalDateTime:        util.JsonDateTime(grpcItemFlightOption.ArrivalDateTime.AsTime()),
				BookingClass:           grpcItemFlightOption.BookingClass,
				FlightNumber:           grpcItemFlightOption.FlightNumber,
				FlightDuration:         grpcItemFlightOption.FlightDuration,
				CurrencyCode:           grpcItemFlightOption.CurrencyCode,
				Price:                  grpcItemFlightOption.Price,
			})
		}
		vbhResItemPassengers := []*customer_api_payloads.VbhResItemPassenger{}
		for _, grpcItemPassenger := range grpcItem.Passengers {
			vbhResItemPassengers = append(vbhResItemPassengers, &customer_api_payloads.VbhResItemPassenger{
				Title:           grpcItemPassenger.Title,
				LastName:        grpcItemPassenger.LastName,
				FirstName:       grpcItemPassenger.FirstName,
				DateOfBirth:     util.JsonDate(grpcItemPassenger.DateOfBirth.AsTime()),
				TravelDocType:   grpcItemPassenger.TravelDocType,
				TravelDocNumber: grpcItemPassenger.TravelDocNumber,
				FfpNumber:       grpcItemPassenger.FfpNumber,
			})
		}
		vbhResItems = append(vbhResItems, &customer_api_payloads.VbhResItem{
			ReservationCode: grpcItem.ReservationCode,
			ContactEmail:    grpcItem.ContactEmail,
			ContactPhone:    grpcItem.ContactPhone,
			CurrencyCode:    grpcItem.CurrencyCode,
			TotalPrice:      grpcItem.TotalPrice,
			FlightOptions:   vbhResItemFlightOptions,
			Passengers:      vbhResItemPassengers,
		})
	}

	response := customer_api_payloads.ViewBookingHistoryResponse{
		Email: email,
		Items: vbhResItems,
	}
	ctx.JSON(http.StatusOK, util.CreateSuccessResponse(response))
}
