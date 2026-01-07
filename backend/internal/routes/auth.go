package routes

import (
	"database/sql"
	"music-migration/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(router fiber.Router, db *sql.DB) {
	authHandler := handlers.NewAuthHandler(db)
	
	auth := router.Group("/auth")
	
	// Spotify OAuth
	auth.Get("/spotify", authHandler.SpotifyLogin)
	auth.Get("/spotify/callback", authHandler.SpotifyCallback)
	
	// Apple Music OAuth
	auth.Get("/apple", authHandler.AppleMusicLogin)
	auth.Get("/apple/callback", authHandler.AppleMusicCallback)
	
	// Auth status
	auth.Get("/status", authHandler.GetAuthStatus)
}
