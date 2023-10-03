package avltree

import "cmp"

// getSmallestSonOnTheRight
// It returns the smallest son on the right of the tree.
// It's assumed that the root of the tree is not nil and has a right son at the first level
func getSmallestSonOnTheRight[T cmp.Ordered](tree *AvlTree[T]) T {
	smallestNode := tree.rightNode
	for ; smallestNode.leftNode != nil; smallestNode = smallestNode.leftNode {
	}
	return smallestNode.value
}
