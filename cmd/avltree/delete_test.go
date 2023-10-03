package avltree

import (
	"reflect"
	"testing"
)

func TestGetSmallestSonOnTheRight(t *testing.T) {
	testCases := []struct {
		input  *AvlTree[int]
		output *AvlTree[int]
	}{
		{
			input: &AvlTree[int]{value: 8, height: 1,
				rightNode: &AvlTree[int]{value: 10, height: 0},
			},
			output: &AvlTree[int]{value: 10, height: 0},
		},
		{
			input: &AvlTree[int]{value: 8, height: 2,
				leftNode: &AvlTree[int]{value: 5, height: 1,
					leftNode: &AvlTree[int]{value: 3, height: 0},
				},
				rightNode: &AvlTree[int]{value: 10, height: 0},
			},
			output: &AvlTree[int]{value: 10, height: 0},
		},
		{
			input: &AvlTree[int]{value: 5, height: 2,
				leftNode: &AvlTree[int]{value: 2, height: 1,
					leftNode: &AvlTree[int]{value: 1, height: 0},
				},
				rightNode: &AvlTree[int]{value: 10, height: 2,
					leftNode: &AvlTree[int]{value: 7, height: 1,
						leftNode: &AvlTree[int]{value: 6, height: 0},
					},
					rightNode: &AvlTree[int]{value: 12, height: 1},
				},
			},
			output: &AvlTree[int]{value: 6, height: 0},
		},
	}

	for index, testCase := range testCases {
		outputReturned := getSmallestSonOnTheRight(testCase.input)
		if !reflect.DeepEqual(outputReturned, testCase.output) {
			t.Errorf("Return case %d\ngot: \n%v\n\nwant: \n%v", index, outputReturned, testCase.output)
		}
	}
}

func TestDeleteHelper(t *testing.T) {
	testCases := []struct {
		input  *AvlTree[int]
		output *AvlTree[int]
	}{
		{ // leaf
			input:  &AvlTree[int]{value: 8, height: 0},
			output: nil,
		},
		{
			// one son on the left
			input: &AvlTree[int]{value: 8, height: 1,
				leftNode: &AvlTree[int]{value: 5, height: 0},
			},
			output: &AvlTree[int]{value: 5, height: 0},
		},
		{
			// one son on the right
			input: &AvlTree[int]{value: 8, height: 1,
				rightNode: &AvlTree[int]{value: 10, height: 0},
			},
			output: &AvlTree[int]{value: 10, height: 0},
		},
		{
			// two sons
			input: &AvlTree[int]{value: 8, height: 2,
				leftNode: &AvlTree[int]{value: 5, height: 1,
					leftNode: &AvlTree[int]{value: 3, height: 0},
				},
				rightNode: &AvlTree[int]{value: 10, height: 0},
			},
			output: &AvlTree[int]{value: 5, height: 1,
				leftNode:  &AvlTree[int]{value: 3, height: 0},
				rightNode: &AvlTree[int]{value: 10, height: 0},
			},
		},
		{
			// two sons
			input: &AvlTree[int]{value: 8, height: 2,
				leftNode: &AvlTree[int]{value: 5, height: 1,
					leftNode: &AvlTree[int]{value: 3, height: 0},
				},
				rightNode: &AvlTree[int]{value: 10, height: 1,
					leftNode: &AvlTree[int]{value: 9, height: 0},
				},
			},
			output: &AvlTree[int]{value: 9, height: 2,
				leftNode: &AvlTree[int]{value: 5, height: 1,
					leftNode: &AvlTree[int]{value: 3, height: 0},
				},
				rightNode: &AvlTree[int]{value: 10, height: 0},
			},
		},
	}
	for index, testCase := range testCases {
		outputReturned := deleteHelper(testCase.input)
		if !reflect.DeepEqual(outputReturned, testCase.output) {
			t.Errorf("Return case %d\ngot: \n%v\n\nwant: \n%v", index, outputReturned, testCase.output)
		}
	}
}

