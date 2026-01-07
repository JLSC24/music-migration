package handlers

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

type MigrationHandler struct {
	db    *sql.DB
	redis *redis.Client
}

func NewMigrationHandler(db *sql.DB, redis *redis.Client) *MigrationHandler {
	return &MigrationHandler{db: db, redis: redis}
}

// StartMigration initiates a new migration job
func (h *MigrationHandler) StartMigration(c *fiber.Ctx) error {
	var req struct {
		SourceProvider string   `json:"source_provider"`
		TargetProvider string   `json:"target_provider"`
		PlaylistIDs    []string `json:"playlist_ids"`
		MigrateLibrary bool     `json:"migrate_library"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Validate providers
	validProviders := map[string]bool{"spotify": true, "apple_music": true}
	if !validProviders[req.SourceProvider] || !validProviders[req.TargetProvider] {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid provider"})
	}

	if req.SourceProvider == req.TargetProvider {
		return c.Status(400).JSON(fiber.Map{"error": "Source and target must be different"})
	}

	userID := 1 // TODO: Get from JWT token

	// Create migration record
	var migrationID int
	err := h.db.QueryRow(`
		INSERT INTO migrations (user_id, source_provider, target_provider, status)
		VALUES ($1, $2, $3, 'pending')
		RETURNING id
	`, userID, req.SourceProvider, req.TargetProvider).Scan(&migrationID)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create migration"})
	}

	// TODO: Start background job for migration processing

	return c.Status(201).JSON(fiber.Map{
		"migration_id": migrationID,
		"status":       "pending",
		"message":      "Migration started",
	})
}

// GetMigrationStatus returns status of a migration job
func (h *MigrationHandler) GetMigrationStatus(c *fiber.Ctx) error {
	migrationID := c.Params("id")

	var migration struct {
		ID             int    `json:"id"`
		Status         string `json:"status"`
		TotalTracks    int    `json:"total_tracks"`
		MatchedTracks  int    `json:"matched_tracks"`
		FailedTracks   int    `json:"failed_tracks"`
		SourceProvider string `json:"source_provider"`
		TargetProvider string `json:"target_provider"`
	}

	err := h.db.QueryRow(`
		SELECT id, status, total_tracks, matched_tracks, failed_tracks, 
		       source_provider, target_provider
		FROM migrations 
		WHERE id = $1
	`, migrationID).Scan(
		&migration.ID,
		&migration.Status,
		&migration.TotalTracks,
		&migration.MatchedTracks,
		&migration.FailedTracks,
		&migration.SourceProvider,
		&migration.TargetProvider,
	)

	if err == sql.ErrNoRows {
		return c.Status(404).JSON(fiber.Map{"error": "Migration not found"})
	}

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
	}

	return c.JSON(migration)
}

// ListMigrations returns all migrations for current user
func (h *MigrationHandler) ListMigrations(c *fiber.Ctx) error {
	userID := 1 // TODO: Get from JWT token

	rows, err := h.db.Query(`
		SELECT id, source_provider, target_provider, status, 
		       total_tracks, matched_tracks, failed_tracks, created_at
		FROM migrations 
		WHERE user_id = $1
		ORDER BY created_at DESC
		LIMIT 50
	`, userID)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
	}
	defer rows.Close()

	migrations := []map[string]interface{}{}
	for rows.Next() {
		var m map[string]interface{} = make(map[string]interface{})
		var id, total, matched, failed int
		var source, target, status, createdAt string
		
		rows.Scan(&id, &source, &target, &status, &total, &matched, &failed, &createdAt)
		
		m["id"] = id
		m["source_provider"] = source
		m["target_provider"] = target
		m["status"] = status
		m["total_tracks"] = total
		m["matched_tracks"] = matched
		m["failed_tracks"] = failed
		m["created_at"] = createdAt
		
		migrations = append(migrations, m)
	}

	return c.JSON(fiber.Map{
		"migrations": migrations,
		"total":      len(migrations),
	})
}
