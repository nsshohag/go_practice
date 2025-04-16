package main

import (
	"fmt"
	"time"
)

// we could use time.Sleep but 1 reason a timer may be useful
// is that we can cancel the timer before it fires meaning we can block that channel
func main() {

	timer1 := time.NewTimer(2 * time.Second)

	fmt.Println("?")
	msg := <-timer1.C // waits 2 second and then fire in the channel (sends time)
	fmt.Println("Timer 1 Fired", msg)

	timer2 := time.NewTimer(time.Second)

	go func() {
		<-timer2.C // The goroutine is still waiting on <-timer2.C, but the timer will never fire because it was stopped.
		fmt.Println("Timer 2 Fired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 Stopped")
	}
	time.Sleep(time.Second * 2)

}
