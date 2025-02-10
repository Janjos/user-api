package external

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type DbConnection struct {
	Db *pgx.Conn
}

const (
	createTables = `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
		email varchar(255) NOT NULL,
		password varchar(72) NOT NULL
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
	pgDb, err := NewPostgresDb("postgres://postgres:123@user-db:5432/postgres")
	// pgDb, err := NewPostgresDb("postgres://postgres:senha123@db:5432/userDatabase")
	if err != nil {
		return nil, err
	}

	return &DbConnection{
		Db: pgDb,
	}, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

var secretKey = []byte("secret-key")

func CreateToken(email string, id int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": email,
			"id":    id,
			"exp":   time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (float64, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return -1, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		fmt.Println(claims["email"], claims["id"])
		return claims["id"].(float64), nil
	}

	return -1, nil
}
