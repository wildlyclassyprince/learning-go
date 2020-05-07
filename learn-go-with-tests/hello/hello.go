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
	prefix := englishHelloPrefix

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("Rita", "Spanish"))
	fmt.Println(Hello("Stan", "French"))
}
