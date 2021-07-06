package main

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	l := LinkedList{
		head:   nil,
		Length: 0,
	}

	l.Push(1)
	l.Push(10)
	l.Push(15)
	l.Push(20)

	l.Insert(2, 0)
	response := l.Traverse()
	testSlice := []int{2, 1, 10, 15, 20}
	t.Run("Check Length", func(t *testing.T) {
		if len(response) != len(testSlice) {
			t.Errorf("Lists lengths are different")
		}

	})

	t.Run("Check Values", func(t *testing.T) {
		for i, v := range response {
			if v != testSlice[i] {
				t.Errorf("Lists values are not equal")
			}
		}
	})

	reversedList := l.Reverse()
	reversedSlice := reversedList.Traverse()

	fmt.Println(reversedSlice)
	testReversedSlice := []int{20, 15, 10, 1, 2}

	t.Run("Check Reverse", func(t *testing.T) {
		for i, v := range reversedSlice {
			if v != testReversedSlice[i] {
				t.Errorf("Reversed Lists values are not equal")
			}
		}
	})

	l.FastReverse()

}
