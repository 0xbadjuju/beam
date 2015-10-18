package main

import (
	"fmt"
	"os"
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