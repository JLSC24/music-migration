package handlers

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

type PlaylistHandler struct {
	db *sql.DB
}

func NewPlaylistHandler(db *sql.DB) *PlaylistHandler {
	return &PlaylistHandler{db: db}
}

// GetPlaylists returns user's playlists from a provider
func (h *PlaylistHandler) GetPlaylists(c *fiber.Ctx) error {
	provider := c.Params("provider")
	
	// TODO: Fetch from Spotify/Apple Music API
	mockPlaylists := []map[string]interface{}{
		{
			"id":          "playlist1",
			"name":        "My Favorite Songs",
			"track_count": 45,
			"provider":    provider,
		},
		{
			"id":          "playlist2",
			"name":        "Workout Mix",
			"track_count": 32,
			"provider":    provider,
		},
	}

	return c.JSON(fiber.Map{
		"playlists": mockPlaylists,
		"provider":  provider,
	})
}

// GetPlaylistTracks returns tracks in a playlist
func (h *PlaylistHandler) GetPlaylistTracks(c *fiber.Ctx) error {
	provider := c.Params("provider")
	playlistID := c.Params("playlistId")

	// TODO: Fetch tracks from API
	mockTracks := []map[string]interface{}{
		{
			"id":      "track1",
			"name":    "Song Name 1",
			"artists": []string{"Artist 1"},
			"album":   "Album Name",
		},
	}

	return c.JSON(fiber.Map{
		"playlist_id": playlistID,
		"provider":    provider,
		"tracks":      mockTracks,
		"total":       len(mockTracks),
	})
}
