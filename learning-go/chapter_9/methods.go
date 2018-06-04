package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Method : function with a receiver argument
func (x Vertex) SS() float64 {
	return math.Sqrt(x.X*x.X + x.Y*x.Y)
}

// Rewritten method with struct as argument
func SS1(x Vertex) float64 {
	return math.Sqrt(x.X*x.X + x.Y*x.Y)
}

func main() {
	v := Vertex{1, 2}
	fmt.Println(v.SS())
	fmt.Println(SS1(v))
}
