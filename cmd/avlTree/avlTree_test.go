package avlTree

import (
	"cmp"
	"reflect"
	"testing"
)

type rotateTestCase[T cmp.Ordered] struct {
	input       *AVLTree[T]
	rotatedTree *AVLTree[T]
}

func getRotateLeftTestCases() []rotateTestCase[int] {
	var (
		testCases                                []rotateTestCase[int]
		node1, node2, node3                      *AVLTree[int]
		rotatedNode1, rotatedNode2, rotatedNode3 *AVLTree[int]
	)

	// test case1
	node1 = &AVLTree[int]{value: 1}
	node2 = &AVLTree[int]{leftNode: node1, value: 2, height: 1}
	node3 = &AVLTree[int]{leftNode: node2, value: 3, height: 2}
	input := node3

	rotatedNode3 = &AVLTree[int]{value: 3, height: 1}
	rotatedNode1 = &AVLTree[int]{value: 1, height: 1}
	rotatedNode2 = &AVLTree[int]{leftNode: rotatedNode1, rightNode: rotatedNode3, value: 2, height: 0}
	rotatedTree := rotatedNode2

	testCases = append(testCases, rotateTestCase[int]{input: input, rotatedTree: rotatedTree})

	return testCases
}

func TestRotateLeft(t *testing.T) {
	for index, testCase := range getRotateLeftTestCases() {
		got := rotateLeft(testCase.input)
		if reflect.DeepEqual(got, testCase.rotatedTree) == false {
			t.Errorf("Case %d", index)
		}
	}
}

func getRotateRightTestCases() []rotateTestCase[int] {
	var (
		testCases                                []rotateTestCase[int]
		node1, node2, node3                      *AVLTree[int]
		rotatedNode1, rotatedNode2, rotatedNode3 *AVLTree[int]
	)

	// test case1
	node3 = &AVLTree[int]{value: 3, height: 2}
	node2 = &AVLTree[int]{rightNode: node3, value: 2, height: 1}
	node1 = &AVLTree[int]{rightNode: node2, value: 1, height: 0}
	input := node1

	rotatedNode3 = &AVLTree[int]{value: 3, height: 1}
	rotatedNode1 = &AVLTree[int]{value: 1, height: 1}
	rotatedNode2 = &AVLTree[int]{leftNode: rotatedNode1, rightNode: rotatedNode3, value: 2, height: 0}
	rotatedTree := rotatedNode2

	testCases = append(testCases, rotateTestCase[int]{input: input, rotatedTree: rotatedTree})

	return testCases
}

type AVLTreeInsertNodeTestCase[T cmp.Ordered] struct {
	input []T
	tree  *AVLTree[T]
}

//func getTestCases() []AVLTreeInsertNodeTestCase[int] {
//	// test case 1 & 2
//	input1 := []int{1, 2, 3, 4, 5, 6}
//	input2 := []int{6, 5, 4, 3, 2, 1}
//
//	treeNode1 := AVLTree[int]{value: 1}
//	treeNode3 := AVLTree[int]{value: 3}
//	treeNode4 := AVLTree[int]{value: 4}
//	treeNode6 := AVLTree[int]{value: 6}
//	treeNode2 := AVLTree[int]{leftNode: &treeNode1, rightNode: &treeNode3, value: 2}
//	treeNode5 := AVLTree[int]{leftNode: &treeNode4, rightNode: &treeNode6, value: 6}
//	tree := AVLTree[int]{leftNode: &treeNode2, rightNode: &treeNode5, value: 3}
//
//	return []AVLTreeInsertNodeTestCase[int]{
//		AVLTreeInsertNodeTestCase[int]{input: input1, tree: &tree},
//		AVLTreeInsertNodeTestCase[int]{input: input2, tree: &tree},
//	}
//}
//
//func TestInsertNode(t *testing.T) {
//	testCases := getTestCases()
//	for index, testCase := range testCases {
//		tree := NewAVLTree[int]()
//		for _, value := range testCase.input {
//			tree.InsertValue(value)
//		}
//		if tree != testCase.tree {
//			t.Errorf("Case %d: %v", index, testCase.input)
//		}
//	}
//}
