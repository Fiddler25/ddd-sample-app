package repository

import (
	"github.com/Fiddler25/ddd-sample-app/domain/model"
	"gorm.io/gorm"
	"log"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewComment(db *gorm.DB) CommentRepository {
	return CommentRepository{db: db}
}

func (r CommentRepository) Create(ret *model.Comment) *model.Comment {
	if err := r.db.Create(&ret).Error; err != nil {
		log.Fatal(err)
	}
	return ret
}
