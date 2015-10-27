package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"
	"regexp"
)

func check_error(err error) {
	if nil != err {
		fmt.Fprintln(os.Stderr, "Error detected\n", err)
	}
}

func check_fatal_error(err error) {
	if nil != err {
		fmt.Fprintln(os.Stderr, "Error detected\n", err)
		fmt.Printf("Bailing out")
		os.Exit(1)
	}
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

func read_input_int(section string) (int) {
	var integer int

	regex, err := regexp.Compile("[0-9]+")
	check_error(err)
	
	fmt.Printf("("+section+") Input: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	check_error(err)
	clean := strings.TrimSpace(input)
	if (0 == len(clean)){
		read_input_int(section)
	}
	if (regex.MatchString(clean)) {
		integer, err := strconv.Atoi(clean)
		check_error(err)
		return integer
	} else {
		fmt.Printf(clean)
		read_input_int(section)
	}
	return integer
}

func confirm()(bool) {
	for {
		fmt.Printf("Continue (Y/N)? ")
		selection := read_input()
		strings.ToUpper(selection)
		if ("Y" == selection) {
			return true
		} else if("N" == selection) {
			return false
		}
	}
}