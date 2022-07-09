package model

import (
	"github.com/Fiddler25/ddd-sample-app/domain/vo"
	"time"
)

type User struct {
	ID        int         `gorm:"primaryKey"`
	Name      string      `gorm:"not null;size:50"`
	Email     string      `gorm:"not null;size:255;unique"`
	Password  vo.Password `gorm:"not null"`
	CreatedAt time.Time   `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time   `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

func NewUser(email string, password vo.Password) *User {
	return &User{
		Email:    email,
		Password: password,
	}
}

func NewUpdateUser(userID int, name string, email string, password vo.Password) *User {
	return &User{
		ID:       userID,
		Name:     name,
		Email:    email,
		Password: password,
	}
}
