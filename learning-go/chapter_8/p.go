package main

import "fmt"

func main(){
	x := 10
	xPtr := &x
	fmt.Println("Value of x", x)
	fmt.Println("Value of *xPtr", *xPtr)
	fmt.Println("Address of x direct", &x)
	fmt.Println("Address of x by xPtr", xPtr)
	fmt.Println("Address of xPtr", &xPtr)
}
