package main

import "fmt"

// multiple argument passing which is passed as slieces and mulitiple return values
func sum_numbers(numbers ...int) (int, []int) {

	// fmt.Println(numbers) // the data type here is slice for multiple argument
	fmt.Printf("%T\n", numbers)
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum, numbers

}

func main() {

	// multiple argument passing and multiple value returns
	sum, slice := sum_numbers(100, 100, 100, 100)
	fmt.Println("The numbers are", slice, "\nAnd he sum is =", sum)
}
