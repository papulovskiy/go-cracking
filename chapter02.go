package main

import (
	"fmt"
)

/*
   Simple implementation of Linked List
*/
type Item struct {
	Value interface{}
	next  *Item
}

type LinkedList struct {
	Head *Item
}

func (l *LinkedList) Add(v interface{}) {
	newItem := &Item{v, nil}
	if l.Head == nil {
		l.Head = newItem
	} else {
		item := l.Head
		for {
			if item.next == nil {
				break
			}
			item = item.next
		}
		item.next = newItem
	}
}

// Just for debug
func (l *LinkedList) Print() {
	item := l.Head
	for {
		fmt.Printf("%+v ", item.Value)
		if item.next == nil {
			break
		}
		item = item.next
	}
	fmt.Println()
}

/*
   Task 2.1
   This implementation uses map to store values.
   Space complexity: O(N).
   Time complexity: O(N).
*/
func (l *LinkedList) RemoveDuplicatesMap() {
	values := make(map[interface{}]bool)
	item := l.Head
	prevItem := l.Head
	for {
		if _, ok := values[item.Value]; ok {
			prevItem.next = item.next
		} else {
			values[item.Value] = true
		}
		if item.next == nil {
			break
		}
		prevItem = item
		item = item.next
	}
}

/*
   Task 2.1
   This implementation iterates over whole list every time.
   Space complexity: O(1).
   Time complexity: O(N²).
*/
func (l *LinkedList) RemoveDuplicates() {
	item := l.Head
	for {
		if item.next == nil {
			break
		}
		sub := item.next
		prevSub := sub
		for {
			// fmt.Printf("%+v %+v\n", item.Value, sub.Value)
			if item.Value == sub.Value {
				prevSub.next = sub.next
			}
			if sub.next == nil {
				break
			}
			if item.Value != sub.Value {
				prevSub = sub
			}
			sub = sub.next
		}
		item = item.next
	}
}
