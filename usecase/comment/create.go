package comment

import (
	"github.com/Fiddler25/ddd-sample-app/domain/model"
	"gorm.io/gorm"
)

type CreateRequest struct {
	Body   string       `json:"body" validate:"required,max=140"`
	UserID model.UserID `json:"user_id" validate:"required"`
}

type CreateUsecase struct {
	db *gorm.DB
}

func NewCreateUsecase(db *gorm.DB) CreateUsecase {
	return CreateUsecase{db: db}
}

func (u CreateUsecase) Execute(req CreateRequest) *model.Comment {
	return &model.Comment{}
}
