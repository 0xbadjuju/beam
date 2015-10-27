package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

func mysql_open_db() {
	open, err := sql.Open("mysql", "")
	check_fatal_error(err)
	err = open.Ping()
	check_fatal_error(err)
	set_connection(open)
}