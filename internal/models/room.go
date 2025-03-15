package models

import (
	"time"
	
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Room struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	CreatedAt time.Time          `bson:"created_at"`
	Users     []User             `bson:"users"`
}

type User struct {
	ID       string    `bson:"id"`
	Username string    `bson:"username"`
	JoinedAt time.Time `bson:"joined_at"`
}

func NewRoom(name string) *Room {
	return &Room{
		ID:        primitive.NewObjectID(),
		Name:      name,
		CreatedAt: time.Now().UTC(),
		Users:     make([]User, 0),
	}
}
