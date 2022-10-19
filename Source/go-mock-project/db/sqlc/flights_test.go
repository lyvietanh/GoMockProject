package db

import (
	"context"
	"database/sql"
	"strings"
	"testing"
	"time"

	util "anhlv11/go-mock-project/util"

	"github.com/stretchr/testify/require"
)

func TestCreateFlight(t *testing.T) {
	arg := CreateFlightParams{
		OriginAirportCode:      strings.ToUpper(util.GenerateRandomString(3)),
		DestinationAirportCode: strings.ToUpper(util.GenerateRandomString(3)),
		DepartureDateTime:      time.Now(),
		BookingClass:           strings.ToUpper(util.GenerateRandomString(1)),
		FlightNumber:           strings.ToUpper("VN" + util.GenerateRandomString(3)),
		FlightDuration:         120,
		Enabled:                true,
		SeatAvailable:          10,
		CurrencyCode:           "VND",
		Price:                  1000000,
		// CreateBy:               sql.NullString{String: "test", Valid: true},
		CreateDate: sql.NullTime{Time: time.Now(), Valid: true},
	}

	flight, err := testQueries.CreateFlight(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, flight)

	require.Equal(t, arg.OriginAirportCode, flight.OriginAirportCode)
	require.Equal(t, arg.DestinationAirportCode, flight.DestinationAirportCode)
}
