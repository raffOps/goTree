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
		isRed: true,
	}

	node1 := &RedBlackTree[int]{
		value:    &nodeValue[int]{1},
		leftNode: node0,
		isRed:    true,
	}

	node2 := &RedBlackTree[int]{
		value:    &nodeValue[int]{2},
		leftNode: node1,
	}

	// want
	wantNode0 := &RedBlackTree[int]{
		value: &nodeValue[int]{0},
		isRed: true,
	}

	wantNode2 := &RedBlackTree[int]{
		value: &nodeValue[int]{2},
		isRed: true,
	}

	wantNode1 := &RedBlackTree[int]{
		value:     &nodeValue[int]{1},
		leftNode:  wantNode0,
		rightNode: wantNode2,
		isRed:     false,
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

	node1 := &RedBlackTree[int]{
		value: &nodeValue[int]{1},
		isRed: true,
	}

	node0 := &RedBlackTree[int]{
		value:     &nodeValue[int]{0},
		isRed:     false,
		rightNode: node1,
	}

	// want
	wantNode0 := &RedBlackTree[int]{
		value: &nodeValue[int]{0},
		isRed: true,
	}

	wantNode1 := &RedBlackTree[int]{
		value:    &nodeValue[int]{1},
		isRed:    false,
		leftNode: wantNode0,
	}

	testCases = append(testCases, RotateTestCase{tree: node0, want: wantNode1})
	return testCases
}
