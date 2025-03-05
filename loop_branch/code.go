package main

import (
	"fmt"
	"time"
)

func main() {

	// looping
	for {
		fmt.Println("Hello")
		break
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for i := range 3 {
		fmt.Println("range", i)
	}

	// here range
	names := []string{"Sadat", "Rony", "Abir"}
	for i, name := range names {
		fmt.Printf("%d: %s\n", i, name)

		if i == 0 {
			break
		}
	}

	// printing only odd numbers
	for i := range 10 {

		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	// branching
	fmt.Println(time.Now())
	fmt.Printf("%T\n", time.Now())

	//Compared to C, JavaScript, and other languages, we don’t need to have a break after each case.
	switch time.Now().Weekday() {
	case time.Friday, time.Saturday:
		fmt.Println("It is weekend")
		fmt.Println("So chill !")
	case time.Tuesday:
		fmt.Println("Home Office !")
	default:
		fmt.Print("You are dead.")
	}

}
