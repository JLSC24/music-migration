package handlers

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	db *sql.DB
}

func NewAuthHandler(db *sql.DB) *AuthHandler {
	return &AuthHandler{db: db}
}

// SpotifyLogin initiates Spotify OAuth flow
func (h *AuthHandler) SpotifyLogin(c *fiber.Ctx) error {
	// TODO: Implement Spotify OAuth flow
	return c.JSON(fiber.Map{
		"message": "Spotify OAuth not yet implemented",
		"auth_url": "https://accounts.spotify.com/authorize",
	})
}

// SpotifyCallback handles Spotify OAuth callback
func (h *AuthHandler) SpotifyCallback(c *fiber.Ctx) error {
	code := c.Query("code")
	if code == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Missing authorization code"})
	}

	// TODO: Exchange code for tokens
	return c.JSON(fiber.Map{
		"message": "Callback received",
		"code": code,
	})
}

// AppleMusicLogin initiates Apple Music OAuth flow
func (h *AuthHandler) AppleMusicLogin(c *fiber.Ctx) error {
	// TODO: Implement Apple Music OAuth flow
	return c.JSON(fiber.Map{
		"message": "Apple Music OAuth not yet implemented",
	})
}

// AppleMusicCallback handles Apple Music OAuth callback
func (h *AuthHandler) AppleMusicCallback(c *fiber.Ctx) error {
	// TODO: Handle Apple Music callback
	return c.JSON(fiber.Map{
		"message": "Apple Music callback not yet implemented",
	})
}

// GetAuthStatus returns current authentication status
func (h *AuthHandler) GetAuthStatus(c *fiber.Ctx) error {
	userID := c.Locals("userID").(int)
	
	rows, err := h.db.Query(`
		SELECT provider, expires_at 
		FROM auth_tokens 
		WHERE user_id = $1 AND expires_at > NOW()
	`, userID)
	
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
	}
	defer rows.Close()

	authenticated := make(map[string]bool)
	for rows.Next() {
		var provider string
		var expiresAt interface{}
		rows.Scan(&provider, &expiresAt)
		authenticated[provider] = true
	}

	return c.JSON(fiber.Map{
		"spotify": authenticated["spotify"],
		"apple_music": authenticated["apple_music"],
	})
}
