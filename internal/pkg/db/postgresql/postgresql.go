package database

import (
	"database/sql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	"log"
)

var DB *sql.DB

func InitDB() {
	db, err := sql.Open("postgres", "postgres://postgres:5432@localhost:5432/template1?sslmode=disable")
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	DB = db
}

func CloseDB() error {
	return DB.Close()
}

func Migrate() {
	if err := DB.Ping(); err != nil {
		log.Fatal(err)
	}
	driver, _ := postgres.WithInstance(DB, &postgres.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://internal/pkg/db/migrations",
		"postgres",
		driver,
	)
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

}
