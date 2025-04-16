// An interface is a type that defines one or more method signatures.
// Methods are not implemented, just their signature: the name, parameter types and return value type.
package main

import "fmt"

// Define an interface

type Greeter interface {
	Greet()
	Farewell()
}

// Define a struct
type Person struct {
	Name string
	age  int
}

// Implement the Greet method for Person
func (p Person) Greet() {
	fmt.Println("Hello, my name is", p.Name)
}

// Implement the Greet method for Person
func (p Person) Farewell() {
	fmt.Println("Bye, my name is", p.Name)
}

// Another struct implementing the same interface
type Dog struct {
	Name string
}

// Implement the Greet method for Dog
func (d Dog) Greet() {
	fmt.Println("Woof! My name is", d.Name, ".And I'm a dog")
}

func (d Dog) Farewell() {
	fmt.Println("Bye! My name is", d.Name, ".And I'm a dog")
}

// Function that accepts any type that implements the Greeter interface
func greetSomeone(g Greeter) {

	fmt.Println(g)

	g.Greet() // Calls the Greet method of the passed object
	g.Farewell()
}

func main() {
	p := Person{Name: "Sadat", age: 24}
	p2 := Person{Name: "Rony"}
	d := Dog{Name: "Alex"}

	// IDK What is happening // Method can be overload // override
	p.Greet()

	greetSomeone(p)
	greetSomeone(p2)
	greetSomeone(d)
}

/*
package main

import (
    "fmt"
    "math"
)

type geometry interface {
    area() float64
    perim() float64
}

type rect struct {
    width, height float64
}
type circle struct {
    radius float64
}

func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func detectCircle(g geometry) {
    if c, ok := g.(circle); ok {
        fmt.Println("circle with radius", c.radius)
    }
}

func main() {
    r := rect{width: 3, height: 4}
    c := circle{radius: 5}

    measure(r)
    measure(c)

    detectCircle(r)
    detectCircle(c)
}
*/
