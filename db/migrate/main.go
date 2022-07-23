package main

import (
	"github.com/Fiddler25/sample-app/db"
	"github.com/Fiddler25/sample-app/db/migrate/query"
)

func main() {
	conn := db.Conn()

	query.CreateTableUsers(conn)
	query.CreateTableComments(conn)
	query.AddColumnRememberDigest(conn)
	query.CreateTableSessions(conn)
	query.DropColumnRememberDigest(conn)
	query.AlterColumnName(conn)
}
