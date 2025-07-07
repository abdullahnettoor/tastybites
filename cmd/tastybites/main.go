package main

import (
	"log"

	"github.com/abdullahnettoor/tastybites/internal/api"
	"github.com/abdullahnettoor/tastybites/internal/config"
	"github.com/abdullahnettoor/tastybites/internal/repo"
	"github.com/abdullahnettoor/tastybites/internal/usecases"
)

func main() {

	// Load the configuration
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	// Set up the database connection
	repository, err := repo.NewRepository(&config.DBConfig)
	if err != nil {
		log.Fatalf("Failed to initialize repository: %v", err)
	}

	// Set up the use cases
	userUsecase := usecases.NewUserUsecase(repository)
	orderUsecase := usecases.NewOrderUsecase(repository)
	menuUsecase := usecases.NewMenuUsecase(repository)
	tableUsecase := usecases.NewTableUsecase(repository)

	// Initialize the application
	app, err := api.NewApp(config, repository)
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	// Initialize the routes
	app.InitializeRoutes(userUsecase, orderUsecase, menuUsecase, tableUsecase)

	if err := app.Start(); err != nil {
		log.Fatalf("Failed to start the application: %v", err)
	}
	log.Printf("Server started on %s:%d", config.ServerConfig.Host, config.ServerConfig.Port)
}
