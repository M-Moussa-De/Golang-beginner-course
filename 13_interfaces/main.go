package main

import (
	"fmt"
	"math"
)

// Define interface => Different methods can be implemented on differnt structs
type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

// Value recieve method
func (c Circle) area() float64 {
  return math.Pi * math.Pow(c.radius, 2)
}

func (r Rectangle) area() float64 {
  return r.width * r.height
}

func getArea (s Shape) float64 {
  return s.area()
}

func main() {

	circle := Circle{x: 0, y: 0, radius: 5}
	rectangle := Rectangle{width: 10, height: 5}

	fmt.Printf("Circle area: %f\n", getArea(circle))
	fmt.Printf("Rectangle area: %f\n", getArea(rectangle))

}