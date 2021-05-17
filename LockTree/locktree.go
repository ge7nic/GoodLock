package locktree

import "fmt"

type LockNode struct {
	key    int
	marked bool
	left   *LockNode
	right  *LockNode
	parent *LockNode
}

type LockTree struct {
	lockSet     map[int]*LockNode
	id          int
	currentNode *LockNode
	root        *LockNode
}

func Test() {
	fmt.Printf("Hallo c:")
}
