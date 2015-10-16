package main

import "fmt"

func main() {
	fmt.Printf("Test\n")
}

func exec_command(cmd,args) {
	proc := exec.Command(cmd,args)
	stdout, stderr := proc.StdoutPipe()
	proc.Start()
	return stdout
}