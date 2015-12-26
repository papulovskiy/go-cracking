package main

import (
// "fmt"
)

/*
   Basic implementation of stack
*/
type StackItem struct {
	Value interface{}
	next  *StackItem
	min   *StackItem
}

type Stack struct {
	top *StackItem
}

func (s *Stack) Push(v interface{}) {
	newItem := &StackItem{Value: v, next: s.top}
	if s.top == nil || s.top.min == nil || newItem.Value.(int) < s.top.min.Value.(int) {
		newItem.min = newItem
	} else {
		newItem.min = s.top.min
	}
	s.top = newItem
}

func (s *Stack) Pop() interface{} {
	if s.top == nil {
		return nil
	}
	item := s.top
	if item.next == nil {
		s.top = nil
	} else {
		s.top = item.next
	}
	return item.Value
}

type QueueItem struct {
	Value interface{}
	next  *QueueItem
}

type Queue struct {
	first, last *QueueItem
}

func (q *Queue) Enqueue(v interface{}) {
	newItem := &QueueItem{Value: v, next: nil}
	if q.last != nil {
		q.last.next = newItem
	}
	if q.first == nil {
		q.first = newItem
	}
	q.last = newItem
}

func (q *Queue) Dequeue() interface{} {
	if q.first == nil {
		return nil
	}
	item := q.first
	q.first = item.next
	return item.Value
}

/*
   Task 3.1
   For three stacks implementation based on just on array, I'd go with
   data structure that stores reference to the previous element of this stack
   and additional three variables should be referencing to three stack's
   top elements.
*/

/*
   Task 3.2
   Implemented in Stack itself - stores reference to min elemenent of every push,
   so in case of pop there is always the previous value.
*/
