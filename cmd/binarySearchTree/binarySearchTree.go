package binarySearchTree

import "cmp"

type Node interface {
	cmp.Ordered
}

type BinarySearchTree[T Node] interface {
	OptimallyInsertNode(T)
}
