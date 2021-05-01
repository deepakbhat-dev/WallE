package models

import (
	"time"

	"gorm.io/gorm"
)

type MonthlyExpense struct {
	gorm.Model
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	User      User      `json:"user"`
	UserID    uint32    `gorm:"uniqueIndex:me_index;not null" json:"user_id"`
	Reason    string    `gorm:"size:255;uniqueIndex:me_index;not null" json:"reason"`
	Amount    int32     `gorm:"not null" json:"amount"`
	EndDate   time.Time `json:"end_date"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
