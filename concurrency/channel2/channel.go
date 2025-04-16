package main

import "fmt"

func f1(c chan string) {

	for {
		msg, open := <-c

		if !open {
			break
		}
		fmt.Println("f1", msg)
	}
	// this will not execute beacuse blocking nature and main will exit
	fmt.Println("---------f1--------")
}
func f2(c chan string) {
	close(c)
	for {
		msg, open := <-c

		if !open {
			break
		}
		fmt.Println("f2", msg)
	}
	// this will not execute beacuse blocking nature and main will exit
	fmt.Println("---------f2--------")

}

func f3(c chan string) {
	for {
		msg, open := <-c

		if !open {
			break
		}

		fmt.Println("f3", msg)
	}
	// this will not execute beacuse blocking nature and main will exit
	fmt.Println("---------f3--------")
}

func main() {
	c := make(chan string)

	go f1(c)
	go f2(c)
	go f3(c)

	for i := 0; i < 100; i++ {
		c <- fmt.Sprintf("M%d FROM Main", i)
	}

}
