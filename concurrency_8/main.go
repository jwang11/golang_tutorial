/*
# https://tour.golang.org/concurrency/8
Exercise: Equivalent Binary Trees
1. Implement the Walk function.

2. Test the Walk function.

The function tree.New(k) constructs a randomly-structured (but always sorted)
 binary tree holding the values k, 2k, 3k, ..., 10k.

Create a new channel ch and kick off the walker:

go Walk(tree.New(1), ch)
Then read and print 10 values from the channel. It should be the numbers 1, 2,
3, ..., 10.

3. Implement the Same function using Walk to determine whether t1 and t2 store
the same values.

4. Test the Same function.

Same(tree.New(1), tree.New(1)) should return true, and Same(tree.New(1), 
tree.New(2)) should return false.

*/
package main

import (
	"golang.org/x/tour/tree"
	"fmt"
	"reflect"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

func GetTree(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1, c2 := make(chan int), make(chan int)
	var a1, a2 []int
	go GetTree(t1, c1)
	for i := range c1 {
		a1 = append(a1, i)
	}
	go GetTree(t2, c2)
	for i := range c2 {
		a2 = append(a2, i)
	}
	fmt.Println(a1, a2)
	return reflect.DeepEqual(a1, a2)
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
