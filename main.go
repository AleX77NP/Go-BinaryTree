package main

import (
	"errors"
	"fmt"
)

//TreeNode ...
type TreeNode struct {
	value int
	left *TreeNode
	right *TreeNode
}

// Insert ...
func (t *TreeNode) Insert(num int) (string, error) {
	if t == nil {
		return "", errors.New("Tree is nil")
	}

	if t.value == num {
		return "", errors.New("This tree has that value")
	}
	if t.value > num {
		if t.left == nil {
			t.left = &TreeNode{value: num}
			return "Inserted on the left side",nil
		}
		return t.left.Insert(num)
	}
	if t.value < num {
		if t.right == nil {
			t.right = &TreeNode{value: num}
			return "Inserted on the right side", nil
		}
		return t.right.Insert(num)
	}
	return "",nil
}


// Find ...
func (t *TreeNode) Find(value int) (string, error) {
	if t == nil {
		return "", errors.New("Tree is nil")
	}
	if t.value == value {
		return "Found value at the root", nil
	} 
	if t.value < value {
		return t.right.Find(value)
	}
	if t.value > value {
		return t.left.Find(value)
	}
	return "", nil
}

// FindMax ...
func (t *TreeNode) FindMax() (int, error) {
	if t == nil {
		return 0, errors.New("Tree is nil")
	}
	if t.right == nil {
		return t.value, nil
	}
	return t.right.FindMax()
}

// FindMin ...
func (t *TreeNode) FindMin() (int, error) {
	if t == nil {
		return 0, errors.New("Tree is nil")
	}
	if t.left == nil {
		return t.value, nil
	}
	return t.left.FindMin()
}

// PrintInorder ...
func (t *TreeNode) PrintInorder() {
	if t == nil {
		return
	}

	t.left.PrintInorder()
	fmt.Println(t.value)
	t.right.PrintInorder()
}

func main() {

	tree := &TreeNode{value: 10}

	tree.Insert(3)
	tree.Insert(34)
	tree.Insert(9)
	tree.Insert(43)
	tree.Insert(15)

	tree.PrintInorder()


}