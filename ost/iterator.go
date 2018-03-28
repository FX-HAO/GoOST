package ost

import "fmt"

// ItemIterator allows callers of Ascend* to iterate in-order over portions of
// the tree.  When this function returns false, iteration will stop and the
// associated Ascend* function will immediately return.
type ItemIterator func(i Item) bool

type direction int

const (
	descend = direction(-1)
	ascend  = direction(+1)
)

// iterate provides a simple method for iterating over elements in the tree.
//
// When ascending, the 'start' should be less than 'stop' and when descending,
// the 'start' should be greater than 'stop'. Setting 'includeStart' to true
// will force the iterator to include the first item when it equals 'start',
// thus creating a "greaterOrEqual" or "lessThanEqual" rather than just a
// "greaterThan" or "lessThan" queries.
func (n *Node) iterate(dir direction, start, stop Item, includeStart bool, iter ItemIterator) (end, terminate bool) {
	if n == nil {
		return false, false
	}

	// var p *Node
	value := n.firstItem()

	switch dir {
	case ascend:
		if start.Greater(stop) {
			panic(fmt.Sprintf("start %v must be less than stop %v.", start, stop))
		}

		if value.Less(start) {
			return false, false
		}

		if value.Less(stop) {
			if includeStart && Equal(start, value) {
				for i := range n.Items {
					if res := iter(n.Items[i]); !res {
						return false, true
					}
				}
			}
			if value.Greater(start) {
				if end, terminate = n.Left.iterate(dir, start, stop, includeStart, iter); end || terminate {
					return end, terminate
				}
				for i := range n.Items {
					if res := iter(n.Items[i]); !res {
						return false, true
					}
				}
			}
			if end, terminate = n.Right.iterate(dir, start, stop, includeStart, iter); end || terminate {
				return end, terminate
			}
		} else {
			if end, terminate = n.Left.iterate(dir, start, stop, includeStart, iter); end || terminate {
				return end, terminate
			}
			return true, false
		}

		return false, false
	case descend:
		if start.Less(stop) {
			panic(fmt.Sprintf("start %v must be greater than stop %v.", start, stop))
		}

		if !value.Greater(stop) {
			if end, terminate = n.Right.iterate(dir, start, stop, includeStart, iter); end || terminate {
				return end, terminate
			}
			return true, false
		}

		if value.Greater(start) {
			if end, terminate = n.Left.iterate(dir, start, stop, includeStart, iter); end || terminate {
				return end, terminate
			}
		} else if includeStart && Equal(start, value) {
			for i := range n.Items {
				if res := iter(n.Items[i]); !res {
					return false, true
				}
			}
		} else if value.Less(start) {
			if end, terminate = n.Right.iterate(dir, start, stop, includeStart, iter); end || terminate {
				return end, terminate
			}
			for i := range n.Items {
				if res := iter(n.Items[i]); !res {
					return false, true
				}
			}
			if end, terminate = n.Left.iterate(dir, start, stop, includeStart, iter); end || terminate {
				return end, terminate
			}
		}

		return false, false
	}

	return true, false
}

func (t *OST) Ascend(start, stop Item, iter ItemIterator) {
	if t.root == nil {
		return
	}
	t.root.iterate(ascend, start, stop, false, iter)
}

func (t *OST) AscendGreaterOrEqual(start, stop Item, iter ItemIterator) {
	if t.root == nil {
		return
	}
	t.root.iterate(ascend, start, stop, true, iter)
}

func (t *OST) Descend(start, stop Item, iter ItemIterator) {
	if t.root == nil {
		return
	}
	t.root.iterate(descend, start, stop, false, iter)
}

func (t *OST) DescendLessOrEqual(start, stop Item, iter ItemIterator) {
	if t.root == nil {
		return
	}
	t.root.iterate(descend, start, stop, true, iter)
}
