package query

import (
	"gorm.io/gorm"
)

func CreateTableUsers(db *gorm.DB) {
	db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id bigint AUTO_INCREMENT,
		name varchar(50) NOT NULL,
		email varchar(255) NOT NULL UNIQUE,
		password varchar(255) NOT NULL,
		created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		PRIMARY KEY (id)
	)`)
}
