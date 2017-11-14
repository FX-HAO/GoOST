package ost

import (
	"testing"
)

func TestAscend(t *testing.T) {
	tree := New()
	order := []Int{
		1, 2, 3, 4, 5, 6, 7, 8, 9,
	}
	for _, i := range order {
		tree.Insert(i)
	}
	k := 0
	tree.Ascend(Int(1), Int(5), func(item Item) bool {
		if k > 3 {
			t.Fatalf("returned more items than expected")
		}
		i1 := order[k+1]
		i2 := item.(Int)
		if i1 != i2 {
			t.Errorf("expecting %v, got %v", i1, i2)
		}
		k++
		return true
	})
	if k != 3 {
		t.Errorf("expecting %v, got %v,", 3, k)
	}
}

func TestAscendGreaterOrEqual(t *testing.T) {
	tree := New()
	order := []Int{
		1, 2, 3, 4, 5, 6, 7, 8, 9,
	}
	for _, i := range order {
		tree.Insert(i)
	}
	// fmt.Println(tree.root)
	// fmt.Println(tree.root.Left)
	// fmt.Println(tree.root.Right)
	// fmt.Println(tree.root.Right.Right)
	// fmt.Println(tree.root.Right.Right.Right)
	k := 0
	tree.AscendGreaterOrEqual(Int(2), Int(5), func(item Item) bool {
		if k > 3 {
			t.Fatalf("returned more items than expected")
		}
		i1 := order[k+1]
		i2 := item.(Int)
		if i1 != i2 {
			t.Errorf("expecting %v, got %v", i1, i2)
		}
		k++
		return true
	})
	if k != 3 {
		t.Errorf("expecting %v, got %v,", 3, k)
	}
}

func TestDescend(t *testing.T) {
	tree := New()
	order := []Int{
		9, 8, 7, 6, 5, 4, 3, 2, 1,
	}
	for _, i := range order {
		tree.Insert(i)
	}
	k := 0
	tree.Descend(Int(9), Int(5), func(item Item) bool {
		if k > 3 {
			t.Fatalf("returned more items than expected")
		}
		i1 := order[k+1]
		i2 := item.(Int)
		if i1 != i2 {
			t.Errorf("expecting %v, got %v", i1, i2)
		}
		k++
		return true
	})
	if k != 3 {
		t.Errorf("expecting %v, got %v,", 3, k)
	}
}

func TestDescendLessOrEqual(t *testing.T) {
	tree := New()
	order := []Int{
		9, 8, 7, 6, 5, 4, 3, 2, 1,
	}
	for _, i := range order {
		tree.Insert(i)
	}
	k := 0
	tree.DescendLessOrEqual(Int(9), Int(6), func(item Item) bool {
		if k > 3 {
			t.Fatalf("returned more items than expected")
		}
		i1 := order[k]
		i2 := item.(Int)
		if i1 != i2 {
			t.Errorf("expecting %v, got %v", i1, i2)
		}
		k++
		return true
	})
	if k != 3 {
		t.Errorf("expecting %v, got %v,", 3, k)
	}
}
