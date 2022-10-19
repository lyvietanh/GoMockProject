package flight_grpc_handlers

import (
	db "anhlv11/go-mock-project/db/sqlc"
	"anhlv11/go-mock-project/pb"
	"anhlv11/go-mock-project/util"
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *FlightGrpcHandler) SearchOneWayFlight(ctx context.Context, in *pb.GrpcSearchOneWayFlightRequest) (*pb.GrpcSearchOneWayFlightResponse, error) {
	if len(in.OriginAirportCode) != 3 {
		return nil, status.Error(codes.Internal, "Origin airport code is not valid")
	}
	if len(in.DestinationAirportCode) != 3 {
		return nil, status.Error(codes.Internal, "Destination airport code is not valid")
	}
	if util.StartOfDay(in.DepartureDate.AsTime()).Unix() < util.StartOfDay(timestamppb.Now().AsTime()).Unix() {
		return nil, status.Error(codes.Internal, "Departure date is not valid")
	}

	outDepartureFlightOptions, err := getGrpcSfrFlightOptions(h, ctx, in.OriginAirportCode, in.DestinationAirportCode, in.DepartureDate.AsTime(), in.BookingClasses)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Departure flight|%v", err.Error()))
	}

	out := &pb.GrpcSearchOneWayFlightResponse{
		DepartureFlightOptions: outDepartureFlightOptions,
	}

	return out, nil
}

func (h *FlightGrpcHandler) SearchRoundTripFlight(ctx context.Context, in *pb.GrpcSearchRoundTripFlightRequest) (*pb.GrpcSearchRoundTripFlightResponse, error) {
	if len(in.OriginAirportCode) != 3 {
		return nil, status.Error(codes.Internal, "Origin airport code is not valid")
	}
	if len(in.DestinationAirportCode) != 3 {
		return nil, status.Error(codes.Internal, "Destination airport code is not valid")
	}
	if util.StartOfDay(in.DepartureDate.AsTime()).Unix() < util.StartOfDay(timestamppb.Now().AsTime()).Unix() {
		return nil, status.Error(codes.Internal, "Departure date is not valid")
	}
	if util.StartOfDay(in.ReturnDate.AsTime()).Unix() < util.StartOfDay(in.DepartureDate.AsTime()).Unix() {
		return nil, status.Error(codes.Internal, "Return date must be greater than Departure date")
	}

	outDepartureFlightOptions, err := getGrpcSfrFlightOptions(h, ctx, in.OriginAirportCode, in.DestinationAirportCode, in.DepartureDate.AsTime(), in.BookingClasses)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Departure flight|%v", err.Error()))
	}
	outReturnFlightOptions, err := getGrpcSfrFlightOptions(h, ctx, in.DestinationAirportCode, in.OriginAirportCode, in.ReturnDate.AsTime(), in.BookingClasses)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Return flight|%v", err.Error()))
	}

	out := &pb.GrpcSearchRoundTripFlightResponse{
		DepartureFlightOptions: outDepartureFlightOptions,
		ReturnFlightOptions:    outReturnFlightOptions,
	}

	return out, nil
}

func getGrpcSfrFlightOptions(h *FlightGrpcHandler, ctx context.Context, originAirportCode string, destinationAirportCode string, departureDate time.Time, bookingClasses string) ([]*pb.GrpcSfResFlightOption, error) {
	getAvailableFlightsByFieldsParams := db.GetAvailableFlightsByFieldsParams{
		Limit:                  100,
		Offset:                 0,
		OriginAirportCode:      originAirportCode,
		DestinationAirportCode: destinationAirportCode,
		DepartureDateTime:      departureDate,
		BookingClasses:         util.SplitStringToArray(bookingClasses),
		FlightNumber:           "",
	}

	flights, err := h.store.Queries.GetAvailableFlightsByFields(ctx, getAvailableFlightsByFieldsParams)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Could not found any flights|%v", err.Error()))
	}
	if flights != nil && len(flights) == 0 {
		return nil, status.Error(codes.Internal, "Could not found any flights")
	}

	grpcSfResFlightOptions := []*pb.GrpcSfResFlightOption{}
	for _, flight := range flights {
		grpcSfResFlightOption := &pb.GrpcSfResFlightOption{
			Id:                     flight.ID,
			OriginAirportCode:      flight.OriginAirportCode,
			DestinationAirportCode: flight.DestinationAirportCode,
			DepartureDateTime:      timestamppb.New(flight.DepartureDateTime),
			ArrivalDateTime:        timestamppb.New(flight.DepartureDateTime.Add(time.Duration(flight.FlightDuration) * time.Minute)),
			BookingClass:           flight.BookingClass,
			FlightNumber:           flight.FlightNumber,
			FlightDuration:         flight.FlightDuration,
			SeatRemaining:          flight.SeatAvailable,
			CurrencyCode:           flight.CurrencyCode,
			Price:                  flight.Price,
		}
		grpcSfResFlightOptions = append(grpcSfResFlightOptions, grpcSfResFlightOption)
	}
	return grpcSfResFlightOptions, nil
}
