// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: booking.sql

package db

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

const createBookingFlight = `-- name: CreateBookingFlight :one
INSERT INTO booking_flights(
    booking_transaction_id, 
    flight_id, 
    price, 
    currency_code
) VALUES (
    $1, $2, $3, $4
)
RETURNING id, booking_transaction_id, flight_id, currency_code, price
`

type CreateBookingFlightParams struct {
	BookingTransactionID sql.NullInt64 `json:"booking_transaction_id"`
	FlightID             sql.NullInt64 `json:"flight_id"`
	Price                float64       `json:"price"`
	CurrencyCode         string        `json:"currency_code"`
}

func (q *Queries) CreateBookingFlight(ctx context.Context, arg CreateBookingFlightParams) (BookingFlight, error) {
	row := q.db.QueryRowContext(ctx, createBookingFlight,
		arg.BookingTransactionID,
		arg.FlightID,
		arg.Price,
		arg.CurrencyCode,
	)
	var i BookingFlight
	err := row.Scan(
		&i.ID,
		&i.BookingTransactionID,
		&i.FlightID,
		&i.CurrencyCode,
		&i.Price,
	)
	return i, err
}

const createBookingPassenger = `-- name: CreateBookingPassenger :one
INSERT INTO booking_passengers(
    booking_transaction_id, 
    title, 
    last_name, 
    first_name,
    date_of_birth, 
    travel_doc_type, 
    travel_doc_number,
    ffp_number
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING id, booking_transaction_id, title, last_name, first_name, date_of_birth, travel_doc_type, travel_doc_number, ffp_number
`

type CreateBookingPassengerParams struct {
	BookingTransactionID sql.NullInt64  `json:"booking_transaction_id"`
	Title                sql.NullString `json:"title"`
	LastName             sql.NullString `json:"last_name"`
	FirstName            sql.NullString `json:"first_name"`
	DateOfBirth          sql.NullTime   `json:"date_of_birth"`
	TravelDocType        sql.NullString `json:"travel_doc_type"`
	TravelDocNumber      sql.NullString `json:"travel_doc_number"`
	FfpNumber            sql.NullString `json:"ffp_number"`
}

func (q *Queries) CreateBookingPassenger(ctx context.Context, arg CreateBookingPassengerParams) (BookingPassenger, error) {
	row := q.db.QueryRowContext(ctx, createBookingPassenger,
		arg.BookingTransactionID,
		arg.Title,
		arg.LastName,
		arg.FirstName,
		arg.DateOfBirth,
		arg.TravelDocType,
		arg.TravelDocNumber,
		arg.FfpNumber,
	)
	var i BookingPassenger
	err := row.Scan(
		&i.ID,
		&i.BookingTransactionID,
		&i.Title,
		&i.LastName,
		&i.FirstName,
		&i.DateOfBirth,
		&i.TravelDocType,
		&i.TravelDocNumber,
		&i.FfpNumber,
	)
	return i, err
}

const createBookingTransaction = `-- name: CreateBookingTransaction :one
INSERT INTO booking_transactions(
    reservation_code, 
    status,
    contact_email,
    contact_phone, 
    create_date
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING id, reservation_code, status, error_code, error_message, contact_email, contact_phone, currency_code, total_price, create_date, modify_date
`

type CreateBookingTransactionParams struct {
	ReservationCode sql.NullString `json:"reservation_code"`
	Status          sql.NullString `json:"status"`
	ContactEmail    string         `json:"contact_email"`
	ContactPhone    string         `json:"contact_phone"`
	CreateDate      sql.NullTime   `json:"create_date"`
}

func (q *Queries) CreateBookingTransaction(ctx context.Context, arg CreateBookingTransactionParams) (BookingTransaction, error) {
	row := q.db.QueryRowContext(ctx, createBookingTransaction,
		arg.ReservationCode,
		arg.Status,
		arg.ContactEmail,
		arg.ContactPhone,
		arg.CreateDate,
	)
	var i BookingTransaction
	err := row.Scan(
		&i.ID,
		&i.ReservationCode,
		&i.Status,
		&i.ErrorCode,
		&i.ErrorMessage,
		&i.ContactEmail,
		&i.ContactPhone,
		&i.CurrencyCode,
		&i.TotalPrice,
		&i.CreateDate,
		&i.ModifyDate,
	)
	return i, err
}

