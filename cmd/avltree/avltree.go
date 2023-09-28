package avltree

import (
	"cmp"
	"fmt"
	"strings"
)

// AvlTree struct for a AVL Tree of any ordered type.
// For more information about theory, see https://en.wikipedia.org/wiki/AVL_tree
type AvlTree[T cmp.Ordered] struct {
	leftNode, rightNode *AvlTree[T]
	value               T
	height              int64
}

// NewAvlTree factory function for avlTree struct.
// Return a nil pointer to avlTree.
func NewAvlTree[T cmp.Ordered]() *avlTree[T] {
	return nil
}

func stringHelper[T cmp.Ordered](tree *avlTree[T], level int) string {
	indentation := strings.Repeat("\t", level)
	if tree == nil {
		return ""
	} else if tree.leftNode == nil && tree.rightNode == nil {
		return fmt.Sprintf("value: %v, height: %d", tree.value, tree.height)
	}
	return fmt.Sprintf("value: %v, height: %d, \n%sleftNode: %v, \n%srightNode: %v",
		tree.value,
		tree.height,
		indentation,
		stringHelper(tree.leftNode, level+1),
		indentation,
		stringHelper(tree.rightNode, level+1),
	)
}

func (tree avlTree[T]) String() string {
	return stringHelper[T](&tree, 0)
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
