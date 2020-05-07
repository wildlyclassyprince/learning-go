package main

import "fmt"

func main() {
	x := 5
	value(&x)
}

func value(xPtr *int) {
	*xPtr = 10
	fmt.Println(*xPtr)
}
