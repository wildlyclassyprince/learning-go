/*
	Stopped at Chapter 6, problem 4.
	How do you count and compare values in a loop?
*/
package main

import "fmt"

func main() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	l := x[0]
	r := x[len(x) - 1]
	for i:=0; i < len(x); i++ {
		if l > r {
			l = x[i]
			fmt.Println(l)
		} else if r > l {
			r = x[len(x) - i]
			fmt.Println(r)
		} else if r == l {
			fmt.Println(r)
		}
	}
}
