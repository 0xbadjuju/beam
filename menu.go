package main

import (
	"fmt"
	"regexp"
	"sync"
)

func main_menu() {
	for {
		fmt.Printf("\n")
		fmt.Printf("1.\tManage Projects\n")
		fmt.Printf("2.\tManage Tools\n")
		fmt.Printf("3.\tManage Macros\n")
		fmt.Printf("4.\tSetup\n")
		fmt.Printf("99.\tExit\n")
		selection := read_input_int("Home")
		switch selection {
		case 1:
			projects()
		case 2:
			tools()
		case 3:
			macros()
		case 4:
			setup()
		case 99:
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
		fmt.Printf("1.\tResume project\n")
		fmt.Printf("2.\tList projects\n")
		fmt.Printf("3.\tAdd project\n")
		fmt.Printf("4.\tDelete project\n")
		fmt.Printf("99.\tReturn\n")
		selection := read_input_int("Projects")
		switch selection {
			case 1:
				resume_project()
			case 2:
				list_projects()
			case 3:
				project_types()
			case 4:
				delete_project()
			case 99:
				return
			default:
				continue
		}
		fmt.Printf("\n")
	}
}

func resume_project() {
	fmt.Printf("\n")
	list_projects()
	project_id := read_input_int("Projects\\Resume")
	resume_scanning(project_id)
	fmt.Printf("\n")
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
	fmt.Printf("1.\tExternal Pen\n")
	fmt.Printf("2.\tInternal Pen\n")
	fmt.Printf("3.\tWeb App Pen\n")
	fmt.Printf("99.\tReturn \n")
	selection := read_input_int("Projects\\Types")
	switch selection {
		case 1:
			create_project(1)
			break
		case 2:
			create_project(2)
			break
		case 3:
			create_project(3)
			break
		case 99:
			return
		default:
			project_types()
	}
	fmt.Printf("\n")
}

func create_project(project_type int) {
	var project_id int
	fmt.Printf("\n")
	fmt.Printf("Client Name: ")
	client_name := read_input()
	if (confirm()) {
		project_id = insert_project(client_name, project_type)
	} else {
		create_project(project_type)
	}
	add_addresss(project_id)
	fmt.Printf("\n")
}

func add_addresss(project_id int) {
	address_regex, _ := regexp.Compile("[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}")
	fmt.Printf("\n")
	for {
		fmt.Printf("Address: ")
		address := read_input()
		if(address_regex.MatchString(address)) {
			assign_tool(project_id, address)
			if (!confirm()) {
				return
			}
		}
	}
	fmt.Printf("\n")
}

func assign_tool(project_id int, address string) {
	fmt.Printf("\n")
	for {
		fmt.Printf("\n")
		fmt.Printf("Assign tool/macro\n")
		fmt.Printf("1.\tAssign tool\n")
		fmt.Printf("2.\tAssign macro\n")
		fmt.Printf("99.\tReturn\n")
		selection := read_input_int("Projects\\Assign")
		switch selection {
			case 1:
				list_tools()
				tool_id := read_input_int("Projects\\Assign\\Tools")
				if (confirm()) {
					insert_scan(project_id, tool_id, address)
				}
			case 2:
				var tool_id int
				list_macros()
				macro := read_input_int("Projects\\Assign\\Macros")
				rows := select_macro(macro)
				for rows.Next() {
					rows.Scan(&tool_id)
				}
			case 99:
				return
			default:
				continue
		}
	}
	fmt.Printf("\n")
}

func delete_project() {
	fmt.Printf("\n")
	list_projects()
	fmt.Printf("\n")
}

func resume_scanning(project_id int) {
	var tool_name, command, arguments, address, start, stop string
	
	result := completed_scans(project_id)
	fmt.Printf("Completed Scans: \n")
	for result.Next() {
		result.Scan(&tool_name, &command, &arguments, &start, &stop, &address)
		fmt.Printf("%s %s %s %s %s\n", tool_name, command, arguments, start, stop)
	}

	result2 := queued_scans(project_id)
	fmt.Printf("Scans to be run: \n")
	for result2.Next() {
		result2.Scan(&tool_name, &command, &arguments, &address)
		fmt.Printf("%s %s %s\n", tool_name, command, arguments)
	}
	if (confirm()) {
		go scan_thread(project_id)
	}
}

func scan_thread(project_id int) {
	var wg sync.WaitGroup
	var tool_name, command, arguments, address string

	wg.Add(1)
	result3 := queued_scans(project_id)
	for result3.Next() {
		result3.Scan(&tool_name, &command, &arguments, &address)
		regex, err := regexp.Compile("$target$")
		check_error(err)
		replaced := regex.ReplaceAllString(arguments, address)
		//start_scan(client_id int, tool_id int)
		stdin, stdout := exec_command(command, replaced)
		add_pipe(project_id, stdin, stdout)
		go read_out(stdout)
		wg.Wait()
		//finish_scan(client_id int, tool_id int)
	}
}

func tools() {
	for {
		fmt.Printf("\n")
		fmt.Printf("1.\tList tools\n")
		fmt.Printf("2.\tAdd tool\n")
		fmt.Printf("3.\tDelete tool\n")
		fmt.Printf("99.\tReturn \n")
		selection := read_input_int("Tools")
		switch selection {
			case 1:
				list_tools()
			case 2:
				add_tool()
			case 3:
				remove_tool()
			case 99:
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
		fmt.Printf("1.\tList Macros\n")
		fmt.Printf("2.\tAdd Macro\n")
		fmt.Printf("3.\tDelete Macro\n")
		fmt.Printf("99.\tReturn\n")
		selection := read_input_int("Macros")
		switch selection {
			case 1:
				list_macros()
			case 2:
				add_macro()
			case 3:
				remove_macro()
			case 99:
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
		fmt.Printf("1.\tAdd Tool to Macro\n")
		fmt.Printf("2.\tDelete Tool from Macro\n")
		fmt.Printf("99.\tReturn\n")
		selection := read_input_int("Macros\\Tools")
		switch selection {
			case 1:
				add_tool_to_macro(macro_id)
			case 2:
				remove_tool_from_macro(macro_id)
			case 99:
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
	position := read_input_int("Macros\\Add")
	fmt.Printf("Tool ID: ")
	tool_id := read_input_int("Macros\\Add")
	if (confirm()) {
		insert_tool_into_macro(macro_id,position,tool_id)
	} else {
		add_tool_to_macro(macro_id)
	}
}

func remove_tool_from_macro(macro_id int) {
	fmt.Printf("Tools to remove: ")
	fmt.Printf("Macro Position: ")
	position := read_input_int("Macros\\Remove")
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
	fmt.Printf("1.\tSQLite\n")
	fmt.Printf("2.\tMySQL\n")
	fmt.Printf("99.\tReturn\n")
	selection := read_input_int("Setup\\SQL")
	switch selection {
		case 1:
			sqlite_create_db()
		case 2:
			mysql_create_db()
		case 99:
			return
		default:
			setup()
	}
	fmt.Printf("\n")
}
	