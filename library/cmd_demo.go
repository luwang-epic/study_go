package main

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
)

type IOWriter struct {
	buffer *string
}

func (IO IOWriter) Write(p []byte) (int, error) {
	*IO.buffer = *IO.buffer + string(p) + "\n"
	return len(p), nil
}

func CmdDemo() {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}

	args := []string{"-c", "ls"}
	command := exec.CommandContext(context.Background(), "powershell", args...)
	command.Dir = "/Users/wanglu51"
	command.Stdout = stdout
	command.Stderr = stderr
	command.Stdin = bytes.NewBuffer([]byte("aaaa"))

	err := command.Run()
	fmt.Printf("err: %v\n", err)
	fmt.Printf("stdout: %s\n", stdout)
	fmt.Printf("stderr: %s\n", stderr)
}
