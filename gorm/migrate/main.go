package main

import (
	"github.com/Fiddler25/ddd-sample-app/gorm"
	"github.com/Fiddler25/ddd-sample-app/gorm/migrate/query"
)

func main() {
	db := gorm.DB()

	query.CreateTableUsers(db)
	query.CreateTableComments(db)
}
