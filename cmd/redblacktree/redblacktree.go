package redblacktree

import "cmp"

// RedBlackTree is a struct for an AVL Tree of any ordered type.
// This is a self-balancing binary tree.
// The difference between the heights of the left and right subtree cannot be more than one for all nodes.
// For more information about theory, see https://en.wikipedia.org/wiki/Red-black_tree.
type RedBlackTree[T cmp.Ordered] struct {
	leftNode, rightNode *RedBlackTree[T]
	value               T
	isRed               bool
}

// NewRedBlackTree factory function that returns a new RedBlackTree
func NewRedBlackTree[T cmp.Ordered]() *RedBlackTree[T] {
	return &RedBlackTree[T]{}
}

// GetValue returns the value stored in the tree
func (tree *RedBlackTree[T]) GetValue() T {
	return tree.value
}

// GetChild returns the child node in the direction specified
func (tree *RedBlackTree[T]) GetChild(direction string) *RedBlackTree[T] {
	if direction == "left" {
		return tree.leftNode
	}
	return tree.rightNode
}

// SetChild sets the child node in the direction specified
func (tree *RedBlackTree[T]) SetChild(direction string, node *RedBlackTree[T]) {
	if direction == "left" {
		tree.leftNode = node
	} else {
		tree.rightNode = node
	}
}
