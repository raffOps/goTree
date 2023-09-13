package tree

import (
	"cmp"
)

type Node interface {
	cmp.Ordered
}

type BinaryTree[T Node] struct {
	leftNode, rightNode *BinaryTree[T]
	value               T
}

func NewNodeTree[T Node](leftNode, rightNode *BinaryTree[T], value T) *BinaryTree[T] {
	tree_ := BinaryTree[T]{leftNode: leftNode, rightNode: rightNode, value: value}
	return &tree_
}

func (t *BinaryTree[T]) GetValue() T {
	return t.value
}

func (t *BinaryTree[T]) InsertNewNode(newValue T) {
	var zeroValue T
	if t.value == zeroValue {
		t.value = newValue
	} else {
		comparison := cmp.Compare(newValue, t.value)
		if comparison == -1 && t.leftNode == nil { // less than, leaf node
			leftNode := NewNodeTree(nil, nil, newValue)
			t.leftNode = leftNode
		} else if comparison == -1 && t.leftNode != nil { // less than, not leaf node
			t.leftNode.InsertNewNode(newValue)
		} else if comparison > -1 && t.rightNode == nil { // greater of equal than, leaf node
			rightNode := NewNodeTree(nil, nil, newValue)
			t.rightNode = rightNode
		} else if comparison > -1 && t.rightNode != nil { // greater of equal than, not leaf node
			t.rightNode.InsertNewNode(newValue)
		}
	}
}
