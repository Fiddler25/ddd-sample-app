package model

type CommentID int

type Comment struct {
	ID     CommentID
	Body   string
	UserID UserID
}

func NewComment(body string, userID UserID) *Comment {
	return &Comment{
		Body:   body,
		UserID: userID,
	}
}
