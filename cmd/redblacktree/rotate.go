package redblacktree

import "cmp"

func rotateToRight[T cmp.Ordered](tree *RedBlackTree[T]) *RedBlackTree[T] {
	parentNode := tree.parentNode
	var side string
	if parentNode.leftNode == tree {
		side = "left"
	} else {
		side = "right"
	}

	rotatedTree := tree.leftNode
	rotatedTree.parentNode = parentNode

	tree.leftNode = rotatedTree.rightNode
	tree.leftNode.parentNode = tree

	rotatedTree.rightNode = tree
	rotatedTree.rightNode.parentNode = rotatedTree

	if side == "left" {
		parentNode.leftNode = rotatedTree
	} else {
		parentNode.rightNode = rotatedTree
	}
	return rotatedTree
}

func rotateToLeft[T cmp.Ordered](tree *RedBlackTree[T]) *RedBlackTree[T] {
	parentNode := tree.parentNode
	var side string
	if parentNode.leftNode == tree {
		side = "left"
	} else {
		side = "right"
	}

	rotatedTree := tree.rightNode
	rotatedTree.parentNode = parentNode

	tree.rightNode = rotatedTree.leftNode
	tree.rightNode.parentNode = tree

	rotatedTree.leftNode = tree
	rotatedTree.leftNode.parentNode = rotatedTree

	if side == "left" {
		parentNode.leftNode = rotatedTree
	} else {
		parentNode.rightNode = rotatedTree
	}
	return rotatedTree
}
