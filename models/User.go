package models

import (
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Accounts        []Account
	MonthlyExpenses []MonthlyExpense
	Investments     []Investment
	ID              uint32    `gorm:"primary_key;auto_increment" json:"id"`
	FullName        string    `gorm:"size:255;not null;unique" json:"full_name"`
	DateOfBirth     time.Time `gorm:"not null" json:"date_of_birth"`
	Email           string    `gorm:"size:100;not null;unique" json:"email"`
	Password        string    `gorm:"size:100;not null;" json:"password"`
	CurrentAge      int32     `gorm:"not null" json:"current_age"`
	RetirementAge   int32     `gorm:"check:RetirementAge < 65;not null" json:"retirement_age"`
	Income          int32     `gorm:"check:Income > 0;not null" json:"income"`
	DaysLeft        int32     `gorm:"not null;" json:"days_left"`
	CreatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	u.CurrentAge = int32(u.DateOfBirth.Year()) - int32(time.Now().Year())
	u.DaysLeft = int32(float32(u.RetirementAge-u.CurrentAge) * 365.25)
	return nil
}

func (u *User) Prepare() {
	u.ID = 0
	u.FullName = html.EscapeString(strings.TrimSpace(u.FullName))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.FullName == "" {
			return errors.New("required fullname")
		}
		if time.Now().Before(u.DateOfBirth) {
			return errors.New("birth date is in the future")
		}
		if u.Password == "" {
			return errors.New("required password")
		}
		if u.Email == "" {
			return errors.New("required email")
		}
		if u.Income == 0 {
			return errors.New("required income")
		}
		if u.RetirementAge == 0 || u.RetirementAge < u.CurrentAge {
			return errors.New("invalid retirement age")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("invalid email")
		}
		return nil

	case "login":
		if u.Password == "" {
			return errors.New("required password")
		}
		if u.Email == "" {
			return errors.New("required email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("invalid email")
		}
		return nil

	default:
		if u.FullName == "" {
			return errors.New("required fullname")
		}
		if time.Now().Before(u.DateOfBirth) {
			return errors.New("birth date is in the future")
		}
		if u.Password == "" {
			return errors.New("required password")
		}
		if u.Email == "" {
			return errors.New("required email")
		}
		if u.Income == 0 {
			return errors.New("required income")
		}
		if u.RetirementAge == 0 || u.RetirementAge < u.CurrentAge {
			return errors.New("invalid retirement age")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("invalid email")
		}
		return nil
	}
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	users := []User{}
	var err error = db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

func (u *User) FindUserByID(db *gorm.DB, uid uint32) (*User, error) {
	var err error = db.Debug().Model(User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &User{}, errors.New("User Not Found")
	}
	return u, err
}

func (u *User) UpdateUser(db *gorm.DB, uid uint32) (*User, error) {
	// To hash the password
	err := u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"password":  u.Password,
			"full_name": u.FullName,
			"email":     u.Email,
			"update_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &User{}, db.Error
	}
	// This is the display the updated user
	err = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) DeleteUser(db *gorm.DB, uid uint32) (int64, error) {
	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).Select("BankAccounts").Delete(&User{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
