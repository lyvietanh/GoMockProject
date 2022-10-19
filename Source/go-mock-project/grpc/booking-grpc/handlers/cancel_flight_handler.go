package booking_grpc_handlers

import (
	db "anhlv11/go-mock-project/db/sqlc"
	"anhlv11/go-mock-project/pb"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *BookingGrpcHandler) CancelBooking(ctx context.Context, in *pb.GrpcCancelBookingRequest) (*pb.GrpcCancelBookingResponse, error) {
	// begin db transaction
	tx, err := h.store.Db.Begin()
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Db transaction can not begin|%v", err))
	}
	defer tx.Rollback()

	txQueries := h.store.Queries.WithTx(tx)

	lookupReservationParams := db.LookupReservationParams{
		Limit:           1,
		Offset:          0,
		ReservationCode: sql.NullString{String: in.ReservationCode, Valid: true},
		LastName:        sql.NullString{String: in.LastName, Valid: true},
		TravelDocNumber: sql.NullString{String: in.TravelDocNumber, Valid: true},
	}
	lookupReservationRows, err := txQueries.LookupReservation(ctx, lookupReservationParams)
	if err != nil {
		return nil, fmt.Errorf("could not get booking transactions by reservationCode=%v lastName=%v travelDocNumber=%v|%v", in.ReservationCode, in.LastName, in.TravelDocNumber, err)
	}
	if len(lookupReservationRows) == 0 {
		return nil, fmt.Errorf("could not found any booking transactions by reservationCode=%v lastName=%v travelDocNumber=%v", in.ReservationCode, in.LastName, in.TravelDocNumber)
	}

	lookupReservationRow := lookupReservationRows[0]

	numberOfPassengers := lookupReservationRow.NumberOfPassengers
	bookingTransaction := db.BookingTransaction{
		ID:              lookupReservationRow.ID,
		ReservationCode: lookupReservationRow.ReservationCode,
		Status:          lookupReservationRow.Status,
		ErrorCode:       lookupReservationRow.ErrorCode,
		ErrorMessage:    lookupReservationRow.ErrorMessage,
		CurrencyCode:    lookupReservationRow.CurrencyCode,
		TotalPrice:      lookupReservationRow.TotalPrice,
		CreateDate:      lookupReservationRow.CreateDate,
		ModifyDate:      lookupReservationRow.ModifyDate,
	}

	if !bookingTransaction.Status.Valid || bookingTransaction.Status.String != BOOKING_TRANSACTION_STATUS_BOOKED {
		return nil, errors.New("this booking transaction status is not valid")
	}

	bookingFlights, err := txQueries.GetBookingFlights(ctx, sql.NullInt64{Int64: bookingTransaction.ID, Valid: true})
	if err != nil {
		return nil, fmt.Errorf("could not get booking flights by bookingTransactionId=%v|%v", bookingTransaction.ID, err)
	}

	h.mutex.Lock()
	for _, bookingFlight := range bookingFlights {
		flight, err := txQueries.GetFlightById(ctx, bookingFlight.FlightID.Int64)
		if err != nil {
			continue
		}
		currentAvailableSeats := flight.SeatAvailable
		newAvailableSeats := currentAvailableSeats + int32(numberOfPassengers)

		updateFlightSeatAvailableParams := db.UpdateFlightSeatAvailableParams{
			ID:            flight.ID,
			SeatAvailable: newAvailableSeats,
			ModifyDate:    sql.NullTime{Time: time.Now(), Valid: true},
		}
		_, err = txQueries.UpdateFlightSeatAvailable(ctx, updateFlightSeatAvailableParams)
		if err != nil {
			return nil, fmt.Errorf("could not update flight seat available with flightId=%v currentAvailableSeats=%v newAvailableSeats=%v|%v", flight.ID, currentAvailableSeats, newAvailableSeats, err)
		}
	}
	defer h.mutex.Unlock()

	cancelDate := time.Now()
	updateBookingTransactionStatusParams := db.UpdateBookingTransactionStatusParams{
		ID:           bookingTransaction.ID,
		Status:       sql.NullString{String: BOOKING_TRANSACTION_STATUS_CANCELED, Valid: true},
		ErrorCode:    sql.NullInt32{Valid: false},
		ErrorMessage: sql.NullString{Valid: false},
		ModifyDate:   sql.NullTime{Time: cancelDate, Valid: true},
	}
	_, err = txQueries.UpdateBookingTransactionStatus(ctx, updateBookingTransactionStatusParams)
	if err != nil {
		return nil, fmt.Errorf("could not update booking transaction status with bookingTransactionId=%v|%v", bookingTransaction.ID, err)
	}

	out := &pb.GrpcCancelBookingResponse{
		ReservationCode: in.ReservationCode,
		CancelDate:      timestamppb.New(cancelDate),
	}

	// commit db transaction
	if errCommit := tx.Commit(); errCommit != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Can not commit|%v", errCommit))
	}

	return out, nil
}
