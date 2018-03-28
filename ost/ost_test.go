package ost

import (
	"testing"
)

// type Int int

// func (i Int) Less(item Item) bool {
// 	return i < item.(Int)
// }

// func (i Int) Greater(item Item) bool {
// 	return i > item.(Int)
// }

// func (i Int) Key() int {
// 	return int(i)
// }

func TestOSTInsert(t *testing.T) {
	tree := New()
	data := []Int{2, 1, 3}
	for _, i := range data {
		tree.Insert(i)
	}
	item := tree.root.firstItem()
	if item.(Int) != data[0] {
		t.Errorf("got %v, expect %v", item.(Int), data[0])
	}
	item = tree.root.Left.firstItem()
	if item.(Int) != data[1] {
		t.Errorf("got %v, expect %v", item.(Int), data[1])
	}
	item = tree.root.Right.firstItem()
	if item.(Int) != data[2] {
		t.Errorf("got %v, expect %v", item.(Int), data[2])
	}
}

func TestDelete(t *testing.T) {
	tree := New()
	data := []Int{1, 2, 3, 4, 5, 6, 7}
	for _, i := range data {
		tree.Insert(i)
	}
	tree.Delete(Int(1))
	tree.Delete(Int(2))
	tree.Delete(Int(3))
	k := 0
	tree.AscendGreaterOrEqual(Int(1), Int(7), func(item Item) bool {
		k++
		if k > 3 {
			t.Fatalf("returned more items than expected")
		}
		i := item.(Int)
		if i != data[2+k] {
			t.Errorf("got %v, expect %v", i, data[2+k])
		}
		return true
	})
	if k != 3 {
		t.Errorf("expecting %v, got %v,", 3, k)
	}
}

func TestRank(t *testing.T) {
	tree := New()
	data := []Int{4, 2, 1, 3, 6, 5, 7}
	for _, i := range data {
		tree.Insert(i)
	}
	for _, i := range data {
		j := tree.Rank(i)
		if int(i) != j {
			t.Errorf("got %v, expect %v", j, i)
		}
	}
}

func TestFindByRank(t *testing.T) {
	tree := New()
	data := []Int{4, 2, 1, 3, 6, 5, 7}
	for _, i := range data {
		tree.Insert(i)
	}
	for _, i := range data {
		j := tree.FindByRank(int(i))
		if i != j[0] {
			t.Errorf("got %v, expect %v", j, i)
		}
	}
}

func TestHeight(t *testing.T) {
	data := map[int][]Int{
		2: {2, 1, 3},
		3: {1, 2, 3, 4, 5, 6, 7},
		4: {1, 2, 3, 4, 5, 6, 7, 8, 9},
	}
	for i, l := range data {
		tree := New()
		for _, v := range l {
			tree.Insert(v)
		}
		h := tree.Height()
		if h != i {
			t.Errorf("got %v, expect %v", h, i)
		}
	}
}

// func TestLeftRotate(t *testing.T) {
// 	tree := New()
// 	data := []Int{2, 1, 4, 3, 5, 6}
// 	for _, i := range data {
// 		tree.Insert(i)
// 	}
// 	tree.PrettyPrint()
// 	tree.root.leftRotate()
// 	tree.PrettyPrint()
// }

// func TestRightRotate(t *testing.T) {
// 	tree := New()
// 	data := []Int{5, 6, 3, 4, 2}
// 	for _, i := range data {
// 		tree.Insert(i)
// 	}
// 	tree.Insert(Int(1))
// 	tree.PrettyPrint()
// }
