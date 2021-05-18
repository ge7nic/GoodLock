package goodmutex

import (
	tree "goodlock/LockTree"
)

type GoodMutex struct {
	l      chan int
	LockID int
}

var id int = 1
var critchan = make(chan int, 1)

func NewMutex() *GoodMutex {
	var ch = make(chan int, 1)
	var gm GoodMutex
	critchan <- 1
	gm = GoodMutex{ch, id}
	id++
	<-critchan
	return &gm
}

func (gm GoodMutex) Lock(tree *tree.LockTree) {
	tree.Lock(gm.LockID)
	gm.l <- 1
}

func (gm GoodMutex) Unlock(tree *tree.LockTree) {
	tree.Unlock(gm.LockID)
	<-gm.l
}
