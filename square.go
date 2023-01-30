package main

import (
	"math"
)

func square(s float64) (float64, float64) {
	area := math.Pow(s, 2)
	perimeter := 4 * s
	return area, perimeter
}
