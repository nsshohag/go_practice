package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	defer fmt.Println("defer")
	sigs := make(chan os.Signal, 1) // os.Signal type channel
	/*
		The signal.Notify function sets up the sigs channel to receive SIGINT and SIGTERM signals.
		When either of these signals is sent to the program, the signal will be delivered to the sigs channel.
		Your program can then read from the sigs channel and handle the signal appropriately
		Such as by cleaning up resources and shutting down gracefully.
	*/
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan bool, 1)

	go func() {
		sig := <-sigs
		fmt.Println()
		time.Sleep(time.Second)
		fmt.Println("Printing, os.Signal =", sig)
		done <- true
	}()

	fmt.Println("awaiting signal")

	<-done
	time.Sleep(time.Second * 2)
	// os.Exit(0) // zodi eida dei taile defer thak zai thak ta print hobe na direct exit korbe
	fmt.Println("exiting")

}
