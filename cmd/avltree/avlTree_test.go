package avltree

import (
	"cmp"
	"reflect"
	"testing"
)

func TestNewAVLTree(t *testing.T) {
	got := NewAVLTree(5)
	if got.value != 5 || got.leftNode != nil || got.rightNode != nil {
		t.Errorf("Case int tree")
	}

	gotString := NewAVLTree("maria")
	if gotString.value != "maria" || gotString.leftNode != nil || gotString.rightNode != nil {
		t.Errorf("Case string tree")
	}
}

func TestAVLTree_GetRootValue(t *testing.T) {
	type TestCase[T cmp.Ordered] struct {
		input  *AVLTree[T]
		output T
	}
	testCases := []TestCase[int]{
		{input: &AVLTree[int]{value: 8}, output: 8},
		{input: &AVLTree[int]{value: 5}, output: 5},
	}

	for index, testCase := range testCases {
		got := testCase.input.GetRootValue()
		if got != testCase.output {
			t.Errorf("Case %d", index)
		}
	}
}

type rotateTestCase[T cmp.Ordered] struct {
	input  *AVLTree[T]
	output *AVLTree[T]
}

func getRotateLeftTestCases() []rotateTestCase[int] {
	var (
		testCases                                                                                  []rotateTestCase[int]
		node0, node1, node2, node3, node4, node5, input                                            *AVLTree[int]
		rotatedNode0, rotatedNode1, rotatedNode2, rotatedNode3, rotatedNode4, rotatedNode5, output *AVLTree[int]
	)

	//test case1
	node1 = &AVLTree[int]{value: 1}
	node2 = &AVLTree[int]{leftNode: node1, value: 2, height: 1}
	node3 = &AVLTree[int]{leftNode: node2, value: 3, height: 2}
	input = node3

	rotatedNode3 = &AVLTree[int]{value: 3}
	rotatedNode1 = &AVLTree[int]{value: 1}
	rotatedNode2 = &AVLTree[int]{leftNode: rotatedNode1, rightNode: rotatedNode3, value: 2, height: 1}
	output = rotatedNode2

	testCases = append(testCases, rotateTestCase[int]{input: input, output: output})

	// test case 2
	node0 = &AVLTree[int]{value: 0}
	node3 = &AVLTree[int]{value: 3}
	node5 = &AVLTree[int]{value: 5}
	node1 = &AVLTree[int]{leftNode: node0, value: 1, height: 1}
	node2 = &AVLTree[int]{leftNode: node1, rightNode: node3, value: 2, height: 2}
	node4 = &AVLTree[int]{leftNode: node2, rightNode: node5, value: 4, height: 3}
	input = node4

	rotatedNode0 = &AVLTree[int]{value: 0}
	rotatedNode3 = &AVLTree[int]{value: 3}
	rotatedNode5 = &AVLTree[int]{value: 5}
	rotatedNode1 = &AVLTree[int]{leftNode: rotatedNode0, value: 1, height: 1}
	rotatedNode4 = &AVLTree[int]{leftNode: rotatedNode3, rightNode: rotatedNode5, value: 4, height: 1}
	rotatedNode2 = &AVLTree[int]{leftNode: rotatedNode1, rightNode: rotatedNode4, value: 2, height: 2}
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
		node1, node2, node3, input                            *AVLTree[int]
		rotatedNode1, rotatedNode2, rotatedNode3, rotatedTree *AVLTree[int]
	)

	//test case1
	node3 = &AVLTree[int]{value: 3}
	node2 = &AVLTree[int]{rightNode: node3, value: 2, height: 1}
	node1 = &AVLTree[int]{rightNode: node2, value: 1, height: 2}
	input = node1

	rotatedNode3 = &AVLTree[int]{value: 3}
	rotatedNode1 = &AVLTree[int]{value: 1}
	rotatedNode2 = &AVLTree[int]{leftNode: rotatedNode1, rightNode: rotatedNode3, value: 2, height: 1}
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
		node1, node2, node3, input                       *AVLTree[int]
		rotatedNode1, rotatedNode2, rotatedNode3, output *AVLTree[int]
	)

	// test case 1
	node2 = &AVLTree[int]{value: 2}
	node1 = &AVLTree[int]{rightNode: node2, value: 1}
	node3 = &AVLTree[int]{leftNode: node1, value: 3}
	input = node3

	rotatedNode3 = &AVLTree[int]{value: 3}
	rotatedNode1 = &AVLTree[int]{value: 1}
	rotatedNode2 = &AVLTree[int]{leftNode: rotatedNode1, rightNode: rotatedNode3, value: 2, height: 1}
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
		node1, node2, node3, input                       *AVLTree[int]
		rotatedNode1, rotatedNode2, rotatedNode3, output *AVLTree[int]
	)

	// test case 1
	node2 = &AVLTree[int]{value: 2}
	node3 = &AVLTree[int]{leftNode: node2, value: 3}
	node1 = &AVLTree[int]{rightNode: node3, value: 1}
	input = node1

	rotatedNode3 = &AVLTree[int]{value: 3}
	rotatedNode1 = &AVLTree[int]{value: 1}
	rotatedNode2 = &AVLTree[int]{leftNode: rotatedNode1, rightNode: rotatedNode3, value: 2, height: 1}
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

