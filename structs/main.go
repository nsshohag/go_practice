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

}
