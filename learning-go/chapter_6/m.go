package main

import "fmt"

func main() {
	var input string
	x := map[string]string{
		"a" : "A",
		"b" : "B",
		"c" : "C",
	}

	fmt.Println("Enter any one of 'a', 'b', 'c': ")
	fmt.Scanf("%s", &input)
	fmt.Println("Value: ", x[input])
}
