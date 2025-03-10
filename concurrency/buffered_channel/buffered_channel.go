package main

import "fmt"

func main() {

	// this is buffered_channel which will not block the channel
	// meaning this will let in this case two communications happen without the receiver

	// but if we put c<- "Third M" then it will panic and cause error

	c := make(chan string, 2) // buffererd_channel zar capacity 2

	c <- "First M"
	c <- "Second M"

	msg := <-c
	fmt.Println(msg)
	msg2 := <-c
	fmt.Println(msg2)
}
