package binaryTreeComparison

import (
	"golang.org/x/tour/tree"
)

func walker(t *tree.Tree, ch chan int) {
	walk(t, ch)
	close(ch)
}

func walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	walk(t.Left, ch)
	ch <- t.Value
	walk(t.Right, ch)
}

func sync(ch1, ch2 chan int, done chan bool) {
	for v1 := range ch1 {
		v2 := <-ch2
		if v1 != v2 {
			done <- false
			break
		}
	}
	done <- true
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	done := make(chan bool)

	go walker(t1, ch1)
	go walker(t2, ch2)
	go sync(ch1, ch2, done)

	return <-done
}
