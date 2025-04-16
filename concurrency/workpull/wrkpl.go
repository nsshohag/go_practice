package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
}

// package main

// import "fmt"

// func main() {

// 	jobs := make(chan int, 100)
// 	results := make(chan int, 100)

// 	// in this case it will take 4 cores
// 	go worker(jobs, results)
// 	go worker(jobs, results)
// 	go worker(jobs, results)
// 	go worker(jobs, results)

// 	for i := 0; i < 100; i++ {
// 		jobs <- i
// 	}
// 	close(jobs)

// 	for j := 0; j < 100; j++ {
// 		fmt.Println(<-results)
// 	}
// }

// func worker(jobs <-chan int, results chan<- int) {

// 	for n := range jobs {
// 		results <- fib(n)
// 	}

// }

// func fib(n int) int {

// 	if n <= 1 {
// 		return n
// 	}

// 	return fib(n-1) + fib(n-2)
// }
