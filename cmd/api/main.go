package main

import (
	"log"
	"net/http"

	"morning-call/internal/handler"
	"morning-call/internal/infrastructure/persistence/inmemory"
	"morning-call/internal/usecase"
)

func main() {
	userRepo := inmemory.NewInMemoryUserRepository()
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	http.HandleFunc("/users", userHandler.Register)

	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("could not listen on port 8080 %v", err)
	}
}
