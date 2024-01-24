package models

import (
	"github.com/google/uuid"
	"time"
)

// New model
type New struct {
	ID        uuid.UUID `json:"id" db:"id" validate:"omitempty,uuid"`
	Title     string    `json:"title" db:"title" validate:"required,gte=3"`
	Content   string    `json:"content" db:"content" validate:"required,gte=10"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// NewsList All News response
type NewsList struct {
	TotalCount int    `json:"total_count"`
	TotalPages int    `json:"total_pages"`
	Page       int    `json:"page"`
	Size       int    `json:"size"`
	HasMore    bool   `json:"has_more"`
	News       []*New `json:"news"`
}

// NewsSwagger Swagger model
type NewsSwagger struct {
	Title   string `json:"title" db:"title" validate:"required,gte=3"`
	Content string `json:"content" db:"content" validate:"required,gte=10"`
}
