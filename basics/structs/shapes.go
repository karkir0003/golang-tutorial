package main

import (
	"fmt"
	"math"
)

// Circle Struct
type Circle struct {
	x, y, r float64
}

// Rectable Struct
type Rectangle struct {
	x, y, width, height float64
}

/*
*
In this version, we are passing a COPY of the whole Circle object to the function and performing
computation. Original variable won't be modified, so in this case, pass the pointer
*/
func circleAreaPassByVal(c Circle) float64 {
	return math.Pi * c.r * c.r
}

// USE this!
func area(c *Circle) float64 {

	// Even though we passed in a pointer, when doing c.r, Go compiler treats this as (*c).r, so less typing
	return math.Pi * c.r * c.r
}

func doubleCircleRadius(c *Circle) {
	c.r *= 2
}

// Circle struct has ability to calculate area
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

// Rectangle struct has ability to calculate area
func (r *Rectangle) area() float64 {
	return r.height * r.width
}

func main() {
	c := Circle{x: 0, y: 0, r: 5}
	r := Rectangle{x: 5, y: 10, height: 2, width: 10}

	// access circle's attributes
	fmt.Println("x-coordinate of center: ", c.x)
	fmt.Println("y-coordinate of center: ", c.y)
	fmt.Println("radius of circle: ", c.r)

	// circle area calculation
	fmt.Println("area of circle: ", circleAreaPassByVal(c))
	fmt.Println("area of circle with ptr: ", area(&c))

	// mutate circle object but persist
	doubleCircleRadius(&c)
	fmt.Println("area of circle with doubled circle radius: ", area(&c))

	// rectangle area
	fmt.Println("Rectangle area: ", r.area())
}
