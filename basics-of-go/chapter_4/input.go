package main

import "fmt"

func main() {
	fmt.Println("Enter a number: ")
	var input float64

	// read and assign input to the 'input' pointer
	fmt.Scanf("%f", &input)

	// do something to the input
	output := input * 2
	fmt.Println(output)
}
