package models

import (
	"time"

	"github.com/google/uuid"
)

// BlogsSwagger Blogs Swagger model
type BlogsSwagger struct {
	Title   string `json:"title" db:"title" validate:"required,gte=3"`
	Content string `json:"content" db:"content" validate:"required,gte=10"`
}

// Blog model
type Blog struct {
	ID        uuid.UUID `json:"id" db:"id" validate:"omitempty,uuid"`
	Title     string    `json:"title" db:"title" validate:"required,gte=3"`
	Content   string    `json:"content" db:"content" validate:"required,gte=10"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// BlogsList All Blogs response
type BlogsList struct {
	TotalCount int     `json:"total_count"`
	TotalPages int     `json:"total_pages"`
	Page       int     `json:"page"`
	Size       int     `json:"size"`
	HasMore    bool    `json:"has_more"`
	Blogs      []*Blog `json:"blogs"`
}
