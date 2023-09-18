package redBlackTree

import (
	"cmp"
)

type RedBlackTree[T cmp.Ordered] struct {
	leftNode, rightNode *RedBlackTree[T]
	value               T
}

func NewNodeTree[T cmp.Ordered](leftNode, rightNode *RedBlackTree[T], value T) *RedBlackTree[T] {
	tree_ := &RedBlackTree[T]{leftNode: leftNode, rightNode: rightNode, value: value}
	return tree_
}

func (t *RedBlackTree[T]) GetValue() T {
	return t.value
}

func (t *RedBlackTree[T]) InsertNode(newValue T) {
	var zeroValue T
	if cmp.Compare(t.GetValue(), zeroValue) == 0 {
		t.value = newValue
	} else {
		comparison := cmp.Compare(newValue, t.value)
		if comparison == -1 && t.leftNode == nil { // less than, leaf node
			leftNode := NewNodeTree(nil, nil, newValue)
			t.leftNode = leftNode
		} else if comparison == -1 && t.leftNode != nil { // less than, not leaf node
			t.leftNode.InsertNode(newValue)
		} else if comparison > -1 && t.rightNode == nil { // greater of equal than, leaf node
			rightNode := NewNodeTree(nil, nil, newValue)
			t.rightNode = rightNode
		} else if comparison > -1 && t.rightNode != nil { // greater of equal than, not leaf node
			t.rightNode.InsertNode(newValue)
		}
	}
}

func (t *RedBlackTree[T]) SearchValue(value T) bool {
	if cmp.Compare(value, t.value) == 0 {
		return true
	}
	if cmp.Compare(value, t.value) == -1 && t.leftNode != nil {
		return t.leftNode.SearchValue(value)
	}
	if cmp.Compare(value, t.value) == 1 && t.rightNode != nil {
		return t.rightNode.SearchValue(value)
	}
	return false
}
