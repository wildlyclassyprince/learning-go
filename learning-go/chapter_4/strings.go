package main

import "fmt"

var x string = "Hello World."

func main() {
	fmt.Println(x)
	x = "I've changed."
	fmt.Println(x)
	f()
	// inferring string type
	y := ":)"
	fmt.Println(y)
}

func f() {
	x += " Again."
	fmt.Println(x)
}
