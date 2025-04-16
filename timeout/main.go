package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string, 1)
	go func() {
		time.Sleep(3 * time.Second)
		c1 <- "1st channel data"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case msg := <-time.After(2 * time.Second): // The time.After function creates a channel that will send the current time after the specified duration (in this case, 1 second
		fmt.Println("timeout 1", msg)
	}

	fmt.Println("Printing C1", <-c1)

	time.Sleep(time.Second * 5)

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	var msg time.Time
	select {
	case res := <-c2:
		fmt.Println(res)
	case msg = <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}

	time.Sleep(time.Second * 4)
	fmt.Println(msg)
}
