package avltree

import (
	"cmp"
	"math"
)

// RotationNotPossible is raised when the tree is nil or when the left/right son of the tree is nil.
type RotationNotPossible struct {
	Message string
}

func (e *RotationNotPossible) Error() string {
	return e.Message
}

// rotateLeft performs a left rotation on the tree. It returns the new root of the tree.
// It panics if the tree is nil or if the left son of the tree is nil.
func rotateLeft[T cmp.Ordered](tree *AvlTree[T]) *AvlTree[T] {
	if tree == nil || tree.leftNode == nil {
		err := RotationNotPossible{"Tree or left son is nil"}
		panic(err)
	}
	rotatedTree := tree.leftNode
	tree.leftNode = rotatedTree.rightNode
	rotatedTree.rightNode = tree

	tree.height = int64(math.Max(float64(getHeight(tree.leftNode)), float64(getHeight(tree.rightNode))) + 1)
	rotatedTree.height = int64(math.Max(float64(getHeight(tree.rightNode)), float64(getHeight(tree))) + 1)
	return rotatedTree
}

// rotateRight performs a right rotation on the tree. It returns the new root of the tree.
// It panics if the tree is nil or if the right son of the tree is nil.
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
