package handlers

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/ThomasHoryn/cyberpunk-scrum-poker/internal/db"
	"github.com/ThomasHoryn/cyberpunk-scrum-poker/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"encoding/json"
	"io"
)

type RoomHandler struct {
	DB *db.MongoDB
}

func NewRoomHandler(db *db.MongoDB) *RoomHandler {
	return &RoomHandler{DB: db}
}

func (h *RoomHandler) CreateRoom(c *gin.Context) {
var request struct {
    Name string `json:"name" binding:"required,min=3,max=50,alphanumspacebrackets"`
}
 

    if err := c.ShouldBindJSON(&request); err != nil {
        // Handle JSON parsing errors first
        if errors.Is(err, io.EOF) {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Empty request body"})
            return
        }

        // Handle validation errors
        if fieldErrors, ok := err.(validator.ValidationErrors); ok {
            for _, fieldErr := range fieldErrors {
                switch fieldErr.Tag() {
                case "required":
                    c.JSON(http.StatusBadRequest, gin.H{"error": "Room name is required"})
                case "min":
                    c.JSON(http.StatusBadRequest, gin.H{"error": "Minimum 3 characters"})
                case "max":
                    c.JSON(http.StatusBadRequest, gin.H{"error": "Maximum 50 characters"})
case "alphanumspacebrackets":
    c.JSON(http.StatusBadRequest, gin.H{
        "error": "Only letters, numbers, spaces, hyphens, underscores and brackets [] allowed",
    })

                default:
                    c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed"})
                }
                return
            }
        }

        // Handle other JSON errors
        if jsonErr, ok := err.(*json.SyntaxError); ok {
            log.Printf("JSON syntax error: %v", jsonErr)
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
            return
        }

        // Fallback error
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
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
