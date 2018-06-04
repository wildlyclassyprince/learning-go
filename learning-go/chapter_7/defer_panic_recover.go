package main

import "fmt"

func main(){
	defer first()
	second()
	panic_recover()
}

func first(){
	fmt.Println("First")
}

func second(){
	fmt.Println("Second")
}

/*
We can handle a run-time panic 
with the built-in 'recover'
function. 'recover' stops the 
panic and returns the value that 
was passed to the call to 'panic'.
*/

func panic_recover(){
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}
