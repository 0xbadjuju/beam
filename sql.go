package main

import (
	"database/sql"
)

type db struct{
	name *sql.DB
}

var connection db

func open_db() {
	open, err := sql.Open("sqlite3", "./foo.db")
	set_connection(open)
	check_fatal_error(err)
}

func set_connection(instance *sql.DB) {
    connection.name = instance
}

func get_connection() *sql.DB {
    return connection.name
}

func add_project() {
	stmt, err := connection.name.Prepare("INSERT INTO projects VALUES(?,?,?,?,?)")
	stmt.Exec()
	check_error(err)
}

func get_projects_list() {
	stmt, err := connection.name.Prepare("SELECT * FROM projects;")
	stmt.Exec()
	check_error(err)
}

func add_tool() {
	stmt, err := connection.name.Prepare("INSERT INTO tools VALUES(?,?,?,?,?)")
	stmt.Exec()
	check_error(err)
}

func get_tools_list() {
	stmt, err := connection.name.Prepare("SELECT tool_name FROM tools;")
	stmt.Exec()
	check_error(err)
}