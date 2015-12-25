package main

import "testing"

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
