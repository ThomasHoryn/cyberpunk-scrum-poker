package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"crypto/tls"
	"github.com/ThomasHoryn/cyberpunk-scrum-poker/internal/config"
)

type MongoDB struct {
	Client *mongo.Client
	Rooms  *mongo.Collection
}

func NewMongoDB(uri string) (*MongoDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Client().
		ApplyURI(uri).
		SetMaxPoolSize(100).
		SetMinPoolSize(4).
		SetServerSelectionTimeout(5 * time.Second)

	// Add TLS configuration only in production
	if config.Load().Environment == "production" {
		opts.SetTLSConfig(&tls.Config{
			InsecureSkipVerify: false,
			MinVersion:         tls.VersionTLS12,
		})
	}

	client, err := mongo.Connect(ctx, opts)	
	if err != nil {
		return nil, fmt.Errorf("mongo connection failed: %w", err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("mongo ping failed: %w", err)
	}

	db := client.Database("cyberpunkscrum")

	return &MongoDB{
		Client: client,
		Rooms:  db.Collection("rooms"),
	}, nil
}

func (m *MongoDB) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	if err := m.Client.Disconnect(ctx); err != nil {
		log.Printf("MongoDB disconnect error: %v", err)
	}
}
