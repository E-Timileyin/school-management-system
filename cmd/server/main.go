package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"

	"school-management-backend/internal/config"
	"school-management-backend/internal/routes"
)

func main() {
	// Load environment variables from .env.local file
	if err := godotenv.Load(".env.local"); err != nil {
		log.Println("Could not load .env.local file, using system environment variables")
	}

	// Connect to the database
	if err := config.ConnectDB(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize the router with all routes
	router := routes.SetupRouter(config.DB)

	// Configure trusted proxies
	if err := router.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		log.Printf("Warning: Failed to set trusted proxies: %v", err)
	}

	// Start the server
	serverAddr := ":8080"
	log.Printf("Server starting on http://localhost%s\n", serverAddr)
	
	// Create a channel to listen for interrupt signals
	server := &http.Server{
		Addr:    serverAddr,
		Handler: router,
	}

	// Start the server in a goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Create a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown the server gracefully
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited properly")
}
