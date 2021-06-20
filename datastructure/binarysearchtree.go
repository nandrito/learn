package main

import "fmt"

var count int

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// Insert node to the tree
func (n *Node) Insert(k int) {
	if n.Key < k {
		// move to the right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		// move to the left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

func (n *Node) Search(k int) bool {
	count++
	if n == nil {
		return false
	}

	if n.Key < k {
		// right
		n.Right.Search(k)
	} else if n.Key > k {
		// left
		n.Left.Search(k)
	}

	return true
}

func main() {
	tree := &Node{Key: 100}
	tree.Insert(50)
	tree.Insert(300)
	tree.Insert(75)
	tree.Insert(80)
	tree.Insert(89)
	tree.Insert(10)
	tree.Insert(19)
	tree.Insert(211)

	fmt.Println(tree)
	fmt.Println(tree.Search(50))
	fmt.Println("This search thru how many node?", count)
}
