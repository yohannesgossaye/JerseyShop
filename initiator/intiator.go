package initiator

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/yohannesgossaye/pkgs/utils/email"
)

// Init is the main entry point to start the server
func Init() {
	// Initialize database
	db := InitDatabase()
	defer db.Close()

	// Initialize repository
	repo := InitRepository(db)

	// Initialize service
	service := InitService(repo, email.NewSMTPSender())

	// Initialize handler
	handler := InitHandler(service)

	// Create a single router instance
	router := chi.NewRouter()

	// Initialize router with middleware and routes
	InitRouter(router, handler)

	// Initialize server
	server := InitServer(router)

	// Start server in a goroutine
	go func() {
		log.Printf("Server starting on %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Server is shutting down...")

	// Give outstanding requests a deadline for completion
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}
