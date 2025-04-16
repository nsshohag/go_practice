package main

import (
	"fmt"
	"time"
)

// tickers are for when you want to do something repeatedly at regular intervals.
// like timer we can also stop that

func main() {

	ticker := time.NewTicker(500 * time.Millisecond)

	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// we can also stop the ticker like timer
	/*
		go func() {
			time.Sleep(time.Millisecond * 1000)
			ticker.Stop()
		}()
	*/
	time.Sleep(2571 * time.Millisecond)
	done <- true

	fmt.Println("Ticker stopped")
}
