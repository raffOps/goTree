package avltree

import (
	"cmp"
	"math"
)

// AVLTree struct for a AVL Tree of any ordered type.
// For more information about theory, see https://en.wikipedia.org/wiki/AVL_tree
type AVLTree[T cmp.Ordered] struct {
	leftNode, rightNode *AVLTree[T]
	value               T
	height              int64
}

// NewAVLTree factory function for AVLTree struct
func NewAVLTree[T cmp.Ordered](rootValue T) *AVLTree[T] {
	return &AVLTree[T]{value: rootValue}
}

func (t *AVLTree[T]) GetRootValue() T {
	return t.value
}

func getHeight[T cmp.Ordered](tree *AVLTree[T]) int64 {
	if tree == nil {
		return -1
	}
	return tree.height
}

// Theory: https://en.wikipedia.org/wiki/AVL_tree#Simple_rotation
func rotateLeft[T cmp.Ordered](tree *AVLTree[T]) *AVLTree[T] {
	leftTree := tree.leftNode
	tree.leftNode = leftTree.rightNode
	leftTree.rightNode = tree

	tree.height = int64(math.Max(float64(getHeight(tree.leftNode)), float64(getHeight(tree.rightNode))) + 1)
	leftTree.height = int64(math.Max(float64(getHeight(tree.rightNode)), float64(getHeight(tree))) + 1)
	return leftTree
}

// Theory: https://en.wikipedia.org/wiki/AVL_tree#Simple_rotation
func rotateRight[T cmp.Ordered](tree *AVLTree[T]) *AVLTree[T] {
	rightTree := tree.rightNode
	tree.rightNode = rightTree.leftNode
	rightTree.leftNode = tree

	tree.height = int64(math.Max(float64(getHeight(tree.leftNode)), float64(getHeight(tree.rightNode))) + 1)
	rightTree.height = int64(math.Max(float64(getHeight(tree.rightNode)), float64(getHeight(tree))) + 1)
	return rightTree
}

// Theory: https://en.wikipedia.org/wiki/AVL_tree#Double_rotation
func doubleRotateLeft[T cmp.Ordered](tree *AVLTree[T]) *AVLTree[T] {
	tree.leftNode = rotateRight(tree.leftNode)
	return rotateLeft(tree)
}

// Theory: https://en.wikipedia.org/wiki/AVL_tree#Double_rotation
func doubleRotateRight[T cmp.Ordered](tree *AVLTree[T]) *AVLTree[T] {
	tree.rightNode = rotateLeft(tree.rightNode)
	return rotateRight(tree)
}
