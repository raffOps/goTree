package avltree

import (
	"cmp"
	"testing"
)

func TestNewAVLTree(t *testing.T) {
	got := NewAVLTree(5)
	if got.value != 5 || got.leftNode != nil || got.rightNode != nil {
		t.Errorf("Case int tree")
	}
}

type getAvlTreeStringTestCase[T cmp.Ordered] struct {
	input  *avlTree[T]
	output string
}

func getAvlTreeStringTestCases() []getAvlTreeStringTestCase[int] {
	testCases := []getAvlTreeStringTestCase[int]{
		{
			input: &avlTree[int]{value: 3, height: 2,
				leftNode: &avlTree[int]{value: 2, height: 1,
					leftNode: &avlTree[int]{value: 1, height: 0},
				},
				rightNode: &avlTree[int]{value: 4, height: 0},
			},
			output: "value: 3, height: 2, \nleftNode: value: 2, height: 1, \n\tleftNode: value: 1, height: 0, \n\trightNode: , \nrightNode: value: 4, height: 0",
		},
	}
}

func TestAvlTreeString(t *testing.T) {
	testCases := getAvlTreeStringTestCases()
	for index, testCase := range testCases {
		got := testCase.input.String()
		if got != testCase.output {
			t.Errorf("Case %d\ngot: \n%v\n\nwant: %v", index, got, testCase.output)
		}
	}
}

// Given a avlTree with a value as rootValue, when GetRootValue is called, then the value is returned.
func TestAVLTree_GetRootValue(t *testing.T) {
	type TestCase[T cmp.Ordered] struct {
		input  *AvlTree[T]
		output T
	}
	testCases := []TestCase[int]{
		{input: &AvlTree[int]{value: 8}, output: 8},
		{input: &AvlTree[int]{value: 5}, output: 5},
	}

	for index, testCase := range testCases {
		got := testCase.input.GetRootValue()
		if got != testCase.output {
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
