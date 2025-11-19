package models

import (
	"time"
)

type Activity struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	Action    string
	ItemID    uint
	CreatedAt time.Time
}
