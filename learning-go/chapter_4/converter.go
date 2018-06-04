package main

import "fmt"

/*
Convert degrees Celsius to Fahrenheit.
*/
func main() {
	fahreinheit_to_celsius()
	feet_to_meters()
}

func fahreinheit_to_celsius() {
	fmt.Println("Enter the temperature in Fahrenheit: ")
	var F float64

	fmt.Scanf("%f", &F)
	C := (F - 32) * 5/9
	fmt.Println("The temperature in Celsius is ", C)
}

func feet_to_meters() {
	fmt.Println("Enter distance in feet: ")
	var feet float64

	fmt.Scanf("%f", &feet)
	meters := 0.3048*feet
	fmt.Println("The distance in meters: ", meters)
}
