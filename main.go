package main

import (
	mutex "goodlock/GoodMutex"
)

func main() {
	mutexOne := mutex.NewMutex()
	mutexTwo := mutex.NewMutex()

	mutexOne.PrintID()
	mutexTwo.PrintID()
}
