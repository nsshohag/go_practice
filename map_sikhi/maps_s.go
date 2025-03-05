package main

import "fmt"

func main() {

	// make a map
	m := make(map[string]string)
	// give input
	m["Sadat"] = "2018331099"
	m["Rony"] = "2018331053"
	m["Preity"] = "2018331053"
	fmt.Println(m)
	fmt.Println("Value for Sadat is", m["Sadat"])
	fmt.Printf("Data type is = %T\n", m)

	// delete from a map
	delete(m, "Preity")
	fmt.Println(m)

	// check if the key was present or not
	// _ is used to ignore value , if we want value then we can use a variable
	_, prs := m["Preity"]
	fmt.Println("Key exist for Preity ?:", prs)

	// delete all
	clear(m)
	fmt.Println(m)

}
