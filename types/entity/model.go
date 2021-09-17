package entity

import (
	"time"
)

type Model struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"createddAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
