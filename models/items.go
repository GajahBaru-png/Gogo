package models

import "time"

type Items struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Quantity    int       `json:"quantity" binding:"gte=0"`
	Description string    `json:"description"`
	AddedBy     string    `json:"added_by"`
	CreatedAt   time.Time `json:"created_at"`
}
