package user_grpc_handlers

import (
	"anhlv11/go-mock-project/pb"
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *UserGrpcHandler) ViewBookingHistory(ctx context.Context, in *pb.GrpcViewBookingHistoryRequest) (*pb.GrpcViewBookingHistoryResponse, error) {
	email := strings.ToLower(strings.TrimSpace(in.Email))
	bookingTransactions, err := h.store.Queries.GetBookingTransactionsByContactEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("could not get booking transaction by email=%v|%v", email, err)
	}
	if len(bookingTransactions) == 0 {
		return nil, fmt.Errorf("could not get booking transaction by email=%v", email)
	}

	bookingTransactionIds := []int64{}
	for _, bookingTransaction := range bookingTransactions {
		bookingTransactionIds = append(bookingTransactionIds, bookingTransaction.ID)
	}

	bookingFlights, err := h.store.Queries.GetBookingFlightsByBookingTransactionIds(ctx, bookingTransactionIds)
	if err != nil {
		return nil, fmt.Errorf("could not get booking flights|%v", err)
	}

	bookingPassengers, err := h.store.Queries.GetBookingPassengersByBookingTransactionIds(ctx, bookingTransactionIds)
	if err != nil {
		return nil, fmt.Errorf("could not get booking passengers|%v", err)
	}

	grpcVbhResItems := []*pb.GrpcVbhResItem{}
	for _, bookingTransaction := range bookingTransactions {
		grpcVbhResItemFlightOptions := []*pb.GrpcVbhResItemFlightOption{}
		for _, bookingFlight := range bookingFlights {
			if !bookingFlight.BookingTransactionID.Valid || bookingFlight.BookingTransactionID.Int64 != bookingTransaction.ID {
				continue
			}
			flight, err := h.store.Queries.GetFlightById(ctx, bookingFlight.FlightID.Int64)
			if err != nil {
				grpcVbhResItemFlightOptions = append(grpcVbhResItemFlightOptions, &pb.GrpcVbhResItemFlightOption{
					Id: bookingFlight.ID,
				})
			} else {
				grpcVbhResItemFlightOptions = append(grpcVbhResItemFlightOptions, &pb.GrpcVbhResItemFlightOption{
					Id:                     bookingFlight.ID,
					OriginAirportCode:      flight.OriginAirportCode,
					DestinationAirportCode: flight.DestinationAirportCode,
					DepartureDateTime:      timestamppb.New(flight.DepartureDateTime),
					ArrivalDateTime:        timestamppb.New(flight.DepartureDateTime.Add(time.Duration(flight.FlightDuration) * time.Minute)),
					BookingClass:           flight.BookingClass,
					FlightNumber:           flight.FlightNumber,
					FlightDuration:         flight.FlightDuration,
					CurrencyCode:           flight.CurrencyCode,
					Price:                  flight.Price,
				})
			}
		}

		grpcVbhResItemPassengers := []*pb.GrpcVbhResItemPassenger{}
		for _, bookingPassenger := range bookingPassengers {
			if !bookingPassenger.BookingTransactionID.Valid || bookingPassenger.BookingTransactionID.Int64 != bookingTransaction.ID {
				continue
			}
			grpcVbhResItemPassengers = append(grpcVbhResItemPassengers, &pb.GrpcVbhResItemPassenger{
				Title:           bookingPassenger.Title.String,
				LastName:        bookingPassenger.LastName.String,
				FirstName:       bookingPassenger.FirstName.String,
				DateOfBirth:     timestamppb.New(bookingPassenger.DateOfBirth.Time),
				TravelDocType:   bookingPassenger.TravelDocType.String,
				TravelDocNumber: bookingPassenger.TravelDocNumber.String,
				FfpNumber:       bookingPassenger.FfpNumber.String,
			})
		}

		grpcVbhResItems = append(grpcVbhResItems, &pb.GrpcVbhResItem{
			ReservationCode: bookingTransaction.ReservationCode.String,
			ContactEmail:    bookingTransaction.ContactEmail,
			ContactPhone:    bookingTransaction.ContactPhone,
			CurrencyCode:    bookingTransaction.CurrencyCode.String,
			TotalPrice:      bookingTransaction.TotalPrice.Float64,
			FlightOptions:   grpcVbhResItemFlightOptions,
			Passengers:      grpcVbhResItemPassengers,
		})
	}

	out := &pb.GrpcViewBookingHistoryResponse{
		Email: email,
		Items: grpcVbhResItems,
	}
	return out, nil
}
