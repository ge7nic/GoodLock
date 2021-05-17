package locktree

import (
	"fmt"
)

type LockNode struct {
	key      int
	marked   bool
	parent   *LockNode
	children []*LockNode
}

type LockTree struct {
	lockSet     map[int]int
	id          int // Propably unneccesary, due to the fact that root.key should be the same thing. Possibly REMOVE.
	currentNode *LockNode
	root        *LockNode
}

var treeList []LockTree // propably also not neccesary, REMOVE LATER

// Create a new LockTree.
func New(id int) *LockTree {
	var lockSet = make(map[int]int)
	var children = make([]*LockNode, 0)
	root := LockNode{id, false, nil, children}

	lt := LockTree{lockSet, id, &root, &root}
	treeList = append(treeList, lt)
	return &lt
}

// Method to check if the splice children has a child with this lockID. Helper function.
func (n *LockNode) hasChild(lockID int) (int, bool) {
	for i, e := range n.children {
		if e.key == lockID {
			return i, true
		}
	}
	return -1, false
}

func (t *LockTree) Lock(lockID int) {
	_, exists := t.lockSet[lockID] // Check if this Thread already owns this lock
	if !exists {
		// if it doesnt, add it to it's lockSet with a value of 0 and check if this lock is a son of current
		t.lockSet[lockID] = 0
		i, exists := t.currentNode.hasChild(lockID) // Check if current has this lockID as a child
		if exists {
			// if it is indeed a child of current
			t.currentNode = t.currentNode.children[i]
		} else {
			// it isn't a child of current, so make a new child and append it as a child of current - this is a new pattern
			newCh := &LockNode{t.root.key, false, t.currentNode, make([]*LockNode, 0)}
			t.currentNode.children = append(t.currentNode.children, newCh)
			t.currentNode = newCh
			fmt.Printf("Found a new Pattern!\n")
		}
	} else {
		// if it already has this lock, update its counter
		t.lockSet[lockID]++
	}
	// fmt.Printf("Key with the ID %d is used %d times.\n", lockID, t.lockSet[lockID])
}

func (t *LockTree) Unlock(lockID int) {
	// Get the counter from this lockID - If it is 0, it is only used once and can be deleted. If not, reduce the counter by one.
	counter := t.lockSet[lockID]
	if counter == 0 {
		t.currentNode = t.currentNode.parent
		delete(t.lockSet, lockID)
		fmt.Printf("Unlocked Lock with ID %d\n", lockID)
	} else {
		t.lockSet[lockID]--
	}
}

func makeN(n *LockNode, m *LockNode) {
	// Implement this recursivly. Add Splice to Parameters. TODO
}

// Helper Method for the needed Set N. Swaps the value of marked of every element from true to false and vice versa.
func swapMark(set []*LockNode) {
	for _, e := range set {
		if e.marked {
			e.marked = false
		} else {
			e.marked = true
		}
	}
}

func (t *LockTree) check(n *LockNode, m *LockNode) {
	// implement this
}

// TODO: FIX METHOD
func (t *LockTree) analyseThis(node *LockNode, secondTree *LockTree) {
	makeN(node, secondTree.root) // Set which contains all Nodes n in LockTree t, for which is n.lock == secondTree.n.lock && n is not below a mark
	for _, e := range n {
		t.check(node, e)
	}
	// Mark all Nodes in N
	swapMark(n)
	for _, child := range node.children {
		t.analyseThis(child, secondTree)
	}
	// Unmark all Nodes in N
	swapMark(n)
}

func (t *LockTree) Analyse(secondTree *LockTree) {
	fmt.Println("----------Analyse------------")
	for _, e := range t.root.children {
		t.analyseThis(e, secondTree)
	}
}
