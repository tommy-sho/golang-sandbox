package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var Conn *sql.DB

func init() {
	user := "root"
	pass := "5251"
	name := "itemsDB"

	db := user + ":" + pass + "@tcp(127.0.0.1:3306)/" + name
	conn, err := sql.Open("mysql", db)
	if err != nil {
		panic(err.Error)
	}
	Conn = conn
}
