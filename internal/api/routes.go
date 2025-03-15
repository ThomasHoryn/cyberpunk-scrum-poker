package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ThomasHoryn/cyberpunk-scrum-poker/internal/config"
	"github.com/ThomasHoryn/cyberpunk-scrum-poker/internal/db"
	"github.com/ThomasHoryn/cyberpunk-scrum-poker/internal/api/handlers"
	"github.com/ThomasHoryn/cyberpunk-scrum-poker/internal/api/validators"
)

func SetupRouter(mongoDB *db.MongoDB) *gin.Engine {
	cfg := config.Load()
	router := gin.Default()
	
	router.Use(
		CORS(cfg),
		SecurityHeaders(),
		RateLimiter(cfg),
		APIKeyAuth(cfg),
		HMACValidation(cfg),
	)
	
	validators.RegisterValidators()
	
	roomHandler := handlers.NewRoomHandler(mongoDB)

	v1 := router.Group("/api/v1")
	{
		rooms := v1.Group("/rooms")
		{
			rooms.POST("", roomHandler.CreateRoom)
			rooms.GET("/:id", roomHandler.GetRoom)
		}
	}

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	return router
}
