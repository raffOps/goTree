package avltree

import (
	"cmp"
	"reflect"
	"testing"
)

func TestNewAvlTree(t *testing.T) {
	got := NewAvlTree[int]()
	if got != nil {
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
	return testCases
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
		input  *avlTree[T]
		output T
	}
	//
	testCases := []TestCase[int]{
		{input: &avlTree[int]{value: 8}, output: 8},
		{input: &avlTree[int]{value: 5}, output: 5},
	}

	for index, testCase := range testCases {
		got := GetRootValue(testCase.input)
		if got != testCase.output {
			t.Errorf("Case %d", index)
		}
	}
}

type insertTestCase[T cmp.Ordered] struct {
	treeInput *avlTree[T]
	value     T
	output    *avlTree[T]
}

func getInsertTestCases() []insertTestCase[int] {
	testCases := []insertTestCase[int]{
		{
			treeInput: nil,
			value:     8,
			output:    &avlTree[int]{value: 8, height: 0},
		},
		{
			treeInput: &avlTree[int]{value: 8, height: 0},
			value:     5,
			output:    &avlTree[int]{value: 8, height: 1, leftNode: &avlTree[int]{value: 5, height: 0}},
		},
		{
			treeInput: &avlTree[int]{value: 8, height: 1, leftNode: &avlTree[int]{value: 5, height: 0}},
			value:     3,
			output:    &avlTree[int]{value: 5, height: 1, leftNode: &avlTree[int]{value: 3, height: 0}, rightNode: &avlTree[int]{value: 8, height: 0}},
		},
		{
			treeInput: &avlTree[int]{value: 5, height: 1, leftNode: &avlTree[int]{value: 3, height: 0}, rightNode: &avlTree[int]{value: 8, height: 0}},
			value:     7,
			output: &avlTree[int]{value: 5, height: 2,
				leftNode: &avlTree[int]{value: 3, height: 0},
				rightNode: &avlTree[int]{value: 8,
					leftNode: &avlTree[int]{value: 7, height: 0},
					height:   1,
				},
			},
		},
		{
			treeInput: &avlTree[int]{value: 5, height: 2,
				leftNode: &avlTree[int]{value: 3, height: 0},
				rightNode: &avlTree[int]{value: 8,
					leftNode: &avlTree[int]{value: 7, height: 0},
					height:   1,
				},
			},
			value: 9,
			output: &avlTree[int]{value: 5, height: 2,
				leftNode: &avlTree[int]{value: 3, height: 0},
				rightNode: &avlTree[int]{value: 8,
					leftNode:  &avlTree[int]{value: 7, height: 0},
					rightNode: &avlTree[int]{value: 9, height: 0},
					height:    1,
				},
			},
		},
		{
			treeInput: &avlTree[int]{value: 5, height: 2,
				leftNode: &avlTree[int]{value: 3, height: 0},
				rightNode: &avlTree[int]{value: 8,
					leftNode:  &avlTree[int]{value: 7, height: 0},
					rightNode: &avlTree[int]{value: 9, height: 0},
					height:    1,
				},
			},
			value: 2,
			output: &avlTree[int]{value: 5, height: 2,
				leftNode: &avlTree[int]{value: 3, height: 1,
					leftNode: &avlTree[int]{value: 2, height: 0},
				},
				rightNode: &avlTree[int]{value: 8,
					leftNode:  &avlTree[int]{value: 7, height: 0},
					rightNode: &avlTree[int]{value: 9, height: 0},
					height:    1,
				},
			},
		},
		{
			treeInput: &avlTree[int]{value: 5, height: 2,
				leftNode: &avlTree[int]{value: 3, height: 1,
					leftNode: &avlTree[int]{value: 2, height: 0},
				},
				rightNode: &avlTree[int]{value: 8,
					leftNode:  &avlTree[int]{value: 7, height: 0},
					rightNode: &avlTree[int]{value: 9, height: 0},
					height:    1,
				},
			},
			value: 1,
			output: &avlTree[int]{value: 5, height: 2,
				leftNode: &avlTree[int]{value: 2, height: 1,
					leftNode:  &avlTree[int]{value: 1, height: 0},
					rightNode: &avlTree[int]{value: 3, height: 0},
				},
				rightNode: &avlTree[int]{value: 8,
					leftNode:  &avlTree[int]{value: 7, height: 0},
					rightNode: &avlTree[int]{value: 9, height: 0},
					height:    1,
				},
			},
		},
	}
	return testCases
}

func TestInsert(t *testing.T) {
	testCases := getInsertTestCases()
	for index, testCase := range testCases {
		got := Insert(testCase.treeInput, testCase.value)
		if !reflect.DeepEqual(got, testCase.output) {
			t.Errorf("Case %d\ngot: \n%v\n\nwant: \n%v", index, got, testCase.output)
		}
	}
}
