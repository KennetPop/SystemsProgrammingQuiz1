package main

import (
	"fmt"
	"math"
)

type triangle struct {
	base   float64
	height float64
}

func (t triangle) area() float64 {
	return t.base * t.height / 2
}

func (t triangle) perimeter() float64 {
	p := math.Sqrt(math.Pow(t.base, 2) + math.Pow(t.height, 2))
	return p + t.base + t.height
}
func main() {

	t := triangle{
		base:   10,
		height: 20,
	}
	fmt.Println("Area of triangle is: ", t.area())
	fmt.Println("Perimeter of triangle is: ", t.perimeter())

	c := circle{
		radius: 5,
	}
	fmt.Println("Area of circle is: ", c.area())
	fmt.Println("Perimeter of circle is: ", c.perimeter())
	
	area, perimeter := square(5)
	fmt.Println("Area of square is: ", area())
	fmt.Println("Perimeter of square is: ",perimeter())
	
}
