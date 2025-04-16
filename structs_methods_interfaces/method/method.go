package main

import (
	"fmt"
	m "freecode/structs_methods_interfaces/method/method_inheritance_overriding"
	"math"
)

type Circle struct {
	radius float64
}

type Rectangle struct {
	width, height float64
}

// Methods in Go cannot have the same name as the struct.
// method passed by value
func (c Circle) Area() float64 {
	c.radius = 100
	return c.radius * c.radius * math.Pi
}

// method passed by value but func (c *Circle) Area() float64 { passed by reference
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func main() {
	c1 := Circle{10}
	c2 := Circle{25}
	r1 := Rectangle{9, 4}
	r2 := Rectangle{12, 2}

	/*
	   One thing that's worth noting is that the method with a dotted line means the receiver is passed by value, not by reference.
	   The difference between them is that a method can change its receiver's values when the receiver is passed by reference
	   and it gets a copy of the receiver when the receiver is passed by value.
	*/

	fmt.Println("Area of c1 is: ", c1.Area()) // passed by value so won't change the value
	fmt.Println("Area of c2 is: ", c2.Area())
	fmt.Println("Area of r1 is: ", r1.Area())
	fmt.Println("Area of r2 is: ", r2.Area())
	fmt.Println(c1.radius)

	fmt.Printf("\nInheritance of Methods\n---------------------------\n")
	m.Method_Inheritance()

	fmt.Printf("\nOverriding of Methods\n---------------------------\n")
	m.Method_Overriding()
}
