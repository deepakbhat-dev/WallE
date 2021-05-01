package models

import (
	"errors"
	"html"
	"strings"
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

func (p *Account) Prepare() {
	p.ID = 0
	p.Title = html.EscapeString(strings.TrimSpace(p.Title))
	p.Content = html.EscapeString(strings.TrimSpace(p.Content))
	p.User = User{}
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Account) Validate() error {

	if p.Title == "" {
		return errors.New("required title")
	}
	if p.Content == "" {
		return errors.New("required content")
	}
	if p.UserID < 1 {
		return errors.New("required user")
	}
	return nil
}

func (p *Account) SavePost(db *gorm.DB) (*Account, error) {
	var err error
	err = db.Debug().Model(&Account{}).Create(&p).Error
	if err != nil {
		return &Account{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
		if err != nil {
			return &Account{}, err
		}
	}
	return p, nil
}

func (p *Account) FindAllPosts(db *gorm.DB) (*[]Account, error) {
	posts := []Account{}
	var err error = db.Debug().Model(&Account{}).Limit(100).Find(&posts).Error
	if err != nil {
		return &[]Account{}, err
	}
	if len(posts) > 0 {
		for i := range posts {
			err := db.Debug().Model(&User{}).Where("id = ?", posts[i].UserID).Take(&posts[i].User).Error
			if err != nil {
				return &[]Account{}, err
			}
		}
	}
	return &posts, nil
}

func (p *Account) FindPostByID(db *gorm.DB, pid uint64) (*Account, error) {
	var err error = db.Debug().Model(&Account{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Account{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
		if err != nil {
			return &Account{}, err
		}
	}
	return p, nil
}

func (p *Account) UpdateAPost(db *gorm.DB) (*Account, error) {

	var err error = db.Debug().Model(&Account{}).Where("id = ?", p.ID).Updates(Account{Title: p.Title, Content: p.Content, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &Account{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
		if err != nil {
			return &Account{}, err
		}
	}
	return p, nil
}

func (p *Account) DeleteAPost(db *gorm.DB, pid uint64, uid uint32) (int64, error) {

	db = db.Debug().Model(&Account{}).Where("id = ? and user_id = ?", pid, uid).Take(&Account{}).Delete(&Account{})

	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			return 0, errors.New("Account not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
