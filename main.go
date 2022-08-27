package main

import (
	"database/sql"
	"embed"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

func main() {

	dbFile := "/tmp/foo.db"

	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatal(fmt.Errorf("opening the db: %w", err))
	}
	defer db.Close()

	if err := applyMigrations(db); err != nil {
		log.Fatal(fmt.Errorf("applying migrations: %w", err))
	}

	// Start up your application, do something with the db

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

//go:embed migrations/*
var migrations embed.FS

func applyMigrations(db *sql.DB) error {
	src, err := iofs.New(migrations, "migrations")
	if err != nil {
		return fmt.Errorf("getting embed instance: %w", err)
	}

	dest, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return fmt.Errorf("getting sqlite3 instance: %w", err)
	}

	m, err := migrate.NewWithInstance("iofs", src, "sqlite3", dest)
	if err != nil {
		return fmt.Errorf("getting migration instance: %w", err)
	}

	err = m.Up()
	if err != nil && err == migrate.ErrNoChange {
		return nil
	}

	return err
}
