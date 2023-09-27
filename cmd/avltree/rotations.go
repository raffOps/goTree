package avltree

import (
	"cmp"
	"math"
)

type RotationNotPossible struct {
	Message string
}

func (e *RotationNotPossible) Error() string {
	return e.Message
}

func rotateLeft[T cmp.Ordered](tree *AvlTree[T]) *AvlTree[T] {
	if tree == nil || tree.leftNode == nil {
		err := RotationNotPossible{"Tree and/or left son is nil"}
		panic(err)
	}
	rotatedTree := tree.leftNode
	tree.leftNode = rotatedTree.rightNode
	rotatedTree.rightNode = tree

	tree.height = int64(math.Max(float64(getHeight(tree.leftNode)), float64(getHeight(tree.rightNode))) + 1)
	rotatedTree.height = int64(math.Max(float64(getHeight(tree.rightNode)), float64(getHeight(tree))) + 1)
	return rotatedTree
}

func rotateRight[T cmp.Ordered](tree *AvlTree[T]) *AvlTree[T] {
	if tree == nil || tree.rightNode == nil {
		err := RotationNotPossible{"Tree and/or right son is nil"}
		panic(err)
	}
	rotatedTree := tree.rightNode
	tree.rightNode = rotatedTree.leftNode
	rotatedTree.leftNode = tree

	tree.height = int64(math.Max(float64(getHeight(tree.leftNode)), float64(getHeight(tree.rightNode))) + 1)
	rotatedTree.height = int64(math.Max(float64(getHeight(tree.rightNode)), float64(getHeight(tree))) + 1)
	return rotatedTree
}

func doubleRotateLeft[T cmp.Ordered](tree *AvlTree[T]) *AvlTree[T] {
	tree.leftNode = rotateRight(tree.leftNode)
	return rotateLeft(tree)
}

func doubleRotateRight[T cmp.Ordered](tree *AvlTree[T]) *AvlTree[T] {
	tree.rightNode = rotateLeft(tree.rightNode)
	return rotateRight(tree)
}
