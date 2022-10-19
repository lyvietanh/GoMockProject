-- name: CreateBookingFlight :one
INSERT INTO booking_flights(
    booking_transaction_id, 
    flight_id, 
    price, 
    currency_code
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;

-- name: CreateBookingPassenger :one
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
) RETURNING *;

-- name: CreateBookingTransaction :one
INSERT INTO booking_transactions(
    reservation_code, 
    status,
    contact_email,
    contact_phone, 
    create_date
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: UpdateBookingTransactionStatus :one
UPDATE booking_transactions
SET 
    status=$2, 
    error_code=$3, 
    error_message=$4, 
    modify_date=$5
WHERE 
    id=$1
RETURNING *;

-- name: UpdateBookingTransactionPriceInformation :one
UPDATE booking_transactions
SET 
    currency_code=$2, 
    total_price=$3, 
    modify_date=$4
WHERE 
    id=$1
RETURNING *;

-- name: GetBookingTransactionsByReservationCode :many
SELECT * FROM booking_transactions
WHERE 
    reservation_code=$3
ORDER BY
    create_date DESC
LIMIT $1
OFFSET $2
;

-- name: LookupReservation :many
SELECT bt.*, (SELECT COUNT(1) FROM booking_passengers WHERE booking_transaction_id=bp.booking_transaction_id) AS number_of_passengers FROM booking_transactions bt
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
;

-- name: GetBookingFlights :many
SELECT * FROM booking_flights
WHERE
    booking_transaction_id=$1
ORDER BY
    id ASC
;

-- name: GetBookingPassengers :many
SELECT * FROM booking_passengers
WHERE
    booking_transaction_id=$1
ORDER BY
    id ASC
;

-- name: GetBookingTransactionsByContactEmail :many
SELECT * FROM booking_transactions
WHERE
    contact_email=$1
ORDER BY
    create_date DESC
;

-- name: GetBookingFlightsByBookingTransactionIds :many
SELECT * FROM booking_flights
WHERE
    booking_transaction_id=ANY(sqlc.narg('BookingTransactionIds')::bigint[]) 
    AND array_length(sqlc.narg('BookingTransactionIds')::bigint[], 1) IS NOT NULL
ORDER BY
    booking_transaction_id ASC
;

-- name: GetBookingPassengersByBookingTransactionIds :many
SELECT * FROM booking_passengers
WHERE
    booking_transaction_id=ANY(sqlc.narg('BookingTransactionIds')::bigint[]) 
    AND array_length(sqlc.narg('BookingTransactionIds')::bigint[], 1) IS NOT NULL
ORDER BY
    booking_transaction_id ASC
;