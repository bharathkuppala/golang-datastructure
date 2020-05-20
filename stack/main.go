package main

import (
	"errors"
	"fmt"
	"sync"
)

// Step1 customStack

type customStack struct {
	Stack []string
	mutex sync.RWMutex
}

// Step2 push ops
func (c *customStack) push(s string) {
	c.Stack = append(c.Stack, s)
}

// Step2 push ops
func (c *customStack) pop() error {
	len := len(c.Stack)
	if len > 0 {
		c.Stack = c.Stack[:len-1]
		return nil
	}
	return errors.New("cannot pop because stack is empty")
}

// Step2 push ops
func (c *customStack) frontEle() (string, error) {
	len := len(c.Stack) - 1
	if len > 0 {
		return c.Stack[len], nil
	}

	return "", errors.New(("stack is empty"))
}

func (c *customStack) emptyStack() bool {
	return len(c.Stack) == 0
}

func (c *customStack) size() int {
	return len(c.Stack)
}

func main() {
	customStack := &customStack{
		Stack: make([]string, 0),
	}
	fmt.Printf("Push: A\n")
	customStack.push("A")
	fmt.Printf("Push: B\n")
	customStack.push("B")
	fmt.Printf("Size: %d\n", customStack.size())
	for customStack.size() > 0 {
		frontVal, _ := customStack.frontEle()
		fmt.Printf("Front: %s\n", frontVal)
		fmt.Printf("Pop: %s\n", frontVal)
		customStack.pop()
	}
	fmt.Printf("Size: %d\n", customStack.size())
}
