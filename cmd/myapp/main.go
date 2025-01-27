package main

import (
	"log"
	"net/http"
	"weatherAPI/internal/config"
	"weatherAPI/internal/database"
	"weatherAPI/internal/handlers"
	"weatherAPI/internal/services"

	"github.com/gorilla/mux"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config: ", err)
	}

	// Connect with the DB
	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	// Create the tables
	if err := database.CreateUsersTable(db); err != nil {
		log.Fatalf("Error creating users table: %v", err)
	}

	// Create the routes
	router := mux.NewRouter()
	handlers.UserRouterHandlers(router, db)

	// Start the notification service
	go services.ScheduleNotifications(db)

	// Start the server
	log.Printf("Server running on %s", cfg.ServerAddress)
	if err := http.ListenAndServe(cfg.ServerAddress, router); err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}
