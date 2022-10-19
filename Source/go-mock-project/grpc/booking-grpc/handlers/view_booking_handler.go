package booking_grpc_handlers

import (
	db "anhlv11/go-mock-project/db/sqlc"
	"anhlv11/go-mock-project/pb"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *BookingGrpcHandler) ViewBooking(ctx context.Context, in *pb.GrpcViewBookingRequest) (*pb.GrpcBookFlightResponse, error) {
	lookupReservationParams := db.LookupReservationParams{
		Limit:           1,
		Offset:          0,
		ReservationCode: sql.NullString{String: in.ReservationCode, Valid: true},
		LastName:        sql.NullString{String: in.LastName, Valid: true},
		TravelDocNumber: sql.NullString{String: in.TravelDocNumber, Valid: true},
	}
	lookupReservationRows, err := h.store.Queries.LookupReservation(ctx, lookupReservationParams)
	if err != nil {
		return nil, fmt.Errorf("could not get booking transactions by reservationCode=%v lastName=%v travelDocNumber=%v|%v", in.ReservationCode, in.LastName, in.TravelDocNumber, err)
	}
	if len(lookupReservationRows) == 0 {
		return nil, fmt.Errorf("could not found any booking transactions by reservationCode=%v lastName=%v travelDocNumber=%v", in.ReservationCode, in.LastName, in.TravelDocNumber)
	}

	lookupReservationRow := lookupReservationRows[0]

	bookingTransaction := db.BookingTransaction{
		ID:              lookupReservationRow.ID,
		ReservationCode: lookupReservationRow.ReservationCode,
		Status:          lookupReservationRow.Status,
		ErrorCode:       lookupReservationRow.ErrorCode,
		ErrorMessage:    lookupReservationRow.ErrorMessage,
		ContactEmail:    lookupReservationRow.ContactEmail,
		ContactPhone:    lookupReservationRow.ContactPhone,
		CurrencyCode:    lookupReservationRow.CurrencyCode,
		TotalPrice:      lookupReservationRow.TotalPrice,
		CreateDate:      lookupReservationRow.CreateDate,
		ModifyDate:      lookupReservationRow.ModifyDate,
	}

	if !bookingTransaction.Status.Valid || bookingTransaction.Status.String != BOOKING_TRANSACTION_STATUS_BOOKED {
		return nil, errors.New("this booking transaction status is not valid")
	}

	bookingFlights, err := h.store.Queries.GetBookingFlights(ctx, sql.NullInt64{Int64: bookingTransaction.ID, Valid: true})
	if err != nil {
		return nil, fmt.Errorf("could not get booking flights by bookingTransactionId=%v|%v", bookingTransaction.ID, err)
	}

	grpcBfResFlightOptions := []*pb.GrpcBfResFlightOption{}
	for _, bookingFlight := range bookingFlights {
		flight, err := h.store.Queries.GetFlightById(ctx, bookingFlight.FlightID.Int64)
		if err != nil {
			return nil, fmt.Errorf("could not get booking flights by bookingTransactionId=%v|%v", bookingTransaction.ID, err)
		}
		grpcBfResFlightOptions = append(grpcBfResFlightOptions, &pb.GrpcBfResFlightOption{
			Id:                     flight.ID,
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

	bookingPassengers, err := h.store.Queries.GetBookingPassengers(ctx, sql.NullInt64{Int64: bookingTransaction.ID, Valid: true})
	if err != nil {
		return nil, fmt.Errorf("could not get booking passengers by bookingTransactionId=%v|%v", bookingTransaction.ID, err)
	}

	grpcBfResPassengers := []*pb.GrpcBfResPassenger{}
	for _, bookingPassenger := range bookingPassengers {
		grpcBfResPassengers = append(grpcBfResPassengers, &pb.GrpcBfResPassenger{
			Title:           bookingPassenger.Title.String,
			LastName:        bookingPassenger.LastName.String,
			FirstName:       bookingPassenger.FirstName.String,
			DateOfBirth:     timestamppb.New(bookingPassenger.DateOfBirth.Time),
			TravelDocType:   bookingPassenger.TravelDocType.String,
			TravelDocNumber: bookingPassenger.TravelDocNumber.String,
			FfpNumber:       bookingPassenger.FfpNumber.String,
		})
	}

	out := &pb.GrpcBookFlightResponse{
		ReservationCode: in.ReservationCode,
		CurrencyCode:    bookingTransaction.CurrencyCode.String,
		TotalPrice:      bookingTransaction.TotalPrice.Float64,
		FlightOptions:   grpcBfResFlightOptions,
		Passengers:      grpcBfResPassengers,
	}
	return out, nil
}
