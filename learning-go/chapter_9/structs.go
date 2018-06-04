package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func main() {
	v := Vertex{1, 2}
	fmt.Println(v)
	fmt.Println(v.X)
	fmt.Println(v.Y)
	v.X = 4
	fmt.Println(v)
	fmt.Println(v.X)
	fmt.Println(v.Y)

	p := Vertex{6, 7}
	q := &p
	fmt.Println(*q)
}
