package main

import (
	// "fmt"
	"math/rand"
	"testing"
	"time"
)

func randInt() int {
	return int(rand.Int31n(100))
}
func TestBinaryTree(t *testing.T) {
	size := 20
	rand.Seed(time.Now().Unix())
	tree := NewBinaryTree(randInt())
	for i := 0; i < size; i++ {
		tree.Insert(randInt())
	}
	tree.Print()
	for i := 0; i < size; i++ {
		// search := randInt()
		// fmt.Printf("%+v %+v %+v\n", i, search, tree.Find(search))
	}

}
