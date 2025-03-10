package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	go count("message 1", c1)

	for {
		msg, open := <-c1
		if !open {
			break
		}
		fmt.Println(msg + " from channel 1")
	}

	// we could change it using range function it would be easier , nice and clean
	/*
		for msg:= range c{
		fmt.Println(msg + " from channel 1")
		}
	*/
	c2 := make(chan string)

	go func() {
		c2 <- "message from channel 2"
	}()
	msg2 := <-c2
	fmt.Println(msg2)
}

func count(thing string, cn chan string) {

	for i := 1; i <= 5; i++ {
		cn <- thing
		time.Sleep(time.Millisecond * 500)
	}

	// sender need to close the channel not the reciever. because sender decides when to send and when finished.
	// if sender does not close the channel receiver will wait for infinte times and cause fatal error
	// and it is runtime error
	close(cn)

}
