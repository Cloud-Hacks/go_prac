package main

import "fmt"

type tree struct {
	key   int
	left  *tree
	right *tree
}

type root struct {
	T *tree
}

func (t *root) initTree(x int) {
	if t.T == nil {
		t.T = &tree{key: x}
	} else {
		t.T.inst(x)
	}
}

func (t *tree) inst(x int) {
	if t.key <= x {
		if t.left == nil {
			t.left = &tree{key: x}
		} else {
			t.left.inst(x)
		}
	} else {
		if t.right == nil {
			t.right = &tree{key: x}
		} else {
			t.right.inst(x)
		}
	}
}

func (t *tree) PreOrder() {
	if t != nil {
		fmt.Println(t.key)
		t.left.PreOrder()
		t.right.PreOrder()
	} else {
		return
	}
}

func main() {
	var t root

	t.initTree(5)
	t.initTree(7)
	t.initTree(4)

	t.T.PreOrder()

}
