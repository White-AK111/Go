package main

import (
	"fmt"

	"gitlab.com/White-AK111/gb-go-level1/lesson2/geometry"
)

// Interface for calc area of shape
type Sizer interface {
	Area() float64
}

// Interface for shapes
type Shaper interface {
	Sizer
	fmt.Stringer
}

// Compare area of shapes
func compareShapesAreas(c geometry.Circle, r geometry.Rectangle) {
	printArea(&c)
	printArea(&r)
	l := less(&c, &r)
	fmt.Printf("%+v is the smallest\n\n", l)
}

// Get smallest area of shapes
func less(s1, s2 Sizer) Sizer {
	if s1.Area() < s2.Area() {
		return s1
	}
	return s2
}

// Func for formatted print by fmt.Stringer
func printArea(s Shaper) {
	fmt.Printf("area of %s is %.2f\n", s.String(), s.Area())
}
