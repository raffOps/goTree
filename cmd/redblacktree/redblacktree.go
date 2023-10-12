package redblacktree

import (
	"cmp"
	"fmt"
	"strings"
)

// nodeValue is a struct that stores the value of a node
type nodeValue[T cmp.Ordered] struct {
	value T
}

// RedBlackTree is a struct for an Red Black Tree of any ordered type.
// This is a self-balancing binary tree.
// The difference between the heights of the left and right subtree cannot be more than one for all nodes.
// For more information about theory, see https://en.wikipedia.org/wiki/Red-black_tree.
type RedBlackTree[T cmp.Ordered] struct {
	leftNode, rightNode, parentNode *RedBlackTree[T]
	value                           *nodeValue[T]
	isRed                           bool
}

// NewRedBlackTree factory function that returns a new RedBlackTree
func NewRedBlackTree[T cmp.Ordered]() *RedBlackTree[T] {
	return &RedBlackTree[T]{}
}

// GetValue returns the value stored in the tree.
// If empty, returns the zero value of the type of tree
func (tree *RedBlackTree[T]) GetValue() T {
	if !tree.IsEmpty() {
		return tree.value.value
	}
	var zeroValue T
	return zeroValue
}

// GetChild returns the child node in the direction specified
func (tree *RedBlackTree[T]) GetChild(direction string) *RedBlackTree[T] {
	if direction == "left" {
		return tree.leftNode
	}
	return tree.rightNode
}

// IsEmpty checks if the tree is empty
func (tree *RedBlackTree[T]) IsEmpty() bool {
	return tree.value == nil
}

// String returns a string representation of the tree.
func (tree *RedBlackTree[T]) String() string {
	return stringHelper[T](tree, 0)
}

func stringHelper[T cmp.Ordered](tree *RedBlackTree[T], level int) string {
	if tree == nil {
		return ""
	}

	indentation := strings.Repeat("  ", level)

	var parentString string
	if tree.parentNode == nil {
		parentString = ""
	} else {
		parentString = fmt.Sprintf("%v", tree.parentNode.GetValue())
	}

	return fmt.Sprintf("%s\nvalue: %v, \n%scolor: %s, \n%sparentNode: %v, \n%sleftNode: %v, \n%srightNode: %v",
		indentation,
		tree.GetValue(),
		indentation,
		tree.getColor(),
		indentation,
		parentString,
		indentation,
		stringHelper(tree.leftNode, level+1),
		indentation,
		stringHelper(tree.rightNode, level+1),
	)
}

func (tree *RedBlackTree[T]) getColor() string {
	if tree.isRed {
		return "red"
	}
	return "black"
}
