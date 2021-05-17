package main

import (
	tree "goodlock/LockTree"
)

func example(threadID int, trees chan *tree.LockTree) {
	lt := tree.New(threadID)
	lt.Lock(threadID)
	lt.Unlock(threadID)
	trees <- lt
}

func main() {
	treeChan := make(chan *tree.LockTree)
	go example(1, treeChan)
	go example(2, treeChan)
}
