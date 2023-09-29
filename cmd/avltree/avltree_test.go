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
			output: "value: 3, height: 2, \nleftNode: value: 2, height: 1, \n  leftNode: value: 1, height: 0, \n  rightNode: , \nrightNode: value: 4, height: 0",
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
func TestGetRootValue(t *testing.T) {
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

type getHeightTestCase[T cmp.Ordered] struct {
	input  *avlTree[T]
	output int64
}

func getGetHeightTestCases() []getHeightTestCase[int] {
	testCases := []getHeightTestCase[int]{
		{input: nil, output: -1},
		{input: &avlTree[int]{value: 8, height: 0}, output: 0},
		{input: &avlTree[int]{value: 8, height: 1}, output: 1},
		{input: &avlTree[int]{value: 8, height: 2}, output: 2},
		{input: &avlTree[int]{value: 8, height: 3}, output: 3},
	}
	return testCases
}

func TestGetHeight(t *testing.T) {
	testCases := getGetHeightTestCases()
	for index, testCase := range testCases {
		got := getHeight(testCase.input)
		if got != testCase.output {
			t.Errorf("Case %d\ngot: \n%v\n\nwant: \n%v", index, got, testCase.output)
		}
	}
}

func TestGetRootValueFromEmptyTree(t *testing.T) {
	intTree := NewAvlTree[int]()
	if GetRootValue(intTree) != 0 {
		t.Errorf("Case int tree")
	}

	stringTree := NewAvlTree[string]()
	if GetRootValue(stringTree) != "" {
		t.Errorf("Case string tree")
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

type rebalanceTreeTestCase[T cmp.Ordered] struct {
	treeInput  *avlTree[T]
	valueInput T
	output     *avlTree[T]
}

func getRebalanceTreeTestCases() []rebalanceTreeTestCase[int] {
	testCases := []rebalanceTreeTestCase[int]{
		{
			// If the balance factor is greater than 1 and the value is less than the root left son,
			// it performs a left rotation.
			treeInput: &avlTree[int]{value: 3, height: 2,
				leftNode: &avlTree[int]{value: 2, height: 1,
					leftNode: &avlTree[int]{value: 1, height: 0},
				},
			},
			valueInput: 1,
			output: &avlTree[int]{value: 2, height: 1,
				leftNode:  &avlTree[int]{value: 1, height: 0},
				rightNode: &avlTree[int]{value: 3, height: 0},
			},
		},
		{
			// If the balance factor is less than -1 and the value is greater than the root right son,
			// it performs a right rotation.
			treeInput: &avlTree[int]{value: 2, height: 2,
				rightNode: &avlTree[int]{value: 3, height: 1,
					rightNode: &avlTree[int]{value: 4, height: 0},
				},
			},
			valueInput: 4,
			output: &avlTree[int]{value: 3, height: 1,
				leftNode:  &avlTree[int]{value: 2, height: 0},
				rightNode: &avlTree[int]{value: 4, height: 0},
			},
		},
		{
			// If the balance factor is greater than 1 and the value is greater than the root left son,
			// it performs a right rotation on the left son and then a left rotation on the root.
			treeInput: &avlTree[int]{value: 5, height: 2,
				leftNode: &avlTree[int]{value: 2, height: 1,
					rightNode: &avlTree[int]{value: 3, height: 0},
				},
			},
			valueInput: 3,
			output: &avlTree[int]{value: 3, height: 1,
				leftNode:  &avlTree[int]{value: 2, height: 0},
				rightNode: &avlTree[int]{value: 5, height: 0},
			},
		},
		{
			// If the balance factor is less than -1 and the value is less than the root right son,
			// it performs a left rotation on the right son and then a right rotation on the root.
			treeInput: &avlTree[int]{value: 3, height: 2,
				rightNode: &avlTree[int]{value: 5, height: 1,
					leftNode: &avlTree[int]{value: 4, height: 0},
				},
			},
			valueInput: 4,
			output: &avlTree[int]{value: 4, height: 1,
				leftNode:  &avlTree[int]{value: 3, height: 0},
				rightNode: &avlTree[int]{value: 5, height: 0},
			},
		},
	}
	return testCases
}

func TestRebalanceTree(t *testing.T) {
	testCases := getRebalanceTreeTestCases()
	for index, testCase := range testCases {
		got := rebalanceTree(testCase.treeInput, testCase.valueInput)
		if !reflect.DeepEqual(got, testCase.output) {
			t.Errorf("Case %d\ngot: \n%v\n\nwant: \n%v", index, got, testCase.output)
		}
	}
}
