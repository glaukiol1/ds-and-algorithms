package main

import "fmt"

// trees

type Node struct {
	data  int
	left  *Node
	right *Node
}

func main() {
	chain := Node{2, &Node{3, &Node{5, nil, nil}, &Node{6, nil, nil}}, &Node{4, nil, nil}}
	fmt.Println("Sum of first tree: " + fmt.Sprint(find_sum(chain)))
}

func find_sum(chain Node) int {
	if chain.left != nil {
		if chain.right != nil {
			return chain.data + (find_sum(*chain.left) + find_sum(*chain.right))
		} else {
			return chain.data + (find_sum(*chain.left) + 0)
		}
	} else {
		return chain.data
	}
}
