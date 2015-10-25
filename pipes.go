package main 

import (
	"io"
)

type pipe struct {
	stdin	io.WriteCloser
	stdout	io.ReadCloser
}

var pipes = map[int]pipe{}

func add_pipe(client_id int, stdin io.WriteCloser, stdout io.ReadCloser) {
	var tmp pipe
	tmp.stdin = stdin
	tmp.stdout = stdout
	pipes[client_id] = tmp
}

func get_stdin(client_id int)(io.WriteCloser) {
	return pipes[client_id].stdin
}

func get_stdout(client_id int)(io.ReadCloser) {
	return pipes[client_id].stdout
}