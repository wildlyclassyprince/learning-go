package main

import "fmt"

const spanish = "Spanish"
const spanishHelloPrefix = "Hola "
const englishHelloPrefix = "Hello "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	if language == spanish {
		return spanishHelloPrefix + name
	}
	if language == "French" {
		return "Bonjour " + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Rita", "Spanish"))
}
