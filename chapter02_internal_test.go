package main

import (
	"testing"
)

func TestList(t *testing.T) {
	var l LinkedList
	l.Add(123)
	l.Add(456)
	// l.Print()

	if l.AsString() != "123456" {
		t.Error("List.AsString problem")
	}
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

/*
   Task 2.4
*/
func TestListDigistSum(t *testing.T) {
	var l1, l2 LinkedList
	l1.Add(5)
	l1.Add(1)

	l2.Add(7)
	l2.Add(8)
	l2.Add(1)

	r := ListDigitsSum(&l1, &l2)
	if r.AsString() != "202" {
		t.Error("Incorrect sum")
	}

	var l3, l4 LinkedList
	l3.Add(5)
	l3.Add(7)
	l3.Add(9)

	l4.Add(9)
	l4.Add(8)
	l4.Add(0)
	l4.Add(5)

	r = ListDigitsSum(&l3, &l4)
	if r.AsString() != "4606" {
		t.Error("Incorrect sum")
	}

}
