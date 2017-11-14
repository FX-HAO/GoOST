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
	data := []Int{4, 2, 1, 3, 6, 5, 7}
	for _, i := range data {
		tree.Insert(i)
	}
	tree.Delete(Int(3))
	tree.Delete(Int(6))
	tree.Delete(Int(4))
	// fmt.Println(tree.root)
	// fmt.Println(tree.root.Left)
	// fmt.Println(tree.root.Left.Left)
	// fmt.Println(tree.root.Left.Right)
	// fmt.Println(tree.root.Right)
	// fmt.Println(tree.root.Right.Left)
	// fmt.Println(tree.root.Right.Right)
	k := 0
	tree.AscendGreaterOrEqual(Int(1), Int(7), func(item Item) bool {
		k++
		if k > 3 {
			t.Fatalf("returned more items than expected")
		}
		i := item.(Int)
		switch k {
		case 1:
			if i != data[2] {
				t.Errorf("got %v, expect %v", i, data[2])
			}
		case 2:
			if i != data[1] {
				t.Errorf("got %v, expect %v", i, data[1])
			}
		case 3:
			if i != data[5] {
				t.Errorf("got %v, expect %v", i, data[5])
			}
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
