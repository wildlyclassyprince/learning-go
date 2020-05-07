package main

import "fmt"

var j, k int = 1, 2

func main(){
	fmt.Println(f(j, k))
}

func f(x int, y int) (int, int){
	return x, y
}
