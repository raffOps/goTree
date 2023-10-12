package avltree

import (
	"cmp"
	"fmt"
	"strings"
)

// AvlTree struct for an AVL Tree of any ordered type.
// This is a self-balancing binary tree.
// The difference between the heights of the left and right subtree cannot be more than one for all nodes.
// For more information about theory, see https://en.wikipedia.org/wiki/AVL_tree
type AvlTree[T cmp.Ordered] struct {
	leftNode, rightNode *AvlTree[T]
	value               T
	height              int64
}

// NewAvlTree is factory function for avlTree struct.
// Return a nil pointer to avlTree.
func NewAvlTree[T cmp.Ordered]() *AvlTree[T] {
	return nil
}

// String returns a string representation of the tree.
func (tree *AvlTree[T]) String() string {
	return stringHelper[T](tree, 0)
}

func stringHelper[T cmp.Ordered](tree *AvlTree[T], level int) string {
	if tree == nil {
		return ""
	} else if tree.leftNode == nil && tree.rightNode == nil {
		return fmt.Sprintf("value: %v, height: %d", tree.value, tree.height)
	}
	indentation := strings.Repeat("  ", level)
	return fmt.Sprintf("value: %v, height: %d, \n%sleftNode: %v, \n%srightNode: %v",
		tree.value,
		tree.height,
		indentation,
		stringHelper(tree.leftNode, level+1),
		indentation,
		stringHelper(tree.rightNode, level+1),
	)
}

// GetRootValue returns the value of the root node. If the tree is nil, returns the zero value of the type.
func GetRootValue[T cmp.Ordered](t *AvlTree[T]) T {
	if t == nil {
		var zeroValue T
		return zeroValue
	}
	return t.value
}

// getHeight
// returns the height of a tree. If the tree is nil, returns -1.
func getHeight[T cmp.Ordered](tree *AvlTree[T]) int64 {
	if tree == nil {
		return -1
	}
	return tree.height
}

// Search returns the node with the value or nil if it is not found.
func Search[T cmp.Ordered](tree *AvlTree[T], value T) *AvlTree[T] {
	if tree == nil || value == tree.value {
		return tree
	}
	if value < tree.value {
		return Search(tree.leftNode, value)
	}
	return Search(tree.rightNode, value)
}

// ToArray returns an array with the values of the tree in order.
func ToArray[T cmp.Ordered](tree *AvlTree[T]) []T {
	arr := []T{}
	return toArrayHelper(tree, arr)
}

func toArrayHelper[T cmp.Ordered](tree *AvlTree[T], array []T) []T {
	if tree == nil {
		return array
	}
	array = toArrayHelper(tree.leftNode, array)
	array = append(array, tree.value)
	array = toArrayHelper(tree.rightNode, array)
	return array
}

// getBalanceFactor
// It returns the difference between the heights of the left and right subtree.
// If the tree is nil, it returns 0.
func getBalanceFactor[T cmp.Ordered](tree *AvlTree[T]) int64 {
	if tree == nil {
		return 0
	}
	balanceFactor := getHeight(tree.leftNode) - getHeight(tree.rightNode)
	return balanceFactor
}
