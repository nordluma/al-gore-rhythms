package main

import "golang.org/x/exp/constraints"

type AVLNode[T constraints.Ordered] struct {
	value  T
	left   *AVLNode[T]
	right  *AVLNode[T]
	height int
}

// Create a binary node with default values
func NewAVLNode[T constraints.Ordered](val T) *AVLNode[T] {
	return &AVLNode[T]{
		value:  val,
		left:   new(AVLNode[T]),
		right:  new(AVLNode[T]),
		height: 0,
	}
}

// Adds a new node to AVL tree with value and rebalance as needed
func (n *AVLNode[T]) Add(val T) *AVLNode[T] {
	newRoot := n

	if val <= n.value {
		n.left = n.addToSubTree(n.left, val)
		if n.heightDifference() == 2 {
			if val <= n.left.value {
				newRoot = n.rotateRight()
			} else {
				newRoot = n.rotateRightLeft()
			}
		}
	} else {
		n.right = n.addToSubTree(n.right, val)
		if n.heightDifference() == -2 {
			if val > n.right.value {
				newRoot = n.rotateLeft()
			} else {
				newRoot = n.rotateLeftRigth()
			}
		}
	}

	newRoot.computHeight()

	return newRoot
}

// Perform right rotation around given node
func (n *AVLNode[T]) rotateRight() *AVLNode[T] {
	newRoot := n.left
	grandson := newRoot.right
	n.left = grandson
	newRoot.right = n

	n.computHeight()

	return newRoot
}

// Perform right, then left rotation around given node
func (n *AVLNode[T]) rotateRightLeft() *AVLNode[T] {
	child := n.right
	newRoot := child.left
	grand1 := newRoot.left
	grand2 := newRoot.right
	child.left = grand2
	n.right = grand1

	newRoot.left = n
	newRoot.right = child

	child.computHeight()
	n.computHeight()

	return newRoot
}

// Perform left rotation around given node
func (n *AVLNode[T]) rotateLeft() *AVLNode[T] {
	newRoot := n.right
	grandson := newRoot.left
	n.right = grandson
	newRoot.left = n

	n.computHeight()

	return newRoot
}

// Perform left, then right rotation around given node
func (n *AVLNode[T]) rotateLeftRigth() *AVLNode[T] {
	child := n.left
	newRoot := child.right
	grand1 := newRoot.left
	grand2 := newRoot.right
	child.right = grand1
	n.left = grand2

	newRoot.left = child
	newRoot.right = n

	child.computHeight()
	n.computHeight()

	return newRoot
}

// Add val to parent subtree (if exists) and return root in case it has changed
// because of rotation
func (n *AVLNode[T]) addToSubTree(parent *AVLNode[T], val T) *AVLNode[T] {
	if parent == nil {
		return NewAVLNode[T](val)
	}

	parent.Add(val)

	return parent
}

// Compute height of node in BST from children
func (n *AVLNode[T]) computHeight() {
	height := -1
	if n.left != nil {
		height = max(height, n.left.height)
	}

	if n.right != nil {
		height = max(height, n.right.height)
	}

	n.height = height + 1
}

// Compute height difference of node's children in BST
func (n *AVLNode[T]) heightDifference() int {
	leftTarget := 0
	rightTarget := 0

	if n.left != nil {
		leftTarget = 1 + n.left.height
	}

	if n.right != nil {
		rightTarget = 1 + n.right.height
	}

	return leftTarget - rightTarget
}

type AVLTree[T constraints.Ordered] struct {
	root *AVLNode[T]
}

func NewAVLTree[T constraints.Ordered]() *AVLTree[T] {
	return &AVLTree[T]{
		root: nil,
	}
}

func (avl *AVLTree[T]) Insert(val T) {
	if avl.root == nil {
		avl.root = NewAVLNode[T](val)
	} else {
		avl.root.Add(val)
	}
}

// Check if the AVL tree contains target value
func (avl *AVLTree[T]) Contains(target T) bool {
	node := avl.root

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