func TestGetBalancedTreeAfterDelete(t *testing.T) {
	testCases := []struct {
		input  *AvlTree[int]
		output *AvlTree[int]
	}{
		{ // nil
			input:  nil,
			output: nil,
		},
		{
			// balanced tree
			input: &AvlTree[int]{value: 8, height: 1,
				leftNode:  &AvlTree[int]{value: 5, height: 0},
				rightNode: &AvlTree[int]{value: 10, height: 0},
			},
			output: &AvlTree[int]{value: 8, height: 1,
				leftNode:  &AvlTree[int]{value: 5, height: 0},
				rightNode: &AvlTree[int]{value: 10, height: 0},
			},
		},
		{
			// unbalanced to the left
			input: &AvlTree[int]{value: 8, height: 2,
				leftNode: &AvlTree[int]{value: 5, height: 1,
					leftNode: &AvlTree[int]{value: 3, height: 0},
				},
			},
			output: &AvlTree[int]{value: 5, height: 1,
				leftNode:  &AvlTree[int]{value: 3, height: 0},
				rightNode: &AvlTree[int]{value: 8, height: 0},
			},
		},
		{
			// unbalanced to the left and his son on the left have greater depth on the right
			input: &AvlTree[int]{value: 8, height: 2,
				leftNode: &AvlTree[int]{value: 5, height: 1,
					rightNode: &AvlTree[int]{value: 7, height: 0},
				},
			},
			output: &AvlTree[int]{value: 7, height: 1,
				leftNode:  &AvlTree[int]{value: 5, height: 0},
				rightNode: &AvlTree[int]{value: 8, height: 0},
			},
		},
		{
			// unbalanced to the right
			input: &AvlTree[int]{value: 8, height: 2,
				rightNode: &AvlTree[int]{value: 10, height: 1,
					rightNode: &AvlTree[int]{value: 12, height: 0},
				},
			},
			output: &AvlTree[int]{value: 10, height: 1,
				leftNode:  &AvlTree[int]{value: 8, height: 0},
				rightNode: &AvlTree[int]{value: 12, height: 0},
			},
		},
		{
			// unbalanced to the right and his son on the right have greater depth on the left
			input: &AvlTree[int]{value: 8, height: 2,
				rightNode: &AvlTree[int]{value: 10, height: 1,
					leftNode: &AvlTree[int]{value: 9, height: 0},
				},
			},
			output: &AvlTree[int]{value: 9, height: 1,
				leftNode:  &AvlTree[int]{value: 8, height: 0},
				rightNode: &AvlTree[int]{value: 10, height: 0},
			},
		},
	}
	for index, testCase := range testCases {
		outputReturned := getBalancedTreeAfterDelete(testCase.input)
		if !reflect.DeepEqual(outputReturned, testCase.output) {
			t.Errorf("Return case %d\ngot: \n%v\n\nwant: \n%v", index, outputReturned, testCase.output)
		}
	}
}

func TestDelete(t *testing.T) {
	testCases := []struct {
		input  *AvlTree[int]
		value  int
		output *AvlTree[int]
	}{
		{ // nil
			input:  nil,
			value:  8,
			output: nil,
		},
		{ // leaf
			input:  &AvlTree[int]{value: 8, height: 0},
			value:  8,
			output: nil,
		},
		{ // one son on the left
			input: &AvlTree[int]{value: 8, height: 1,
				leftNode: &AvlTree[int]{value: 5, height: 0},
			},
			value:  8,
			output: &AvlTree[int]{value: 5, height: 0},
		},
		{ // one son on the right
			input: &AvlTree[int]{value: 8, height: 1,
				rightNode: &AvlTree[int]{value: 10, height: 0},
			},
			value:  8,
			output: &AvlTree[int]{value: 10, height: 0},
		},
		{ // two sons
			input: &AvlTree[int]{value: 8, height: 2,
				leftNode: &AvlTree[int]{value: 5, height: 1,
					leftNode: &AvlTree[int]{value: 3, height: 0},
				},
				rightNode: &AvlTree[int]{value: 10, height: 0},
			},
			value: 8,
			output: &AvlTree[int]{value: 5, height: 1,
				leftNode:  &AvlTree[int]{value: 3, height: 0},
				rightNode: &AvlTree[int]{value: 10, height: 0},
			},
		},
		{ // two sons, value on the middle
			input: &AvlTree[int]{value: 8, height: 2,
				leftNode: &AvlTree[int]{value: 5, height: 1,
					leftNode: &AvlTree[int]{value: 3, height: 0},
				},
				rightNode: &AvlTree[int]{value: 10, height: 1,
					leftNode: &AvlTree[int]{value: 9, height: 0},
				},
			},
			value: 5,
			output: &AvlTree[int]{value: 8, height: 2,
				leftNode: &AvlTree[int]{value: 3, height: 0},
				rightNode: &AvlTree[int]{value: 10, height: 1,
					leftNode: &AvlTree[int]{value: 9, height: 0},
				},
			},
		},
	}
	for index, testCase := range testCases {
		outputReturned := Delete(testCase.input, testCase.value)
		if !reflect.DeepEqual(outputReturned, testCase.output) {
			t.Errorf("Return case %d\ngot: \n%v\n\nwant: \n%v", index, outputReturned, testCase.output)
		}
	}
}
