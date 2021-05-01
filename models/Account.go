package models

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Investments []Investment
	ID          uint32    `gorm:"primary_key;auto_increment" json:"id"`
	User        User      `json:"user"`
	Bank        Bank      `json:"bank"`
	UserID      uint32    `gorm:"not null" json:"user_id"`
	BankID      uint32    `gorm:"not null" json:"bank_id"`
	Number      string    `gorm:"size:255;not null;unique" json:"number"`
	Balance     int32     `gorm:"not null" json:"balance"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
