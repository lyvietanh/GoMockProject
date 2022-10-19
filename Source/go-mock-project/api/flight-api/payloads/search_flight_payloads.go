package flight_api_payloads

import "anhlv11/go-mock-project/util"

type SearchOneWayFlightRequest struct {
	OriginDestination string `uri:"originDestination" json:"originDestination" binding:"required,len=7"`
	DepartureDate     string `uri:"departureDate" json:"departureDate" binding:"required,len=10"`
	Seats             int    `form:"seats,default=1" json:"seats"`
	BookingClasses    string `form:"bookingClasses" json:"bookingClasses"`
}

type SearchRoundTripFlightRequest struct {
	OriginDestination string `uri:"originDestination" json:"originDestination" binding:"required,len=7"`
	DepartureDate     string `uri:"departureDate" json:"departureDate" binding:"required,len=10"`
	ReturnDate        string `uri:"returnDate" json:"returnDate" binding:"required,len=10"`
	Seats             int    `form:"seats,default=1" json:"seats"`
	BookingClasses    string `form:"bookingClasses" json:"bookingClasses"`
}

type SearchOneWayFlightResponse struct {
	DepartureFlightOptions []SfrFlightOption `json:"departureFlightOptions"`
	RequestData            interface{}       `json:"requestData"`
}

type SearchRoundTripFlightResponse struct {
	DepartureFlightOptions []SfrFlightOption `json:"departureFlightOptions"`
	ReturnFlightOptions    []SfrFlightOption `json:"returnFlightOptions"`
	RequestData            interface{}       `json:"requestData"`
}

type SfrFlightOption struct {
	Id                     int64             `json:"id"`
	OriginAirportCode      string            `json:"originAirportCode"`
	DestinationAirportCode string            `json:"destinationAirportCode"`
	DepartureDateTime      util.JsonDateTime `json:"departureDateTime"`
	ArrivalDateTime        util.JsonDateTime `json:"arrivalDateTime"`
	BookingClass           string            `json:"bookingClass"`
	FlightNumber           string            `json:"flightNumber"`
	FlightDuration         int32             `json:"flightDuration"`
	SeatRemaining          int32             `json:"seatRemaining"`
	CurrencyCode           string            `json:"currencyCode"`
	Price                  float64           `json:"price"`
}
