package main

import (
	"testing"
)

func TestList(t *testing.T) {
	var l LinkedList
	l.Add(123)
	l.Add(456)
	l.Print()
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
