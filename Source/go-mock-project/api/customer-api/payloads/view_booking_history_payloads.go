package customer_api_payloads

import "anhlv11/go-mock-project/util"

type ViewBookingHistoryRequest struct {
}

type ViewBookingHistoryResponse struct {
	Email string        `json:"email"`
	Items []*VbhResItem `json:"items"`
}

type VbhResItem struct {
	ReservationCode string                    `json:"reservationCode"`
	ContactEmail    string                    `json:"contactEmail"`
	ContactPhone    string                    `json:"contactPhone"`
	CurrencyCode    string                    `json:"currencyCode"`
	TotalPrice      float64                   `json:"totalPrice"`
	Passengers      []*VbhResItemPassenger    `json:"passengers"`
	FlightOptions   []*VbhResItemFlightOption `json:"flightOptions"`
}

type VbhResItemFlightOption struct {
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

type VbhResItemPassenger struct {
	Title           string        `json:"title"`
	LastName        string        `json:"lastName"`
	FirstName       string        `json:"firstName"`
	DateOfBirth     util.JsonDate `json:"dateOfBirth"`
	TravelDocType   string        `json:"travelDocType"`
	TravelDocNumber string        `json:"travelDocNumber"`
	FfpNumber       string        `json:"ffpNumber"`
}
