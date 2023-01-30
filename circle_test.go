package main

import (
	"testing"
)

func TestCircleArea(t *testing.T) {
	c1 := circle{
		radius: 5,
	}
	got := c1.area()
	if got != 78.53981633974483 {
		t.Errorf("Expected area of 78.54, got %v", got)
	}
}

func TestCirclePerimeter(t *testing.T) {
	c2 := circle{radius: 5}
	got := c2.perimeter()
	if got != 31.415 {
		t.Errorf("Expected area of 31.415, got %v", got)
	}
}
