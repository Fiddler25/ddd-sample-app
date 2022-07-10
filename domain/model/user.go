package model

import (
	"github.com/Fiddler25/ddd-sample-app/domain/vo"
	"time"
)

type UserID int

type User struct {
	ID        UserID      `gorm:"primaryKey"`
	Name      string      `gorm:"not null;size:50"`
	Email     string      `gorm:"not null;size:255;unique"`
	Password  vo.Password `gorm:"not null"`
	Comments  []Comment   `gorm:"constraint:OnDelete:CASCADE"`
	CreatedAt time.Time   `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time   `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

func NewCreateUser(email string, password vo.Password) *User {
	return &User{
		Email:    email,
		Password: password,
	}
}

func NewUpdateUser(userID UserID, name string, email string, password vo.Password) *User {
	return &User{
		ID:       userID,
		Name:     name,
		Email:    email,
		Password: password,
	}
}
