package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDb *sql.DB

var (
	dbDriver = "postgres"
	dbSource = "postgresql://root:1234@localhost:5432/simple_bank?sslmode=disable"
)

func TestMain(m *testing.M) {
	var err error
	// Initialize the test database connection
	testDb, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}

	testQueries = New(testDb)
	os.Exit(m.Run())
}
