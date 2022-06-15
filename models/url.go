package models

import "time"

type URL struct {
	Hash           *string    `json:"hash"`
	OriginalURL    string     `json:"original_url"`
	CreationDate   *time.Time `json:"creation_date"`
	ExpirationDate *time.Time `json:"expiration_date"`
	UserID         *int       `json:"user_id"`
}
