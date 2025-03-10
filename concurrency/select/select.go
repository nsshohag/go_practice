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
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}

	}()

	go func() {

		for {
			c2 <- "Every two seconds"
			time.Sleep(time.Second * 2)
		}

	}()

	// but here one after another block slower one . So we have to wait for slower one
	/*
		for {
			fmt.Println(<-c1)
			fmt.Println(<-c2)
		}
	*/

	// here we can use select as a result which ever channel is ready it will print out
	for {

		select {

		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}

	}

}
