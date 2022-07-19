package comment

import "github.com/Fiddler25/ddd-sample-app/domain/user"

type CommentID int

type Comment struct {
	ID     CommentID
	Body   string
	UserID user.UserID
}

func New(body string, userID user.UserID) *Comment {
	return &Comment{
		Body:   body,
		UserID: userID,
	}
}
