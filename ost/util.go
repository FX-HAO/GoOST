package ost

func Equal(a, b Item) bool {
	return !a.Less(b) && !b.Less(a)
}

type Int int

func (i Int) Less(item Item) bool {
	return i < item.(Int)
}

func (i Int) Greater(item Item) bool {
	return i > item.(Int)
}

func (i Int) Equal(item Item) bool {
	return i == item.(Int)
}

func (i Int) Key() int {
	return int(i)
}
