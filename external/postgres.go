package external

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func NewPostgresConnection(cfg Config) (*sql.DB, error) {

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("Error on PostgreSQL connection: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("Failed PostgreSQL connection test: %w", err)
	}

	log.Println("Suscessfully connected with postgress")
	return db, nil
}
