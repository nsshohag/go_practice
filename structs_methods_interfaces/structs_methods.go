package main

import "fmt"

type myType struct {
	name string
	reg  int
}

func f(name string) *myType {
	s4 := myType{name, 2018331000}
	return &s4
}

type in int

// mane ekekta datatype er sathe ekekta function assign kortechi za method name porichito
// tobe mul kotha holo custom data type use korte hobe
// can not use built in data type need to use type
// methods (You can assign a function to a struct,dataype and in this case we call it a method )
// method two types 1. value reciever and pointer reciever

// Methods in Go cannot have the same name as the struct.
func (p myType) value_reciever_method() {
	fmt.Println(p.name, "showing name from value reciever methods")
}

func (p *myType) pointer_reciever_method() {
	fmt.Println(p.name, "showing name from pointer reciever methods")
}

func (m in) method() {
	fmt.Println("Showing the integer number from method instance", m)
}

type S struct {
	Name       string
	red_marked bool
}

// Constructor function with default values
func Constructor_Zero_Arg() S {
	return S{
		Name:       "Not Provided",
		red_marked: false,
	}
}

// Constructor function with values
func Constructor_2_Arg(name string, r bool) S {
	return S{
		Name:       name,
		red_marked: r,
	}
}

// Constructor function with values
func Constructor_1_Arg(name string) S {
	return S{
		Name:       name,
		red_marked: false,
	}
}

func main() {

	var s0 myType
	s1 := myType{"Sadat", 2018331099}
	var s2 = myType{"Sadat", 2018331099}

	var s3 = myType{}
	s3.name = "Sohan"

	s4 := f("Muhit")
	s5 := *s4
	fmt.Println(s0, s1, s2, s3)

	fmt.Printf("Type = %T and the Data is %s\n", s5, s5.name)

	s1.value_reciever_method()
	s1.pointer_reciever_method()

	// could not use built in int needed to use custom type int
	// to assgin to method
	var m in = 10
	m.method()

	// Using Constructor Function

	st1 := Constructor_Zero_Arg()
	st2 := S{"Sadat", false}
	st3 := Constructor_2_Arg("Rony", true)
	st4 := Constructor_2_Arg("Preity", false)
	st5 := Constructor_1_Arg("Tusher")
	fmt.Println(st1, st2, st3, st4, st5)
}