const getBookingFlights = `-- name: GetBookingFlights :many
SELECT id, booking_transaction_id, flight_id, currency_code, price FROM booking_flights
WHERE
    booking_transaction_id=$1
ORDER BY
    id ASC
`

func (q *Queries) GetBookingFlights(ctx context.Context, bookingTransactionID sql.NullInt64) ([]BookingFlight, error) {
	rows, err := q.db.QueryContext(ctx, getBookingFlights, bookingTransactionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []BookingFlight{}
	for rows.Next() {
		var i BookingFlight
		if err := rows.Scan(
			&i.ID,
			&i.BookingTransactionID,
			&i.FlightID,
			&i.CurrencyCode,
			&i.Price,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getBookingFlightsByBookingTransactionIds = `-- name: GetBookingFlightsByBookingTransactionIds :many
SELECT id, booking_transaction_id, flight_id, currency_code, price FROM booking_flights
WHERE
    booking_transaction_id=ANY($1::bigint[]) 
    AND array_length($1::bigint[], 1) IS NOT NULL
ORDER BY
    booking_transaction_id ASC
`

func (q *Queries) GetBookingFlightsByBookingTransactionIds(ctx context.Context, bookingtransactionids []int64) ([]BookingFlight, error) {
	rows, err := q.db.QueryContext(ctx, getBookingFlightsByBookingTransactionIds, pq.Array(bookingtransactionids))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []BookingFlight{}
	for rows.Next() {
		var i BookingFlight
		if err := rows.Scan(
			&i.ID,
			&i.BookingTransactionID,
			&i.FlightID,
			&i.CurrencyCode,
			&i.Price,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getBookingPassengers = `-- name: GetBookingPassengers :many
SELECT id, booking_transaction_id, title, last_name, first_name, date_of_birth, travel_doc_type, travel_doc_number, ffp_number FROM booking_passengers
WHERE
    booking_transaction_id=$1
ORDER BY
    id ASC
`

func (q *Queries) GetBookingPassengers(ctx context.Context, bookingTransactionID sql.NullInt64) ([]BookingPassenger, error) {
	rows, err := q.db.QueryContext(ctx, getBookingPassengers, bookingTransactionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []BookingPassenger{}
	for rows.Next() {
		var i BookingPassenger
		if err := rows.Scan(
			&i.ID,
			&i.BookingTransactionID,
			&i.Title,
			&i.LastName,
			&i.FirstName,
			&i.DateOfBirth,
			&i.TravelDocType,
			&i.TravelDocNumber,
			&i.FfpNumber,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getBookingPassengersByBookingTransactionIds = `-- name: GetBookingPassengersByBookingTransactionIds :many
SELECT id, booking_transaction_id, title, last_name, first_name, date_of_birth, travel_doc_type, travel_doc_number, ffp_number FROM booking_passengers
WHERE
    booking_transaction_id=ANY($1::bigint[]) 
    AND array_length($1::bigint[], 1) IS NOT NULL
ORDER BY
    booking_transaction_id ASC
`

func (q *Queries) GetBookingPassengersByBookingTransactionIds(ctx context.Context, bookingtransactionids []int64) ([]BookingPassenger, error) {
	rows, err := q.db.QueryContext(ctx, getBookingPassengersByBookingTransactionIds, pq.Array(bookingtransactionids))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []BookingPassenger{}
	for rows.Next() {
		var i BookingPassenger
		if err := rows.Scan(
			&i.ID,
			&i.BookingTransactionID,
			&i.Title,
			&i.LastName,
			&i.FirstName,
			&i.DateOfBirth,
			&i.TravelDocType,
			&i.TravelDocNumber,
			&i.FfpNumber,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getBookingTransactionsByContactEmail = `-- name: GetBookingTransactionsByContactEmail :many
SELECT id, reservation_code, status, error_code, error_message, contact_email, contact_phone, currency_code, total_price, create_date, modify_date FROM booking_transactions
WHERE
    contact_email=$1
ORDER BY
    create_date DESC
`

func (q *Queries) GetBookingTransactionsByContactEmail(ctx context.Context, contactEmail string) ([]BookingTransaction, error) {
	rows, err := q.db.QueryContext(ctx, getBookingTransactionsByContactEmail, contactEmail)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []BookingTransaction{}
	for rows.Next() {
		var i BookingTransaction
		if err := rows.Scan(
			&i.ID,
			&i.ReservationCode,
			&i.Status,
			&i.ErrorCode,
			&i.ErrorMessage,
			&i.ContactEmail,
			&i.ContactPhone,
			&i.CurrencyCode,
			&i.TotalPrice,
			&i.CreateDate,
			&i.ModifyDate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getBookingTransactionsByReservationCode = `-- name: GetBookingTransactionsByReservationCode :many
SELECT id, reservation_code, status, error_code, error_message, contact_email, contact_phone, currency_code, total_price, create_date, modify_date FROM booking_transactions
WHERE 
    reservation_code=$3
ORDER BY
    create_date DESC
LIMIT $1
OFFSET $2
`

type GetBookingTransactionsByReservationCodeParams struct {
	Limit           int32          `json:"limit"`
	Offset          int32          `json:"offset"`
	ReservationCode sql.NullString `json:"reservation_code"`
}

func (q *Queries) GetBookingTransactionsByReservationCode(ctx context.Context, arg GetBookingTransactionsByReservationCodeParams) ([]BookingTransaction, error) {
	rows, err := q.db.QueryContext(ctx, getBookingTransactionsByReservationCode, arg.Limit, arg.Offset, arg.ReservationCode)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []BookingTransaction{}
	for rows.Next() {
		var i BookingTransaction
		if err := rows.Scan(
			&i.ID,
			&i.ReservationCode,
			&i.Status,
			&i.ErrorCode,
			&i.ErrorMessage,
			&i.ContactEmail,
			&i.ContactPhone,
			&i.CurrencyCode,
			&i.TotalPrice,
			&i.CreateDate,
			&i.ModifyDate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const lookupReservation = `-- name: LookupReservation :many
SELECT bt.id, bt.reservation_code, bt.status, bt.error_code, bt.error_message, bt.contact_email, bt.contact_phone, bt.currency_code, bt.total_price, bt.create_date, bt.modify_date, (SELECT COUNT(1) FROM booking_passengers WHERE booking_transaction_id=bp.booking_transaction_id) AS number_of_passengers FROM booking_transactions bt
LEFT JOIN booking_passengers bp
ON bt.id=bp.booking_transaction_id
WHERE
    bt.reservation_code=$3
    AND bp.last_name=$4
    AND bp.travel_doc_number=$5
ORDER BY
    bt.create_date DESC
LIMIT $1
OFFSET $2
`

type LookupReservationParams struct {
	Limit           int32          `json:"limit"`
	Offset          int32          `json:"offset"`
	ReservationCode sql.NullString `json:"reservation_code"`
	LastName        sql.NullString `json:"last_name"`
	TravelDocNumber sql.NullString `json:"travel_doc_number"`
}

type LookupReservationRow struct {
	ID                 int64           `json:"id"`
	ReservationCode    sql.NullString  `json:"reservation_code"`
	Status             sql.NullString  `json:"status"`
	ErrorCode          sql.NullInt32   `json:"error_code"`
	ErrorMessage       sql.NullString  `json:"error_message"`
	ContactEmail       string          `json:"contact_email"`
	ContactPhone       string          `json:"contact_phone"`
	CurrencyCode       sql.NullString  `json:"currency_code"`
	TotalPrice         sql.NullFloat64 `json:"total_price"`
	CreateDate         sql.NullTime    `json:"create_date"`
	ModifyDate         sql.NullTime    `json:"modify_date"`
	NumberOfPassengers int64           `json:"number_of_passengers"`
}

func (q *Queries) LookupReservation(ctx context.Context, arg LookupReservationParams) ([]LookupReservationRow, error) {
	rows, err := q.db.QueryContext(ctx, lookupReservation,
		arg.Limit,
		arg.Offset,
		arg.ReservationCode,
		arg.LastName,
		arg.TravelDocNumber,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []LookupReservationRow{}
	for rows.Next() {
		var i LookupReservationRow
		if err := rows.Scan(
			&i.ID,
			&i.ReservationCode,
			&i.Status,
			&i.ErrorCode,
			&i.ErrorMessage,
			&i.ContactEmail,
			&i.ContactPhone,
			&i.CurrencyCode,
			&i.TotalPrice,
			&i.CreateDate,
			&i.ModifyDate,
			&i.NumberOfPassengers,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateBookingTransactionPriceInformation = `-- name: UpdateBookingTransactionPriceInformation :one
UPDATE booking_transactions
SET 
    currency_code=$2, 
    total_price=$3, 
    modify_date=$4
WHERE 
    id=$1
RETURNING id, reservation_code, status, error_code, error_message, contact_email, contact_phone, currency_code, total_price, create_date, modify_date
`

type UpdateBookingTransactionPriceInformationParams struct {
	ID           int64           `json:"id"`
	CurrencyCode sql.NullString  `json:"currency_code"`
	TotalPrice   sql.NullFloat64 `json:"total_price"`
	ModifyDate   sql.NullTime    `json:"modify_date"`
}

func (q *Queries) UpdateBookingTransactionPriceInformation(ctx context.Context, arg UpdateBookingTransactionPriceInformationParams) (BookingTransaction, error) {
	row := q.db.QueryRowContext(ctx, updateBookingTransactionPriceInformation,
		arg.ID,
		arg.CurrencyCode,
		arg.TotalPrice,
		arg.ModifyDate,
	)
	var i BookingTransaction
	err := row.Scan(
		&i.ID,
		&i.ReservationCode,
		&i.Status,
		&i.ErrorCode,
		&i.ErrorMessage,
		&i.ContactEmail,
		&i.ContactPhone,
		&i.CurrencyCode,
		&i.TotalPrice,
		&i.CreateDate,
		&i.ModifyDate,
	)
	return i, err
}

const updateBookingTransactionStatus = `-- name: UpdateBookingTransactionStatus :one
UPDATE booking_transactions
SET 
    status=$2, 
    error_code=$3, 
    error_message=$4, 
    modify_date=$5
WHERE 
    id=$1
RETURNING id, reservation_code, status, error_code, error_message, contact_email, contact_phone, currency_code, total_price, create_date, modify_date
`

type UpdateBookingTransactionStatusParams struct {
	ID           int64          `json:"id"`
	Status       sql.NullString `json:"status"`
	ErrorCode    sql.NullInt32  `json:"error_code"`
	ErrorMessage sql.NullString `json:"error_message"`
	ModifyDate   sql.NullTime   `json:"modify_date"`
}

func (q *Queries) UpdateBookingTransactionStatus(ctx context.Context, arg UpdateBookingTransactionStatusParams) (BookingTransaction, error) {
	row := q.db.QueryRowContext(ctx, updateBookingTransactionStatus,
		arg.ID,
		arg.Status,
		arg.ErrorCode,
		arg.ErrorMessage,
		arg.ModifyDate,
	)
	var i BookingTransaction
	err := row.Scan(
		&i.ID,
		&i.ReservationCode,
		&i.Status,
		&i.ErrorCode,
		&i.ErrorMessage,
		&i.ContactEmail,
		&i.ContactPhone,
		&i.CurrencyCode,
		&i.TotalPrice,
		&i.CreateDate,
		&i.ModifyDate,
	)
	return i, err
}
