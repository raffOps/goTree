package redblacktree

import "cmp"

// rotateToRight rotates the tree to the right.
// Returns the new root of the tree.
func rotateToRight[T cmp.Ordered](tree *RedBlackTree[T]) *RedBlackTree[T] {
	rotatedTree := tree.leftNode
	tree.leftNode = rotatedTree.rightNode
	rotatedTree.rightNode = tree

	rotatedTree.isRed = tree.isRed
	tree.isRed = true

	return rotatedTree
}

// rotateToLeft rotates the tree to the left.
// Returns the new root of the tree.
func rotateToLeft[T cmp.Ordered](tree *RedBlackTree[T]) *RedBlackTree[T] {
	rotatedTree := tree.rightNode
	tree.rightNode = rotatedTree.leftNode
	rotatedTree.leftNode = tree

	rotatedTree.isRed = tree.isRed
	tree.isRed = true

	return rotatedTree
}
