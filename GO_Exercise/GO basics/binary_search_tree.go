package main

import "fmt"

type Tree struct {
	root *tree
}

type tree struct {
	x     string
	left  *tree
	right *tree
}

// Tree Creation
func (t *Tree) add(data string) {
	if t.root == nil {
		t.root = &tree{x: data}
	} else {
		t.root.insert(data)
	}
}

// tree insertion
func (n *tree) insert(data string) {
	// fmt.Println(n.x, data)
	if data <= n.x {
		if n.left == nil {
			n.left = &tree{x: data}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &tree{x: data}
		} else {
			n.right.insert(data)
		}
	}
}

// PreOrder Way
func printPreOrder(n *tree) {
	if n == nil {
		return
	} else {
		fmt.Printf("%s ", n.x)
		printPreOrder(n.left)
		printPreOrder(n.right)
	}
}

func main() {
	var t Tree

	t.add("D")
	t.add("C")
	t.add("E")
	t.add("G")
	t.add("H")
	t.add("I")

	printPreOrder(t.root)
}
