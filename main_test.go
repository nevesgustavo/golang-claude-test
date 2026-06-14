package main

import (
	"errors"
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	result := add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("add(2, 3) = %d; want %d", result, expected)
	}
}

func TestSubtract(t *testing.T) {
	result := subtract(5, 3)
	expected := 2
	if result != expected {
		t.Errorf("subtract(5, 3) = %d; want %d", result, expected)
	}
}

func TestDivide(t *testing.T) {
	cases := []struct {
		name string
		a, b int
		want int
	}{
		{"positive", 6, 3, 2},
		{"truncation", 7, 2, 3},
		{"negative dividend", -6, 3, -2},
		{"negative divisor", 6, -3, -2},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := divide(tc.a, tc.b)
			if err != nil {
				t.Fatalf("divide(%d, %d) returned unexpected error: %v", tc.a, tc.b, err)
			}
			if got != tc.want {
				t.Errorf("divide(%d, %d) = %d; want %d", tc.a, tc.b, got, tc.want)
			}
		})
	}
}

func TestDivideByZero(t *testing.T) {
	got, err := divide(1, 0)
	if !errors.Is(err, ErrDivideByZero) {
		t.Errorf("divide(1, 0) error = %v; want ErrDivideByZero", err)
	}
	if got != 0 {
		t.Errorf("divide(1, 0) = %d; want 0 on error", got)
	}
}

func TestDivideOverflow(t *testing.T) {
	got, err := divide(math.MinInt, -1)
	if !errors.Is(err, ErrDivideOverflow) {
		t.Errorf("divide(math.MinInt, -1) error = %v; want ErrDivideOverflow", err)
	}
	if got != 0 {
		t.Errorf("divide(math.MinInt, -1) = %d; want 0 on error", got)
	}
}
