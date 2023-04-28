package models

import "time"

type News struct {
	ID          int
	Title       string `gorm:"size:255"`
	Description string `gorm:"size:255"`
	Author      string `gorm:"size:255"`
	PhoneNumber int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
