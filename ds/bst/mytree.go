//This tree uses recursion to Insert values vs the solution provided. This looks like hte more elegant solution.

package bst

import (
	"fmt"
	"io"
)

type Node struct {
	Left, Right *Node
	Value       int
}

type Tree struct {
	Root *Node
}

func InitNode(value int) (n *Node) {
	n = &Node{
		Left:  nil,
		Right: nil,
		Value: value,
	}
	return
}

func (t *Tree) Insert(value int) {
	if t.Root == nil {
		t.Root = InitNode(value)
	} else {
		t.Root.Insert(value)
	}
	return
}

func (n *Node) Insert(value int) {
	if n == nil {
		return
	} else if value <= n.Value {
		if n.Left == nil {
			n.Left = InitNode(value)
		} else {
			n.Left.Insert(value)
		}
	} else {
		if n.Right == nil {
			n.Right = InitNode(value)
		} else {
			n.Right.Insert(value)
		}
	}
}

func (n *Node) PrintTree(w io.Writer, ns int, ch rune) {
	if n == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, n.Value)
	n.Left.PrintTree(w, ns+2, 'L')
	n.Right.PrintTree(w, ns+2, 'R')
}

func (t *Tree) Lookup(value int) bool {
	if t.Root == nil {
		return false
	} else {
		if t.Root.Value == value {
			return true
		} else {
			return t.Root.Lookup(value)
		}
	}
}

func (n *Node) Lookup(value int) bool {
	if n == nil {
		return false
	} else {
		if n.Value == value {
			return true
		} else {
			if value < n.Value {
				if n.Left == nil {
					return false
				} else {
					return n.Left.Lookup(value)
				}
			} else {
				if n.Right == nil {
					return false
				} else {
					return n.Right.Lookup(value)
				}
			}
		}
	}
}
