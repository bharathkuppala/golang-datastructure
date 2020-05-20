package main

import (
	"errors"
	"fmt"
	"sync"
)

// Step1 customStack

type customQueue struct {
	Queue []string
	mutex sync.RWMutex
}

// Step2 push ops
func (c *customQueue) enqueue(s string) {
	c.Queue = append(c.Queue, s)
}

// Step2 push ops
func (c *customQueue) dequeue() error {
	len := len(c.Queue)
	if len > 0 {
		c.Queue = c.Queue[1:]
		return nil
	}
	return errors.New("cannot pop because stack is empty")
}

// Step2 push ops
func (c *customQueue) frontEle() (string, error) {
	len := len(c.Queue)
	if len > 0 {
		return c.Queue[0], nil
	}

	return "", errors.New(("stack is empty"))
}

func (c *customQueue) emptyStack() bool {
	return len(c.Queue) == 0
}

func (c *customQueue) size() int {
	return len(c.Queue)
}

func main() {
	customQueue := &customQueue{
		Queue: make([]string, 0),
	}
	fmt.Printf("enqueue: A\n")
	customQueue.enqueue("A")
	fmt.Printf("emqueue: B\n")
	customQueue.enqueue("B")
	fmt.Printf("Size: %d\n", customQueue.size())
	for customQueue.size() > 0 {
		frontVal, _ := customQueue.frontEle()
		fmt.Printf("Front: %s\n", frontVal)
		fmt.Printf("dequeue: %s\n", frontVal)
		customQueue.dequeue()
		fmt.Println(customQueue.size())
	}
	fmt.Printf("Size: %d\n", customQueue.size())
}
