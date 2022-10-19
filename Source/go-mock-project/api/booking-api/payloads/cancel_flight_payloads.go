package booking_api_payloads

import "anhlv11/go-mock-project/util"

type CancelFlightRequest struct {
	ReservationCode string `json:"reservationCode" binding:"required,len=6"`
	LastName        string `json:"lastName" binding:"required,alpha,min=1,max=25"`
	TravelDocNumber string `json:"travelDocNumber" binding:"min=1,max=25"`
}

type CancelFlightResponse struct {
	ReservationCode string            `json:"reservationCode"`
	CancelDate      util.JsonDateTime `json:"cancelDate"`
	RequestData     interface{}       `json:"requestData"`
}
