package main

import (
	"fmt"
	"time"
	"math/rand"
)


// a simple function illustrating goroutines.
// it's got some random delay so we can see they execute simultaneously
func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}


// pinger, ponger, & printer communicate over a channel.
func pinger(c chan<- string) {
	for i := 0;; i++ {
		c <- "ping" // send a message to `printer` (blocking)
	}
}

func ponger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func printer(c <-chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

// NOTE: main is an implicit goroutine
func main() {
	/*
	// Simple goroutines.
	for i:=0; i < 10; i++ {
		go f(i)  // another goroutine due to 'go f()'
	}
	*/

	// A Simple channel.
	//var c chan string = make(chan string)
	//go pinger(c)
	//go ponger(c)
	//go printer(c)

	// Select.
	// Note; the 2nd param (a buffer size) changes the channel from a
	// *synchronous* one to an *asynchronous* one.
	c1 := make(chan string)
	c2 := make(chan string)
	//c1 := make(chan string, 4)
	//c2 := make(chan string, 4)

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 1)
		}
	}()

	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			// select picks the first channel that is ready, then reads from
			// (or writes to) it.
			select {
				case msg1 := <-c1:
					fmt.Println(msg1)
				case msg2 := <-c2:
					fmt.Println(msg2)
				//default:
					//fmt.Println("nothing ready")
			}
		}
	}()

	// Wait for input so main doesn't stop.
	var input string
	fmt.Scanln(&input)
	fmt.Println(input)
}
