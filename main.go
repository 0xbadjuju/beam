package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Beam, because sometimes a 2x4 isn't enough\n\n")
	sqlite()
	main_menu()
	close_db()
}

func mysql() {
	mysql_open_db()
	mysql_create_db()
}

func sqlite() {
	sqlite_open_db()
	sqlite_create_db()
}