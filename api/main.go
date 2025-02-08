package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Janjos/user-api/controllers"
	"github.com/Janjos/user-api/interfaces/repositories"
	"github.com/Janjos/user-api/usecases"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres password=senha123 dbname=meubanco sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userRepo := repositories.NewUserRepositoryImpl(db)

	userUsecase := usecases.NewUserUsecase(userRepo)

	userController := controllers.NewUserController(userUsecase)

	http.HandleFunc("/users", userController.CreateUser)
	http.HandleFunc("/users/get", userController.GetUserByID)

	log.Println("Running on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
