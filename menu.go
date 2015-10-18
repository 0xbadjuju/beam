package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader := bufio.NewReader(os.Stdin)

func main_menu() {
	for {
		fmt.Printf("1. Manage Projects")
		fmt.Printf("2. Manage Tools")
		fmt.Print("Input: ")
		input, _ := reader.ReadString('\n')
		switch input {
		case 1:
			tools()
		case 2:
			projects()
		}
	}
}

func projects() {
	for {
		fmt.Printf("1. List projects")
		fmt.Printf("2. Add project")
		fmt.Printf("3. Delete project")
		fmt.Printf("Intput: ")
		input, _ := reader.ReadString('\n')
		switch input {
		case 1:
			list_projects()
		case 2:
			project_types()
		case 3:
		}
	}
}

func project_types() {
	for {
		fmt.Printf("1. External Pen")
		fmt.Printf("2. Internal Pen")
		fmt.Printf("3. Web App Pen")
		fmt.Printf("Intput: ")
		input, _ := reader.ReadString('\n')
		switch input {
		case 1:
			add_project(External Pen)
			break
		case 2:
			add_project(Internal Pen)
			break
		case 3:
			add_project(Web App Pen)
			break
		}
	}
}

func add_project(project_type string) {

}

func delete_project() {
	list_projects()
}

func tools() {
	fmt.Printf("1. List tools")
	fmt.Printf("2. Add tool")
	fmt.Printf("2. Delete tool")
	fmt.Printf("Intput: ")
	input, _ := reader.ReadString('\n')
	switch input {
	case 1:
		list_tools()
	case 2:
		add_tool()
	case 3:
		delete_tool()
	}
}

func list_tools() {

}

func add_tool() {
	fmt.Printf("Tool name: ")
	tool_name, _ := reader.ReadString('\n')
	fmt.Printf("Binary file: ")
	command, _ := reader.ReadString('\n')
	fmt.Printf("Arguements/Flags: ")
	arguements, _ := reader.ReadString('\n')
	fmt.Printf("%s: \n",tool_name)
	fmt.Printf("Command: %s %s\n", command, arguements)
	if (confirm()) {
		//insert
	} else {
		add_tool()
	}
}

func delete_tool() {
	list_tools()
	fmt.Printf("Tool ID: ")
	tool_id, _ := reader.ReadString('\n')
	fmt.Printf
	if (confirm()) {
		//delete
	} else {
		delete_tool()
	}
}

func confirm()(bool) {
	for {
		fmt.Printf("Continue (Y/N)?\n")
		cont := reader.ReadString('\n')
		strings.ToUpper(cont)
		if ("Y" == cont) {
			return true
		} else if("N" == cont) {
			return false
		}
	}
}