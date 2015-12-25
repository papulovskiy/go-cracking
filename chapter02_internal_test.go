package main

import (
	"testing"
)

func TestList(t *testing.T) {
	var l LinkedList
	l.Add(123)
	l.Add(456)
	// l.Print()
}

func listSize(l *LinkedList) int {
	i := 0
	item := l.Head
	for {
		i++
		if item.next == nil {
			break
		}
		item = item.next
	}
	return i
}

/*
   Task 2.1
*/
func TestRemoveDuplicatesMap(t *testing.T) {
	var l LinkedList
	l.Add(1)
	l.Add(2)
	l.Add(3)

	if listSize(&l) != 3 {
		t.Error("Incorrect list size")
	}

	l.Add(4)
	l.Add(2)
	l.Add(5)
	l.Add(1)

	if listSize(&l) != 7 {
		t.Error("Incorrect list size")
	}

	l.RemoveDuplicatesMap()
	if listSize(&l) != 5 {
		t.Error("Incorrect list size")
	}
}

func TestRemoveDuplicates(t *testing.T) {
	var l LinkedList
	l.Add(1)
	l.Add(2)
	l.Add(3)

	if listSize(&l) != 3 {
		t.Error("Incorrect list size")
	}

	l.Add(4)
	l.Add(2)
	l.Add(2)
	l.Add(5)
	l.Add(1)

	if listSize(&l) != 8 {
		t.Error("Incorrect list size")
	}

	l.RemoveDuplicates()
	if listSize(&l) != 5 {
		t.Error("Incorrect list size")
	}
}

/*
   Task 2.2
*/
func TestNthToLast(t *testing.T) {
	var l LinkedList
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)
	l.Add(5)

	if l.NthToLast(1) != 4 {
		t.Error("Incorrect nth to last element")
	}
	if l.NthToLast(2) != 3 {
		t.Error("Incorrect nth to last element")
	}
	if l.NthToLast(0) != 5 {
		t.Error("Incorrect nth to last element")
	}
}

/*
   Task 2.3
*/
func TestDeleteItem(t *testing.T) {
	var l LinkedList
	l.Add(1)
	m := l.Add(2)
	l.Add(3)
	l.Add(4)

	DeleteItem(m)
	if listSize(&l) != 3 {
		t.Error("Incorrect list size")
	}

	l.Add(5)
	m = l.Add(6)
	if listSize(&l) != 5 {
		t.Error("Incorrect list size")
	}

	// It's better to do proper testing here, but it requires aditional
	// methods to be implemeted in LinkedList.
}
