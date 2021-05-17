package mutex

import "fmt"

type GoodMutex struct {
	l  chan int
	id int
}

var id int = 1

func NewMutex() GoodMutex {
	var ch = make(chan int, 1)
	gm := GoodMutex{ch, id}
	id++
	return gm
}

func (gm GoodMutex) Lock() {
	// insert here function for locktree
	gm.l <- 1
}

func (gm GoodMutex) Unlock() {
	// insert here function for locktree
	<-gm.l
}

func (gm GoodMutex) PrintID() {
	fmt.Printf("Hello %d!\n", gm.id)
}
