package comment

import (
	"gorm.io/gorm"
	"log"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{db: db}
}

func (r Repository) Create(ret *Comment) *Comment {
	if err := r.db.Create(&ret).Error; err != nil {
		log.Fatal(err)
	}
	return ret
}
