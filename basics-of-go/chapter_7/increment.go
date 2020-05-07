package main

import "fmt"

/*
A function like this together with the non-local
variables it references is known as a closure. In
this case 'increment' and the variable 'x' form
the closure.

One way to use a closure is by writing a function
which returns another function which - when called 
- can genereate a sequence of numbers.
*/
func main(){
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())

	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())

	nextOdd := makeOddGenerator()
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
}

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

/*
A function is able to call itself - Recursion.
*/
func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x - 1)
}
