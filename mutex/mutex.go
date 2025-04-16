/*
A mutex (short for "mutual exclusion") is a synchronization primitive
used in concurrent programming to ensure that only one goroutine can access
a particular section of code or a shared resource at a time.
*/

/*
Deadlocks: Improper use of mutexes can lead to deadlocks
where two or more goroutines are waiting indefinitely for each other to release locks.
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

// Shared counter
var counter int

// Mutex to protect the counter
var mutex sync.Mutex

func increment() {
	for i := 0; i < 10; i++ {
		// Lock the mutex before accessing the shared resource
		mutex.Lock()
		counter++
		fmt.Println("Counter:", counter)
		// Unlock the mutex after accessing the shared resource
		mutex.Unlock()
		time.Sleep(time.Millisecond * 100) // Simulate work
	}
}

func main() {
	// Start multiple goroutines to increment the counter
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}
