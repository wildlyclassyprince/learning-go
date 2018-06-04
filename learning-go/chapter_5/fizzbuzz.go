/*
	FizzBuzz
*/

package main

import "fmt"

func main() {
	var (
		x string = "Fizz"
		y string = "Buzz"
	)

	for i:=1; i <= 100; i++ {
		if (i % 3 == 0) && (i % 5 != 0) {
			fmt.Println(x)
		} else if (i % 5 == 0) && (i % 3 != 0) {
			fmt.Println(y)
		} else if (i % 3 ==0) && (i % 5 ==0 ) {
			fmt.Println(x + y)
		} else {
			fmt.Println(i)
		}
	}
}
