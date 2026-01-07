package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type AuthToken struct {
	ID           int       `json:"id"`
	UserID       int       `json:"user_id"`
	Provider     string    `json:"provider"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token,omitempty"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type Migration struct {
	ID             int       `json:"id"`
	UserID         int       `json:"user_id"`
	SourceProvider string    `json:"source_provider"`
	TargetProvider string    `json:"target_provider"`
	Status         string    `json:"status"`
	TotalTracks    int       `json:"total_tracks"`
	MatchedTracks  int       `json:"matched_tracks"`
	FailedTracks   int       `json:"failed_tracks"`
	CreatedAt      time.Time `json:"created_at"`
	CompletedAt    *time.Time `json:"completed_at,omitempty"`
}

type TrackMapping struct {
	ID              int       `json:"id"`
	MigrationID     int       `json:"migration_id"`
	SourceTrackID   string    `json:"source_track_id"`
	SourceTrackName string    `json:"source_track_name"`
	SourceArtist    string    `json:"source_artist"`
	TargetTrackID   string    `json:"target_track_id,omitempty"`
	Matched         bool      `json:"matched"`
	CreatedAt       time.Time `json:"created_at"`
}

type Track struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Artists  []string `json:"artists"`
	Album    string   `json:"album"`
	ISRC     string   `json:"isrc,omitempty"`
	Duration int      `json:"duration_ms"`
}

type Playlist struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description,omitempty"`
	TrackCount  int      `json:"track_count"`
	Tracks      []Track  `json:"tracks,omitempty"`
}
