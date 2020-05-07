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
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Rita", "Spanish"))
	fmt.Println(Hello("Stan", "French"))
}
