package main

import "golang.org/x/exp/constraints"

type BinaryNode[T constraints.Ordered] struct {
	value T
	left  *BinaryNode[T]
	right *BinaryNode[T]
}

// Create a binary node
func NewBinaryNode[T constraints.Ordered](value T) *BinaryNode[T] {
	return &BinaryNode[T]{
		value: value,
		left:  nil,
		right: nil,
	}
}

// Add a new node to BST with the given value
func (n *BinaryNode[T]) Add(val T) {
	if val <= n.value {
		if n.left != nil {
			n.left.Add(val)
		} else {
			n.left = NewBinaryNode[T](val)
		}
	} else {
		if n.right != nil {
			n.right.Add(val)
		} else {
			n.right = NewBinaryNode[T](val)
		}
	}
}

type BinaryTree[T constraints.Ordered] struct {
	root *BinaryNode[T]
}

// Create an empty BST
func NewBinaryTree[T constraints.Ordered]() *BinaryTree[T] {
	return &BinaryTree[T]{
		root: nil,
	}
}

// Insert value into proper location in BST
func (bst *BinaryTree[T]) Insert(val T) {
	if bst.root == nil {
		bst.root = NewBinaryNode[T](val)
	} else {
		bst.root.Add(val)
	}
}

// Check wether BST contains target value
func (bst *BinaryTree[T]) Contains(target T) bool {
	node := bst.root

	for node != nil {
		if target < node.value {
			node = node.left
		} else if target > node.value {
			node = node.right
		} else {
			return true
		}
	}

	return false
}
