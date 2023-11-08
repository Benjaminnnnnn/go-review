package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	defer close(ch) // <- closes the channel when this function returns
	var walk func(t *tree.Tree)
	walk = func(t *tree.Tree) {
		if t == nil {
			return
		}
		walk(t.Left)
		ch <- t.Value
		walk(t.Right)
	}
	walk(t)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if v1 != v2 || ok1 != ok2 {
			return false
		}

		if !ok1 {
			break
		}
	}
	return true
}

func test(n1, n2 int) {
	fmt.Printf("Testing tree with %v & %v values\n", n1, n2)
	t1 := tree.New(n1)
	t2 := tree.New(n2)

	fmt.Println(Same(t1, t2))
	fmt.Println("===============")
}

func main() {
	test(1, 1)
	test(5, 5)
	test(5, 2)
	// ch := make(chan int)

	// go func() {
	// 	Walk(tree.New(1), ch)
	// 	close(ch)
	// }()
	// for v, ok := <-ch; ok; v, ok = <-ch {
	// 	fmt.Println(v)
	// }
}
