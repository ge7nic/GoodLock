package main

import (
	mutex "goodlock/GoodMutex"
)

func main() {
	mutexOne := mutex.NewMutex(1)
	mutexOne.Lock()
	mutexTwo := mutex.NewMutex(2)
	mutexTwo.Lock()
}
