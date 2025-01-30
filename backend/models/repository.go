package models

import "time"

type Repository struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	LastUpdate  time.Time `json:"last_update"`
	IsPrivate   bool      `json:"is_private"`
	Owner       string    `json:"owner"`
}
