package routes

import (
	"database/sql"
	"music-migration/internal/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func SetupMigrationRoutes(router fiber.Router, db *sql.DB, redis *redis.Client) {
	migrationHandler := handlers.NewMigrationHandler(db, redis)
	
	migration := router.Group("/migrations")
	
	migration.Post("/", migrationHandler.StartMigration)
	migration.Get("/", migrationHandler.ListMigrations)
	migration.Get("/:id", migrationHandler.GetMigrationStatus)
}
