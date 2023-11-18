package main

import (
	"fmt"
	"sync"
)

func main() {
	concurrentDemo()
	channelDemo()
	selectDemo()
	channelLoopDemo()
}

func concurrentDemo() {
	var wg sync.WaitGroup

	wg.Add(1)   // add 1 work value
	go func() { //goroutine
		fmt.Println("This happens asynchronous")
		wg.Done() //decrement task
	}()

	fmt.Println("This happens synchronous")
	wg.Wait() //wait until wg is decremented to 0
}

func channelDemo() {
	var wg sync.WaitGroup
	ch := make(chan string) //channel
	wg.Add(1)

	go func() {
		ch <- "The message"
	}()

	go func() {
		fmt.Println(<-ch)
		wg.Done() //decrement task
	}()

	wg.Wait()
}

func selectDemo() {
	ch1, ch2 := make(chan string), make(chan string)

	go func() {
		ch1 <- "The message to channel 1"
	}()

	go func() {
		ch2 <- "The message to channel 2"
	}()

	select { //case will be chosen at random, as per go guidelines
	case msg := <-ch1:
		fmt.Println(msg)
	case msg := <-ch2:
		fmt.Println(msg)
	}
}

func channelLoopDemo() {
	ch := make(chan int) //channel

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	for msg := range ch {
		fmt.Println(msg)
	}

	fmt.Println("Done with channel loop")
}
