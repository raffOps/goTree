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
	input  *AvlTree[T]
	output string
}

func getAvlTreeStringTestCases() []getAvlTreeStringTestCase[int] {
	testCases := []getAvlTreeStringTestCase[int]{
		{
			input: &AvlTree[int]{value: 3, height: 2,
				leftNode: &AvlTree[int]{value: 2, height: 1,
					leftNode: &AvlTree[int]{value: 1, height: 0},
				},
				rightNode: &AvlTree[int]{value: 4, height: 0},
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
		input  *AvlTree[T]
		output T
	}
	//
	testCases := []TestCase[int]{
		{input: &AvlTree[int]{value: 8}, output: 8},
		{input: &AvlTree[int]{value: 5}, output: 5},
	}

	for index, testCase := range testCases {
		got := GetRootValue(testCase.input)
		if got != testCase.output {
			t.Errorf("Case %d", index)
		}
	}
}

type getHeightTestCase[T cmp.Ordered] struct {
	input  *AvlTree[T]
	output int64
}

func getGetHeightTestCases() []getHeightTestCase[int] {
	testCases := []getHeightTestCase[int]{
		{input: nil, output: -1},
		{input: &AvlTree[int]{value: 8, height: 0}, output: 0},
		{input: &AvlTree[int]{value: 8, height: 1}, output: 1},
		{input: &AvlTree[int]{value: 8, height: 2}, output: 2},
		{input: &AvlTree[int]{value: 8, height: 3}, output: 3},
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
	treeInput *AvlTree[T]
	value     T
	output    *AvlTree[T]
}

func getInsertTestCases() []insertTestCase[int] {
	testCases := []insertTestCase[int]{
		{
			treeInput: nil,
			value:     8,
			output:    &AvlTree[int]{value: 8, height: 0},
		},
		{
			treeInput: &AvlTree[int]{value: 8, height: 0},
			value:     8,
			output:    &AvlTree[int]{value: 8, height: 0},
		},
		{
			treeInput: &AvlTree[int]{value: 8, height: 0},
			value:     5,
			output:    &AvlTree[int]{value: 8, height: 1, leftNode: &AvlTree[int]{value: 5, height: 0}},
		},
		{
			treeInput: &AvlTree[int]{value: 8, height: 1, leftNode: &AvlTree[int]{value: 5, height: 0}},
			value:     3,
			output:    &AvlTree[int]{value: 5, height: 1, leftNode: &AvlTree[int]{value: 3, height: 0}, rightNode: &AvlTree[int]{value: 8, height: 0}},
		},
		{
			treeInput: &AvlTree[int]{value: 5, height: 1, leftNode: &AvlTree[int]{value: 3, height: 0}, rightNode: &AvlTree[int]{value: 8, height: 0}},
			value:     7,
			output: &AvlTree[int]{value: 5, height: 2,
				leftNode: &AvlTree[int]{value: 3, height: 0},
				rightNode: &AvlTree[int]{value: 8,
					leftNode: &AvlTree[int]{value: 7, height: 0},
					height:   1,
				},
			},
		},
		{
			treeInput: &AvlTree[int]{value: 5, height: 2,
				leftNode: &AvlTree[int]{value: 3, height: 0},
				rightNode: &AvlTree[int]{value: 8,
					leftNode: &AvlTree[int]{value: 7, height: 0},
					height:   1,
				},
			},
			value: 9,
			output: &AvlTree[int]{value: 5, height: 2,
				leftNode: &AvlTree[int]{value: 3, height: 0},
				rightNode: &AvlTree[int]{value: 8,
					leftNode:  &AvlTree[int]{value: 7, height: 0},
					rightNode: &AvlTree[int]{value: 9, height: 0},
					height:    1,
				},
			},
		},
		{
			treeInput: &AvlTree[int]{value: 5, height: 2,
				leftNode: &AvlTree[int]{value: 3, height: 0},
				rightNode: &AvlTree[int]{value: 8,
					leftNode:  &AvlTree[int]{value: 7, height: 0},
					rightNode: &AvlTree[int]{value: 9, height: 0},
					height:    1,
				},
			},
			value: 2,
			output: &AvlTree[int]{value: 5, height: 2,
				leftNode: &AvlTree[int]{value: 3, height: 1,
					leftNode: &AvlTree[int]{value: 2, height: 0},
				},
				rightNode: &AvlTree[int]{value: 8,
					leftNode:  &AvlTree[int]{value: 7, height: 0},
					rightNode: &AvlTree[int]{value: 9, height: 0},
					height:    1,
				},
			},
		},
		{
			treeInput: &AvlTree[int]{value: 5, height: 2,
				leftNode: &AvlTree[int]{value: 3, height: 1,
					leftNode: &AvlTree[int]{value: 2, height: 0},
				},
				rightNode: &AvlTree[int]{value: 8,
					leftNode:  &AvlTree[int]{value: 7, height: 0},
					rightNode: &AvlTree[int]{value: 9, height: 0},
					height:    1,
				},
			},
			value: 1,
			output: &AvlTree[int]{value: 5, height: 2,
				leftNode: &AvlTree[int]{value: 2, height: 1,
					leftNode:  &AvlTree[int]{value: 1, height: 0},
					rightNode: &AvlTree[int]{value: 3, height: 0},
				},
				rightNode: &AvlTree[int]{value: 8,
					leftNode:  &AvlTree[int]{value: 7, height: 0},
					rightNode: &AvlTree[int]{value: 9, height: 0},
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
	treeInput  *AvlTree[T]
	valueInput T
	output     *AvlTree[T]
}

func getRebalanceTreeTestCases() []rebalanceTreeTestCase[int] {
	testCases := []rebalanceTreeTestCase[int]{
		{
			// If the balance factor is greater than 1 and the value is less than the root left son,
			// it performs a left rotation.
			treeInput: &AvlTree[int]{value: 3, height: 2,
				leftNode: &AvlTree[int]{value: 2, height: 1,
					leftNode: &AvlTree[int]{value: 1, height: 0},
				},
			},
			valueInput: 1,
			output: &AvlTree[int]{value: 2, height: 1,
				leftNode:  &AvlTree[int]{value: 1, height: 0},
				rightNode: &AvlTree[int]{value: 3, height: 0},
			},
		},
		{
			// If the balance factor is less than -1 and the value is greater than the root right son,
			// it performs a right rotation.
			treeInput: &AvlTree[int]{value: 2, height: 2,
				rightNode: &AvlTree[int]{value: 3, height: 1,
					rightNode: &AvlTree[int]{value: 4, height: 0},
				},
			},
			valueInput: 4,
			output: &AvlTree[int]{value: 3, height: 1,
				leftNode:  &AvlTree[int]{value: 2, height: 0},
				rightNode: &AvlTree[int]{value: 4, height: 0},
			},
		},
		{
			// If the balance factor is greater than 1 and the value is greater than the root left son,
			// it performs a right rotation on the left son and then a left rotation on the root.
			treeInput: &AvlTree[int]{value: 5, height: 2,
				leftNode: &AvlTree[int]{value: 2, height: 1,
					rightNode: &AvlTree[int]{value: 3, height: 0},
				},
			},
			valueInput: 3,
			output: &AvlTree[int]{value: 3, height: 1,
				leftNode:  &AvlTree[int]{value: 2, height: 0},
				rightNode: &AvlTree[int]{value: 5, height: 0},
			},
		},
		{
			// If the balance factor is less than -1 and the value is less than the root right son,
			// it performs a left rotation on the right son and then a right rotation on the root.
			treeInput: &AvlTree[int]{value: 3, height: 2,
				rightNode: &AvlTree[int]{value: 5, height: 1,
					leftNode: &AvlTree[int]{value: 4, height: 0},
				},
			},
			valueInput: 4,
			output: &AvlTree[int]{value: 4, height: 1,
				leftNode:  &AvlTree[int]{value: 3, height: 0},
				rightNode: &AvlTree[int]{value: 5, height: 0},
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
