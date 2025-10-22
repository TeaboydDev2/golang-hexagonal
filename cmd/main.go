package main

import (
	"fmt"
	"golang-hexagonal/internal/core/services"
	"golang-hexagonal/internal/infastructure/json/repository"
	"golang-hexagonal/internal/presentations/handler"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now().Format(time.RFC3339))
}

func main() {
	userRepository := repository.NewJsonUserRepository()

	userService := services.NewUserServices(userRepository)

	userHandler := handler.NewUserHandler(userService)

	mux := http.NewServeMux()

	mux.HandleFunc("/", greet)

	mux.HandleFunc("/users", userHandler.GetUser)

	fmt.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Printf("Server failed: %v\n", err)
	}
}
