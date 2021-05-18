package main

import (
	"fmt"
	tree "goodlock/LockTree"
)

type test struct {
	key  int
	mark bool
}

func example(threadID int, trees chan *tree.LockTree) {
	lt := tree.New(threadID)
	lt.Lock(threadID)
	lt.Unlock(threadID)
	trees <- lt
}

func something(x int, set *[]*test, flag bool) {
	if x != 0 {
		*set = append(*set, &test{x, flag})
		x = x - 1
		if flag {
			flag = false
		} else {
			flag = true
		}
		something(x, set, flag)
	}
}

func swap(set []*test) {
	for _, e := range set {
		if e.mark {
			e.mark = false
		} else {
			e.mark = true
		}
	}
}

func main() {
	/*treeChan := make(chan *tree.LockTree)
	go example(1, treeChan)
	go example(2, treeChan)*/

	//-----------------------//
	var z = make([]*test, 0)

	something(5, &z, true)

	for _, e := range z {
		fmt.Printf("%t\n", e.mark)
	}
	swap(z)
}
