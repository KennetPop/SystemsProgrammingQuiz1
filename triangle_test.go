package main

import (
	"testing"
)

func TestArea(t *testing.T) {
	t1 := triangle{
		base:   10,
		height: 20,
	}

	result := t1.area()

	if result != 100 {
		t.Errorf(" expected %v", result)
	}

}

func TestPerimeter(t *testing.T) {
	t2 := triangle{
		base:   10,
		height: 20,
	}
	got := t2.perimeter()
	if got != 200 {
		t.Errorf("Expected 200, got %v ", got)
	}
}
