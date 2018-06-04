/*
This program prints "from 1" every 2 seconds and "from 2" every 3 seconds.
'select' picks the first channel that is ready and receives it (or sends to it).
If more than one of the channels are ready, then it randomly picks which one to 
receive from. If none of the channels are ready, the statement blocks until one
becomes available.

The 'select' statement is often used to implement a timeout. 'time.After' creates
a channel and after the given duration will send the current time on it (we aren't
interested in the time so we don't store it in a variable). We can also specify a
default case which happens immediately if none of the channels are ready.
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()
	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()
	go func () {
		for {
			select {
			case msg1 := <- c1:
				fmt.Println("Message 1", msg1)
			case msg2 := <- c2:
				fmt.Println("Message 2", msg2)
			case <- time.After(time.Second):
				fmt.Println("Timeout.")
			//default:
			//	fmt.Println("Nothing ready. Waiting ...") // Prints default select non-stop if uncommented
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}
