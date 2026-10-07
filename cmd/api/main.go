package main

import (
	"log"
	"net/http"

	"github.com/lucassresende/wallet-service/internal/database"
	"github.com/lucassresende/wallet-service/internal/handler"
	"github.com/lucassresende/wallet-service/internal/repository"
	"github.com/lucassresende/wallet-service/internal/routes"
	"github.com/lucassresende/wallet-service/internal/service"
)

func main() {

	db, err := database.NewPostgres()

	if err != nil {
		log.Fatal(err)
	}

	repo :=
		repository.NewPostgresWalletRepository(db)

	service :=
		service.NewWalletService(repo)

	handler :=
		handler.NewWalletHandler(service)

	routes.RegisterRoutes(handler)

	log.Println("server running on :8080")

	http.ListenAndServe(":8080", nil)
}