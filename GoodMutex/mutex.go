package mutex

type GoodMutex struct {
	l  chan int
	id int
}

func NewMutex(id int) GoodMutex {
	var ch = make(chan int, 1)
	gm := GoodMutex{ch, id}
	return gm
}

func (gm GoodMutex) Lock() {
	gm.l <- 1
}

func (gm GoodMutex) Unlock() {
	<-gm.l
}
