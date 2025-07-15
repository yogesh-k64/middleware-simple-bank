package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/yogesh-k64/middleware-simple-bank/utils"
)

var testQueries *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {
	var err error
	// Initialize the test database connection

	config, err := utils.LoadConfig("../..")

	if err != nil {
		log.Fatalf("failed to load config %v\n", err)
	}

	testDb, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}

	testQueries = New(testDb)
	os.Exit(m.Run())
}
