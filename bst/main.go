package main

import "fmt"

type tree struct {
	item *node
}

type node struct {
	value int
	left  *node
	right *node
}

func (t *tree) insert(val int) *tree {
	node := &node{
		value: val,
	}
	// if tree is empty and root node yet to determined
	if t.item == nil {
		t.item = node
	} else {
		t.item.insert(val)
	}
	return t
}

func (n *node) insert(value int) {
	node := &node{value: value}
	// if condition satisfies then we will move to left side of the tree else right side
	if value <= n.value {
		if n.left == nil {
			n.left = node
		} else {
			n.left.insert(value)
		}
	} else {
		if n.right == nil {
			n.right = node
		} else {
			n.right.insert(value)
		}
	}
}

func pritnNode(n *node) {
	if n == nil {
		return
	}

	fmt.Println(n.value)
	pritnNode(n.left)
	pritnNode(n.right)
}

func main() {
	tree := &tree{}

	tree.insert(10).
		insert(8).
		insert(20).
		insert(9).
		insert(0).
		insert(15)

	pritnNode(tree.item)
}
