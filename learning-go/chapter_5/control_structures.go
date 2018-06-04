/* 
	Control Flow: 'for' loop and 'if' conditional
*/

package main

import "fmt"

func main() {
	version_1()
	version_2()
}

func version_1() {
	// Counting from 1 to 10
	fmt.Println("Version 1 (long form)")
	// initialize
	i := 0
	// loop
	for i <= 10 {
		if i % 2 == 0 {
			fmt.Println(i, "is even.")
			i++
		} else {
			fmt.Println(i, "is odd.")
			i++
		}
	}
}


func version_2() {
	fmt.Println("Version 2 (short form, C-like)")
	for  i := 0; i <= 10; i++ {
		if i % 2 == 0 {
			fmt.Println(i, "is even.")
		} else {
			fmt.Println(i, "is odd.")
		}
	}
}
