package user

import (
	"github.com/Fiddler25/sample-app/domain/user"
	"gorm.io/gorm"
)

type UpdateRequest struct {
	ID   user.UserID `json:"id" validate:"required"`
	Name string      `json:"name" validate:"max=50"`
}

type UpdateUsecase struct {
	db *gorm.DB
}

func NewUpdateUsecase(db *gorm.DB) UpdateUsecase {
	return UpdateUsecase{db: db}
}

func (uc UpdateUsecase) Execute(req UpdateRequest) (*user.User, error) {
	usr := user.New(req.ID, req.Name)

	uRepo := user.NewRepository(uc.db)
	if _, err := uRepo.GetByUserID(usr.ID); err != nil {
		return nil, err
	}
	uRepo.Update(usr)

	return usr, nil
}
