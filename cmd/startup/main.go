package main

import (
	fmt"
	os/exec"
)

func main() {
	cmd := exec.Command("tmux", "bind-key", "T", "run-shell", "echo 'hello'")
	cmd.Start()
	fmt.Println("hello world")
}
