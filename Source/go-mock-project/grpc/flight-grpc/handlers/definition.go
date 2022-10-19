package flight_grpc_handlers

import (
	db "anhlv11/go-mock-project/db/sqlc"
	"anhlv11/go-mock-project/pb"
	"anhlv11/go-mock-project/util"
)

type FlightGrpcHandler struct {
	pb.UnimplementedFlightGrpcServer
	config util.GrpcConfig
	store  *db.Store
}

func NewFlightGrpcHandler(config util.GrpcConfig, store *db.Store) (*FlightGrpcHandler, error) {
	return &FlightGrpcHandler{
		config: config,
		store:  store,
	}, nil
}
