package ost

import (
	"fmt"
)

type side int

const (
	left  = side(-1)
	right = side(+1)
)

type Item interface {
	Less(item Item) bool
	Greater(item Item) bool
	Equal(item Item) bool
	Key() int
}

type Node struct {
	Items       []Item
	count       int
	height      int
	Left, Right *Node
}

func (n *Node) firstItem() Item {
	return n.Items[0]
}

func newNode(item Item) *Node {
	return &Node{Items: []Item{item}, count: 1, height: 1}
}

func (n *Node) append(item Item) {
	if item.Less(n.firstItem()) {
		if n.Left == nil {
			n.Left = newNode(item)
			// n.count++
		} else {
			n.Left.append(item)
			// n.count++
		}
	} else if item.Greater(n.firstItem()) {
		if n.Right == nil {
			n.Right = newNode(item)
			// n.count++
		} else {
			n.Right.append(item)
			// n.count++
		}
	} else {
		n.Items = append(n.Items, item)
		// n.count++
	}
	n.updateHeight()
	n.count++
	n.rebalance()
}

func (n *Node) getHeight() int {
	if n == nil {
		return 0
	}
	return n.height
}

func (n *Node) updateHeight() {
	n.height = 1 + max(n.Left.getHeight(), n.Right.getHeight())
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (n *Node) getBalance() int {
	if n == nil {
		return 0
	}
	return n.Left.getHeight() - n.Right.getHeight()
}

func (n *Node) rebalance() {
	balance := n.getBalance()
	// Left left case
	if balance > 1 && n.Left.getBalance() >= 0 {
		n.rightRotate()
	}
	// Right Right case
	if balance < -1 && n.Right.getBalance() <= 0 {
		n.leftRotate()
	}
	// Left Right case
	if balance > 1 && n.Left.getBalance() < 0 {
		n.Left.leftRotate()
		n.rightRotate()
	}
	// Right Left case
	if balance < -1 && n.Right.getBalance() > 0 {
		n.Right.rightRotate()
		n.leftRotate()
	}
}

func (n *Node) exchange(n2 *Node) {
	n.Items, n2.Items = n2.Items, n.Items
}

func (n *Node) getCount() int {
	if n == nil {
		return 0
	}
	return n.count
}

func (n *Node) recalculate() {
	n.updateHeight()
	n.count = len(n.Items) + n.Left.getCount() + n.Right.getCount()
}

func (n *Node) rightRotate() {
	right := n.Left
	n.exchange(right)

	n.Left = right.Left
	right.Left = right.Right
	right.Right = n.Right
	n.Right = right

	right.recalculate()
	n.recalculate()
}

func (n *Node) leftRotate() {
	left := n.Right
	n.exchange(left)

	n.Right = left.Right
	left.Right = left.Left
	left.Left = n.Left
	n.Left = left

	left.recalculate()
	n.recalculate()
}

func (n *Node) removeItem(item Item) (bool, *Node) {
	var newRoot *Node
	if ok, pos := n.include(item); ok {
		n.Items = append(n.Items[:pos], n.Items[pos+1:]...)
	}
	if len(n.Items) == 0 {
		var rootSide side
		if n.Left.getHeight() > n.Right.getHeight() {
			rootSide = left
		} else {
			rootSide = right
		}
		switch rootSide {
		case left:
			newRoot = n.Left.deleteMaximum()
		case right:
			newRoot = n.Right.deleteMinimum()
		}
		if newRoot != nil {
			if !n.Left.equal(newRoot) {
				newRoot.Left = n.Left
			}
			if !n.Right.equal(newRoot) {
				newRoot.Right = n.Right
			}
			newRoot.recalculate()
			newRoot.rebalance()
		}
		return true, newRoot
	}
	return false, nil
}

func (n *Node) remove(item Item) (isDel bool, subRoot *Node) {
	if n == nil {
		return false, nil
	}

	if n.firstItem().Greater(item) {
		if removed, leftRoot := n.Left.remove(item); removed {
			isDel = removed
			n.Left = leftRoot
		}
	} else if n.firstItem().Less(item) {
		if removed, rightRoot := n.Right.remove(item); removed {
			isDel = removed
			n.Right = rightRoot
		}
	} else {
		return n.removeItem(item)
	}

	if isDel {
		n.recalculate()
		n.rebalance()
	}

	return isDel, n
}

func (n *Node) equal(other *Node) bool {
	if n != nil && other != nil && n.firstItem().Equal(other.firstItem()) {
		return true
	}
	return false
}

func (n *Node) include(item Item) (bool, int) {
	for i := range n.Items {
		if n.Items[i].Key() == item.Key() {
			return true, i
		}
	}
	return false, -1
}

// minimum returns the minimum node of subtree
func (n *Node) minimum() *Node {
	if n.Left != nil {
		return n.Left.minimum()
	}
	return n
}

// deleteMinimum removes the minimum node of subtree
func (n *Node) deleteMinimum() *Node {
	if n == nil {
		return nil
	}

	if n.Left != nil {
		defer n.recalculate()
		if n.Left.Left == nil {
			minimum := n.Left
			n.Left = nil
			return minimum
		}
		return n.Left.deleteMinimum()
	}
	return n
}

// minimum returns the maximum node of subtree
func (n *Node) maximum() *Node {
	if n.Right != nil {
		return n.Right.maximum()
	}
	return n
}

// deleteMaximum removes the maximum node of subtree
func (n *Node) deleteMaximum() *Node {
	if n == nil {
		return nil
	}

	if n.Right != nil {
		defer n.recalculate()
		// n.count--
		if n.Right.Right == nil {
			maximum := n.Right
			n.Right = nil
			return maximum
		}
		return n.Right.deleteMaximum()
	}
	return n
}

func (n *Node) rank(item Item) int {
	p := n
	count := 0

	for {
		if ok, _ := p.include(item); ok {
			count++
			if p.Left != nil {
				count += p.Left.count
			}
			return count
		}
		if item.Less(p.firstItem()) {
			if p.Left == nil {
				return -1
			}
			p = p.Left
			continue
		}
		if item.Greater(p.firstItem()) {
			if p.Right == nil {
				return -1
			}
			count = count + p.count - p.Right.count
			p = p.Right
			continue
		}
	}
}

func (n *Node) findByRank(rank int) []Item {
	leftcount := 0
	if n.Left != nil {
		leftcount = n.Left.count
	}

	if rank <= leftcount {
		return n.Left.findByRank(rank)
	} else if rank == (leftcount + 1) {
		return n.Items
	} else if rank > (leftcount + 1) {
		return n.Right.findByRank(rank - leftcount - 1)
	}
	return nil
}

func (n *Node) PrettyPrint() {
	height := n.getHeight()
	lineNum := 1<<uint(height) - 1
	var s [][]string
	for i := 0; i < height; i++ {
		var l []string
		for j := 0; j < lineNum; j++ {
			l = append(l, "\"\"")
		}
		s = append(s, l)
	}

	var helper func(node *Node, d, pos int)
	helper = func(node *Node, d, pos int) {
		if node == nil {
			return
		}
		s[d-1][pos] = fmt.Sprintf("%2.1v", node.firstItem())
		helper(node.Left, d+1, pos-(1<<uint(height-d-1)))
		helper(node.Right, d+1, pos+(1<<uint(height-d-1)))
	}
	helper(n, 1, (1<<uint(height-1))-1)
	for i := 0; i < height; i++ {
		fmt.Println(s[i])
	}
}

type OST struct {
	count int
	root  *Node
}

func New() *OST {
	return new(OST)
}

func (t *OST) Insert(item Item) {
	if t.root == nil {
		t.root = newNode(item)
		return
	}
	t.root.append(item)
}

func (t *OST) Delete(item Item) {
	if t.root == nil {
		return
	}

	if isDel, newRoot := t.root.remove(item); isDel && newRoot != nil {
		t.root = newRoot
		t.root.rebalance()
	}
}

func (t *OST) Rank(item Item) int {
	if t.root == nil {
		return -1
	}
	return t.root.rank(item)
}

func (t *OST) FindByRank(rank int) []Item {
	if t.root == nil {
		return nil
	}
	if t.root.count < rank {
		return nil
	}
	return t.root.findByRank(rank)
}

func (t *OST) Height() int {
	if t.root == nil {
		return 0
	}

	return t.root.getHeight()
}

func (t *OST) PrettyPrint() {
	if t.root == nil {
		return
	}
	t.root.PrettyPrint()
}
