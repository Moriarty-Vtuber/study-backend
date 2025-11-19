package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/yourname/my-study-space/internal/db"
	"github.com/yourname/my-study-space/internal/youtube"
)

func main() {
	// 1. Connect to Supabase
	connStr := os.Getenv("DATABASE_URL")
	database, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	dbClient := db.New(database)

	// 2. Start YouTube Monitor (Background Worker)
	apiKey := os.Getenv("YOUTUBE_API_KEY")
	videoID := os.Getenv("VIDEO_ID") // The ID of your active live stream
	
    // Run monitor in a goroutine so it doesn't block the server
	go youtube.StartMonitor(apiKey, videoID, dbClient)

	// 3. Start Web Server (Gin)
	r := gin.Default()
	
    // Serve static files (The OBS Overlay)
	r.Static("/overlay", "./public")

	// API for the Overlay to fetch room state
	r.GET("/api/room", func(c *gin.Context) {
		users, err := dbClient.GetCurrentUsers()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"users": users})
	})

    // Keep-alive endpoint for free tier
    r.GET("/ping", func(c *gin.Context) {
        c.String(200, "pong")
    })

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
