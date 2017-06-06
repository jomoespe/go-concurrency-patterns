package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string) <-chan string { // Returns receive-only channel of strings
	c := make(chan string)
	go func() { // we launch the gorutine from inside the function
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c // returns the channel to the caller
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() { for { c <- <-input1 } }()
	go func() { for { c <- <-input2 } }()
	return c
}

func fanInWithSelect(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() { 
		for {
			select {
				case s := <-input1: c <-s
				case s := <-input2: c <-s
			}
		}
	}()
	return c
}

func fanIn3(input []<-chan string) <-chan string {
	c := make(chan string)
	go func() { for { c <- <-input[0] } }()
	return c
}


func main() {
	//c := fanIn(boring("joe"), boring("ann"))
	c := fanInWithSelect(boring("joe"), boring("ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're both are boring. Leaving.")
}
