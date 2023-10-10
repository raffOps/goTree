package redblacktree

import (
	"reflect"
	"testing"
)

func TestNewRedBlackTree(t *testing.T) {
	tree := NewRedBlackTree[int]()
	if tree.value != 0 ||
		tree.leftNode != nil ||
		tree.rightNode != nil ||
		tree.isRed != false {
		t.Errorf("Returned %v", tree)
	}
}

func TestRedBlackTree_GetValue(t *testing.T) {
	testCases := []struct {
		tree *RedBlackTree[int]
		want int
	}{
		{
			tree: &RedBlackTree[int]{},
			want: 0,
		},
		{
			tree: &RedBlackTree[int]{value: 3},
			want: 3,
		},
		{
			tree: &RedBlackTree[int]{value: 4,
				leftNode: &RedBlackTree[int]{value: 1},
			},
			want: 4,
		},
	}
	for index, testCase := range testCases {
		got := testCase.tree.GetValue()
		if got != testCase.want {
			t.Errorf("Test %d: \nWanted %v\nGot %v",
				index,
				testCase.want,
				got)
		}
	}
}

func TestRedBlackTree_GetChild(t *testing.T) {
	testCases := []struct {
		tree      *RedBlackTree[int]
		direction string
		want      *RedBlackTree[int]
	}{
		{
			direction: "left",
			tree:      &RedBlackTree[int]{},
			want:      nil,
		},

		{
			direction: "right",
			tree:      &RedBlackTree[int]{value: 3},
			want:      nil,
		},

		{
			direction: "left",
			tree: &RedBlackTree[int]{value: 3,
				leftNode: &RedBlackTree[int]{value: 1},
			},
			want: &RedBlackTree[int]{value: 1},
		},
		{
			direction: "right",
			tree: &RedBlackTree[int]{value: 3,
				rightNode: &RedBlackTree[int]{value: 5},
			},
			want: &RedBlackTree[int]{value: 5},
		},
	}
	for index, testCase := range testCases {
		got := testCase.tree.GetChild(testCase.direction)
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("Test %d: \nWanted %v\nGot %v",
				index,
				testCase.want,
				got,
			)
		}
	}
}

func TestRedBlackTree_SetChild(t *testing.T) {
	testCases := []struct {
		tree      *RedBlackTree[int]
		direction string
		child     *RedBlackTree[int]
		want      *RedBlackTree[int]
	}{
		{
			direction: "left",
			tree:      &RedBlackTree[int]{},
			child:     &RedBlackTree[int]{value: -1},
			want: &RedBlackTree[int]{value: 0,
				leftNode: &RedBlackTree[int]{value: -1},
			},
		},

		{
			direction: "right",
			tree:      &RedBlackTree[int]{value: 3},
			child:     &RedBlackTree[int]{value: 5},
			want: &RedBlackTree[int]{value: 3,
				rightNode: &RedBlackTree[int]{value: 5},
			},
		},
	}
	for index, testCase := range testCases {
		testCase.tree.SetChild(testCase.direction, testCase.child)
		got := testCase.tree
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("Test %d: \nWanted %v\nGot %v",
				index,
				testCase.want,
				got,
			)
		}
	}
}
