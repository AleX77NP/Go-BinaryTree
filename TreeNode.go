package binarytree

import (
	"errors"
)

//TreeNode ...
type TreeNode struct {
	value int
	left *TreeNode
	right *TreeNode
}

//Insert ...
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