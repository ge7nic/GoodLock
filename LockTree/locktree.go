package tree

import "fmt"

type LockNode struct {
	key      int
	marked   bool
	parent   *LockNode
	children []*LockNode
}

type LockTree struct {
	lockSet     map[int]int
	currentNode *LockNode
	root        *LockNode
}

func (t *LockTree) PrintLockSet() {
	for k, e := range t.lockSet {
		fmt.Println("Key: ", k, "=>", "Element:", e)
	}
}

// Create a new LockTree.
func New(id int) *LockTree {
	var lockSet = make(map[int]int)
	var children = make([]*LockNode, 0)
	root := LockNode{id, false, nil, children}
	lt := LockTree{lockSet, &root, &root}
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
	} else {
		t.lockSet[lockID]--
	}
}

// Helper Method for the needed Set N. Makes the Set according to n_t.lock == n.lock && n is not below a mark
func makeN(n *LockNode, m *LockNode, set *[]*LockNode) {
	for _, e := range m.children {
		// if it has the same lock as in n, add it to the set
		if e.key == n.key {
			*set = append(*set, e)
		}
		if !e.marked {
			// iterate recursivly through all children e
			makeN(n, e, set)
		}
	}
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

// Checks if node n is above node m.
func (n *LockNode) isAbove(m *LockNode) bool {
	// iterate through all the children of n - If m is one of them, n is above m.
	for _, e := range n.children {
		if e.key == m.key {
			return true
		}
	}
	return false
}

// Outputs an Error message to the user - A Deadlock has been detected.
func (t *LockTree) conflict(n *LockNode, m *LockNode) {
	fmt.Printf("Conflict Detected between Lock %d and Lock %d detected!\n", n.key, m.key)
}

func (t *LockTree) check(n *LockNode, m *LockNode) {
	for _, e := range n.children {
		if m.isAbove(e) {
			t.conflict(e, m)
		} else {
			t.check(e, m) // iterate through
		}
	}
}

func (t *LockTree) analyseThis(node *LockNode, secondTree *LockTree) {
	var n = make([]*LockNode, 0)

	makeN(node, secondTree.root, &n) // Set which contains all Nodes n in LockTree t, for which is n.lock == secondTree.n.lock && n is not below a mark
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
