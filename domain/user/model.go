package user

import "time"

type User struct {
	ID    int    `gorm:"primaryKey"`
	Name  string `gorm:"not null;size:50"`
	Email string `gorm:"not null;size:255;unique"`
	CreatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
