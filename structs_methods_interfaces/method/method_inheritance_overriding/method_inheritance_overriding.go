package m

import "fmt"

type Human struct {
	name    string
	age     int
	overlap string
}

type Student struct {
	Human
	school  string
	overlap string
}

// Methods in Go cannot have the same name as the struct.
func (h Human) Human_M() {
	fmt.Println("From Human Method")
}

func (s Student) Student() {

	fmt.Println("From Student Method")
}

// Method and Overridden Method for Student and Human

func (h Human) Method() {
	fmt.Println("This method (Human) is overridden. This is Parent")
}

func (h Student) Method() {
	fmt.Println("This method (Student) is overiding. This is Child")
}

// exported method name must start with Capital Alphabet
func Method_Inheritance() {

	sadat := Student{Human{"Sadat", 24, "overlap_from Human"}, "BGPSC", "overlap_from Student"}
	sadat.Student()

	fmt.Println(sadat.Human.name, sadat.name)       // here in this case no overlap so inherited
	fmt.Println(sadat.overlap, sadat.Human.overlap) // here in this case overlapped so accessed through different referenc

	// inheritance of methods
	sadat.Student()
	sadat.Human_M()

}

// exported method name must start with Capital alphabet
func Method_Overriding() {
	sadat := Student{Human{"Sadat", 24, ""}, "BGPSC", ""}

	sadat.Method()

	sadat.Human.Method() // To access overridden method call it via it's parent reference
}
