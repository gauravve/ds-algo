package main

import "fmt"

func main() {
	l := LinkedList{
		head:   nil,
		Length: 0,
	}

	l.Push(1)
	l.Push(10)
	l.Push(15)
	l.Push(20)

	l.Insert(2, 0)
	fmt.Println(l.Traverse())

	l.FastReverse()

	fmt.Println(l.Traverse())
}
