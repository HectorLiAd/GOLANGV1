package BD

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func InitBD() *sql.DB {
	connectionStr := "root:hector@tcp(localhost:3306)/bdgolang"

	databaseConnection, err := sql.Open("mysql", connectionStr)

	if err != nil || databaseConnection == nil {
		panic(err.Error())
	}
	return databaseConnection
}
