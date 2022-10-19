package booking_grpc_handlers

import (
	db "anhlv11/go-mock-project/db/sqlc"
	"anhlv11/go-mock-project/pb"
	"anhlv11/go-mock-project/util"
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *BookingGrpcHandler) BookFlight(ctx context.Context, in *pb.GrpcBookFlightRequest) (*pb.GrpcBookFlightResponse, error) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	// begin db transaction
	tx, err := h.store.Db.Begin()
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Db transaction can not begin|%v", err))
	}
	defer tx.Rollback()

	txQueries := h.store.Queries.WithTx(tx)

	// create booking transaction
	reservationCode, err := generateReservationCode(ctx, txQueries)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Can not generate reservation code|%v", err))
	}
	createBookingTransactionParams := db.CreateBookingTransactionParams{
		ReservationCode: sql.NullString{String: *reservationCode, Valid: true},
		Status:          sql.NullString{String: BOOKING_TRANSACTION_STATUS_BOOKED, Valid: true},
		ContactEmail:    in.ContactEmail,
		ContactPhone:    in.ContactPhone,
		CreateDate:      sql.NullTime{Time: time.Now(), Valid: true},
	}
	bookingTransaction, err := txQueries.CreateBookingTransaction(ctx, createBookingTransactionParams)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Rollback because can not create booking transaction|%v", err))
	}

	// create booking flights
	outFlightOptions := []*pb.GrpcBfResFlightOption{}
	totalPrice := 0.0
	for _, flightId := range in.FlightOptionIds {
		existedFlight, err := txQueries.GetFlightById(ctx, flightId)
		if err != nil {
			return nil, status.Error(codes.Internal, fmt.Sprintf("Can not found flight with id=%v|%v", flightId, err))
		}
		if !strings.EqualFold(existedFlight.CurrencyCode, in.CurrencyCode) {
			return nil, status.Error(codes.Internal, fmt.Sprintf("Flight with id=%v not matched currency code|%v", flightId, err))
		}

		// update seat available
		seatAvailable := existedFlight.SeatAvailable - int32(len(in.Passengers))
		log.Printf("seatAvailable=%v\n", seatAvailable)
		if seatAvailable < 0 {
			log.Println("Out of available seats!!!")
			return nil, status.Error(codes.Internal, "Out of available seats")
		}

		time.Sleep(1 * time.Second)

		updateFlightSeatAvailableParams := db.UpdateFlightSeatAvailableParams{
			ID:            existedFlight.ID,
			SeatAvailable: seatAvailable,
			ModifyDate:    sql.NullTime{Time: time.Now(), Valid: true},
		}
		_, err = txQueries.UpdateFlightSeatAvailable(ctx, updateFlightSeatAvailableParams)
		if err != nil {
			return nil, status.Error(codes.Internal, fmt.Sprintf("Rollback because can not decrease seat available of flight with id=%v|%v", existedFlight.ID, err))
		}

		time.Sleep(1 * time.Second)

		arrivalDateTime := existedFlight.DepartureDateTime.Add(time.Duration(existedFlight.FlightDuration) * time.Minute)
		createBookingFlightParams := db.CreateBookingFlightParams{
			BookingTransactionID: sql.NullInt64{Int64: bookingTransaction.ID, Valid: true},
			FlightID:             sql.NullInt64{Int64: flightId, Valid: true},
			CurrencyCode:         existedFlight.CurrencyCode,
			Price:                existedFlight.Price,
		}
		_, err = txQueries.CreateBookingFlight(ctx, createBookingFlightParams)
		if err != nil {
			return nil, status.Error(codes.Internal, fmt.Sprintf("Rollback because can not create booking flight with id=%v|%v", existedFlight.ID, err))
		}
		totalPrice += existedFlight.Price
		outFlightOptions = append(outFlightOptions, &pb.GrpcBfResFlightOption{
			Id:                     existedFlight.ID,
			OriginAirportCode:      existedFlight.OriginAirportCode,
			DestinationAirportCode: existedFlight.DestinationAirportCode,
			DepartureDateTime:      timestamppb.New(existedFlight.DepartureDateTime),
			ArrivalDateTime:        timestamppb.New(arrivalDateTime),
			BookingClass:           existedFlight.BookingClass,
			FlightNumber:           existedFlight.FlightNumber,
			FlightDuration:         existedFlight.FlightDuration,
			CurrencyCode:           existedFlight.CurrencyCode,
			Price:                  existedFlight.Price,
		})
	}

	totalPrice = totalPrice * float64(len(in.Passengers))

	// update booking transaction price information
	updateBookingTransactionPriceInformationParams := db.UpdateBookingTransactionPriceInformationParams{
		ID:           bookingTransaction.ID,
		CurrencyCode: sql.NullString{String: in.CurrencyCode, Valid: true},
		TotalPrice:   sql.NullFloat64{Float64: totalPrice, Valid: true},
		ModifyDate:   sql.NullTime{Time: time.Now(), Valid: true},
	}
	bookingTransaction, err = txQueries.UpdateBookingTransactionPriceInformation(ctx, updateBookingTransactionPriceInformationParams)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Rollback because can not update booking transaction price information|%v", err))
	}

	// create booking passengers
	outPassengers := []*pb.GrpcBfResPassenger{}
	for _, inPassenger := range in.Passengers {
		createBookingPassengerParams := db.CreateBookingPassengerParams{
			BookingTransactionID: sql.NullInt64{Int64: bookingTransaction.ID, Valid: true},
			Title:                sql.NullString{String: inPassenger.Title, Valid: true},
			LastName:             sql.NullString{String: inPassenger.LastName, Valid: true},
			FirstName:            sql.NullString{String: inPassenger.FirstName, Valid: true},
			DateOfBirth:          sql.NullTime{Time: inPassenger.DateOfBirth.AsTime(), Valid: true},
			TravelDocType:        sql.NullString{String: inPassenger.TravelDocType, Valid: true},
			TravelDocNumber:      sql.NullString{String: inPassenger.TravelDocNumber, Valid: true},
			FfpNumber:            sql.NullString{String: inPassenger.FfpNumber, Valid: true},
		}
		bookingPassenger, err := txQueries.CreateBookingPassenger(ctx, createBookingPassengerParams)
		if err != nil {
			return nil, status.Error(codes.Internal, fmt.Sprintf("Rollback because can not create booking passenger with title=%v lastName=%v firstName=%v|%v",
				inPassenger.Title, inPassenger.LastName, inPassenger.FirstName, err))
		}
		outPassengers = append(outPassengers, &pb.GrpcBfResPassenger{
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
		ReservationCode: *reservationCode,
		CurrencyCode:    in.CurrencyCode,
		TotalPrice:      totalPrice,
		Passengers:      outPassengers,
		FlightOptions:   outFlightOptions,
	}

	// commit db transaction
	if errCommit := tx.Commit(); errCommit != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Can not commit|%v", errCommit))
	}

	return out, nil
}

func generateReservationCode(ctx context.Context, queries *db.Queries) (*string, error) {
	reservationCode := ""
	for {
		reservationCode = strings.ToUpper(util.GenerateRandomString(6))
		getBookingTransactionsByReservationCodeParams := db.GetBookingTransactionsByReservationCodeParams{
			Limit:           1,
			Offset:          0,
			ReservationCode: sql.NullString{String: reservationCode, Valid: true},
		}
		bookingTransactions, err := queries.GetBookingTransactionsByReservationCode(ctx, getBookingTransactionsByReservationCodeParams)
		if err != nil {
			return nil, fmt.Errorf("could not get booking transaction by reservation code: %v|%v", reservationCode, err)
		}
		if len(bookingTransactions) == 0 {
			break
		}
	}
	return &reservationCode, nil
}
