package main

import (
	_ "github.com/mxk/go-sqlite/sqlite3"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)

type db struct{
	name *sql.DB
}

var connection db

func mysql_open_db() {
	open, err := sql.Open("mysql", "")
	check_fatal_error(err)
	err = open.Ping()
	check_fatal_error(err)
	set_connection(open)
}

func sqlite_open_db() {
	open, err := sql.Open("sqlite3", "./foo.sqlite")
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

func insert_project(client_name string, project_type int)(int) {
	var project_id int
	stmt, err := connection.name.Prepare("INSERT INTO projects VALUES(?,?,?)")
	check_error(err)
	result, err := stmt.Exec(nil, client_name, project_type)
	check_error(err)
	check_result(result)
	stmt2, err := connection.name.Prepare(`
		SELECT project_id 
		FROM projects 
		WHERE client_name LIKE ?
		AND project_type LIKE ?
	`)
	check_error(err)
	result2, err := stmt2.Query(client_name, project_type)
	check_error(err)
	for result2.Next() {
		result2.Scan(&project_id)
	}
	return project_id
}

func get_projects_list()(*sql.Rows) {
	stmt, err := connection.name.Prepare("SELECT * FROM projects;")
	check_error(err)
	result, err := stmt.Query()
	check_error(err)
	return result
}

func insert_tool(tool_name string, command string, arguments string) {
	stmt, err := connection.name.Prepare("INSERT INTO tools VALUES(?,?,?,?)")
	check_error(err)
	result, err := stmt.Exec(nil, tool_name, command, arguments)
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
	stmt, err := connection.name.Prepare("SELECT tool_id,tool_name FROM tools;")
	check_error(err)
	result, err := stmt.Query()
	check_error(err)
	return result
}

func create_macro(macro_name string)(int){
	stmt, err := connection.name.Prepare("INSERT INTO macro_names VALUES(?,?);")
	check_error(err)
	result, err := stmt.Exec(nil, macro_name)
	check_error(err)
	check_result(result)
	return get_macro_id(macro_name)
}

func select_macro(macro_id int)(*sql.Rows) {
	stmt, err := connection.name.Prepare("SELECT * FROM macros WHERE macro_id LIKE ?")
	check_error(err)
	result, err := stmt.Query(macro_id)
	check_error(err)
	return result
}

func delete_macro(tool_id string) {
	stmt, err := connection.name.Prepare("DELETE FROM macros WHERE macro_id LIKE ?;")
	check_error(err)
	result, err := stmt.Exec(tool_id)
	check_error(err)
	check_result(result)
}

func get_macro(macro_id int)(*sql.Rows) {
	stmt, err := connection.name.Prepare(`
		SELECT * FROM macros
		INNER JOIN tools
		ON macros.tool_id = tools.tool_id
		WHERE macro_id LIKE ? 
		;`)
	check_error(err)
	result, err := stmt.Query(macro_id)
	check_error(err)
	return result
}

func get_macros_list()(*sql.Rows) {
	stmt, err := connection.name.Prepare("SELECT macro_id,macro_name FROM macro_names;")
	check_error(err)
	result, err := stmt.Query()
	check_error(err)
	return result
}

func insert_tool_into_macro(macro_id int, sequence int, tool_id int)(*sql.Rows) {
	stmt, err := connection.name.Prepare(`
		INSERT OR REPLACE INTO macros VALUES(?,?,?);`)
	check_error(err)
	result, err := stmt.Query(macro_id, sequence, tool_id)
	check_error(err)
	return result
}

func delete_tool_from_macro(macro_id int, sequence int)(*sql.Rows) {
	stmt, err := connection.name.Prepare(`
		DELETE FROM macros 
		WHERE macro_id LIKE ? 
		AND sequence LIKE ?
	;`)
	check_error(err)
	result, err := stmt.Query(macro_id, sequence)
	check_error(err)
	return result
}

func get_macro_id(macro_name string)(int) {
	var macro_id = 0
	stmt, err := connection.name.Prepare(`
		SELECT macro_id 
		FROM macro_names 
		WHERE macro_name LIKE ?
	;`)
	check_error(err)
	result, err := stmt.Query(macro_name)
	check_error(err)

	for result.Next() {
		check_error(err)
		result.Scan(&macro_id)
		check_error(err)
	}
	return macro_id
}

func insert_scan(project_id int, tool_id int) {
	stmt, err := connection.name.Prepare(`
		INSERT INTO project_status 
		VALUES(NULL,?,?,NULL,NULL);
	`)
	check_error(err)
	result, err := stmt.Exec(tool_id,project_id)
	check_error(err)
	check_result(result)
}

func start_scan(client_id int, tool_id int) {
	stmt, err := connection.name.Prepare(`
		UPDATE project_status 
		SET start = datetime('now'),
		stop = NULL
		WHERE project_id = ?
		AND tool_id = ?;
	`)
	check_error(err)
	result, err := stmt.Exec(client_id, tool_id)
	check_error(err)
	check_result(result)
}

func finish_scan(client_id int, tool_id int) {
	stmt, err := connection.name.Prepare(`
		UPDATE project_status 
		SET stop = datetime('now')
		WHERE project_id = ?
		AND tool_id = ?;
	`)
	check_error(err)
	result, err := stmt.Exec(client_id, tool_id)
	check_error(err)
	check_result(result)
}

func get_scans(client_id int)(*sql.Rows) {
	stmt, err := connection.name.Prepare(`
		SELECT tool_id, start, stop 
		FROM project_status
		WHERE project_id LIKE ?;
	`)
	check_error(err)
	result, err := stmt.Query(client_id)
	check_error(err)
	return result
}

func resume_scanning(project_id int) {
	var tool_name, command, arguments, start, stop string
	stmt, err := connection.name.Prepare(`
		SELECT tool_name, command, arguments 
		FROM project_status
		INNER JOIN tools
		ON project_status.tool_id = tools.tool_id
		WHERE project_id LIKE ?
		AND stop IS NOT NULL
		ORDER BY scan_id ASC;
	`)
	check_error(err)
	result, err := stmt.Query(project_id)
	check_error(err)
	fmt.Printf("Completed Scans: \n")
	for result.Next() {
		result.Scan(&tool_name, &command, &arguments, &start, &stop)
		fmt.Printf("%s %s %s %s %s\n", tool_name, command, arguments, start, stop)
	}

	stmt2, err := connection.name.Prepare(`
		SELECT tool_name, command, arguments 
		FROM project_status
		INNER JOIN tools
		ON project_status.tool_id = tools.tool_id
		WHERE project_id LIKE ?
		AND stop IS NULL
		ORDER BY scan_id ASC;
	`)
	check_error(err)
	result2, err := stmt2.Query(project_id)
	check_error(err)
	fmt.Printf("Scans to be run: \n")
	for result2.Next() {
		result2.Scan(&tool_name, &command, &arguments)
		fmt.Printf("%s %s %s\n", tool_name, command, arguments)
	}

	if (confirm()) {
		result3, err := stmt2.Query(project_id)
		check_error(err)
		for result3.Next() {
			result3.Scan(&tool_name, &command, &arguments)
			stdin, stdout := exec_command(command, arguments)
			add_pipe(project_id, stdin, stdout)
			go read_out(stdout)
		}
	}
}