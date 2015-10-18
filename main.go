package main

import (
	"fmt"
)

func main() {
	fmt.Printf("test")
	output, input := exec_command("nmap.exe", "-p- 192.168.0.1")
	go read_out(output)
	write_out(input)
}

