package user

type UserID int

type User struct {
	ID   UserID
	Name string
}

func New(userID UserID, name string) *User {
	return &User{
		ID:   userID,
		Name: name,
	}
}
