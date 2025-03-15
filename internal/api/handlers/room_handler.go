package handlers

import (
	"context"
	"errors"
	"net/http"
	"time"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ThomasHoryn/cyberpunk-scrum-poker/internal/db"
	"github.com/ThomasHoryn/cyberpunk-scrum-poker/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoomHandler struct {
	DB *db.MongoDB
}

func NewRoomHandler(db *db.MongoDB) *RoomHandler {
	return &RoomHandler{DB: db}
}

func (h *RoomHandler) CreateRoom(c *gin.Context) {
	var request struct {
		Name string `json:"name" binding:"required,min=3,max=50,alphanumspace"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room name, only alphanumeric characters and spaces allowed (3-50 characters)"})
		return
	}

	newRoom := models.NewRoom(request.Name)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err := h.DB.Rooms.InsertOne(ctx, newRoom); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create room"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         newRoom.ID.Hex(),
		"name":       newRoom.Name,
		"created_at": newRoom.CreatedAt,
	})
}

func (h *RoomHandler) GetRoom(c *gin.Context) {
	roomID := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(roomID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room ID format"})
		return
	}

	var room models.Room
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = h.DB.Rooms.FindOne(ctx, bson.M{"_id": objID}).Decode(&room)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
		} else {
			log.Printf("Database error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         room.ID.Hex(),
		"name":       room.Name,
		"created_at": room.CreatedAt,
		"users":      room.Users,
	})
}
