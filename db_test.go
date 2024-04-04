package main

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "modernc.org/sqlite"
)

var db *sql.DB // Глобальная переменная для доступа к БД

func TestMain(m *testing.M) {
	var err error
	db, err = sql.Open("sqlite", ":memory:")
	if err != nil {
		log.Fatalf("could not open db: %v", err)
	}
	defer db.Close()

	if err := createTestTable(db); err != nil {
		log.Fatalf("could not create test table: %v", err)
	}

	os.Exit(m.Run())
}

func createTestTable(db *sql.DB) error {
	const tableCreationQuery = `CREATE TABLE IF NOT EXISTS parcel (
		number INTEGER PRIMARY KEY AUTOINCREMENT,
		client INTEGER NOT NULL,
		status TEXT NOT NULL,
		address TEXT NOT NULL,
		created_at TEXT NOT NULL
	)`
	_, err := db.Exec(tableCreationQuery)
	return err
}
