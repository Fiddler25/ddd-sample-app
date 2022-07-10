package comment

import "github.com/Fiddler25/ddd-sample-app/domain/model"

type CreateRequest struct {
	Body   string       `json:"body" validate:"required,max=140"`
	UserID model.UserID `json:"user_id" validate:"required"`
}
