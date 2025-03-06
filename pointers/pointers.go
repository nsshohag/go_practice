package main

import "fmt"

// pass by reference
func p(ptr *int) {
	*ptr = 0
}

// pass by value
func f(b int) {
	b = 0
}

func main() {
	a := 10
	b := 20
	f(b)
	p(&a)
	// here you will see that value is changed
	fmt.Println(a, b)
}
