package main

import "fmt"


func main(){
	version_1()
	version_2()
}

func version_1(){
	var total float64 = 0
	x := [5]float64{ 99,
					 78,
					 89,
					 100,
					 900,
					}
	for _, value := range x {
		total += value
	}
fmt.Println(total / float64(len(x)))
}

func version_2(){
	var total float64 = 0
	x := [5]float64{ 99, 78, 89, 100, 900 }
	for _, value := range x {
		total += value
	}
fmt.Println(total / float64(len(x)))
}
