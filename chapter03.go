package main

import (
	"fmt"
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
	top  *StackItem
	size int
}

func (s *Stack) Push(v interface{}) {
	newItem := &StackItem{Value: v, next: s.top}
	if s.top == nil || s.top.min == nil || newItem.Value.(int) < s.top.min.Value.(int) {
		newItem.min = newItem
	} else {
		newItem.min = s.top.min
	}
	s.top = newItem
	s.size++
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
	s.size--
	return item.Value
}

func (s *Stack) AsString() string {
	str := ""
	item := s.top
	if item == nil {
		return ""
	}
	for {
		str = str + fmt.Sprintf("%+v", item.Value)
		if item.next == nil {
			break
		}
		item = item.next
	}
	return str

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

/*
   Task 3.3
*/

type SetOfStacks struct {
	stacks       []*Stack
	size         int
	maxStackSize int
}

func NewSetOfStacks(maxStackSize int) *SetOfStacks {
	sos := new(SetOfStacks)
	sos.maxStackSize = maxStackSize
	sos.stacks = make([]*Stack, 0)
	sos.stacks = append(sos.stacks, new(Stack))
	sos.size++
	return sos
}

func (sos *SetOfStacks) Push(v interface{}) {
	if sos.stacks[sos.size-1].size >= sos.maxStackSize {
		sos.stacks = append(sos.stacks, new(Stack))
		sos.size++
	}
	sos.stacks[sos.size-1].Push(v)
}

func (sos *SetOfStacks) Pop() interface{} {
	item := sos.stacks[sos.size-1].Pop()
	if item == nil && sos.size > 1 {
		sos.size--
		item = sos.Pop()
	}
	return item
}

func (sos *SetOfStacks) PopAt(n int) interface{} {
	if n >= sos.size {
		// Of course it's better to raise an error
		return nil
	}
	return sos.stacks[n].Pop()
}

/*
   Task 3.4
*/
func HanoiMoveN(n int, from, middle, to *Stack) {
	if n >= 1 {
		HanoiMoveN(n-1, from, to, middle)
		to.Push(from.Pop())
		HanoiMoveN(n-1, middle, from, to)
	}
}
