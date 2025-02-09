package main

import (
	"log"
	"net/http"

	"github.com/janjos/user-api/controllers"
	"github.com/janjos/user-api/external"
	"github.com/janjos/user-api/interfaces/repositories"
	"github.com/janjos/user-api/useCases"

	_ "github.com/lib/pq"
)

func main() {
	dbConnection, err := external.NewDbs()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	userRepo := repositories.NewUserRepositoryImpl(dbConnection)

	userUsecase := useCases.NewUserUsecase(userRepo)

	userController := controllers.NewUserController(userUsecase)

	http.HandleFunc("/users", userController.CreateUser)
	http.HandleFunc("/users/get", userController.GetUserByID)
	http.HandleFunc("/users/login", userController.LogIn)

	log.Println("Running on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
