package main

import (
	"fmt"
	"strconv"
	"sync"
)

func f(m string, c chan string) {
	c <- m
	wg.Done()
}

var wg = sync.WaitGroup{}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	for i := 0; i < 10; i++ {

		wg.Add(2)
		go f("f1 Go Routine No "+strconv.Itoa(i), c1)
		go f("f2 GO Routine No "+strconv.Itoa(i), c2)
	}

	// Close channels after all goroutines are done -- this is deeeeep think
	// as we must have to close the channel as we used for loop
	go func() {

		wg.Wait()
		close(c1)
		close(c2)
	}()

	for c1 != nil || c2 != nil {
		select {
		case m1, open := <-c1:
			if open {
				fmt.Println(m1)
			} else {
				c1 = nil
			}
		case m2, open := <-c2:
			if open {
				fmt.Println(m2)
			} else {
				c2 = nil
			}

		}

	}
}
