package main

import "fmt"

func main() {
	x := []float64{ 999, 99, 9, }
	y := make([]float64, 15)
	y[0] = 1
	y[5] = 10
	y[7] = 14
	fmt.Println(len(x))
	fmt.Println(y)
}
