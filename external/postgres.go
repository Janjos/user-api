package external

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type DbConnection struct {
	Db *pgx.Conn
}

const (
	createTables = `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
		email varchar(255) NOT NULL,
		password varchar(50) NOT NULL
    );
    `
)

func NewPostgresDb(url string) (*pgx.Conn, error) {
	config, err := pgx.ParseConfig(url)
	if err != nil {
		fmt.Println("Error parsing config", err)
		return nil, err
	}

	db, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		fmt.Println("Error creating database connection", err)
		return nil, err
	}

	if _, err := db.Exec(context.Background(), createTables); err != nil {
		fmt.Println("Error creating table users", err)
		return nil, err
	}

	return db, nil
}

func NewDbs() (*DbConnection, error) {
	pgDb, err := NewPostgresDb("postgres://postgres:senha123@db:5432/userDatabase")
	if err != nil {
		return nil, err
	}

	return &DbConnection{
		Db: pgDb,
	}, nil
}
