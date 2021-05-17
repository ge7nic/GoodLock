package main

import (
	"fmt"
	tree "goodlock/LockTree"
)

func example(threadID int, trees chan *tree.LockTree) {
	lt := tree.New(threadID)
	lt.Lock(threadID)
	lt.Unlock(threadID)
	trees <- lt
}

func something(i int, x int, set *[]int) {
	*set = append(*set, i)
	i++
	if x != 0 {
		*set = append(*set, x)
		x--
		something(i, x, set)
	}
}

func main() {
	treeChan := make(chan *tree.LockTree)
	go example(1, treeChan)
	go example(2, treeChan)

	//-----------------------//
	var z = make([]int, 0)

	something(0, 5, &z)

	for _, e := range z {
		fmt.Printf("%d\n", e)
	}
}
