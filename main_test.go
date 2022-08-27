package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

// setupTestDatabase sets up a temporary database. Returns a handle to the DB
// and a cleanup function that you can defer in your tests.
func setupTestDatabase(t *testing.T) (*sql.DB, func()) {
	dir, err := os.MkdirTemp("", "")
	require.NoError(t, err)

	db, err := sql.Open("sqlite3", filepath.Join(dir, "test.db"))
	if err != nil {
		os.RemoveAll(dir)
		t.Fatal(fmt.Errorf("opening the db: %w", err))
	}

	fn := func() {
		db.Close()
		os.RemoveAll(dir)
	}

	if err := applyMigrations(db); err != nil {
		fn()
		t.Fatal(fmt.Errorf("applying migrations: %w", err))
	}

	return db, fn
}

func TestSomething(t *testing.T) {
	db, cleanup := setupTestDatabase(t)
	defer cleanup()

	// This is a contrived example, but here you could do something to test against
	// your database.

	row, err := db.Query("SELECT name FROM people ORDER BY name")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var name string
		row.Scan(&name)
		log.Println(name)
	}
}
