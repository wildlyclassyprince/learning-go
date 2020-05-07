package main

import "fmt"

func main(){
	x := 5
	if half(x) % 2 == 0 {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}

func half(x int) int {
	y := x/2
	return y
}
