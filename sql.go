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

func create_db() {
	stmt, err := connection.name.Prepare("
		CREATE TABLE IF NOT EXISTS projects (
		project_id	INTEGER AUTOINCREMENT PRIMARY KEY,
		client_name	TEXT,
		type 		TEXT
		);")
	stmt.Exec()
	check_fatal_error(err)

	stmt, err := connection.name.Prepare("
		CREATE TABLE IF NOT EXISTS project_status (
		scan_id		INTEGER AUTOINCREMENT PRIMARY KEY,
		client_name	TEXT,
		scan 		TEXT,
		start 		TEXT,
		stop 		TEXT
		);")
	stmt.Exec()
	check_fatal_error(err)

	stmt, err := connection.name.Prepare("
		CREATE TABLE IF NOT EXISTS tools (
		tool_id		INTEGER AUTOINCREMENT PRIMARY KEY,
		tool_name	TEXT PRIMARY KEY,
		command		TEXT,
		arguments	TEXT
		);")
	stmt.Exec()
	check_fatal_error(err)
}

func add_project() {
	stmt, err := connection.name.Prepare("INSERT INTO projects VALUES(?,?,?)")
	check_error(err)
	result, err := stmt.Query()
	check_error(err)
	return result
}

func get_projects_list() {
	stmt, err := connection.name.Prepare("SELECT * FROM projects;")
	check_error(err)
	result, err := stmt.Query()
	check_error(err)
	return result
}

func start_scan() {
	stmt, err := connection.name.Prepare("INSERT INTO tools VALUES(?,?,?,?,?)")
	check_error(err)
	result, err := stmt.Query()
	check_error(err)
	return result
}

func stop_scan() {
	stmt, err := connection.name.Prepare("INSERT INTO tools VALUES(?,?,?,?,?)")
	stmt.Exec()
	check_error(err)
	return result
}

func add_tool(tool_name string, command string, arguments string) {
	stmt, err := connection.name.Prepare("INSERT INTO tools VALUES(?,?,?)")
	check_error(err)
	result, err := stmt.Query(tool_name, command, arguments)
	check_error(err)
	return result
}

func delete_tool(tool_id) {
	stmt, err := connection.name.Prepare("DELETE FROM tools WHERE tool_id LIKE ?;")
	check_error(err)
	result, err := stmt.Query(tool_id)
	check_error(err)
}

func get_tools_list() {
	stmt, err := connection.name.Prepare("SELECT tool_name FROM tools;")
	check_error(err)
	result, err := stmt.Query()
	check_error(err)
	return result
}