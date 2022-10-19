package booking_api_payloads

type ViewBookingRequest struct {
	ReservationCode string `uri:"reservationCode" json:"reservationCode" binding:"required,len=6"`
	LastName        string `form:"lastName" json:"lastName"`
	TravelDocNumber string `form:"travelDocNumber" json:"travelDocNumber"`
}
