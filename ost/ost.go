package ost

import (
	"errors"
)

type Item interface {
	Less(item Item) bool
	Greater(item Item) bool
	Key() int
}

type Node struct {
	Items       []Item
	count       int
	Left, Right *Node
}

func (n *Node) firstItem() Item {
	return n.Items[0]
}

func newNode(item Item) *Node {
	return &Node{Items: []Item{item}, count: 1}
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
	n.count++
}

func (n *Node) remove(item Item) bool {
	var delNode *Node

	if n.firstItem().Greater(item) {
		if n.Left != nil {
			if ok := n.Left.remove(item); ok {
				delNode = n.Left
				if len(delNode.Items) == 0 {
					if delNode.Left == nil {
						n.Left = delNode.Right
						return false
					}
					if delNode.Right == nil {
						n.Left = delNode.Left
						return false
					}
					if delNode.Left.count > delNode.Right.count {
						maxNode, err := delNode.Left.deleteMaximum()
						if err == ErrDeleteAtMinimum {
							maxNode.Right = delNode.Right
						} else {
							maxNode.Left = delNode.Left
							maxNode.Right = delNode.Right
						}
						maxNode.count = delNode.count - 1
						n.Left = maxNode
					} else {
						minNode, err := delNode.Right.deleteMinimum()
						if err == ErrDeleteAtMinimum {
							minNode.Left = delNode.Left
						} else {
							minNode.Left = delNode.Left
							minNode.Right = delNode.Right
						}
						minNode.count = delNode.count - 1
						n.Left = minNode
					}
				}
				n.count--
			}
		} else {
			return false
		}
	} else if n.firstItem().Less(item) {
		if n.Right != nil {
			if ok := n.Right.remove(item); ok {
				delNode := n.Right
				if len(delNode.Items) == 0 {
					if delNode.Left == nil {
						n.Right = delNode.Right
						return false
					}
					if delNode.Right == nil {
						n.Right = delNode.Left
						return false
					}
					if delNode.Left.count > delNode.Right.count {
						maxNode, err := delNode.Left.deleteMaximum()
						if err == ErrDeleteAtMinimum {
							maxNode.Right = delNode.Right
						} else {
							maxNode.Left = delNode.Left
							maxNode.Right = delNode.Right
						}
						maxNode.count = delNode.count - 1
						n.Right = maxNode
					} else {
						minNode, err := delNode.Right.deleteMinimum()
						if err == ErrDeleteAtMinimum {
							minNode.Left = delNode.Left
						} else {
							minNode.Left = delNode.Left
							minNode.Right = delNode.Right
						}
						minNode.count = delNode.count - 1
						n.Right = minNode
					}
				}
				n.count--
			}
		} else {
			return false
		}
	} else {
		if ok, pos := n.include(item); ok {
			n.Items = append(n.Items[:pos], n.Items[pos+1:]...)
			return true
		}
	}
	return false
}

func (n *Node) promote() {

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

var ErrDeleteAtMinimum = errors.New("The root you want to delete is the minimum node of the subtree")

// deleteMinimum removes the minimum node of subtree
func (n *Node) deleteMinimum() (*Node, error) {
	if n.Left != nil {
		n.count--
		if n.Left.Left == nil {
			minimum := n.Left
			n.Left = nil
			return minimum, nil
		}
		return n.Left.deleteMinimum()
	}
	return n, ErrDeleteAtMinimum
}

// minimum returns the maximum node of subtree
func (n *Node) maximum() *Node {
	if n.Right != nil {
		return n.Right.maximum()
	}
	return n
}

var ErrDeleteAtMaximum = errors.New("The root you want to delete is the maximum node of the subtree")

// deleteMaximum removes the maximum node of subtree
func (n *Node) deleteMaximum() (*Node, error) {
	if n.Right != nil {
		n.count--
		if n.Right.Right == nil {
			maximum := n.Right
			n.Right = nil
			return maximum, nil
		}
		return n.Right.deleteMaximum()
	}
	return n, ErrDeleteAtMaximum
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
	var delNode, minNode, maxNode *Node
	var err error
	if ok, pos := t.root.include(item); ok {
		delNode = t.root
		delNode.Items = append(delNode.Items[:pos], delNode.Items[pos+1:]...)
		if len(delNode.Items) == 0 {
			if delNode.Left == nil {
				t.root = delNode.Right
			}
			if delNode.Right == nil {
				t.root = delNode.Left
			}
			if delNode.Left.count > delNode.Right.count {
				maxNode, err = delNode.Left.deleteMaximum()
				if err == ErrDeleteAtMaximum {
					maxNode.Right = delNode.Right
				} else {
					maxNode.Left = delNode.Left
					maxNode.Right = delNode.Right
				}
				maxNode.count = delNode.count - 1
				t.root = maxNode
			} else {
				minNode, err = delNode.Right.deleteMinimum()
				if err == ErrDeleteAtMinimum {
					minNode.Left = delNode.Left
				} else {
					minNode.Left = delNode.Left
					minNode.Right = delNode.Right
				}
				minNode.count = delNode.count - 1
				t.root = minNode
			}
		}
		return
	}
	t.root.remove(item)
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
