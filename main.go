package main

import (
	"fmt"
	"strings"
)

var reg uint64 = 2018331099
var name string = "Nazmus Sadat"
var c byte = '3'
var F_32 float32
var var_float float64 = 5.6

func ff() string {
	y := ", is my name."
	x := "Sadat"
	x += y
	return x
}

func main() {
	// f()
	// basic data type
	x := ff()
	F_32 := var_float
	fmt.Println("Hello World")
	fmt.Println(reg)
	fmt.Println(x[0:])
	fmt.Println(strings.HasPrefix(x, "sadat"))
	fmt.Println(c)
	fmt.Println(x[0])
	fmt.Println(var_float)
	fmt.Println(F_32)

	// strings in go data type string
	fmt.Println(strings.Contains("Sadatx", "Sadat"))

	// slice in go  data type []string
	fmt.Println(strings.Split(x, ", "))
	fmt.Printf("%T\n", reg)
	fmt.Printf("%T\n", x)
	fmt.Printf("Data Type for Slice is %T\n", strings.Split(x, ", "))

	// arrays in go
	var myArray1 [3]string
	var myArray2 = [...]string{"P", "S"}
	myArray3 := myArray2
	myArray3[1] = "wirhu43wrh"
	fmt.Printf("%d\n", len(myArray1[1]))
	fmt.Println((myArray2))
	fmt.Println(myArray3)

	// slice in go like array just don't neeed to define size
	var slice1 = []string{"string1", "string2"}

	slice2 := []string{"s1", "s2", "s3"}

	slice3 := make([]string, 3)
	copy(slice2, slice1)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)

	// initialize slice from array
	myArray22 := [2]string{"1", "1"} // slice manei reference
	slice4 := myArray22[:]
	myArray22[0] = "1111111"
	fmt.Printf("%T and data is %s\n", slice4, slice4)

	// slice manei reference
	slice5 := []string{"1", "2", "3"}
	slice6 := slice5
	slice6[0] = "100"
	slice_temp := slice6
	slice6 = append(slice6, "????????")
	slice_temp[1] = "fefje"
	fmt.Println(slice5, slice6)

	// more of slice
	//newSlice1 := make([]string, 0, 10) // an empty slice with capacity 10
	newSlice1 := slice6[:1]
	slice6[0] = "cganged"
	fmt.Println(newSlice1)

	//Maps in Go

}
