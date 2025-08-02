package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/kento/tranzure/internal/config"
	"github.com/kento/tranzure/internal/database"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	fmt.Printf("Starting Payment Service on %s:%d\n", cfg.Server.Host, cfg.Server.Port)
	fmt.Printf("Environment: %s\n", cfg.App.Environment)

	// Connect to MongoDB
	mongoDB, err := database.NewMongoDB(&cfg.Database.MongoDB)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := mongoDB.Close(ctx); err != nil {
			log.Printf("Error closing MongoDB connection: %v", err)
		}
	}()

	fmt.Println("Successfully connected to MongoDB!")
	fmt.Printf("Server running on http://%s:%d\n", cfg.Server.Host, cfg.Server.Port)

	// Keep the application running
	select {}
}
