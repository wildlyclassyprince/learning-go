package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

// Countdown - prints a value and waits one second before printing the next
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}

// Sleeper adds a delay to on our countdown
type Sleeper interface {
	Sleep()
}

// DefaultSleeper struct
type DefaultSleeper struct{}

// Sleep method implements Sleeper interface
func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}
