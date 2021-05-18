package main

import (
	"fmt"
	mutex "goodlock/GoodMutex"
	tree "goodlock/LockTree"
)

func exampleThread(id int, locks []*mutex.GoodMutex, stopchan chan string, treechan chan *tree.LockTree) {
	lockTree := tree.New(id)
	go func() {
		for {
			// Content here!
			locks[0].Lock(lockTree)
			locks[2].Lock(lockTree)
			locks[1].Lock(lockTree)
			locks[1].Unlock(lockTree)
			locks[3].Lock(lockTree)
			locks[3].Unlock(lockTree)
			locks[0].Unlock(lockTree)
			locks[3].Lock(lockTree)
			locks[1].Lock(lockTree)
			locks[2].Lock(lockTree)
			locks[2].Unlock(lockTree)
			locks[1].Unlock(lockTree)
			locks[3].Unlock(lockTree)
		}
	}()
	<-stopchan
	treechan <- lockTree
}

func main() {
	var stop string
	availableLocks := []*mutex.GoodMutex{mutex.NewMutex(), mutex.NewMutex(), mutex.NewMutex(), mutex.NewMutex()}
	var trees []*tree.LockTree
	stopchan := make(chan string)
	treechan := make(chan *tree.LockTree)
	go exampleThread(1, availableLocks, stopchan, treechan)

	fmt.Scanf("%s", &stop)
	go func() {
		for {
			stopchan <- stop
		}
	}()
	trees = append(trees, <-treechan)
	for _, e := range trees {
		e.PrintLockSet()
	}
}
