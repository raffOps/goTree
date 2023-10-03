package avltree

import (
	"cmp"
	"math"
)

// Delete the value in the tree.
// If the tree is nil, it returns nil.
// If the value is equal to the root, it deletes the root.
// If the value is less than the root, it deletes in the left subtree.
// If the value is greater than the root, it deletes in the right subtree.
// After deleting, it rebalances the tree.
// It returns the new root of the tree.
func Delete[T cmp.Ordered](tree *AvlTree[T], value T) *AvlTree[T] {
	if tree == nil {
		return nil
	}
	if value == tree.value {
		tree = deleteHelper(tree)
		return tree
	}

	if value < tree.value {
		tree.leftNode = Delete(tree.leftNode, value)
	} else {
		tree.rightNode = Delete(tree.rightNode, value)
	}
	tree.height = int64(math.Max(float64(getHeight(tree.leftNode)), float64(getHeight(tree.rightNode))) + 1)
	return tree
}

// deleteHelper
// Remove the root of the tree.
// If the tree is a leaf, it returns nil.
// If the tree has only one son, it returns the son.
// If the tree has two sons, it returns the value smallest son on the right as root of a balanced tree.
func deleteHelper[T cmp.Ordered](tree *AvlTree[T]) *AvlTree[T] {
	if tree.leftNode == nil && tree.rightNode == nil {
		return nil
	}
	if tree.leftNode != nil && tree.rightNode == nil {
		return tree.leftNode
	}
	if tree.leftNode == nil && tree.rightNode != nil {
		return tree.rightNode
	}
	newRootTree := getSmallestSonOnTheRight(tree)
	tree = Delete(tree, newRootTree.value)
	newRootTree.leftNode = tree.leftNode
	newRootTree.rightNode = tree.rightNode
	newRootTree.height = int64(math.Max(float64(getHeight(tree.leftNode)), float64(getHeight(tree.rightNode))) + 1)

	balancedTree := getBalancedTreeAfterDelete(newRootTree)
	return balancedTree
}

// getBalancedTreeAfterDelete
// If tree is balanced or nil,
//
//	it returns tree.
//
// If tree is unbalanced to the left and his son on the left have greater depth on the left,
//
//	it returns the tree with a left rotation.
//
// If tree is unbalanced to the left and his son on the left have greater depth on the right,
//
//	it returns the tree with a double left rotation.
//
// If tree is unbalanced to the right and his son on the right have greater depth on the right,
//
//	it returns the tree with a right rotation.
//
// If tree is unbalanced to right and his son on the right have greater depth on the left,
//
//	it returns the tree with a double right rotation.
func getBalancedTreeAfterDelete[T cmp.Ordered](tree *AvlTree[T]) *AvlTree[T] {
	if tree == nil {
		return tree
	}
	balanceFactor := getBalanceFactor(tree)
	if balanceFactor > 1 && getBalanceFactor(tree.leftNode) == 1 {
		return rotateLeft(tree)
	}
	if balanceFactor > 1 && getBalanceFactor(tree.leftNode) == -1 {
		tree.leftNode = rotateRight(tree.leftNode)
		return rotateLeft(tree)
	}
	if balanceFactor < -1 && getBalanceFactor(tree.rightNode) == -1 {
		return rotateRight(tree)
	}
	if balanceFactor < -1 && getBalanceFactor(tree.rightNode) == 1 {
		tree.rightNode = rotateLeft(tree.rightNode)
		return rotateRight(tree)
	}
	return tree
}

// getSmallestSonOnTheRight
// It returns the smallest son on the right of the tree.
// It's assumed that the root of the tree is not nil and has a right son at the first level
func getSmallestSonOnTheRight[T cmp.Ordered](tree *AvlTree[T]) *AvlTree[T] {
	smallestNode := tree.rightNode
	for ; smallestNode.leftNode != nil; smallestNode = smallestNode.leftNode {
	}
	return smallestNode
}
