package models

import "time"

type Repository struct {
	ID           int64         `json:"id"`
	Name         string        `json:"name"`
	FullName     string        `json:"full_name"`
	Description  string        `json:"description,omitempty"`
	Private      bool          `json:"private"`
	UpdatedAt    time.Time     `json:"updated_at"`
	Language     string        `json:"language,omitempty"`
	Organization *Organization `json:"organization,omitempty"`
}

type Organization struct {
	Login     string `json:"login"`
	AvatarURL string `json:"avatar_url"`
}

type PaginatedResponse struct {
	Data       []Repository `json:"data"`
	Total      int          `json:"total"`
	Page       int          `json:"page"`
	PerPage    int          `json:"per_page"`
	TotalPages int          `json:"total_pages"`
}
