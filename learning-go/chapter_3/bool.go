package main

import "fmt"

func main() {
	fmt.Println("True and True ->", true && true)
	fmt.Println("True and False->", true && false)
	fmt.Println("True or True ->", true || true)
	fmt.Println("True or False ->", true || false)
	fmt.Println("Not True ->", !true)
}
