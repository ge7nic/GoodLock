package main

import (
	"fmt"
	mutex "goodlock/GoodMutex"
	tree "goodlock/LockTree"
)

type algo struct {
	trees *[]*tree.LockTree
}

func New(trees *[]*tree.LockTree) *algo {
	return &algo{trees}
}

func (a *algo) Analyze() {
	for i, e := range *a.trees {
		fmt.Printf("%d: %v\n", i, e)
	}

}

func exampleThread1(id int, locks []*mutex.GoodMutex, stopchan chan string, treechan chan *tree.LockTree) {
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

func exampleThread2(id int, locks []*mutex.GoodMutex, stopchan chan string, treechan chan *tree.LockTree) {
	lockTree := tree.New(id)
	go func() {
		for {
			// Content here!
			locks[0].Lock(lockTree)
			locks[1].Lock(lockTree)
			locks[2].Lock(lockTree)
			locks[2].Unlock(lockTree)
			locks[1].Unlock(lockTree)
			locks[0].Unlock(lockTree)
			locks[3].Lock(lockTree)
			locks[2].Lock(lockTree)
			locks[1].Lock(lockTree)
			locks[1].Unlock(lockTree)
			locks[2].Unlock(lockTree)
			locks[3].Unlock(lockTree)
		}
	}()
	<-stopchan
	treechan <- lockTree
}

// TODO: Fix. Help.
func main() {
	var stop string
	availableLocks := []*mutex.GoodMutex{mutex.NewMutex(), mutex.NewMutex(), mutex.NewMutex(), mutex.NewMutex()}
	var trees []*tree.LockTree
	stopchan := make(chan string)
	treechan := make(chan *tree.LockTree)
	go exampleThread1(1, availableLocks, stopchan, treechan)
	go exampleThread2(2, availableLocks, stopchan, treechan)

	fmt.Scanf("%s", &stop)
	go func() {
		for {
			stopchan <- stop
		}
	}()
	algo := New(&trees)
	trees = append(trees, <-treechan)
	trees = append(trees, <-treechan)
	algo.Analyze()
}
