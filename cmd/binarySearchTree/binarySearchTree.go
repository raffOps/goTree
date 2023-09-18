package binarySearchTree

import "cmp"

type BinarySearchTree[T cmp.Ordered] interface {
	GetValue() T
	OptimallyInsertNode(T)
}
