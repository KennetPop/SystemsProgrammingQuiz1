package main

import (
	"testing"
)

func TestSquare(t *testing.T) {
	a, p := square(4)
	if a != 16 {
		t.Errorf("Expected 24, got %v", a)
	}
	if p != 20 {
		t.Errorf("Expected 20,got %v", p)
	}
}
