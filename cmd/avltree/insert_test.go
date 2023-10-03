package avltree

import (
	"cmp"
	"reflect"
	"testing"
)

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

type getBalancedTreeTestCase[T cmp.Ordered] struct {
	treeInput  *AvlTree[T]
	valueInput T
	output     *AvlTree[T]
}

func getBalancedTreeTestCases() []getBalancedTreeTestCase[int] {
	testCases := []getBalancedTreeTestCase[int]{
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

func TestGetBalancedTree(t *testing.T) {
	testCases := getBalancedTreeTestCases()
	for index, testCase := range testCases {
		got := getBalancedTree(testCase.treeInput, testCase.valueInput)
		if !reflect.DeepEqual(got, testCase.output) {
			t.Errorf("Case %d\ngot: \n%v\n\nwant: \n%v", index, got, testCase.output)
		}
	}
}
