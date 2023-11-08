package redblacktree

import (
	"reflect"
	"testing"
)

type RotateTestCase struct {
	tree *RedBlackTree[int]
	want *RedBlackTree[int]
}

func TestRedBlackTree_rotateToRight(t *testing.T) {

	testCases := getRotateToRightTestCases()

	// assertions
	for index, testCase := range testCases {
		rotatedTree := rotateToRight(testCase.tree)
		if !reflect.DeepEqual(rotatedTree, testCase.want) {
			t.Errorf("test case %d failed\nwant: %v\n\ngot: %v",
				index,
				testCase.want,
				rotatedTree,
			)
		}

	}
}

func getRotateToRightTestCases() []RotateTestCase {
	// test cases definition
	var testCases []RotateTestCase

	// test case 1: root is right child of his parent

	// tree
	node0 := &RedBlackTree[int]{}
	node1 := &RedBlackTree[int]{}
	node2 := &RedBlackTree[int]{}
	node3 := &RedBlackTree[int]{}
	node4 := &RedBlackTree[int]{}
	node5 := &RedBlackTree[int]{}

	*node1 = RedBlackTree[int]{
		value:      &nodeValue[int]{1},
		parentNode: node2,
	}

	*node2 = RedBlackTree[int]{
		value:      &nodeValue[int]{2},
		leftNode:   node1,
		rightNode:  node3,
		parentNode: node4,
	}

	*node3 = RedBlackTree[int]{
		value:      &nodeValue[int]{3},
		parentNode: node2,
	}

	*node4 = RedBlackTree[int]{
		value:      &nodeValue[int]{4},
		leftNode:   node2,
		rightNode:  node5,
		parentNode: node0,
	}

	*node5 = RedBlackTree[int]{
		value:      &nodeValue[int]{5},
		parentNode: node4,
	}

	*node0 = RedBlackTree[int]{
		value:      &nodeValue[int]{0},
		rightNode:  node4,
		parentNode: nil,
	}
	// want
	wantNode0 := &RedBlackTree[int]{}
	wantNode1 := &RedBlackTree[int]{}
	wantNode2 := &RedBlackTree[int]{}
	wantNode3 := &RedBlackTree[int]{}
	wantNode4 := &RedBlackTree[int]{}
	wantNode5 := &RedBlackTree[int]{}

	*wantNode1 = RedBlackTree[int]{
		value:      &nodeValue[int]{1},
		parentNode: wantNode2,
	}

	*wantNode3 = RedBlackTree[int]{
		value:      &nodeValue[int]{3},
		parentNode: wantNode4,
	}

	*wantNode5 = RedBlackTree[int]{
		value:      &nodeValue[int]{5},
		parentNode: wantNode4,
	}

	*wantNode4 = RedBlackTree[int]{
		value:      &nodeValue[int]{4},
		leftNode:   wantNode3,
		rightNode:  wantNode5,
		parentNode: wantNode2,
	}

	*wantNode2 = RedBlackTree[int]{
		value:      &nodeValue[int]{2},
		leftNode:   wantNode1,
		rightNode:  wantNode4,
		parentNode: wantNode0,
	}

	*wantNode0 = RedBlackTree[int]{
		value:      &nodeValue[int]{0},
		rightNode:  wantNode2,
		parentNode: nil,
	}

	testCases = append(testCases, RotateTestCase{tree: node4, want: wantNode2})

	// test case 2: root is left child of his parent

	// tree

	node8 := &RedBlackTree[int]{}
	node9 := &RedBlackTree[int]{}
	node10 := &RedBlackTree[int]{}
	node11 := &RedBlackTree[int]{}
	node12 := &RedBlackTree[int]{}

	*node8 = RedBlackTree[int]{
		value:      &nodeValue[int]{8},
		parentNode: node9,
	}

	*node9 = RedBlackTree[int]{
		value:      &nodeValue[int]{9},
		leftNode:   node8,
		rightNode:  node10,
		parentNode: node11,
	}

	*node10 = RedBlackTree[int]{
		value:      &nodeValue[int]{10},
		parentNode: node9,
	}

	*node11 = RedBlackTree[int]{
		value:      &nodeValue[int]{11},
		leftNode:   node9,
		parentNode: node12,
	}

	*node12 = RedBlackTree[int]{
		value:    &nodeValue[int]{12},
		leftNode: node11,
	}

	// want
	wantNode8 := &RedBlackTree[int]{}
	wantNode9 := &RedBlackTree[int]{}
	wantNode10 := &RedBlackTree[int]{}
	wantNode11 := &RedBlackTree[int]{}
	wantNode12 := &RedBlackTree[int]{}

	*wantNode8 = RedBlackTree[int]{
		value:      &nodeValue[int]{8},
		parentNode: wantNode9,
	}

	*wantNode10 = RedBlackTree[int]{
		value:      &nodeValue[int]{10},
		parentNode: wantNode11,
	}

	*wantNode11 = RedBlackTree[int]{
		value:      &nodeValue[int]{11},
		leftNode:   node10,
		parentNode: wantNode9,
	}

	*wantNode9 = RedBlackTree[int]{
		value:      &nodeValue[int]{9},
		leftNode:   wantNode8,
		rightNode:  wantNode11,
		parentNode: wantNode12,
	}

	*wantNode12 = RedBlackTree[int]{
		value:    &nodeValue[int]{12},
		leftNode: wantNode9,
	}

	testCases = append(testCases, RotateTestCase{tree: node11, want: wantNode9})
	return testCases
}

