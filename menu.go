package main

import (
	"fmt"
)

func main_menu() {
	for {
		fmt.Printf("\n")
		fmt.Printf("1. Manage Projects\n")
		fmt.Printf("2. Manage Tools\n")
		fmt.Printf("3. Manage Macros")
		fmt.Printf("4. Setup\n")
		fmt.Printf("5. Exit\n")
		fmt.Printf("Input: ")
		selection := read_input()
		switch selection {
		case "1":
			projects()
		case "2":
			tools()
		case "3":
			setup()
		case "4":
			macros()
		case "5":
			return
		default:
			main_menu()
		}
		fmt.Printf("\n")
	}
}

func projects() {
	for {
		fmt.Printf("\n")
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
		fmt.Printf("\n")
	}
}

func list_projects() {
	fmt.Printf("\n")
	fmt.Printf("\n")
}

func project_types() {
	fmt.Printf("\n")
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
			project_types()
	}
	fmt.Printf("\n")
}

func create_project(project_type string) {

}

func delete_project() {
	fmt.Printf("\n")
	list_projects()
	fmt.Printf("\n")
}

func tools() {
	for {
		fmt.Printf("\n")
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
		fmt.Printf("\n")
	}
}

func list_tools() {
	fmt.Printf("\n")
	var (
		tool_id int
		tool_name string
		)
	tools := get_tools_list()
	for tools.Next() {
		err := tools.Scan(&tool_id, &tool_name)
		check_error(err)
		fmt.Printf("(%d) %s\n", tool_id, tool_name)
	}
	fmt.Printf("\n")
}

func add_tool() {
	fmt.Printf("\n")
	fmt.Printf("Tool name: ")
	tool_name := read_input()
	fmt.Printf("Binary file: ")
	command := read_input()
	fmt.Printf("Arguements/Flags: ")
	arguments := read_input()
	fmt.Printf("%s: \n",tool_name)
	fmt.Printf("Command: %s %s\n", command, arguments)
	if (confirm()) {
		insert_tool(tool_name, command, arguments)
	} else {
		add_tool()
	}
	fmt.Printf("\n")
}

func remove_tool() {
	fmt.Printf("\n")
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
	fmt.Printf("\n")
}

func macros() {
	for {
		fmt.Printf("\n")
		fmt.Printf("1. List Macros\n")
		fmt.Printf("2. Add Macro\n")
		fmt.Printf("3. Delete Macro\n")
		fmt.Printf("Input: ")
		selection := read_input()
		switch selection {
			case "1":
				list_macro()
			case "2":
				add_macro()
			case "3":
				delete_macro()
			case "4":
				return
			default:
				continue
		}
		fmt.Printf("\n")
	}
}

func list_macro() {
	
}

func add_macro() {
	
}

func delete_macro() {
	
}

func setup() {
	fmt.Printf("\n")
	fmt.Printf("1. SQLite\n")
	fmt.Printf("2. MySQL\n")
	fmt.Printf("3. Return\n")
	fmt.Printf("Input: ")
	selection := read_input()
	switch selection {
		case "1":
			sqlite_create_db()
		case "2":
			mysql_create_db()
		case "3":
			return
		default:
			setup()
	}
	fmt.Printf("\n")
}
	