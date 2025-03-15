package handlers

import (
	"net/http"
	"time"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/ThomasHoryn/cyberpunk-scrum-poker/internal/db"
	"github.com/ThomasHoryn/cyberpunk-scrum-poker/internal/models"
	"go.mongodb.org/mongo-driver/bson"
)

type RoomHandler struct {
	DB *db.MongoDB
}

func NewRoomHandler(db *db.MongoDB) *RoomHandler {
	return &RoomHandler{DB: db}
}

func (h *RoomHandler) CreateRoom(c *gin.Context) {
	var request struct {
		Name string `json:"name" binding:"required,min=3,max=50"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room name"})
		return
	}

	newRoom := models.NewRoom(request.Name)
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	if _, err := h.DB.Rooms.InsertOne(ctx, newRoom); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create room"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         newRoom.ID,
		"name":       newRoom.Name,
		"created_at": newRoom.CreatedAt,
	})

	
}

func (h *RoomHandler) GetRoom(c *gin.Context) {
	roomID := c.Param("id")
	
	var room models.Room
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	err := h.DB.Rooms.FindOne(ctx, bson.M{"_id": roomID}).Decode(&room)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         room.ID,
		"name":       room.Name,
		"created_at": room.CreatedAt,
		"users":      room.Users,
	})
}
