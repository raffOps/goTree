package avltree

import (
	"cmp"
	"reflect"
	"testing"
)

type rotateTestCase[T cmp.Ordered] struct {
	input  *avlTree[T]
	output *avlTree[T]
}

func getRotateLeftTestCases() []rotateTestCase[int] {
	testCases := []rotateTestCase[int]{
		{
			input: &avlTree[int]{value: 3, height: 2,
				leftNode: &avlTree[int]{value: 2, height: 1,
					leftNode: &avlTree[int]{value: 1, height: 0},
				},
			},
			output: &avlTree[int]{value: 2, height: 1,
				leftNode:  &avlTree[int]{value: 1, height: 0},
				rightNode: &avlTree[int]{value: 3, height: 0},
			},
		},
		{
			input: &avlTree[int]{value: 3, height: 2,
				leftNode: &avlTree[int]{value: 2, height: 1,
					leftNode: &avlTree[int]{value: 1, height: 0},
				},
				rightNode: &avlTree[int]{value: 4, height: 0},
			},
			output: &avlTree[int]{value: 2, height: 2,
				leftNode: &avlTree[int]{value: 1, height: 0},
				rightNode: &avlTree[int]{value: 3, height: 1,
					rightNode: &avlTree[int]{value: 4, height: 0},
				},
			},
		},
	}
	return testCases
}

func TestRotateLeft(t *testing.T) {
	for index, testCase := range getRotateLeftTestCases() {
		got := rotateLeft(testCase.input)
		if !reflect.DeepEqual(got, testCase.output) {
			t.Errorf("Case %d\ngot: \n%v\n\nwant: \n%v", index, got, testCase.output)
		}
	}
}

func getRotateRightTestCases() []rotateTestCase[int] {
	var (
		testCases                                             []rotateTestCase[int]
		node1, node2, node3, input                            *avlTree[int]
		rotatedNode1, rotatedNode2, rotatedNode3, rotatedTree *avlTree[int]
	)

	//test case1
	node3 = &avlTree[int]{value: 3}
	node2 = &avlTree[int]{rightNode: node3, value: 2, height: 1}
	node1 = &avlTree[int]{rightNode: node2, value: 1, height: 2}
	input = node1

	rotatedNode3 = &avlTree[int]{value: 3}
	rotatedNode1 = &avlTree[int]{value: 1}
	rotatedNode2 = &avlTree[int]{leftNode: rotatedNode1, rightNode: rotatedNode3, value: 2, height: 1}
	rotatedTree = rotatedNode2

	testCases = append(testCases, rotateTestCase[int]{input: input, output: rotatedTree})

	return testCases
}

func TestRotateRight(t *testing.T) {
	for index, testCase := range getRotateRightTestCases() {
		got := rotateRight(testCase.input)
		if !reflect.DeepEqual(got, testCase.output) {
			t.Errorf("Case %d\ngot: \n%v\n\nwant: \n%v", index, got, testCase.output)
		}
	}
}

func getDoubleRotateLeftTestCases() []rotateTestCase[int] {
	var (
		testCases                                        []rotateTestCase[int]
		node1, node2, node3, input                       *avlTree[int]
		rotatedNode1, rotatedNode2, rotatedNode3, output *avlTree[int]
	)

	// test case 1
	node2 = &avlTree[int]{value: 2}
	node1 = &avlTree[int]{rightNode: node2, value: 1}
	node3 = &avlTree[int]{leftNode: node1, value: 3}
	input = node3

	rotatedNode3 = &avlTree[int]{value: 3}
	rotatedNode1 = &avlTree[int]{value: 1}
	rotatedNode2 = &avlTree[int]{leftNode: rotatedNode1, rightNode: rotatedNode3, value: 2, height: 1}
	output = rotatedNode2

	testCases = append(testCases, rotateTestCase[int]{input: input, output: output})

	return testCases
}

func TestDoubleRotateLeft(t *testing.T) {
	for index, testCase := range getDoubleRotateLeftTestCases() {
		got := doubleRotateLeft(testCase.input)
		if !reflect.DeepEqual(got, testCase.output) {
			t.Errorf("Case %d\ngot: \n%v\n\nwant: \n%v", index, got, testCase.output)

		}
	}
}

func getDoubleRotateRightTestCases() []rotateTestCase[int] {
	var (
		testCases                                        []rotateTestCase[int]
		node1, node2, node3, input                       *avlTree[int]
		rotatedNode1, rotatedNode2, rotatedNode3, output *avlTree[int]
	)

	// test case 1
	node2 = &avlTree[int]{value: 2}
	node3 = &avlTree[int]{leftNode: node2, value: 3}
	node1 = &avlTree[int]{rightNode: node3, value: 1}
	input = node1

	rotatedNode3 = &avlTree[int]{value: 3}
	rotatedNode1 = &avlTree[int]{value: 1}
	rotatedNode2 = &avlTree[int]{leftNode: rotatedNode1, rightNode: rotatedNode3, value: 2, height: 1}
	output = rotatedNode2

	testCases = append(testCases, rotateTestCase[int]{input: input, output: output})

	return testCases
}

func TestDoubleRotateRight(t *testing.T) {
	for index, testCase := range getDoubleRotateRightTestCases() {
		got := doubleRotateRight(testCase.input)
		if !reflect.DeepEqual(got, testCase.output) {
			t.Errorf("Case %d\ngot: \n%v\n\nwant: \n%v", index, got, testCase.output)
		}
	}
}
