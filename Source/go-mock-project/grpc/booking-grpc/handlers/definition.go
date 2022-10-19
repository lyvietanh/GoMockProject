package booking_grpc_handlers

import (
	db "anhlv11/go-mock-project/db/sqlc"
	"anhlv11/go-mock-project/pb"
	"anhlv11/go-mock-project/util"
	"sync"
)

type BookingGrpcHandler struct {
	pb.UnimplementedBookingGrpcServer
	config util.GrpcConfig
	store  *db.Store
	mutex  *sync.Mutex
}

func NewBookingGrpcHandler(config util.GrpcConfig, store *db.Store) (*BookingGrpcHandler, error) {
	return &BookingGrpcHandler{
		config: config,
		store:  store,
		mutex:  &sync.Mutex{},
	}, nil
}

const (
	BOOKING_TRANSACTION_STATUS_BOOKED   string = "BOOKED"
	BOOKING_TRANSACTION_STATUS_CANCELED string = "CANCELED"
)
