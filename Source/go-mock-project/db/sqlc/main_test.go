package db

import (
	util "anhlv11/go-mock-project/util"
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	config, err := util.LoadGrpcConfig("../../.")
	if err != nil {
		log.Fatal("Can not load config:", err)
	}
	// conn, err := sql.Open(config.Database.DriverName, config.Database.DataSourceName)
	// conn, err := sql.Open("postgres", "postgresql://127.0.0.1:5432/db_go_mock_project?sslmode=disable&user=postgres&password=sql#123456;")
	conn, err := sql.Open("postgres", util.GetDatabaseSourceName(config.Database))
	if err != nil {
		log.Fatal("Can not connect to db:", err)
	}
	log.Println("Test is beginning...")
	testQueries = New(conn)
	os.Exit(m.Run())
}
