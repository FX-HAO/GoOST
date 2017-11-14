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
		fmt.Println("n is null")
		return false, false
	}

	var p *Node
	switch dir {
	case ascend:
		if start.Greater(stop) {
			panic(fmt.Sprintf("start %v must be less than stop %v.", start, stop))
		}

		p = n

		// this node is before start
		for {
			if p == nil {
				return true, false
			}
			if p.firstItem().Less(start) {
				p = p.Right
			} else {
				break
			}
		}

		// equals
		if Equal(start, p.firstItem()) {
			if includeStart {
				for i := range p.Items {
					if res := iter(p.Items[i]); !res {
						return false, true
					}
				}
			}
			if p.Right != nil {
				if end, terminate = p.Right.iterate(dir, start, stop, includeStart, iter); end || terminate {
					return end, terminate
				}
			} else {
				return false, false
			}
		}

		// this node is between start and stop
		if start.Less(p.firstItem()) && p.Left != nil {
			if end, terminate = p.Left.iterate(dir, start, stop, includeStart, iter); end || terminate {
				return end, terminate
			}
		}
		if stop.Greater(p.firstItem()) {
			for i := range p.Items {
				if res := iter(p.Items[i]); !res {
					return false, true
				}
			}
			if p.Right != nil {
				if end, terminate = p.Right.iterate(dir, start, stop, includeStart, iter); end || terminate {
					return end, terminate
				}
			}
		} else {
			return true, false
		}

		// if p.firstItem().Less(stop) {
		// if p.Left != nil {
		// 	if end, terminate = p.Left.iterate(dir, start, stop, includeStart, iter); terminate {
		// 		return end, terminate
		// 	}
		// }
		// if start.Less(p.firstItem()) && stop.Greater(p.firstItem()) {
		// 	for i := range p.Items {
		// 		if res := iter(p.Items[i]); !res {
		// 			return false, true
		// 		}
		// 	}
		// }
		// if p.Right != nil {
		// 	if end, terminate = p.Right.iterate(dir, start, stop, includeStart, iter); end || terminate {
		// 		return end, terminate
		// 	}
		// }
		// }

		// this node is between start and stop
		// if start.Less(p.firstItem()) && stop.Greater(p.firstItem()) {
		// 	if p.Left != nil {
		// 		if end, terminate = p.Left.iterate(dir, start, stop, includeStart, iter); terminate {
		// 			return end, terminate
		// 		}
		// 	}
		// 	for i := range p.Items {
		// 		if res := iter(p.Items[i]); !res {
		// 			return false, true
		// 		}
		// 	}
		// 	if p.Right != nil {
		// 		if end, terminate = p.Right.iterate(dir, start, stop, includeStart, iter); end || terminate {
		// 			return end, terminate
		// 		}
		// 	}
		// }

		return false, false
	case descend:
		if start.Less(stop) {
			panic(fmt.Sprintf("start %v must be greater than stop %v.", start, stop))
		}

		p = n

		for {
			if p == nil {
				return true, false
			}
			if p.firstItem().Greater(start) {
				p = p.Left
			} else {
				break
			}
		}

		// equals
		if Equal(start, p.firstItem()) {
			if includeStart {
				for i := range p.Items {
					if res := iter(p.Items[i]); !res {
						return false, true
					}
				}
			}
			if end, terminate = p.Left.iterate(dir, start, stop, includeStart, iter); terminate {
				return end, terminate
			}
		}

		// this node is between start and stop
		if start.Greater(p.firstItem()) && stop.Less(p.firstItem()) {
			if p.Right != nil {
				if end, terminate = p.Right.iterate(dir, start, stop, includeStart, iter); terminate {
					return end, terminate
				}
			}
			for i := range p.Items {
				if res := iter(p.Items[i]); !res {
					return false, true
				}
			}
			if p.Left != nil {
				if end, terminate = p.Left.iterate(dir, start, stop, includeStart, iter); end || terminate {
					return end, terminate
				}
			}
		}

		return true, false
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
