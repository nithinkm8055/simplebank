package db

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/lib/pq"
)

const (
	driverName = "postgres"
	sourceName = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var (
	testQueries *Queries
	testDB      *sql.DB
)

func TestMain(m *testing.M) {
	var err error

	testDB, err = sql.Open(driverName, sourceName)
	if err != nil {
		log.Fatal("cannot open db connection", err)
	}

	testQueries = New(testDB)

	m.Run()
}
