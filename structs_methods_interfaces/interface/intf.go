// An interface is a type that defines one or more method signatures.
// Methods are not implemented, just their signature: the name, parameter types and return value type.
package main

import "fmt"

// Define an interface
// same greet method is used twice
type Greeter interface {
	Greet()
}

// Define a struct
type Person struct {
	Name string
}

// Implement the Greet method for Person
func (p Person) Greet() {
	fmt.Println("Hello, my name is", p.Name)
}

// Another struct implementing the same interface
type Dog struct {
	Name string
}

// Implement the Greet method for Dog
func (d Dog) Greet() {
	fmt.Println("Woof! My name is", d.Name, ".And I'm a dog")
}

// Function that accepts any type that implements the Greeter interface
func greetSomeone(g Greeter) {
	g.Greet() // Calls the Greet method of the passed object
}

func main() {
	p := Person{Name: "Sadat"}
	p2 := Person{Name: "Rony"}
	d := Dog{Name: "Alex"}

	greetSomeone(p)
	greetSomeone(p2)
	greetSomeone(d)
}
