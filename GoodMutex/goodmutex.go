package goodmutex

import (
	tree "goodlock/LockTree"
)

type GoodMutex struct {
	l      chan int
	lockID int
}

var id int = 1

func NewMutex() *GoodMutex {
	var ch = make(chan int, 1)
	gm := GoodMutex{ch, id}
	id++
	return &gm
}

func (gm GoodMutex) Lock(tree *tree.LockTree) {
	tree.Lock(gm.lockID)
	gm.l <- 1
}

func (gm GoodMutex) Unlock(tree *tree.LockTree) {
	tree.Unlock(gm.lockID)
	<-gm.l
}
