package main

import(
	"os"
)

func sqlite_create_db() {
	file, err := os.Create("./foo.sqlite")
    check_fatal_error(err)
    file.Close()

	stmt, err := connection.name.Prepare(`
		CREATE TABLE IF NOT EXISTS projects (
		project_id	INTEGER PRIMARY KEY,
		client_name	TEXT,
		type 		TEXT
		);`)
	check_fatal_error(err)
	_, err = stmt.Exec()
	check_fatal_error(err)

	stmt2, err := connection.name.Prepare(`
		CREATE TABLE IF NOT EXISTS project_status (
		scan_id		INTEGER PRIMARY KEY,
		client_name	TEXT,
		scan 		TEXT,
		start 		TEXT,
		stop 		TEXT
		);`)
	stmt2.Exec()
	check_fatal_error(err)

	stmt3, err := connection.name.Prepare(`
		CREATE TABLE IF NOT EXISTS tools (
		tool_id		INTEGER PRIMARY KEY,
		tool_name	TEXT,
		command		TEXT,
		arguments	TEXT
		);`)
	stmt3.Exec()
	check_fatal_error(err)

	stmt4, err := connection.name.Prepare(`
		CREATE TABLE IF NOT EXISTS macros (
		macro_id	INTEGER PRIMARY KEY,
		sequence	INTEGER,
		tool_id		INTEGER,
		FOREIGN KEY(macro_id) REFERENCES macro_names(macro_id),
		FOREIGN KEY(tool_id) REFERENCES tools(tool_id),
		UNIQUE(macro_id,sequence)
		);
		`)
	stmt4.Exec()
	check_fatal_error(err)

	stmt5, err := connection.name.Prepare(`
		CREATE TABLE IF NOT EXISTS macro_names (
		macro_id	INTEGER PRIMARY KEY,
		macro_name	TEXT
		);`)
	stmt5.Exec()
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

	stmt4, err := connection.name.Prepare(`
		CREATE TABLE IF NOT EXISTS macros (
		macro_id	INTEGER PRIMARY KEY,
		sequence	INT,
		tool_id		INT,
		
		CONSTRAINT macro_id
		FOREIGN KEY foreign_key_name (macro_id)
		REFERENCES tools(macro_id)
		ON DELETE CASCADE
		ON UPDATE CASCADE,

		CONSTRAINT tool_id
		FOREIGN KEY foreign_key_name (tool_id)
		REFERENCES tools(tool_id)
		ON DELETE CASCADE
		ON UPDATE CASCADE,
		UNIQUE(macro_id,sequence)
		);`)
	stmt4.Exec()
	check_fatal_error(err)

	stmt5, err := connection.name.Prepare(`
		CREATE TABLE IF NOT EXISTS macro_names (
		macro_id	INT PRIMARY KEY,
		macro_name	VARCHAR(50)
		);`)
	stmt5.Exec()
	check_fatal_error(err)
}