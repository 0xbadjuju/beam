package main

import (
	"fmt"
)

func main_menu() {
	for {
		fmt.Printf("\n")
		fmt.Printf("1. Manage Projects\n")
		fmt.Printf("2. Manage Tools\n")
		fmt.Printf("3. Manage Macros\n")
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
			macros()
		case "4":
			setup()
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
		fmt.Printf("1. Resume project\n")
		fmt.Printf("2. List projects\n")
		fmt.Printf("3. Add project\n")
		fmt.Printf("4. Delete project\n")
		fmt.Printf("5. Return\n")
		fmt.Printf("Input: ")
		selection := read_input()
		switch selection {
			case "1":
				resume_project()
			case "2":
				list_projects()
			case "3":
				project_types()
			case "4":
				delete_project()
			case "5":
				return
			default:
				continue
		}
		fmt.Printf("\n")
	}
}

func resume_project() {
	var scan_id int
	var tool, start, stop string
	list_projects()
	fmt.Printf("Project to resume: ")
	project := read_input_int()
	scans := get_scans(project)
	for scans.Next() {
		scans.Scan(&scan_id,&tool,&start,&stop)
		fmt.Printf("%s %s %s %s", scan_id, tool, start, stop)
	}
}

func open_project() {

}

func list_projects() {
	fmt.Printf("\n")
	var (
		project_id int
		client_name string
		project_type string
		)
	tools := get_projects_list()
	for tools.Next() {
		err := tools.Scan(&project_id, &client_name, &project_type)
		check_error(err)
		fmt.Printf("(%d) %s\t%s\n", project_id, client_name, project_type)
	}
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
	fmt.Printf("\n")
	fmt.Printf("Client Name: ")
	client_name := read_input()
	fmt.Printf("Client Name: %s\n", client_name)
	fmt.Printf("Project Type: %s\n", project_type)
	if (confirm()) {
		insert_project(client_name, project_type)
	} else {
		create_project(project_type)
	}
	for {
		fmt.Printf("\n")
		fmt.Printf("Assign tool/macro: to project\n")
		fmt.Printf("1. Assign tool\n")
		fmt.Printf("2. Assign macro\n")
		fmt.Printf("Input: ")
		selection := read_input_int()
		switch selection {
			case 1:
				list_tools()
				tool := read_input_int()
			case 2:
				list_macros()
				macro := read_input_int()
			case 3:
				return
			default:
				continue
		}
	}
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
		fmt.Printf("4. Return\n")
		fmt.Printf("Input: ")
		selection := read_input()
		switch selection {
			case "1":
				list_macros()
			case "2":
				add_macro()
			case "3":
				remove_macro()
			case "4":
				return
			default:
				continue
		}
		fmt.Printf("\n")
	}
}

func list_macros() {
	fmt.Printf("\n")
	var (
		macro_id int
		macro_name string
		)
	macros := get_macros_list()
	for macros.Next() {
		err := macros.Scan(&macro_id, &macro_name)
		check_error(err)
		fmt.Printf("(%d) %s\n", macro_id, macro_name)
	}
	fmt.Printf("\n")
}

func add_macro() {
	fmt.Printf("\n")
	fmt.Printf("Macro name: ")
	macro_name := read_input()
	fmt.Printf("Create Macro %s\n", macro_name)
	if (confirm()) {
		macro_id := create_macro(macro_name)
		edit_macro(macro_name, macro_id)
	} else {
		add_macro()
	}
	fmt.Printf("\n")
}

func edit_macro(macro_name string, macro_id int) {
	var sequence, tool_id int
	fmt.Printf("\n")
	fmt.Printf("Macro %s:\n", macro_name)
	macro := get_macro(macro_id)
	for macro.Next() {
		macro.Scan(&macro_id,&macro_name,&sequence,&tool_id)
	}
	for {
		fmt.Printf("\n")
		fmt.Printf("1. Add Tool to Macro\n")
		fmt.Printf("2. Delete Tool from Macro\n")
		fmt.Printf("3. Return\n")
		fmt.Printf("Input: ")
		selection := read_input()
		switch selection {
			case "1":
				add_tool_to_macro(macro_id)
			case "2":
				remove_tool_from_macro(macro_id)
			case "3":
				return
			default:
				continue
			}
		fmt.Printf("\n")
	}
}

func add_tool_to_macro(macro_id int) {
	fmt.Printf("Tools to add: ")
	list_tools()
	fmt.Printf("Macro Position: ")
	position := read_input_int()
	fmt.Printf("Tool ID: ")
	tool_id := read_input_int()
	if (confirm()) {
		insert_tool_into_macro(macro_id,position,tool_id)
	} else {
		add_tool_to_macro(macro_id)
	}
}

func remove_tool_from_macro(macro_id int) {
	fmt.Printf("Tools to remove: ")
	fmt.Printf("Macro Position: ")
	position := read_input_int()
	if (confirm()) {
		delete_tool_from_macro(macro_id,position)
	} else {
		remove_tool_from_macro(macro_id)
	}
}

func remove_macro() {

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
	