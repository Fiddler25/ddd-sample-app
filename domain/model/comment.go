package model

import "time"

type CommentID int

type Comment struct {
	ID        CommentID `gorm:"primaryKey"`
	Body      string    `gorm:"not null"`
	UserID    UserID    `gorm:"not null;index:idx_user_id_created_at"`
	User      User
	CreatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP;index:idx_user_id_created_at"`
	UpdatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

func NewComment(body string, userID UserID) *Comment {
	return &Comment{
		Body:   body,
		UserID: userID,
	}
}
