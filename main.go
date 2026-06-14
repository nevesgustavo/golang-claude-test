package main

import (
	"errors"
	"fmt"
	"math"
	"time"
)

// ErrDivideByZero is returned by divide when the divisor is zero.
var ErrDivideByZero = errors.New("division by zero is undefined")

// ErrDivideOverflow is returned by divide when the division would overflow,
// which happens only for math.MinInt / -1.
var ErrDivideOverflow = errors.New("division overflow")

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

func multiply(a, b int) int {
	return a * b
}

// divide returns a / b. It returns ErrDivideByZero if b is zero, and
// ErrDivideOverflow for the math.MinInt / -1 case, which would otherwise
// silently wrap to an incorrect value.
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	if a == math.MinInt && b == -1 {
		return 0, ErrDivideOverflow
	}
	return a / b, nil
}
