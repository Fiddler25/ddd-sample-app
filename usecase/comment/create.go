package comment

import (
	"github.com/Fiddler25/sample-app/domain/comment"
	"github.com/Fiddler25/sample-app/domain/user"
	"gorm.io/gorm"
)

type CreateRequest struct {
	Body   string      `json:"body" validate:"required,max=140"`
	UserID user.UserID `json:"user_id" validate:"required"`
}

type CreateUsecase struct {
	db *gorm.DB
}

func NewCreateUsecase(db *gorm.DB) CreateUsecase {
	return CreateUsecase{db: db}
}

func (uc CreateUsecase) Execute(req CreateRequest) *comment.Comment {
	c := comment.New(req.Body, req.UserID)
	comment.NewRepository(uc.db).Create(c)

	return c
}
