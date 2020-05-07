package main

import "fmt"

func main() {
	x := []int{
		6, 2, 3, 4, 5,
	}

	l := x[0]
	r := x[len(x) - 1]

	for i:=0; i < len(x); i++ {
		if l > r {
			l = x[i]
			fmt.Println(r)
		} else if r == l {
			fmt.Println(l)
		} else {
			r = x[len(x) - 2]
			fmt.Println(r)
		}
	}
}
