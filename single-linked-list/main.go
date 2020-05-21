package main

import "fmt"

type node struct {
	value string
	next  *node
}

type linkedList struct {
	head *node
	tail *node
}

func initList() *linkedList {
	return &linkedList{}
}

func (l *linkedList) addValue(s string) {
	node := &node{
		value: s,
	}

	// no elements in the list
	if l.head == nil {
		l.head = node
	} else {
		l.tail.next = node
	}
	// this will point to the last element in the node
	l.tail = node
}

func (l *linkedList) removeFront() {
	if l.head == nil {
		fmt.Println("list is empty")
		return
	}
}

func main() {
	ll := initList()
	ll.addValue("Bharath")
	ll.addValue("suumanth")
	ll.addValue("abhilash")
	ll.addValue("hema")

	current := ll.head
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
}