//func getRotateRightTestCases() []rotateTestCase[int] {
//	var (
//		testCases                                []rotateTestCase[int]
//		node1, node2, node3                      *avlTree[int]
//		rotatedNode1, rotatedNode2, rotatedNode3 *avlTree[int]
//	)
//
//	// test case1
//	node3 = &avlTree[int]{value: 3, height: 2}
//	node2 = &avlTree[int]{rightNode: node3, value: 2, height: 1}
//	node1 = &avlTree[int]{rightNode: node2, value: 1}
//	input := node1
//
//	rotatedNode3 = &avlTree[int]{value: 3, height: 1}
//	rotatedNode1 = &avlTree[int]{value: 1, height: 1}
//	rotatedNode2 = &avlTree[int]{leftNode: rotatedNode1, rightNode: rotatedNode3, value: 2}
//	output := rotatedNode2
//
//	testCases = append(testCases, rotateTestCase[int]{input: input, output: output})
//
//	return testCases
//}
//
//type AVLTreeInsertNodeTestCase[T cmp.Ordered] struct {
//	input []T
//	input  *avlTree[T]
//}

//func getTestCases() []AVLTreeInsertNodeTestCase[int] {
//	// test case 1 & 2
//	input1 := []int{1, 2, 3, 4, 5, 6}
//	input2 := []int{6, 5, 4, 3, 2, 1}
//
//	treeNode1 := avlTree[int]{value: 1}
//	treeNode3 := avlTree[int]{value: 3}
//	treeNode4 := avlTree[int]{value: 4}
//	treeNode6 := avlTree[int]{value: 6}
//	treeNode2 := avlTree[int]{leftNode: &treeNode1, rightNode: &treeNode3, value: 2}
//	treeNode5 := avlTree[int]{leftNode: &treeNode4, rightNode: &treeNode6, value: 6}
//	input := avlTree[int]{leftNode: &treeNode2, rightNode: &treeNode5, value: 3}
//
//	return []AVLTreeInsertNodeTestCase[int]{
//		AVLTreeInsertNodeTestCase[int]{input: input1, input: &input},
//		AVLTreeInsertNodeTestCase[int]{input: input2, input: &input},
//	}
//}
//
//func TestInsertNode(t *testing.T) {
//	testCases := getTestCases()
//	for index, testCase := range testCases {
//		input := NewAVLTree[int]()
//		for _, value := range testCase.input {
//			input.InsertValue(value)
//		}
//		if input != testCase.input {
//			t.Errorf("Case %d: %v", index, testCase.input)
//		}
//	}
//}
