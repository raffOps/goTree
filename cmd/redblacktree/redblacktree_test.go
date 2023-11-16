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
			t.Errorf("Test %d: \nGot %v\nWanted %v",
				index,
				got,
				testCase.want,
			)
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
			t.Errorf("Test %d: \nGot %v\nWanted %v",
				index,
				got,
				testCase.want,
			)
		}
	}
}

type stringTestCase struct {
	tree *RedBlackTree[int]
	want string
}

func TestRedBlackTree_String(t *testing.T) {
	testCases := getStringTestCases()
	for index, testCase := range testCases {
		got := testCase.tree.String()
		if got != testCase.want {
			t.Errorf("Test %d: \nGot \n%v\nWanted \n%v",
				index,
				got,
				testCase.want,
			)
		}
	}
}

func getStringTestCases() []stringTestCase {
	var testCases []stringTestCase

	// test case 1: root without parent

	node := &RedBlackTree[int]{value: &nodeValue[int]{1}}
	want := "\nvalue: 1, \ncolor: black, \nleftNode: , \nrightNode: "
	testCases = append(testCases, stringTestCase{tree: node, want: want})

	// test case 2: root with parent

	node0 := &RedBlackTree[int]{
		value: &nodeValue[int]{1},
	}

	node2 := &RedBlackTree[int]{
		value: &nodeValue[int]{3},
	}

	node1 := &RedBlackTree[int]{
		value:     &nodeValue[int]{2},
		leftNode:  node0,
		rightNode: node2,
	}

	node3 := &RedBlackTree[int]{
		value:    &nodeValue[int]{4},
		leftNode: node1,
	}

	want = "\nvalue: 4, \ncolor: black, \nleftNode: \n\tvalue: 2, \n\tcolor: black, \n\tleftNode: \n\t\tvalue: 1, \n\t\tcolor: black, \n\t\tleftNode: , \n\t\trightNode: , \n\trightNode: \n\t\tvalue: 3, \n\t\tcolor: black, \n\t\tleftNode: , \n\t\trightNode: , \nrightNode: "
	testCases = append(testCases, stringTestCase{tree: node3, want: want})
	return testCases
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
			t.Errorf("Test %d: \nGot %v\nWanted %v",
				index,
				got,
				testCase.want,
			)
		}
	}
}
