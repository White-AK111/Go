package geometry

import (
	"errors"
	"fmt"
	"math"
	"os"
)

// Text for custom error for invalid input parameters
const ErrInputText = "incorrect parameter entered"

// Custom error for invalid input parameters
var ErrInput = errors.New(ErrInputText)

type Circle struct {
	radius float64
}

// Calc area of circle
func (c *Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

// Calc diameter of circle, using area func
func (c *Circle) Diameter() float64 {
	return 2 * math.Sqrt(c.Area()/math.Pi)
}

// Calc length of circle, using diameter func
func (c *Circle) Length() float64 {
	return math.Pi * c.Diameter()
}

// Func for formatted print
func (c *Circle) String() string {
	return fmt.Sprintf("Circle {Radius: %.2f}", c.radius)
}

// Create new circle entity
func NewCircle() Circle {
	var rad float64

	fmt.Print("Enter radius of circle: ")
	_, err := fmt.Scan(&rad)
	if err != nil {
		fmt.Printf("Error: %v\n\n", ErrInput)
		os.Exit(1)
	}

	return Circle{rad}
}
