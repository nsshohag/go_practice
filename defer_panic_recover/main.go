package main

import "fmt"

func f1() {
	fmt.Println("1st")

}

func f2() {
	fmt.Println("2nd")
	num := 0
	println(0 / num)
	panic("my panic")
}

func f3() {
	fmt.Println("3rd")
}
func f4() {
	fmt.Println("4th")
}

func main() {
	// here defer runs at last like finally in java
	// it's needed especilly for files closing at last whatever happens files get closed

	defer f3()
	defer func() {
		str := recover()
		fmt.Println(str)
	}()

	f4()
	num := 0
	println(0 / num)
	// sohoj kothay panic paile er pore ar kichui execute korbe na
	f1()
	defer f2()

	fmt.Println("After Panic") // unrechable beacuse paniced has happened

}
