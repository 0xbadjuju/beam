package main

import (
	"fmt"
	"os"
	"strings"
)

func main_menu() {
	for {
		fmt.Printf("1. Manage Projects\n")
		fmt.Printf("2. Manage Tools\n")
		fmt.Printf("3. Exit\n")
		fmt.Printf("Input: ")
		selection := read_input()
		switch selection {
		case "1":
			projects()
		case "2":
			tools()
		case "3":
			os.Exit(0)
		default:
			continue
		}
	}
}

func projects() {
	for {
		fmt.Printf("1. List projects\n")
		fmt.Printf("2. Add project\n")
		fmt.Printf("3. Delete project\n")
		fmt.Printf("4. Return\n")
		fmt.Printf("Input: ")
		selection := read_input()
		switch selection {
			case "1":
				list_projects()
			case "2":
				project_types()
			case "3":
				delete_project()
			case "4":
				return
			default:
				continue
			}
	}
}

func list_projects() {

}

func project_types() {
	for {
		fmt.Printf("1. External Pen\n")
		fmt.Printf("2. Internal Pen\n")
		fmt.Printf("3. Web App Pen\n")
		fmt.Printf("4. Return \n")
		fmt.Printf("Input: ")
		selection := read_input()
		switch selection {
			case "1":
				create_project("External Pen")
				break
			case "2":
				create_project("Internal Pen")
				break
			case "3":
				create_project("Web App Pen")
				break
			case "4":
				return
			default:
				continue
		}
	}
}

func create_project(project_type string) {

}

func delete_project() {
	list_projects()
}

func tools() {
	for {
		fmt.Printf("1. List tools\n")
		fmt.Printf("2. Add tool\n")
		fmt.Printf("3. Delete tool\n")
		fmt.Printf("4. Return \n")
		fmt.Printf("Input: ")
		selection := read_input()
		switch selection {
		case "1":
			list_tools()
		case "2":
			add_tool()
		case "3":
			remove_tool()
		case "4":
			return
		default:
			continue
		}
	}
}

func list_tools() {

}

func add_tool() {
	fmt.Printf("Tool name: ")
	tool_name := read_input()
	fmt.Printf("Binary file: ")
	command := read_input()
	fmt.Printf("Arguements/Flags: ")
	arguements := read_input()
	fmt.Printf("%s: \n",tool_name)
	fmt.Printf("Command: %s %s\n", command, arguements)
	if (confirm()) {
		//insert
	} else {
		add_tool()
	}
}

func remove_tool() {
	list_tools()
	fmt.Printf("Tool ID: ")
	tool_id := read_input()
	tool_row := select_tool(tool_id)
	defer tool_row.Close()
	fmt.Printf("Removing %s",tool_row)
	if (confirm()) {
		delete_tool(tool_id)
	} else {
		remove_tool()
	}
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