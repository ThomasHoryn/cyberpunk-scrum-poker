package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ThomasHoryn/cyberpunk-scrum-poker/internal/api"
	"github.com/ThomasHoryn/cyberpunk-scrum-poker/internal/api/validators"
	"github.com/ThomasHoryn/cyberpunk-scrum-poker/internal/config"
	"github.com/ThomasHoryn/cyberpunk-scrum-poker/internal/db"
)

func main() {
	cfg := config.Load()
	if cfg.MongoURI == "" {
		log.Fatal("MONGO_URI environment variable is required")
	}
	validators.RegisterValidators()

	mongoDB, err := db.NewMongoDB(cfg.MongoURI)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer mongoDB.Close()

	router := api.SetupRouter(mongoDB)

	server := &http.Server{
		Addr:    ":" + cfg.ServerPort,
		Handler: router,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("🚀 Server running on port %s", cfg.ServerPort)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	<-quit
	log.Println("🛑 Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
}
