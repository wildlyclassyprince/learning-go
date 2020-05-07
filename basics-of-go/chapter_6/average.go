package main

import "fmt"

var	total float64 = 0

func main() {
	version_1()
	version_2()
}

func version_1() {
	x := [5]float64{ 98, 93, 77, 82, 83, }
	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	fmt.Println(total/float64(len(x)))
}

func version_2() {
	x := [5]float64{ 98, 93, 77, 82, 83, }
	for _, value := range x {
		total += value + 1
	}
	fmt.Println(total/float64(len(x)))
}
