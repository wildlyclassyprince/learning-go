package main

import "math"

// Perimeter calculates the perimeter of a rectangle given a height and width.
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Rectangle encapsulates the properties of a rectangle.
type Rectangle struct {
	Width  float64
	Height float64
}

// Area method calculates the area of a rectangle given the width and height
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle encapsulates the properties of a circle
type Circle struct {
	Radius float64
}

// Area method calculates the area of a circle given the radius
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
