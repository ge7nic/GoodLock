package main

import (
	tree "goodlock/LockTree"
)

func example(threadID int) {
	lt := tree.New(threadID)
	lt.Lock(1)
	lt.Unlock(1)
}

func main() {
	example(1)
}
