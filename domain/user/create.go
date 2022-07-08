package user

import (
	"gorm.io/gorm"
	"log"
)

type CreateRepository struct {
	db *gorm.DB
}

func NewCreateRepository(db *gorm.DB) CreateRepository {
	return CreateRepository{db: db}
}

func (r CreateRepository) Create(ret User) User {
	if err := r.db.Create(&ret).Error; err != nil {
		log.Fatal(err)
	}

	return ret
}
