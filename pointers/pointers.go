// pointer arithmetic not allowed in like in c++ --- p++; or *(p+1)
// null is nil here
package main

import "fmt"

// pass by reference
func p(ptr *int) {
	*ptr = 0
}

type myType struct {
	node_val    int
	left_child  *myType
	right_child *myType
}

// pass by value
func f(b int) {
	b = 0
}

// this is method - value reciever
func DFS(v *myType) {
	if v != nil {
		fmt.Println("Node", v.node_val)
		DFS(v.left_child)
		DFS(v.right_child)
	}
}

func main() {
	a := 10
	b := 20
	f(b)
	p(&a)
	// here you will see that value is changed
	fmt.Println(a, b)

	node1 := myType{1, nil, nil}
	node2 := myType{2, nil, nil}
	node3 := myType{3, nil, nil}
	node4 := myType{4, nil, nil}

	node1.left_child = &node2
	node1.right_child = &node3
	node2.left_child = &node4

	fmt.Println("DFS Traversal of Nodes")
	DFS(&node1)

}
