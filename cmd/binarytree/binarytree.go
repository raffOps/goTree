package binarytree

import (
	"cmp"
)

// BinaryTree is an interface for binary tree
type BinaryTree[T cmp.Ordered] interface {
	GetValue() T
	GetChild(direction string) *BinaryTree[T]
}

// SelfBalancingBinaryTree is an interface for self-balancing binary tree
type SelfBalancingBinaryTree[T cmp.Ordered] interface {
	InsertWithSelfBalancing(value T)
}
