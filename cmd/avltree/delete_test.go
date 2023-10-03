package avltree

import (
	"reflect"
	"testing"
)

func TestGetSmallestSonOnTheRight(t *testing.T) {
	testCases := []struct {
		input  *AvlTree[int]
		output int
	}{
		{
			input: &AvlTree[int]{value: 8, height: 1,
				rightNode: &AvlTree[int]{value: 10, height: 0},
			},
			output: 10,
		},
		{
			input: &AvlTree[int]{value: 8, height: 2,
				leftNode: &AvlTree[int]{value: 5, height: 0},
				rightNode: &AvlTree[int]{value: 10, height: 1,
					leftNode: &AvlTree[int]{value: 9, height: 0},
				},
			},
			output: 9,
		},
		{
			input: &AvlTree[int]{value: 8, height: 2,
				leftNode: &AvlTree[int]{value: 5, height: 1,
					leftNode: &AvlTree[int]{value: 3, height: 0},
				},
				rightNode: &AvlTree[int]{value: 10, height: 0},
			},
			output: 10,
		},
		{
			input: &AvlTree[int]{value: 8, height: 2,
				leftNode: &AvlTree[int]{value: 5, height: 1},
				rightNode: &AvlTree[int]{value: 10, height: 1,
					rightNode: &AvlTree[int]{value: 12, height: 0},
				},
			},
			output: 10,
		},
	}

	for index, testCase := range testCases {
		outputReturned := getSmallestSonOnTheRight(testCase.input)
		if !reflect.DeepEqual(outputReturned, testCase.output) {
			t.Errorf("Return case %d\ngot: \n%v\n\nwant: \n%v", index, outputReturned, testCase.output)
		}
	}
}
