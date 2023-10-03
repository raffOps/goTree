package avltree

import (
	"cmp"
	"reflect"
	"testing"
)

type rotateTestCase[T cmp.Ordered] struct {
	input  *AvlTree[T]
	output *AvlTree[T]
}

func getRotateLeftTestCases() []rotateTestCase[int] {
	testCases := []rotateTestCase[int]{
		{
			input: &AvlTree[int]{value: 3, height: 2,
				leftNode: &AvlTree[int]{value: 2, height: 1,
					leftNode: &AvlTree[int]{value: 1, height: 0},
				},
			},
			output: &AvlTree[int]{value: 2, height: 1,
				leftNode:  &AvlTree[int]{value: 1, height: 0},
				rightNode: &AvlTree[int]{value: 3, height: 0},
			},
		},
		{
			input: &AvlTree[int]{value: 3, height: 2,
				leftNode: &AvlTree[int]{value: 2, height: 1,
					leftNode: &AvlTree[int]{value: 1, height: 0},
				},
				rightNode: &AvlTree[int]{value: 4, height: 0},
			},
			output: &AvlTree[int]{value: 2, height: 2,
				leftNode: &AvlTree[int]{value: 1, height: 0},
				rightNode: &AvlTree[int]{value: 3, height: 1,
					rightNode: &AvlTree[int]{value: 4, height: 0},
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
	testCases := []rotateTestCase[int]{
		{
			input: &AvlTree[int]{value: 2, height: 2,
				rightNode: &AvlTree[int]{value: 3, height: 1,
					rightNode: &AvlTree[int]{value: 4, height: 0},
				},
			},
			output: &AvlTree[int]{value: 3, height: 1,
				leftNode:  &AvlTree[int]{value: 2, height: 0},
				rightNode: &AvlTree[int]{value: 4, height: 0},
			},
		},
		{
			input: &AvlTree[int]{value: 1, height: 2,
				rightNode: &AvlTree[int]{value: 3, height: 1,
					leftNode:  &AvlTree[int]{value: 2, height: 0},
					rightNode: &AvlTree[int]{value: 4, height: 0},
				},
			},
			output: &AvlTree[int]{value: 3, height: 2,
				leftNode: &AvlTree[int]{value: 1, height: 1,
					rightNode: &AvlTree[int]{value: 2, height: 0},
				},
				rightNode: &AvlTree[int]{value: 4, height: 0},
			},
		},
	}
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
