package main

import(
	"bytes"
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func exec_command(command string, args string) (io.ReadCloser, io.WriteCloser) {
	fmt.Printf("Command = " + command + " " + args + "\n")
	proc := exec.Command(command,args)
	
	stdout, err := proc.StdoutPipe()
	check_error(err)

	stdin, err := proc.StdinPipe()
	check_error(err)

	err = proc.Start()
	check_error(err)

	/*
	err = proc.Wait()
	check_error(err)
	*/
	
	return stdout, stdin
}

func read_out(stdout io.ReadCloser) {
	read := bufio.NewScanner(stdout)
	for read.Scan() {
		fmt.Printf("Command output: %s\n", read.Text())
	}
}

func write_out(stdin io.WriteCloser) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
	io.Copy(stdin, bytes.NewBufferString(text))
}