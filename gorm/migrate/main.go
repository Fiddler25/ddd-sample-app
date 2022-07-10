package main

import (
	"github.com/Fiddler25/ddd-sample-app/gorm"
	"github.com/Fiddler25/ddd-sample-app/gorm/migrate/comment"
	"github.com/Fiddler25/ddd-sample-app/gorm/migrate/user"
)

func main() {
	db := gorm.DB()

	user.Create(db)
	user.AddPassword(db)
	comment.Create(db)
}
