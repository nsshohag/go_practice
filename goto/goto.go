package main

import "fmt"

// the main work of goto is to jump

func main() {

	i := 0

Here: // this is the label where find goto it will jump here
	if i == 10 {
		return
	}
	fmt.Println(i)
	i++
	goto Here

	fmt.Println("Not") // unreachable in this case

}
