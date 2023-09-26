package avltree

import (
	"cmp"
)

// AvlTree struct for a AVL Tree of any ordered type.
// For more information about theory, see https://en.wikipedia.org/wiki/AVL_tree
type AvlTree[T cmp.Ordered] struct {
	leftNode, rightNode *AvlTree[T]
	value               T
	height              int64
}

// NewAVLTree factory function for avlTree struct
func NewAVLTree[T cmp.Ordered](rootValue T) *AvlTree[T] {
	return &AvlTree[T]{value: rootValue}
}

func (t *AvlTree[T]) GetRootValue() T {
	return t.value
}

func getHeight[T cmp.Ordered](tree *AvlTree[T]) int64 {
	if tree == nil {
		return -1
	}
	return tree.height
}
