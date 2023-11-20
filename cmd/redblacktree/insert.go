package redblacktree

import "cmp"

// getRecoloredTree
func getRecoloredTree[T cmp.Ordered](tree *RedBlackTree[T]) *RedBlackTree[T] {
	tree.isRed = !tree.isRed
	tree.leftNode.isRed = !tree.leftNode.isRed
	tree.rightNode.isRed = !tree.rightNode.isRed
	return tree
}

func insertHelper[T cmp.Ordered](tree *RedBlackTree[T], value T) *RedBlackTree[T] {
	if tree.IsEmpty() {
		tree.value = &nodeValue[T]{value}
		tree.isRed = true
		tree.leftNode = NewRedBlackTree[T]()
		tree.rightNode = NewRedBlackTree[T]()
		return tree
	}
	if value == tree.GetValue() {
		return tree
	}

	if value < tree.GetValue() {
		tree.leftNode = insertHelper(tree.leftNode, value)
	} else {
		tree.rightNode = insertHelper(tree.rightNode, value)
	}

	if tree.rightNode.isRed {
		tree = rotateToLeft(tree)
	}

	if tree.leftNode.isRed && tree.leftNode.leftNode != nil && tree.leftNode.leftNode.isRed {
		tree = rotateToRight(tree)
	}

	if tree.leftNode.isRed && tree.rightNode.isRed {
		tree = getRecoloredTree(tree)
	}

	return tree
}

// Insert a value into the tree.
// If the tree is nil, it returns a new tree with the value.
// If the value is already in the tree, it returns the tree.
// If the value is less than the root value, it inserts the value in the left subtree.
// If the value is greater than the root value, it inserts the value in the right subtree.
// After inserting, it rebalances the tree.
// It returns the new root of the tree.
func Insert[T cmp.Ordered](tree *RedBlackTree[T], value T) *RedBlackTree[T] {
	tree = insertHelper(tree, value)
	tree.isRed = false
	return tree
}
