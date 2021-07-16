package stack

import "fmt"

type Element struct {
	next  *Element
	value interface{}
}

type Stack struct {
	Top, Bottom *Element
	Len         int
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
	if stack.Bottom == nil {
		newNode := Init(value)
		stack.Bottom = newNode
		stack.Top = newNode
		stack.Len++
	} else {
		currentNode := stack.Top
		newNode := Init(value)
		stack.Top = newNode
		currentNode.next = newNode
		stack.Len++
	}
}

func (stack *Stack) Traverse() {
	currentNode := stack.Bottom
	for currentNode != nil {
		fmt.Println(currentNode.value)
		currentNode = currentNode.next
	}
}

func (stack *Stack) Pop() (value interface{}) {
	//First entry in the stack
	poppedValue := stack.Top.value
	currentNode := stack.Bottom

	//Run the loop to the Top element keeping the previous element in memory
	for currentNode.next != stack.Top {
		currentNode = currentNode.next
	}

	//Assign Top to the previous element
	stack.Top = currentNode
	currentNode.next = nil

	return poppedValue

}
