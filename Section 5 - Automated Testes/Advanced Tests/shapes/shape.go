package shapes

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

func WriteArea(s Shape) {
	fmt.Printf("The area of the shape is %0.2f\n", s.area())
}

type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) area() float64 {
	return r.Height * r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func main() {
	r := Rectangle{10, 15}
	WriteArea(r)

	c := Circle{10}
	WriteArea(c)
}
