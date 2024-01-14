package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
}

// Constructor
func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	// Verify the connection
	if err := db.Ping(); err != nil {
		return nil, err
	}
	log.Println("Successful connection to the database")

	return &PostgresRepository{db}, nil
}

func (repo *PostgresRepository) Close() error {
	return repo.db.Close()
}
