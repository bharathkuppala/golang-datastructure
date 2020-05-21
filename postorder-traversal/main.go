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
	log.Println("Postorder Traversal Implementation (LEFT - RIGHT - ROOT)")

	root := &Tree{1, nil, nil}

	root.Left = &Tree{2, nil, nil}
	root.Right = &Tree{3, nil, nil}
	root.Left.Left = &Tree{4, nil, nil}
	root.Left.Right = &Tree{5, nil, nil}
	// root.Left.Left.Left = &Tree{6, nil, nil}
	// root.Right.Right = &Tree{7, nil, nil}

	root.inorderTraversal(root)
}

func (t *Tree) inorderTraversal(root *Tree) {
	if root == nil {
		return
	}

	t.inorderTraversal(root.Left)
	t.inorderTraversal(root.Right)

	fmt.Printf("%d \t", root.Data)
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
