package redblacktree

import (
	"reflect"
	"testing"
)

type RotateTestCase struct {
	tree *RedBlackTree[int]
	want *RedBlackTree[int]
}

func TestRotateToRight(t *testing.T) {
	testCases := getRotateToRightTestCases()

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
	var testCases []RotateTestCase

	// tree
	node0 := &RedBlackTree[int]{
		value: &nodeValue[int]{0},
	}

	node1 := &RedBlackTree[int]{
		value:    &nodeValue[int]{1},
		leftNode: node0,
	}

	node3 := &RedBlackTree[int]{
		value: &nodeValue[int]{3},
	}

	node2 := &RedBlackTree[int]{
		value:     &nodeValue[int]{2},
		leftNode:  node1,
		rightNode: node3,
	}

	// want
	wantNode0 := &RedBlackTree[int]{
		value: &nodeValue[int]{0},
	}

	wantNode3 := &RedBlackTree[int]{
		value: &nodeValue[int]{3},
	}

	wantNode2 := &RedBlackTree[int]{
		value:     &nodeValue[int]{2},
		rightNode: wantNode3,
	}

	wantNode1 := &RedBlackTree[int]{
		value:     &nodeValue[int]{1},
		leftNode:  wantNode0,
		rightNode: wantNode2,
	}

	testCases = append(testCases, RotateTestCase{tree: node2, want: wantNode1})

	return testCases
}

func TestRotateToLeft(t *testing.T) {
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

	// tree
	node0 := &RedBlackTree[int]{
		value: &nodeValue[int]{0},
	}

	node3 := &RedBlackTree[int]{
		value: &nodeValue[int]{3},
	}

	node2 := &RedBlackTree[int]{
		value:     &nodeValue[int]{2},
		rightNode: node3,
	}

	node1 := &RedBlackTree[int]{
		value:     &nodeValue[int]{1},
		leftNode:  node0,
		rightNode: node2,
	}

	// want
	wantNode0 := &RedBlackTree[int]{
		value: &nodeValue[int]{0},
	}

	wantNode1 := &RedBlackTree[int]{
		value:    &nodeValue[int]{1},
		leftNode: wantNode0,
	}

	wantNode3 := &RedBlackTree[int]{
		value: &nodeValue[int]{3},
	}

	wantNode2 := &RedBlackTree[int]{
		value:     &nodeValue[int]{2},
		leftNode:  wantNode1,
		rightNode: wantNode3,
	}

	testCases = append(testCases, RotateTestCase{tree: node1, want: wantNode2})
	return testCases
}
