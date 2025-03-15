package api

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"time"
	
	"github.com/gin-gonic/gin"
	"github.com/ThomasHoryn/cyberpunk-scrum-poker/internal/config"
	"golang.org/x/time/rate"
)

func SecurityHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-Frame-Options", "DENY")
		c.Header("Content-Security-Policy", "default-src 'self'")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Next()
	}
}

func CORS(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		for _, host := range cfg.AllowedHosts {
			if host == origin {
				c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
				break
			}
		}
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-API-Key, X-Signature, X-Timestamp")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func APIKeyAuth(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-Key")
		if apiKey != cfg.APIKey {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid API credentials",
			})
			return
		}
		c.Next()
	}
}

func HMACValidation(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		signature := c.GetHeader("X-Signature")
		timestamp := c.GetHeader("X-Timestamp")
		
		if timestamp == "" || signature == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Missing security headers"})
			return
		}
		
		ts, err := time.Parse(time.RFC3339, timestamp)
		if err != nil || time.Since(ts) > 5*time.Minute {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid timestamp"})
			return
		}
		
		body, _ := c.GetRawData()
		mac := hmac.New(sha256.New, []byte(cfg.HMACSecret))
		mac.Write(body)
		mac.Write([]byte(timestamp))
		expectedMAC := hex.EncodeToString(mac.Sum(nil))
		
		if !hmac.Equal([]byte(signature), []byte(expectedMAC)) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid signature"})
			return
		}
		
		c.Next()
	}
}

func RateLimiter(cfg *config.Config) gin.HandlerFunc {
	limiter := rate.NewLimiter(rate.Every(cfg.RateLimitWindow), cfg.RateLimit)
	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many requests",
			})
			return
		}
		c.Next()
	}
}
