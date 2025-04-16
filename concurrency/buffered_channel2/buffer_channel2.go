package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
)

// Buffered Channels Only Avoid Blocking on Send (if Not Full), Not on Receive:
func f(c chan string) {
	fmt.Println("Before Request Get")
	time.Sleep(time.Second * 3)
	fmt.Println("Num of CPU", runtime.NumCPU())
	fmt.Println("Num of goroutine", runtime.NumGoroutine())
	fmt.Println(<-c)
	fmt.Println("After Channel")
	wg.Done()
}

var wg = sync.WaitGroup{}

func main() {

	c := make(chan string, 100)
	for i := 0; i < 100; i++ {
		wg.Add(1)

		go f(c)
		c <- strconv.Itoa(i)

	}
	wg.Wait()
}
