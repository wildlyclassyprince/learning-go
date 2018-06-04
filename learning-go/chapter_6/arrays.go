package main

import "fmt"

func main() {
	var x [5]int
	x[2] = 22
	x[4] = 100
	fmt.Println(x)
	fmt.Println(x[2])
	fmt.Println(x[4])
}
