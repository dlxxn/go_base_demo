package examples

import (
	"fmt"
	"time"
)

func TestChannels() {
	message := make(chan string)
	go func() {
		message <- "ping"
	}()

	msg := <-message
	fmt.Println(msg)
}

/**
Channel Synchronization
*/
func Worker(done chan bool) {
	fmt.Println("working.....")
	time.Sleep(time.Minute)
	fmt.Println("Done")

	done <- true
}

/**
When using channels as function parameters, you can specify if a channel is meant to only send or
receive values. This specificity increases the type-safety of the program.
*/

func ping(pings chan<- string, msg string) {
	pings <- msg
}
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
func TestPingPong() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

/**
Go’s select lets you wait on multiple channel operations. Combining goroutines and channels
with select is a powerful feature of Go.

Note that the total execution time is only ~2 seconds since both the 1 and 2 second Sleeps
execute concurrently.
*/
func TestSelect() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 2)
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}
