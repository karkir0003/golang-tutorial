package main

import (
	"fmt"
	"math"
)

// Shape interface to define the set of methods the concrete structs must implement
type Shape interface {
	// an entry in the "method set".
	// The function name and return type need to match for Go to determine the interface is
	// "implemented" by the concrete struct
	area() float64
}

type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

type Rectangle struct {
	x, y, height, width float64
}

func (r *Rectangle) area() float64 {
	return r.height * r.width
}

/**
 * Calculate total area of all shapes
 * Here, we have a variadic function that takes in
 * any number of shapes and then call the .area() function on
 * all the Shape objects.
 *
 * Since Circle and Rectangle both implement their area function with the
 * return of float64, they implement the "Shape interface" without us "stating that
 * explicitly"
 */
func totalArea(shapes ...Shape) float64 {
	area := 0.0
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func main() {
	circle := Circle{x: 0, y: 0, r: 5}
	rectangle := Rectangle{x: 5, y: 10, height: 2, width: 10}
	fmt.Println("Total area of the shapes: ", totalArea(&circle, &rectangle))
}
