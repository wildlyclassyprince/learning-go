package main

import "fmt"

func main(){
	var x string = "Home"

	for i:=0; i < len(x); i++ {
		fmt.Println(string(x[i]))
	}
}
