package flight_api_payloads

import (
	"anhlv11/go-mock-project/util"
	"time"
)

type CreateFlightRequest struct {
	OriginAirportCode      string            `json:"originAirportCode" binding:"required,alpha,len=3"`
	DestinationAirportCode string            `json:"destinationAirportCode" binding:"required,alpha,len=3"`
	DepartureDateTime      util.JsonDateTime `json:"departureDateTime" binding:"required"`
	BookingClass           string            `json:"bookingClass" binding:"required,oneof=BUSINESS PREMIUM_ECONOMY ECONOMY"`
	FlightNumber           string            `json:"flightNumber" binding:"required,alphanum,max=8"`
	FlightDuration         int32             `json:"flightDuration" binding:"required,min=30,max=1440"`
	SeatRemaining          int32             `json:"seatRemaining" binding:"required,min=0,max=200"`
	CurrencyCode           string            `json:"currencyCode" binding:"required,alpha,len=3"`
	Price                  float64           `json:"price" binding:"required,min=1,max=99999999"`
}

type CreateFlightResponse struct {
	Id          int64       `json:"id"`
	CreateBy    string      `json:"createBy"`
	CreateAt    time.Time   `json:"createAt"`
	RequestData interface{} `json:"requestData"`
}
