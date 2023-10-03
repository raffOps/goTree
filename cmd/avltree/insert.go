package avltree

import (
	"cmp"
	"math"
)

// Insert a value in the tree.
// If the tree is nil, it creates a new node with the value.
// If the value is already in the tree, it does nothing.
// If the value is less than the root, it inserts in the left subtree.
// If the value is greater than the root, it inserts in the right subtree.
// After inserting, it rebalances the tree.
// It returns the new root of the tree.
func Insert[T cmp.Ordered](tree *AvlTree[T], value T) *AvlTree[T] {
	if tree == nil {
		return &AvlTree[T]{value: value}
	}
	if value == tree.value {
		return tree
	}

	if value < tree.value {
		tree.leftNode = Insert(tree.leftNode, value)
	} else {
		tree.rightNode = Insert(tree.rightNode, value)
	}

	tree.height = int64(math.Max(float64(getHeight(tree.leftNode)), float64(getHeight(tree.rightNode))) + 1)

	return getBalancedTree(tree, value)
}

// getBalancedTree
// If the tree is nil, it returns nil.
// It uses the balance factor to know if the tree is balanced or not.
// If the balance factor is greater than 1 and the value is less than the root left son,
//
//	it returns a tree with left rotation.
//
// If the balance factor is greater than 1 and the value is greater than the root left son,
//
//	it returns a tree with right rotation on the left son and then a left rotation on the root.
//
// If the balance factor is less than -1 and the value is greater than the root right son,
//
//	it returns a tree with right rotation.
//
// If the balance factor is less than -1 and the value is less than the root right son,
//
//	it returns a tree with left rotation on the right son and then a right rotation on the root.
//
// It returns the new root of the tree.
func getBalancedTree[T cmp.Ordered](tree *AvlTree[T], value T) *AvlTree[T] {
	balanceFactor := getBalanceFactor(tree)
	if balanceFactor > 1 && value < tree.leftNode.value {
		return rotateLeft(tree)
	}
	if balanceFactor > 1 && value > tree.leftNode.value {
		tree.leftNode = rotateRight(tree.leftNode)
		return rotateLeft(tree)
	}
	if balanceFactor < -1 && value > tree.rightNode.value {
		return rotateRight(tree)
	}

	if balanceFactor < -1 && value < tree.rightNode.value {
		tree.rightNode = rotateLeft(tree.rightNode)
		return rotateRight(tree)
	}
	return tree
}
