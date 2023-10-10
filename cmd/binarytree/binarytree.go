package binarytree

import (
	"cmp"
)

// SelfBalancingBinaryTree is an interface for self-balancing binary tree
type SelfBalancingBinaryTree[T cmp.Ordered] interface {
	InsertWithSelfBalancing(value T)
}
