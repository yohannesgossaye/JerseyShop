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
	"github.com/yohannesgossaye/pkgs/logger"
	"github.com/yohannesgossaye/pkgs/utils/email"
)

func Init() {

	logger := logger.NewLogger()

	// Initialize database
	logger.Infof("Initializing database connection...")
	db := InitDatabase()
	defer db.Close()

	logger.Infof("Database connection initialized")
	repo := InitRepository(db)

	// Initialize service
	logger.Infof("Initializing service...")
	service := InitService(repo, email.NewSMTPSender())

	// Initialize handler
	logger.Infof("Initializing handler...")
	handler := InitHandler(service, logger)

	router := chi.NewRouter()

	// Initialize router with middleware and routes
	InitRouter(router, handler)

	// Initialize server
	logger.Infof("Initializing server...")
	server := InitServer(router)

	// Start server in a goroutine
	go func() {

		port := server.Addr
		if len(port) > 0 && port[0] == ':' {
			port = port[1:]
		}

		logger.Infof("🔥 Server is running on port %s", port)
		log.Printf("Server starting on %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := server.Shutdown(ctx); err != nil {
		logger.Errorf("Server shutdown failed: %v", err)
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}
