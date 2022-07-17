package comment

import entity "github.com/Fiddler25/ddd-sample-app/domain/user"

type CommentID int

type Comment struct {
	ID     CommentID
	Body   string
	UserID entity.UserID
}

func NewEntity(body string, userID entity.UserID) *Comment {
	return &Comment{
		Body:   body,
		UserID: userID,
	}
}
