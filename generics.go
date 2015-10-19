package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"
)

func check_error(err error) {
	if nil != err {
		fmt.Fprintln(os.Stderr, "Error detected\n", err)
	}
}

func check_fatal_error(err error) {
	if nil != err {
		fmt.Fprintln(os.Stderr, "Error detected\n", err)
	}
	fmt.Printf("Bailing out")
	os.Exit(1)
}

func check_result(result sql.Result){
	rows, err := result.RowsAffected()
	check_error(err)
	if(rows == 0) {
		fmt.Printf("No rows affected by query")
	}
}

func read_input() (string) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	check_error(err)
	return strings.TrimSpace(input)
}