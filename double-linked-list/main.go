package main

import "fmt"

type dLinkedList struct {
	head *node
	tail *node
}

type node struct {
	value string
	next  *node
	prev  *node
}

func initdLinkedList() *dLinkedList {
	return &dLinkedList{}
}

func (dll *dLinkedList) push(s string) {
	node := &node{
		value: s,
	}

	if dll.head == nil {
		dll.head = node
	} else {
		dll.tail.next = node
		node.prev = dll.tail
	}

	dll.tail = node
}

func (dll *dLinkedList) traversal() {
	currentNode := dll.tail
	for currentNode != nil {
		fmt.Println(currentNode.value)
		currentNode = currentNode.prev
	}
}

func main() {
	ill := initdLinkedList()
	ill.push("Bharath")
	ill.push("Suumanth")
	ill.push("Cyril")

	ill.traversal()
}
