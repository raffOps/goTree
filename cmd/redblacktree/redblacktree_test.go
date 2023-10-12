package redblacktree

import (
	"reflect"
	"testing"
)

func TestNewRedBlackTree(t *testing.T) {
	tree := NewRedBlackTree[int]()
	if tree.value != nil ||
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
			tree: &RedBlackTree[int]{value: &nodeValue[int]{3}},
			want: 3,
		},
		{
			tree: &RedBlackTree[int]{value: &nodeValue[int]{4},
				leftNode: &RedBlackTree[int]{value: &nodeValue[int]{1}},
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
			tree:      &RedBlackTree[int]{value: &nodeValue[int]{3}},
			want:      nil,
		},

		{
			direction: "left",
			tree: &RedBlackTree[int]{value: &nodeValue[int]{3},
				leftNode: &RedBlackTree[int]{value: &nodeValue[int]{1}},
			},
			want: &RedBlackTree[int]{value: &nodeValue[int]{1}},
		},
		{
			direction: "right",
			tree: &RedBlackTree[int]{value: &nodeValue[int]{3},
				rightNode: &RedBlackTree[int]{value: &nodeValue[int]{5}},
			},
			want: &RedBlackTree[int]{value: &nodeValue[int]{5}},
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

func TestRedBlackTree_String(t *testing.T) {
	// test  cases definition
	type testCase struct {
		tree *RedBlackTree[int]
		want string
	}
	var testCases []testCase

	// test case 1: root without parent

	node := &RedBlackTree[int]{value: &nodeValue[int]{1}}
	want := "\nvalue: 1, \ncolor: black, \nparentNode: , \nleftNode: , \nrightNode: "
	testCases = append(testCases, testCase{tree: node, want: want})

	// test case 2: root with parent

	node0 := &RedBlackTree[int]{}
	node1 := &RedBlackTree[int]{}
	node2 := &RedBlackTree[int]{}
	node3 := &RedBlackTree[int]{}

	*node1 = RedBlackTree[int]{
		value:      &nodeValue[int]{1},
		parentNode: node2,
	}

	*node2 = RedBlackTree[int]{
		value:      &nodeValue[int]{2},
		leftNode:   node1,
		rightNode:  node3,
		parentNode: node0,
	}

	*node3 = RedBlackTree[int]{
		value:      &nodeValue[int]{3},
		parentNode: node2,
	}

	*node0 = RedBlackTree[int]{
		value:      &nodeValue[int]{0},
		rightNode:  node2,
		parentNode: nil,
	}

	want = "\nvalue: 2, \ncolor: black, \nparentNode: 0, \nleftNode:   \nvalue: 1, \n  color: black, \n  parentNode: 2, \n  leftNode: , \n  rightNode: , \nrightNode:   \nvalue: 3, \n  color: black, \n  parentNode: 2, \n  leftNode: , \n  rightNode: "
	testCases = append(testCases, testCase{tree: node2, want: want})

	for index, testCase := range testCases {
		got := testCase.tree.String()
		if got != testCase.want {
			t.Errorf("Test %d: \nWanted:\n %v\nGot:\n %v",
				index,
				testCase.want,
				got,
			)
		}
	}
}

func TestRedBlackTree_getColor(t *testing.T) {
	testCases := []struct {
		tree *RedBlackTree[int]
		want string
	}{
		{
			tree: &RedBlackTree[int]{},
			want: "black",
		},
		{
			tree: &RedBlackTree[int]{isRed: true},
			want: "red",
		},
	}
	for index, testCase := range testCases {
		got := testCase.tree.getColor()
		if got != testCase.want {
			t.Errorf("Test %d: \nWanted %v\nGot %v",
				index,
				testCase.want,
				got)
		}
	}
}
