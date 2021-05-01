package models

import (
	"time"

	"gorm.io/gorm"
)

type Bank struct {
	gorm.Model
	Accounts     []Account
	ID           uint32    `gorm:"primary_key;auto_increment" json:"id"`
	User         User      `json:"user"`
	UserID       uint32    `gorm:"not null" json:"user_id"`
	Name         banks     `gorm:"size:255;not null;unique" json:"name"`
	Interest     float32   `gorm: "not null" json:"interest"`
	InterestOnFd float32   `gorm: "not null" json:"interest_on_fd"`
	InterestOnRd float32   `gorm: "not null" json:"interest_on_rd"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
