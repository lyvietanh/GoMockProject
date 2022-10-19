package booking_api_handlers

import (
	"anhlv11/go-mock-project/pb"
	"anhlv11/go-mock-project/util"
)

type BookingApiHandler struct {
	GrpcClient pb.BookingGrpcClient
}

func NewBookingApiHandler(grpcClient pb.BookingGrpcClient) BookingApiHandler {
	return BookingApiHandler{
		GrpcClient: grpcClient,
	}
}

var (
	BAD_REQUEST_ERROR    = util.NewApiError(400, "Bad Request")
	BUSINESS_LOGIC_ERROR = util.NewApiError(500, "Business Logic Error")
)
