package avltree

import (
	"cmp"
	"fmt"
	"math"
	"strings"
)

// avlTree struct for an AVL Tree of any ordered type.
// This is a self-balancing binary tree.
// The difference between the heights of the left and right subtree cannot be more than one for all nodes.
// For more information about theory, see https://en.wikipedia.org/wiki/AVL_tree
type avlTree[T cmp.Ordered] struct {
	leftNode, rightNode *avlTree[T]
	value               T
	height              int64
}

// NewAvlTree factory function for avlTree struct.
// Return a nil pointer to avlTree.
func NewAvlTree[T cmp.Ordered]() *avlTree[T] {
	return nil
}

func stringHelper[T cmp.Ordered](tree *avlTree[T], level int) string {
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

func (tree avlTree[T]) String() string {
	return stringHelper[T](&tree, 0)
}

// GetRootValue returns the value of the root node.
// If the tree is nil, returns the zero value of the type.
func GetRootValue[T cmp.Ordered](t *avlTree[T]) T {
	if t == nil {
		var zeroValue T
		return zeroValue
	}
	return t.value
}

// getHeight returns the height of a tree.
// If the tree is nil, returns -1.
func getHeight[T cmp.Ordered](tree *avlTree[T]) int64 {
	if tree == nil {
		return -1
	}
	return tree.height
}

// Insert  inserts a value in the tree.
// If the tree is nil, it creates a new node with the value.
// If the value is already in the tree, it does nothing.
// If the value is less than the root, it inserts in the left subtree.
// If the value is greater than the root, it inserts in the right subtree.
// After inserting, it rebalances the tree.
// It returns the new root of the tree.
func Insert[T cmp.Ordered](tree *avlTree[T], value T) *avlTree[T] {
	if tree == nil {
		return &avlTree[T]{value: value}
	}
	if cmp.Compare(value, tree.value) == 0 {
		return tree
	}

	if cmp.Compare(value, tree.value) < 0 {
		tree.leftNode = Insert(tree.leftNode, value)
	} else {
		tree.rightNode = Insert(tree.rightNode, value)
	}

	tree.height = int64(math.Max(float64(getHeight(tree.leftNode)), float64(getHeight(tree.rightNode))) + 1)

	return rebalanceTree(tree, value)
}

// rebalanceTree rebalances the tree.
// If the tree is nil, it returns nil.
// It uses the balance factor to know if the tree is balanced or not.
// If the balance factor is greater than 1 and the value is less than the root left son, it performs a left rotation.
// If the balance factor is less than -1 and the value is greater than the root right son, it performs a right rotation.
// If the balance factor is greater than 1 and the value is greater than the root left son, it performs a right rotation on the left son and then a left rotation on the root.
// If the balance factor is less than -1 and the value is less than the root right son, it performs a left rotation on the right son and then a right rotation on the root.
// It returns the new root of the tree.
func rebalanceTree[T cmp.Ordered](unbalancedTree *avlTree[T], value T) *avlTree[T] {
	balanceFactor := getHeight(unbalancedTree.leftNode) - getHeight(unbalancedTree.rightNode)
	if balanceFactor > 1 && cmp.Compare(value, unbalancedTree.leftNode.value) < 0 {
		return rotateLeft(unbalancedTree)
	}
	if balanceFactor < -1 && cmp.Compare(value, unbalancedTree.rightNode.value) > 0 {
		return rotateRight(unbalancedTree)
	}
	if balanceFactor > 1 && cmp.Compare(value, unbalancedTree.leftNode.value) > 0 {
		unbalancedTree.leftNode = rotateRight(unbalancedTree.leftNode)
		return rotateLeft(unbalancedTree)
	}
	if balanceFactor < -1 && cmp.Compare(value, unbalancedTree.rightNode.value) < 0 {
		unbalancedTree.rightNode = rotateLeft(unbalancedTree.rightNode)
		return rotateRight(unbalancedTree)
	}
	return unbalancedTree
}
