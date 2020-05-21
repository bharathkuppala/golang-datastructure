package main

import (
	"fmt"
	"log"
)

type (
	// Tree ...
	Tree struct {
		Data  int
		Left  *Tree
		Right *Tree
	}
)

func main() {
	log.SetFlags(log.Lshortfile)
	log.Println("Inorder Traversal Implementation (LEFT - ROOT - RIGHT)")

	root := &Tree{10, nil, nil}

	root.Left = &Tree{20, nil, nil}
	root.Right = &Tree{30, nil, nil}
	root.Left.Left = &Tree{40, nil, nil}
	root.Left.Right = &Tree{50, nil, nil}
	root.Left.Left.Left = &Tree{70, nil, nil}
	root.Right.Right = &Tree{60, nil, nil}

	root.inorderTraversal(root)
}

func (t *Tree) inorderTraversal(root *Tree) {
	if root == nil {
		return
	}
	t.inorderTraversal(root.Left)
	fmt.Printf("%d \t", root.Data)
	t.inorderTraversal(root.Right)
}

/*
			10
		   /  \
		 20	   30
		/ \      \
	   40  50     60
	  /
	 70
*/
