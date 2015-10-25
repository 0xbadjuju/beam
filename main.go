package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var config = map[string]string{}

func main() {
	fmt.Printf("Beam, because sometimes a 2x4 isn't enough\n\n")
	read_config()
	if ("mysql" == config["DATABASE_TYPE"]) {
		mysql()
	} else if("sqlite" == config["DATABASE_TYPE"]) {
		sqlite()
	} else {
		fmt.Printf("%s\n", config["DATABASE_TYPE"])
		fmt.Printf("Bailing out\n")
		os.Exit(1)
	}
	main_menu()
	close_db()
}

func read_config() {
	conf, err := os.Open("beam.conf")
	check_fatal_error(err)
	defer conf.Close()
	scanner := bufio.NewScanner(conf)
	for scanner.Scan() {
		line := scanner.Text()
		if (3 < len(line)){
			split_string := strings.Split(line, "=")
			if (2 == len(split_string)) {
				clean_part_0 := strings.TrimSpace(split_string[0])
				clean_part_1 := strings.TrimSpace(split_string[1])
				config[clean_part_0] = clean_part_1
			}
		}
	}
}

func mysql() {
	mysql_open_db()
}

func sqlite() {
	sqlite_open_db()
}