package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const frenchHelloPrefix = "Bonjour "
const spanishHelloPrefix = "Hola "
const englishHelloPrefix = "Hello "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	if language == spanish {
		return spanishHelloPrefix + name
	}
	if language == french {
		return frenchHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Rita", "Spanish"))
	fmt.Println(Hello("Stan", "French"))
}
