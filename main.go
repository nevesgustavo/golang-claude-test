package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Golang Claude Test!")
	greet("Developer")
	displayCurrentTime()
}

func greet(name string) {
	fmt.Printf("Hello, %s! This is a test project for Claude Code integration.\n", name)
}

func displayCurrentTime() {
	currentTime := time.Now()
	fmt.Printf("Current time: %s\n", currentTime.Format("2006-01-02 15:04:05"))
}

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}
