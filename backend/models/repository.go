package models

import "time"

type Repository struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	LastUpdate  time.Time `json:"last_update"`
	IsPrivate   bool      `json:"is_private"`
	Owner       string    `json:"owner"`
}

type PaginatedResponse struct {
	Repositories interface{} `json:"repositories"`
	CurrentPage  int         `json:"current_page"`
	TotalPages   int         `json:"total_pages"`
	TotalItems   int         `json:"total_items"`
	ItemsPerPage int         `json:"items_per_page"`
}
