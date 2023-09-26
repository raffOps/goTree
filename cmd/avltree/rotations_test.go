package avltree

import (
	"cmp"
	"reflect"
	"testing"
)

type rotateTestCase[T cmp.Ordered] struct {
	input  *AvlTree[T]
	output *AvlTree[T]
}

func getRotateLeftTestCases() []rotateTestCase[int] {
	var (
		testCases                                                                                  []rotateTestCase[int]
		node0, node1, node2, node3, node4, node5, input                                            *AvlTree[int]
		rotatedNode0, rotatedNode1, rotatedNode2, rotatedNode3, rotatedNode4, rotatedNode5, output *AvlTree[int]
	)

	//test case1
	node1 = &AvlTree[int]{value: 1}
	node2 = &AvlTree[int]{leftNode: node1, value: 2, height: 1}
	node3 = &AvlTree[int]{leftNode: node2, value: 3, height: 2}
	input = node3

	rotatedNode3 = &AvlTree[int]{value: 3}
	rotatedNode1 = &AvlTree[int]{value: 1}
	rotatedNode2 = &AvlTree[int]{leftNode: rotatedNode1, rightNode: rotatedNode3, value: 2, height: 1}
	output = rotatedNode2

	testCases = append(testCases, rotateTestCase[int]{input: input, output: output})

	// test case 2
	node0 = &AvlTree[int]{value: 0}
	node3 = &AvlTree[int]{value: 3}
	node5 = &AvlTree[int]{value: 5}
	node1 = &AvlTree[int]{leftNode: node0, value: 1, height: 1}
	node2 = &AvlTree[int]{leftNode: node1, rightNode: node3, value: 2, height: 2}
	node4 = &AvlTree[int]{leftNode: node2, rightNode: node5, value: 4, height: 3}
	input = node4

	rotatedNode0 = &AvlTree[int]{value: 0}
	rotatedNode3 = &AvlTree[int]{value: 3}
	rotatedNode5 = &AvlTree[int]{value: 5}
	rotatedNode1 = &AvlTree[int]{leftNode: rotatedNode0, value: 1, height: 1}
	rotatedNode4 = &AvlTree[int]{leftNode: rotatedNode3, rightNode: rotatedNode5, value: 4, height: 1}
	rotatedNode2 = &AvlTree[int]{leftNode: rotatedNode1, rightNode: rotatedNode4, value: 2, height: 2}
	output = rotatedNode2

	testCases = append(testCases, rotateTestCase[int]{input: input, output: output})

	return testCases
}

func TestRotateLeft(t *testing.T) {
	for index, testCase := range getRotateLeftTestCases() {
		got := rotateLeft(testCase.input)
		if reflect.DeepEqual(got, testCase.output) == false {
			t.Errorf("Case %d", index)
		}
	}
}

func getRotateRightTestCases() []rotateTestCase[int] {
	var (
		testCases                                             []rotateTestCase[int]
		node1, node2, node3, input                            *AvlTree[int]
		rotatedNode1, rotatedNode2, rotatedNode3, rotatedTree *AvlTree[int]
	)

	//test case1
	node3 = &AvlTree[int]{value: 3}
	node2 = &AvlTree[int]{rightNode: node3, value: 2, height: 1}
	node1 = &AvlTree[int]{rightNode: node2, value: 1, height: 2}
	input = node1

	rotatedNode3 = &AvlTree[int]{value: 3}
	rotatedNode1 = &AvlTree[int]{value: 1}
	rotatedNode2 = &AvlTree[int]{leftNode: rotatedNode1, rightNode: rotatedNode3, value: 2, height: 1}
	rotatedTree = rotatedNode2

	testCases = append(testCases, rotateTestCase[int]{input: input, output: rotatedTree})

	return testCases
}

func TestRotateRight(t *testing.T) {
	for index, testCase := range getRotateRightTestCases() {
		got := rotateRight(testCase.input)
		if reflect.DeepEqual(got, testCase.output) == false {
			t.Errorf("Case %d", index)
		}
	}
}

func getDoubleRotateLeftTestCases() []rotateTestCase[int] {
	var (
		testCases                                        []rotateTestCase[int]
		node1, node2, node3, input                       *AvlTree[int]
		rotatedNode1, rotatedNode2, rotatedNode3, output *AvlTree[int]
	)

	// test case 1
	node2 = &AvlTree[int]{value: 2}
	node1 = &AvlTree[int]{rightNode: node2, value: 1}
	node3 = &AvlTree[int]{leftNode: node1, value: 3}
	input = node3

	rotatedNode3 = &AvlTree[int]{value: 3}
	rotatedNode1 = &AvlTree[int]{value: 1}
	rotatedNode2 = &AvlTree[int]{leftNode: rotatedNode1, rightNode: rotatedNode3, value: 2, height: 1}
	output = rotatedNode2

	testCases = append(testCases, rotateTestCase[int]{input: input, output: output})

	return testCases
}

func TestDoubleRotateLeft(t *testing.T) {
	for index, testCase := range getDoubleRotateLeftTestCases() {
		got := doubleRotateLeft(testCase.input)
		if reflect.DeepEqual(got, testCase.output) == false {
			t.Errorf("Case %d", index)
		}
	}
}

func getDoubleRotateRightTestCases() []rotateTestCase[int] {
	var (
		testCases                                        []rotateTestCase[int]
		node1, node2, node3, input                       *AvlTree[int]
		rotatedNode1, rotatedNode2, rotatedNode3, output *AvlTree[int]
	)

	// test case 1
	node2 = &AvlTree[int]{value: 2}
	node3 = &AvlTree[int]{leftNode: node2, value: 3}
	node1 = &AvlTree[int]{rightNode: node3, value: 1}
	input = node1

	rotatedNode3 = &AvlTree[int]{value: 3}
	rotatedNode1 = &AvlTree[int]{value: 1}
	rotatedNode2 = &AvlTree[int]{leftNode: rotatedNode1, rightNode: rotatedNode3, value: 2, height: 1}
	output = rotatedNode2

	testCases = append(testCases, rotateTestCase[int]{input: input, output: output})

	return testCases
}

func TestDoubleRotateRight(t *testing.T) {
	for index, testCase := range getDoubleRotateRightTestCases() {
		got := doubleRotateRight(testCase.input)
		if reflect.DeepEqual(got, testCase.output) == false {
			t.Errorf("Case %d", index)
		}
	}
}
