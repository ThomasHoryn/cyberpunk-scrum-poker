package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"net/http"

	"github.com/ThomasHoryn/cyberpunk-scrum-poker/internal/api"
	"github.com/ThomasHoryn/cyberpunk-scrum-poker/internal/config"
	"github.com/ThomasHoryn/cyberpunk-scrum-poker/internal/db"
)

func main() {
	cfg := config.Load()
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
}
