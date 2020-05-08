package main

import (
	"bytes"
	"fmt"
)

// Greet prints the name given
func Greet(writer *bytes.Buffer, name string) {
	fmt.Printf("Hello, %s", name)
}