func TestRedBlackTree_rotateToLeft(t *testing.T) {
	testCases := getRotateToLeftTestCases()

	// assertions
	for index, testCase := range testCases {
		rotatedTree := rotateToLeft(testCase.tree)
		if !reflect.DeepEqual(rotatedTree, testCase.want) {
			t.Errorf("test case %d failed\nwant: %v\n\ngot: %v",
				index,
				testCase.want,
				rotatedTree,
			)
		}
	}
}

func getRotateToLeftTestCases() []RotateTestCase {
	// test cases definition
	var testCases []RotateTestCase

	// test case 1: root is right child of his parent

	// tree
	node0 := &RedBlackTree[int]{}
	node1 := &RedBlackTree[int]{}
	node2 := &RedBlackTree[int]{}
	node3 := &RedBlackTree[int]{}
	node4 := &RedBlackTree[int]{}
	node5 := &RedBlackTree[int]{}

	*node1 = RedBlackTree[int]{
		value:      &nodeValue[int]{1},
		parentNode: node2,
	}

	*node3 = RedBlackTree[int]{
		value:      &nodeValue[int]{3},
		parentNode: node4,
	}

	*node5 = RedBlackTree[int]{
		value:      &nodeValue[int]{5},
		parentNode: node4,
	}

	*node4 = RedBlackTree[int]{
		value:      &nodeValue[int]{4},
		leftNode:   node3,
		rightNode:  node5,
		parentNode: node2,
	}

	*node2 = RedBlackTree[int]{
		value:      &nodeValue[int]{2},
		leftNode:   node1,
		rightNode:  node4,
		parentNode: node0,
	}

	*node0 = RedBlackTree[int]{
		value:     &nodeValue[int]{0},
		rightNode: node2,
	}

	// want
	wantNode0 := &RedBlackTree[int]{}
	wantNode1 := &RedBlackTree[int]{}
	wantNode2 := &RedBlackTree[int]{}
	wantNode3 := &RedBlackTree[int]{}
	wantNode4 := &RedBlackTree[int]{}
	wantNode5 := &RedBlackTree[int]{}

	*wantNode1 = RedBlackTree[int]{
		value:      &nodeValue[int]{1},
		parentNode: wantNode2,
	}

	*wantNode3 = RedBlackTree[int]{
		value:      &nodeValue[int]{3},
		parentNode: wantNode2,
	}

	*wantNode2 = RedBlackTree[int]{
		value:      &nodeValue[int]{2},
		leftNode:   wantNode1,
		rightNode:  wantNode3,
		parentNode: wantNode4,
	}

	*wantNode5 = RedBlackTree[int]{
		value:      &nodeValue[int]{5},
		parentNode: wantNode4,
	}

	*wantNode4 = RedBlackTree[int]{
		value:      &nodeValue[int]{4},
		leftNode:   wantNode2,
		rightNode:  wantNode5,
		parentNode: wantNode0,
	}

	*wantNode0 = RedBlackTree[int]{
		value:      &nodeValue[int]{0},
		rightNode:  wantNode4,
		parentNode: nil,
	}

	testCases = append(testCases, RotateTestCase{tree: node2, want: wantNode4})

	// test case 2: root is left child of his parent
	// tree
	node7 := &RedBlackTree[int]{}
	node4 = &RedBlackTree[int]{}
	node2 = &RedBlackTree[int]{}
	node5 = &RedBlackTree[int]{}
	node3 = &RedBlackTree[int]{}
	node1 = &RedBlackTree[int]{}

	*node1 = RedBlackTree[int]{
		value:      &nodeValue[int]{1},
		parentNode: node2,
	}

	*node3 = RedBlackTree[int]{
		value:      &nodeValue[int]{3},
		parentNode: node4,
	}

	*node5 = RedBlackTree[int]{
		value:      &nodeValue[int]{5},
		parentNode: node4,
	}

	*node4 = RedBlackTree[int]{
		value:      &nodeValue[int]{4},
		leftNode:   node3,
		rightNode:  node5,
		parentNode: node2,
	}

	*node2 = RedBlackTree[int]{
		value:      &nodeValue[int]{2},
		leftNode:   node1,
		rightNode:  node4,
		parentNode: node7,
	}

	*node7 = RedBlackTree[int]{
		value:      &nodeValue[int]{7},
		leftNode:   node2,
		parentNode: nil,
	}

	// want
	wantNode7 := &RedBlackTree[int]{}
	wantNode4 = &RedBlackTree[int]{}
	wantNode2 = &RedBlackTree[int]{}
	wantNode5 = &RedBlackTree[int]{}
	wantNode3 = &RedBlackTree[int]{}
	wantNode1 = &RedBlackTree[int]{}

	*wantNode1 = RedBlackTree[int]{
		value:      &nodeValue[int]{1},
		parentNode: wantNode2,
	}

	*wantNode3 = RedBlackTree[int]{
		value:      &nodeValue[int]{3},
		parentNode: wantNode2,
	}

	*wantNode2 = RedBlackTree[int]{
		value:      &nodeValue[int]{2},
		leftNode:   wantNode1,
		rightNode:  wantNode3,
		parentNode: wantNode4,
	}

	*wantNode5 = RedBlackTree[int]{
		value:      &nodeValue[int]{5},
		parentNode: wantNode4,
	}

	*wantNode4 = RedBlackTree[int]{
		value:      &nodeValue[int]{4},
		leftNode:   wantNode2,
		rightNode:  wantNode5,
		parentNode: wantNode7,
	}

	*wantNode7 = RedBlackTree[int]{
		value:      &nodeValue[int]{7},
		leftNode:   wantNode4,
		parentNode: nil,
	}

	testCases = append(testCases, RotateTestCase{tree: node2, want: wantNode4})
	return testCases
}
