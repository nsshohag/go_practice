package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {

		for {
			c1 <- "Every 500ms"
			fmt.Println(runtime.NumCPU())
			fmt.Println(runtime.NumGoroutine())
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

		/*
			select is blocking by default and it continues to execute only when one of channels has data to send or receive.
			If several channels are ready to use at the same time, select chooses which to execute randomly
		*/

		select {

		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
			// default if aboved channel are blocked then deafult will execute
			//default:
			///fmt.Println("CCC")
		}

	}

}
