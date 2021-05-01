package models

import (
	"time"

	"gorm.io/gorm"
)

type Investment struct {
	gorm.Model
	ID          uint32      `gorm:"primary_key;auto_increment" json:"id"`
	User        User        `json:"user"`
	Account     Account     `json:account`
	UserID      uint32      `gorm:"not null" json:"user_id"`
	AccountID   uint32      `gorm:"not null" json:"account_id"`
	Type        investments `gorm:"size:255;not null" json:"type"`
	Amount      int32       `gorm:"not null" json:"amount"`
	EndDate     time.Time   `gorm:"not null" json: "end_date"`
	TotalReturn int32       `gorm:"not null" json: "total_return"`
	TotalProfit int32       `gorm:"not null" json: "total_profit"`
	Description string      `json: "description"`
	CreatedAt   time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
