package models

import (
	"github.com/google/uuid"
	"time"
)

type Product struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Price       int       `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Images      []string  `json:"images"`
	CategoryID  uuid.UUID `json:"category_id"`
	Description string    `json:"description"`
	Version     uint      `json:"version"`
}
