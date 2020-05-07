package main

import "fmt"

func main() {
	x := 1
	y := 2
	swap(&x, &y)
}

func swap(xPtr *int, yPtr *int) {
	*xPtr = *yPtr
	defer *yPtr = *xPtr
	fmt.Println("x =", xPtr)
	fmt.Println("y =", yPtr)
}
