package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"bufio"
)

func main() {
	output, _ := exec_command("cmd.exe","/c dir")
	read_out(output)
}

func exec_command(command string, args string) (io.ReadCloser, io.WriteCloser) {
	fmt.Printf("Command = " + command + " " + args + "\n")
	proc := exec.Command(command,args)
	stdout, err := proc.StdoutPipe()
	if nil != err {
		fmt.Printf("Error creating StdoutPipe\n")
	}
	stdin, err := proc.StdinPipe()
	if nil != err {
		fmt.Printf("Error creating StdinPipe\n")
	}

	read := bufio.NewScanner(stdout)
	go func() {
		for read.Scan() {
			fmt.Printf("Command output: %s\n", read.Text())
		}
	}()

	err = proc.Start()
	if nil != err {
		fmt.Fprintln(os.Stderr, "Error with Start\n", err)
	}

	err = proc.Wait()
	if nil != err {
		fmt.Fprintln(os.Stderr, "Error with Wait\n", err)
	}
	return stdout, stdin
}

func read_out(stdout io.ReadCloser) {
	read := bufio.NewScanner(stdout)
	for read.Scan() {
		fmt.Printf(read.Text())
	}
}