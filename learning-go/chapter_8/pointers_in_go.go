package main

import "fmt"

func main(){
	i, j := 42, 2701

	p := &i				// point to i
	fmt.Println(*p)		// read i through the pointer p

	*p = 21				// set i through the pointer p
	fmt.Println(i)		// print new value of i

	p = &j				// point to j
	*p = *p / 37		// divide j through the pointer
	fmt.Println(j)		// print new value of j
}
