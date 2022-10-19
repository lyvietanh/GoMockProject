package booking_api_payloads

import "anhlv11/go-mock-project/util"

type BookFlightRequest struct {
	CurrencyCode    string            `json:"currencyCode" binding:"required,alpha,len=3"`
	ContactEmail    string            `json:"contactEmail" binding:"min=1,max=25"`
	ContactPhone    string            `json:"contactPhone" binding:"min=1,max=25"`
	FlightOptionIds []int64           `json:"flightOptionIds" binding:"required,min=1,max=6"`
	Passengers      []*BfReqPassenger `json:"passengers" binding:"required,min=1,max=9"`
}

type BfReqPassenger struct {
	Title           string        `json:"title" binding:"required,alpha,min=2,max=8"`
	LastName        string        `json:"lastName" binding:"required,alpha,min=1,max=25"`
	FirstName       string        `json:"firstName" binding:"required,alpha,min=1,max=25"`
	DateOfBirth     util.JsonDate `json:"dateOfBirth" binding:"required"`
	TravelDocType   string        `json:"travelDocType" binding:"alpha,len=1,oneof=P I"`
	TravelDocNumber string        `json:"travelDocNumber" binding:"min=1,max=25"`
	FfpNumber       string        `json:"ffpNumber" binding:"min=1,max=20"`
}

type BookFlightResponse struct {
	ReservationCode string               `json:"reservationCode"`
	ContactEmail    string               `json:"contactEmail"`
	ContactPhone    string               `json:"contactPhone"`
	CurrencyCode    string               `json:"currencyCode"`
	TotalPrice      float64              `json:"totalPrice"`
	Passengers      []*BfResPassenger    `json:"passengers"`
	FlightOptions   []*BfResFlightOption `json:"flightOptions"`
	RequestData     interface{}          `json:"requestData"`
}

type BfResFlightOption struct {
	Id                     int64             `json:"id"`
	OriginAirportCode      string            `json:"originAirportCode"`
	DestinationAirportCode string            `json:"destinationAirportCode"`
	DepartureDateTime      util.JsonDateTime `json:"departureDateTime"`
	ArrivalDateTime        util.JsonDateTime `json:"arrivalDateTime"`
	BookingClass           string            `json:"bookingClass"`
	FlightNumber           string            `json:"flightNumber"`
	FlightDuration         int32             `json:"flightDuration"`
	CurrencyCode           string            `json:"currencyCode"`
	Price                  float64           `json:"price"`
}

type BfResPassenger struct {
	Title           string        `json:"title"`
	LastName        string        `json:"lastName"`
	FirstName       string        `json:"firstName"`
	DateOfBirth     util.JsonDate `json:"dateOfBirth"`
	TravelDocType   string        `json:"travelDocType"`
	TravelDocNumber string        `json:"travelDocNumber"`
	FfpNumber       string        `json:"ffpNumber"`
}
