package main

import (
	"fmt"
	"sync"
	"time"
)

func sum_seq() int {
	sum := 0
	for i := 1; i <= 10e6; i++ {
		sum += i
	}
	return sum
}

func sum_for() int {
	n := int(10e6)
	sum := (n * (n + 1)) / 2
	return sum
}

var wg = sync.WaitGroup{}

func worker(start int, end int, resultChannel chan<- int) {
	sum := 0
	for i := start; i <= end; i++ {
		sum += i
	}
	resultChannel <- sum
	wg.Done()
}

func sum_goroutine() int {
	Worker := 10
	N := int(10e6)
	resultChan := make(chan int, Worker)
	chunk := N / Worker

	for i := 0; i < Worker; i++ {

		wg.Add(1)
		Start := i*chunk + 1
		End := (i + 1) * chunk
		if i == Worker-1 {
			End = N
		}

		go worker(Start, End, resultChan)
	}

	wg.Wait()
	close(resultChan)

	total_sum := 0
	for partial_sum := range resultChan {
		total_sum += partial_sum
	}
	return total_sum
}

func main() {

	// sequential
	sum := 0
	it := 50
	start := time.Now()
	for i := 0; i < it; i++ {
		sum = sum_seq()
	}
	total := time.Since(start)
	fmt.Printf("Sum = %d Total Time Using Seq %s\n", sum, total/time.Duration(it))

	// formula
	start = time.Now()
	for i := 0; i < it; i++ {
		sum = sum_for()
	}
	total = time.Since(start)
	fmt.Printf("Sum = %d Total Time Using For %s\n", sum, total/time.Duration(it))

	// goroutine
	start = time.Now()
	for i := 0; i < it; i++ {
		sum = sum_goroutine()
	}
	total = time.Since(start)
	fmt.Printf("Sum = %d Total Time Using GoRoutine %s\n", sum, total/time.Duration(it))
}
