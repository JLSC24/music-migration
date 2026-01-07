package routes

import (
	"database/sql"
	"music-migration/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupPlaylistRoutes(router fiber.Router, db *sql.DB) {
	playlistHandler := handlers.NewPlaylistHandler(db)
	
	playlists := router.Group("/playlists")
	
	playlists.Get("/:provider", playlistHandler.GetPlaylists)
	playlists.Get("/:provider/:playlistId", playlistHandler.GetPlaylistTracks)
}
