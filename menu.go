package main

import (
	"bufio"
	"fmt"
	"os"
)

func main_menu() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("1. Manage Projects")
	fmt.Printf("2. Manage Tools")
	fmt.Print("Input: ")
	text, _ := reader.ReadString('\n')
	if ("1" == text) {
		projects()
	}
}

func projects() {
	fmt.Printf("1. List projects")
	fmt.Printf("2. Add project")
	fmt.Printf("3. Delete project")
}

func project_types() {
	fmt.Printf("1. External Pen")
	fmt.Printf("2. Internal Pen")
	fmt.Printf("3. Web App Pen")
}

func tools() {
	fmt.Printf("1. List tools")
	fmt.Printf("2. Add tool")
	fmt.Printf("2. Delete tool")
}