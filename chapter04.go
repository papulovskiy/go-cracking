package main

import (
	"fmt"
)

/*
   Basic data structures
*/
type BinaryTree struct {
	Value interface{}
	Left  *BinaryTree
	Right *BinaryTree
}

func NewBinaryTree(v interface{}) *BinaryTree {
	return &BinaryTree{Value: v}
}

func (t *BinaryTree) Insert(v interface{}) {
	if v.(int) < t.Value.(int) {
		if t.Left == nil {
			t.Left = NewBinaryTree(v)
		} else {
			t.Left.Insert(v)
		}
	} else if v.(int) > t.Value.(int) {
		if t.Right == nil {
			t.Right = NewBinaryTree(v)
		} else {
			t.Right.Insert(v)
		}
	}
}

func (t *BinaryTree) Find(v interface{}) *BinaryTree {
	if t.Value == v {
		return t
	}
	if t.Left != nil {
		if left := t.Left.Find(v); left != nil {
			return left
		}
	}
	if t.Right != nil {
		if right := t.Right.Find(v); right != nil {
			return right
		}
	}
	return nil
}

func (t *BinaryTree) PrintSub(level, offset int) {
	if t.Right != nil {
		t.Right.PrintSub(level+1, offset)
	}
	for i := 0; i < level; i++ {
		fmt.Print("    ")
	}
	fmt.Printf("%+v\n", t.Value)
	if t.Left != nil {
		t.Left.PrintSub(level+1, offset)
	}
}
func (t *BinaryTree) Print() {
	t.PrintSub(0, 0)
}
