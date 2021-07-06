package main

import "fmt"

type Element struct {
	next  *Element
	value interface{}
}

type Stack struct {
	top, bottom *Element
	len         int
}

func Init(value interface{}) (e *Element) {
	e = &Element{
		next:  nil,
		value: value,
	}
	return
}

func (stack *Stack) Push(value interface{}) {
	//First entry in the stack
	if stack.bottom == nil {
		newNode := Init(value)
		stack.bottom = newNode
		stack.top = newNode
		stack.len++
	} else {
		currentNode := stack.top
		newNode := Init(value)
		stack.top = newNode
		currentNode.next = newNode
		stack.len++
	}
}

func (stack *Stack) Traverse() {
	currentNode := stack.bottom
	for currentNode != nil {
		fmt.Println(currentNode.value)
		currentNode = currentNode.next
	}
}

func (stack *Stack) Pop() (value interface{}) {
	//First entry in the stack
	poppedValue := stack.top.value
	currentNode := stack.bottom

	//Run the loop to the top element keeping the previous element in memory
	for currentNode.next != stack.top {
		currentNode = currentNode.next
	}

	//Assign top to the previous element
	stack.top = currentNode
	currentNode.next = nil

	return poppedValue

}
