package main

import (
	_ "github.com/mxk/go-sqlite/sqlite3"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"os"
)

type db struct{
	name *sql.DB
}

var connection db

func mysql_open_db() {
	fmt.Printf("Opening SQL Connection\n")
	open, err := sql.Open("mysql", "")
	check_fatal_error(err)
	err = open.Ping()
	check_fatal_error(err)
	set_connection(open)
}

func sqlite_open_db() {
	fmt.Printf("Opening SQL Connection\n")
	open, err := sql.Open("sqlite3", "")
	check_fatal_error(err)
	err = open.Ping()
	check_fatal_error(err)
	set_connection(open)
}

func close_db() {
	get_connection().Close()
}

func set_connection(instance *sql.DB) {
    connection.name = instance
}

func get_connection() *sql.DB {
    return connection.name
}

func sqlite_create_db() {
	file, err := os.Create("./foo.db")
    check_fatal_error(err)
    file.Close()

	stmt, err := connection.name.Prepare(`
		CREATE TABLE IF NOT EXISTS projects (
		project_id	INTEGER AUTOINCREMENT PRIMARY KEY,
		client_name	TEXT,
		type 		TEXT
		);`)
	stmt.Exec()
	check_fatal_error(err)

	stmt2, err := connection.name.Prepare(`
		CREATE TABLE IF NOT EXISTS project_status (
		scan_id		INTEGER AUTOINCREMENT PRIMARY KEY,
		client_name	TEXT,
		scan 		TEXT,
		start 		TEXT,
		stop 		TEXT
		);`)
	stmt2.Exec()
	check_fatal_error(err)

	stmt3, err := connection.name.Prepare(`
		CREATE TABLE IF NOT EXISTS tools (
		tool_id		INTEGER AUTOINCREMENT PRIMARY KEY,
		tool_name	TEXT,
		command		TEXT,
		arguments	TEXT
		);`)
	stmt3.Exec()
	check_fatal_error(err)
}

func mysql_create_db() {
	file, err := os.Create("./foo.db")
    check_fatal_error(err)
    file.Close()

	stmt, err := connection.name.Prepare(`
		CREATE TABLE IF NOT EXISTS projects (
		project_id	INT AUTOINCREMENT PRIMARY KEY,
		client_name	VARCHAR(50),
		type 		VARCHAR(50)
		);`)
	stmt.Exec()
	check_fatal_error(err)

	stmt2, err := connection.name.Prepare(`
		CREATE TABLE IF NOT EXISTS project_status (
		scan_id		INT AUTOINCREMENT PRIMARY KEY,
		client_name	VARCHAR(50),
		scan 		VARCHAR(50),
		start 		VARCHAR(50),
		stop 		VARCHAR(50)
		);`)
	stmt2.Exec()
	check_fatal_error(err)

	stmt3, err := connection.name.Prepare(`
		CREATE TABLE IF NOT EXISTS tools (
		tool_id		INT AUTOINCREMENT PRIMARY KEY,
		tool_name	VARCHAR(50),
		command		VARCHAR(50),
		arguments	VARCHAR(50)
		);`)
	stmt3.Exec()
	check_fatal_error(err)
}

func insert_project() {
	stmt, err := connection.name.Prepare("INSERT INTO projects VALUES(?,?,?)")
	check_error(err)
	result, err := stmt.Exec()
	check_error(err)
	check_result(result)
}

func get_projects_list()(*sql.Rows) {
	stmt, err := connection.name.Prepare("SELECT * FROM projects;")
	check_error(err)
	result, err := stmt.Query()
	check_error(err)
	return result
}

func start_scan() {
	stmt, err := connection.name.Prepare("INSERT INTO tools VALUES(?,?,?,?,?)")
	check_error(err)
	result, err := stmt.Exec()
	check_error(err)
	check_result(result)
}

func stop_scan() {
	stmt, err := connection.name.Prepare("INSERT INTO tools VALUES(?,?,?,?,?)")
	check_error(err)
	stmt.Exec()
	check_error(err)
}

func insert_tool(tool_name string, command string, arguments string) {
	stmt, err := connection.name.Prepare("INSERT INTO tools VALUES(?,?,?)")
	check_error(err)
	result, err := stmt.Exec(tool_name, command, arguments)
	check_error(err)
	check_result(result)
}

func select_tool(tool_id string)(*sql.Rows) {
	stmt, err := connection.name.Prepare("SELECT * FROM tools WHERE tool_id LIKE ?")
	check_error(err)
	result, err := stmt.Query(tool_id)
	check_error(err)
	return result
}

func delete_tool(tool_id string) {
	stmt, err := connection.name.Prepare("DELETE FROM tools WHERE tool_id LIKE ?;")
	check_error(err)
	result, err := stmt.Exec(tool_id)
	check_error(err)
	check_result(result)
}

func get_tools_list()(*sql.Rows) {
	stmt, err := connection.name.Prepare("SELECT tool_name FROM tools;")
	check_error(err)
	result, err := stmt.Query()
	check_error(err)
	return result
}