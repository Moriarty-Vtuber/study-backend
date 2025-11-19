package db

import (
	"database/sql"
	"log"
	"time"
)

type Client struct {
	db *sql.DB
}

type User struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	DisplayName     string `json:"display_name"`
	ProfileImageURL string `json:"profile_image_url"`
}

func New(db *sql.DB) *Client {
	return &Client{db: db}
}

func (c *Client) EnterRoom(userID, userName, userImage string) {
	// 1. Upsert User
	_, err := c.db.Exec(`
		INSERT INTO users (id, name, display_name, profile_image_url, last_entered_at)
		VALUES ($1, $2, $2, $3, NOW())
		ON CONFLICT (id) DO UPDATE SET
			name = EXCLUDED.name,
			display_name = EXCLUDED.display_name,
			profile_image_url = EXCLUDED.profile_image_url,
			last_entered_at = NOW();
	`, userID, userName, userImage)
	if err != nil {
		log.Printf("Error upserting user: %v", err)
		return
	}

	// 2. Add to Room State
	// Assign a seat_id if not exists (simplified logic: just insert)
	// In a real app, you'd find the first available seat.
	_, err = c.db.Exec(`
		INSERT INTO room_state (user_id, status, entered_at)
		VALUES ($1, 'studying', NOW())
		ON CONFLICT (user_id) DO UPDATE SET
			status = 'studying',
			entered_at = NOW();
	`, userID)
	if err != nil {
		log.Printf("Error entering room: %v", err)
	}
}

func (c *Client) ExitRoom(userID string) {
	// Remove from room_state
	_, err := c.db.Exec(`DELETE FROM room_state WHERE user_id = $1`, userID)
	if err != nil {
		log.Printf("Error exiting room: %v", err)
	}
	
	// Optional: Log the exit in study_logs
}

func (c *Client) GetCurrentUsers() ([]User, error) {
	rows, err := c.db.Query(`
		SELECT u.id, u.name, u.display_name, u.profile_image_url
		FROM room_state r
		JOIN users u ON r.user_id = u.id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.DisplayName, &u.ProfileImageURL); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}
