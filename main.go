package main

import (
	mutex "goodlock/GoodLock"
	tree "goodlock/LockTree"
)

func test() {
	tree.Test()
}

func main() {
	mutexOne := mutex.NewMutex()
	mutexTwo := mutex.NewMutex()

	mutexOne.PrintID()
	mutexTwo.PrintID()

	test()
}
