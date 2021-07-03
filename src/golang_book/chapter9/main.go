package main

import (
	"fmt"
	"math"
)

//Структуры и интерфейсы

type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (r *Rectangle) perimeter() float64 {
	p := 2 * ((r.x2 - r.x1) + (r.y2 - r.y1))
	return p
}

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
	Person
	Model string
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func circleArea(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

type Shape interface {
	area() float64
}

type MultiShape struct {
	shapes []Shape
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

func main() {
	c := Circle{0, 0, 5}
	fmt.Println(circleArea(&c))

	fmt.Println(c.area())
	fmt.Println(c.perimeter())

	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())
	fmt.Println(r.perimeter())

	a := new(Android)
	a.Name = "Vasya"
	a.Talk()

	fmt.Println(totalArea(&c, &r))
}
