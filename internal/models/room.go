package models

import (
	"time"
	
	"github.com/google/uuid"
)

type Room struct {
	ID        string    `json:"id" bson:"_id"`
	Name      string    `json:"name" bson:"name"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	Users     []User    `json:"users" bson:"users"`
}

type User struct {
	ID       string    `json:"id" bson:"id"`
	Username string    `json:"username" bson:"username"`
	JoinedAt time.Time `json:"joined_at" bson:"joined_at"`
}

func NewRoom(name string) *Room {
	return &Room{
		ID:        uuid.New().String(),
		Name:      name,
		CreatedAt: time.Now().UTC(),
		Users:     make([]User, 0),
	}
}
