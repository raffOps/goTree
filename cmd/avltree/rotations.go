package avltree

import (
	"cmp"
	"math"
)

func rotateLeft[T cmp.Ordered](tree *AvlTree[T]) *AvlTree[T] {
	leftTree := tree.leftNode
	tree.leftNode = leftTree.rightNode
	leftTree.rightNode = tree

	tree.height = int64(math.Max(float64(getHeight(tree.leftNode)), float64(getHeight(tree.rightNode))) + 1)
	leftTree.height = int64(math.Max(float64(getHeight(tree.rightNode)), float64(getHeight(tree))) + 1)
	return leftTree
}

func rotateRight[T cmp.Ordered](tree *AvlTree[T]) *AvlTree[T] {
	rightTree := tree.rightNode
	tree.rightNode = rightTree.leftNode
	rightTree.leftNode = tree

	tree.height = int64(math.Max(float64(getHeight(tree.leftNode)), float64(getHeight(tree.rightNode))) + 1)
	rightTree.height = int64(math.Max(float64(getHeight(tree.rightNode)), float64(getHeight(tree))) + 1)
	return rightTree
}

func doubleRotateLeft[T cmp.Ordered](tree *AvlTree[T]) *AvlTree[T] {
	tree.leftNode = rotateRight(tree.leftNode)
	return rotateLeft(tree)
}

func doubleRotateRight[T cmp.Ordered](tree *AvlTree[T]) *AvlTree[T] {
	tree.rightNode = rotateLeft(tree.rightNode)
	return rotateRight(tree)
}
