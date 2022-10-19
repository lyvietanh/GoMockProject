package flight_grpc_handlers

import (
	db "anhlv11/go-mock-project/db/sqlc"
	"anhlv11/go-mock-project/pb"
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *FlightGrpcHandler) CreateFlight(ctx context.Context, in *pb.GrpcCreateFlightRequest) (*pb.GrpcCreateFlightResponse, error) {
	getFlightsByFieldsParams := db.GetFlightsByFieldsParams{
		Offset:                 0,
		Limit:                  10,
		OriginAirportCode:      in.OriginAirportCode,
		DestinationAirportCode: in.DestinationAirportCode,
		DepartureDateTime:      time.Time(in.DepartureDateTime.AsTime()),
		BookingClass:           in.BookingClass,
		FlightNumber:           in.FlightNumber,
	}

	existedFlights, err := h.store.Queries.GetFlightsByFields(ctx, getFlightsByFieldsParams)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Can not lookup existed flight before creating|%v", err))
	}

	if len(existedFlights) > 0 {
		return nil, status.Error(codes.Internal, "This flight has existed")
	}

	createAt := time.Now()

	createFlightParams := db.CreateFlightParams{
		OriginAirportCode:      strings.ToUpper(in.OriginAirportCode),
		DestinationAirportCode: strings.ToUpper(in.DestinationAirportCode),
		DepartureDateTime:      time.Time(in.DepartureDateTime.AsTime()),
		BookingClass:           strings.ToUpper(in.BookingClass),
		FlightNumber:           strings.ToUpper(in.FlightNumber),
		FlightDuration:         in.FlightDuration,
		Enabled:                true,
		SeatAvailable:          in.SeatRemaining,
		CurrencyCode:           in.CurrencyCode,
		Price:                  in.Price,
		CreateDate:             sql.NullTime{Time: createAt, Valid: true},
	}

	flight, err := h.store.Queries.CreateFlight(ctx, createFlightParams)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Can not create flight|%v", err))
	}

	out := &pb.GrpcCreateFlightResponse{
		Id:       flight.ID,
		CreateAt: timestamppb.New(createAt),
	}

	return out, nil
}
