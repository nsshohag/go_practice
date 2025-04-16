package main

import "fmt"

func fibo(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibo(n-1) + fibo(n-2)
}

// multiple argument passing
func sum_numbers(numbers ...int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

// add two numbers // can not overload // But method can be overload
func sum(a int, b int) int {
	return a + b
}

func main() {
	//
	fmt.Println("7th Fibo =", fibo(7))
	// multiple argument passing
	fmt.Println(sum_numbers(1, 2, 3, 4, 5, 6))
	fmt.Println("Sum of 1+1 =", sum(1, 1))
}
