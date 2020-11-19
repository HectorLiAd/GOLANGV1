package BD

import (
	"database/sql"

	// _ "github.com/mattn/go-oci8"
	_ "github.com/go-sql-driver/mysql"
)

func InitBD() *sql.DB {
	connectionStr := "root:hector@tcp(localhost:3306)/pruebago"
	// connectionStr := "USER_01/123@localhost:1521/ORCL"
	databaseConnection, err := sql.Open("mysql", connectionStr)

	if err != nil || databaseConnection == nil {
		panic(err.Error())
	}
	return databaseConnection
}
