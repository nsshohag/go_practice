package main

import (
	"fmt"
	"sync"
)

func f(from string) {
	for i := 0; i < 10; i++ {
		fmt.Println(from, i)
	}
	wg.Done()
}
func f2(from string) {
	for i := 0; i < 10; i++ {
		fmt.Println(from, i)
	}
}

func direct_f(from string) {
	for i := 0; i < 500; i++ {
		fmt.Println(from, i)
	}
}

var wg = sync.WaitGroup{}

func main() {

	//f execute and then other later functions will start creating goroutine adn doing stuff
	direct_f("direct")

	wg.Add(3)
	go f("goroutine 1")
	go f("goroutine 2")
	go f("goroutine 3")

	// this may partially or not execeute because the main goroutine may exit before this go routine finish executing
	// as we have not added it to the WaitGroup
	go f2("from f2 goroutine 2")

	go func(msg string) {
		fmt.Println(msg)
	}("annonymous function")

	wg.Wait()

	//time.Sleep(time.Millisecond) //eida basically main go routine ke running rakhe tachara other goroutine/thread execute howar agei eida exit hoye mane main program e exit hoye zabe

}
