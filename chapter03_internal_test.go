package main

import "testing"

// import "fmt"

/*
   Tests for basic data structures
*/
func TestStack(t *testing.T) {
	var s Stack

	if s.Pop() != nil {
		t.Error("Stack is not empty")
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	if s.Pop() != 3 {
		t.Error("Incorrect stack order")
	}
	if s.Pop() != 2 {
		t.Error("Incorrect stack order")
	}
	if s.Pop() != 1 {
		t.Error("Incorrect stack order")
	}

	if s.Pop() != nil {
		t.Error("Stack is not empty")
	}
}

func TestQueue(t *testing.T) {
	var q Queue

	if q.Dequeue() != nil {
		t.Error("Queue is not empty")
	}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if q.Dequeue() != 1 {
		t.Error("Incorrect queue order")
	}
	if q.Dequeue() != 2 {
		t.Error("Incorrect queue order")
	}
	if q.Dequeue() != 3 {
		t.Error("Incorrect queue order")
	}

	if q.Dequeue() != nil {
		t.Error("Queue is not empty")
	}
}

/*
	Task 3.2
*/
func TestStackMin(t *testing.T) {
	var s Stack

	if s.Pop() != nil {
		t.Error("Stack is not empty")
	}

	s.Push(10)
	s.Push(20)
	s.Push(30)
	if s.top.min.Value != 10 {
		t.Error("Incorrect min value")
	}

	s.Push(5)
	if s.top.min.Value != 5 {
		t.Error("Incorrect min value")
	}

	s.Pop()
	if s.top.min.Value != 10 {
		t.Error("Incorrect min value")
	}
}

/*
	Task 3.3
*/
func TestSetOfStacks(t *testing.T) {
	s := NewSetOfStacks(3)

	if s.Pop() != nil {
		t.Error("Stack is not empty")
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	s.Push(6)
	s.Push(7)
	s.Push(8)
	s.Push(9)

	if s.Pop() != 9 {
		t.Error("Incorrect stack order")
	}
	s.Pop()
	s.Pop()
	if s.Pop() != 6 {
		t.Error("Incorrect stack order")
	}
	s.Pop()
	s.Pop()
	if s.Pop() != 3 {
		t.Error("Incorrect stack order")
	}
	s.Pop()
	if s.Pop() != 1 {
		t.Error("Incorrect stack order")
	}

	if s.Pop() != nil {
		t.Error("Stack is not empty")
	}
}

func TestSetOfStacksPopAt(t *testing.T) {
	s := NewSetOfStacks(2)

	if s.Pop() != nil {
		t.Error("Stack is not empty")
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	if s.PopAt(0) != 2 {
		t.Error("Incorrect stack order")
	}
	if s.PopAt(1) != 4 {
		t.Error("Incorrect stack order")
	}
}

/*
	Task 3.4
*/
func TestHanoi(t *testing.T) {
	size := 5
	var s1, s2, s3 Stack
	for i := 0; i < size; i++ {
		s1.Push(size - i)
	}

	HanoiMoveN(size, &s1, &s2, &s3)
	if s1.size != 0 {
		t.Error("Not all elements were moved")
	}
	if s2.size != 0 {
		t.Error("Not all elements were moved")
	}
	if s3.size != size {
		t.Error("Not all elements were moved")
	}
}
