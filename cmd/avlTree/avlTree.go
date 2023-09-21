package avlTree

import (
	"cmp"
)

type AVLTree[T cmp.Ordered] struct {
	leftNode, rightNode *AVLTree[T]
	value               T
	height              int64
}

func NewAVLTree[T cmp.Ordered]() *AVLTree[T] {
	return &AVLTree[T]{}
}

func (t *AVLTree[T]) IsEmpty() bool {
	var zeroValue T
	return t.value == zeroValue && t.leftNode == nil && t.rightNode == nil
}

func (t *AVLTree[T]) GetRootValue() T {
	return t.value
}

func (t *AVLTree[T]) setHeight(height int64) {
	t.height = height
}

func GetHeight[T cmp.Ordered](tree *AVLTree[T]) int64 {
	if tree == nil {
		return -1
	}
	return tree.height
}

func maxHeight(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

func rotateLeft[T cmp.Ordered](tree *AVLTree[T]) *AVLTree[T] {
	leftTree := tree.leftNode
	tree.leftNode = leftTree.rightNode
	leftTree.rightNode = tree

	tree.setHeight(maxHeight(GetHeight(tree.leftNode), GetHeight(tree.rightNode)) + 1)
	leftTree.setHeight(maxHeight(GetHeight(tree.rightNode), GetHeight(tree)) + 1)
	return leftTree
}

func doubleRotateLeft[T cmp.Ordered](tree *AVLTree[T]) *AVLTree[T] {
	return nil
}

func rotateRight[T cmp.Ordered](tree *AVLTree[T]) *AVLTree[T] {
	return nil
}

func doubleRotateRight[T cmp.Ordered](tree *AVLTree[T]) *AVLTree[T] {
	return nil
}

//func insertValueOnTheLeft[T cmp.Ordered](value T, tree, fatherTree *AVLTree[T]) *AVLTree[T] {
//	if tree.leftNode.GetHeight()-tree.GetHeight() > 1 {
//		if value < tree.leftNode.value {
//			rotatedTree := rotateLeft(tree)
//		} else {
//			rotatedTree := doubleRotateLeft(tree)
//		}
//
//	}
//}

//
//
//func insertValueOnTheRight[T cmp.Ordered](value T, tree, fatherTree *AVLTree[T]) *AVLTree[T] {
//	return nil
//}
//
//func insertValue[T cmp.Ordered](value T, tree, fatherTree *AVLTree[T]) {
//	if tree.IsEmpty() { // empty tree
//		tree.value = value
//		return
//	}
//	valueIsLesserOrGreaterThanRootValue := cmp.Compare(value, tree.value)
//	if valueIsLesserOrGreaterThanRootValue == -1 {
//		tree.leftNode = insertValueOnTheLeft[T](value, tree.leftNode, tree)
//	} else {
//		tree.rightNode = insertValueOnTheRight[T](value, tree.leftNode, tree)
//	}
//}
//
//func (t *AVLTree[T]) InsertValue(value T) {
//	insertValue(value, t, nil)
//}
