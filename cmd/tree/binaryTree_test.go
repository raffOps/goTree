package tree

import (
	"reflect"
	"testing"
)

func TestNewBinaryTree(t *testing.T) {
	value := 3
	rootNode := NewBinaryTree(nil, nil, value)

	got := rootNode.GetValue()
	if got != value {
		t.Errorf("got %d, wanted %d", got, value)
	}
}

func TestInsertNewNodeToTheLeft(t *testing.T) {
	rootNode := NewBinaryTree(nil, nil, 3)

	newValue := 1
	rootNode.InsertNewNode(newValue)
	got := (*rootNode.leftNode).GetValue()
	if got != newValue {
		t.Errorf("got %d, wanted %d", got, newValue)
	}
}

func TestInsertNewNodeToTheRight(t *testing.T) {
	rootNode := NewBinaryTree(nil, nil, 3)

	newValue := 5
	rootNode.InsertNewNode(newValue)
	got := (*rootNode.rightNode).GetValue()
	if got != newValue {
		t.Errorf("got %d, wanted %d", got, newValue)
	}
}

type insertNodesTestCase[T Node] struct {
	input []T
	want  *BinaryTree[T]
}

func getInsertNodesIntegerTestCases[T Node]() []insertNodesTestCase[int] {
	valueListTestCase1 := []int{25, 22, 40, 30, 45, 27, 20, 21, 48}

	case1Node21 := NewBinaryTree(nil, nil, 21)
	case1Node20 := NewBinaryTree(nil, case1Node21, 20)
	case1Node22 := NewBinaryTree(case1Node20, nil, 22)
	case1Node27 := NewBinaryTree(nil, nil, 27)
	case1Node30 := NewBinaryTree(case1Node27, nil, 30)
	case1Node48 := NewBinaryTree(nil, nil, 48)
	case1Node45 := NewBinaryTree(nil, case1Node48, 45)
	case1node40 := NewBinaryTree(case1Node30, case1Node45, 40)
	treeTestCase1 := NewBinaryTree(case1Node22, case1node40, 25)

	insertNodeTestCase1 := insertNodesTestCase[int]{input: valueListTestCase1, want: treeTestCase1}

	// test case 2
	valueListTestCase2 := []int{40, 25, 20, 30, 45, 27, 22, 21, 48}
	case2Node21 := NewBinaryTree(nil, nil, 21)
	case2Node22 := NewBinaryTree(case2Node21, nil, 22)
	case2Node20 := NewBinaryTree(nil, case2Node22, 20)
	case2Node27 := NewBinaryTree(nil, nil, 27)
	case2Node30 := NewBinaryTree(case2Node27, nil, 30)
	case2Node25 := NewBinaryTree(case2Node20, case2Node30, 25)
	case2Node48 := NewBinaryTree(nil, nil, 48)
	case2Node45 := NewBinaryTree(nil, case2Node48, 45)
	treeTestCase2 := NewBinaryTree(case2Node25, case2Node45, 40)

	insertNodeTestCase2 := insertNodesTestCase[int]{valueListTestCase2, treeTestCase2}

	// final list
	insertNodeTests := []insertNodesTestCase[int]{
		insertNodeTestCase1,
		insertNodeTestCase2,
	}
	return insertNodeTests
}

func TestInsertIntegerNodes(t *testing.T) {
	insertNodeTests := getInsertNodesIntegerTestCases[int]()
	for _, testCase := range insertNodeTests {
		tree_ := NewBinaryTree(nil, nil, 0)
		for _, value := range testCase.input {
			tree_.InsertNewNode(value)
		}
		if reflect.DeepEqual(testCase.want, tree_) == false {
			t.Fail()
		}
		tree_ = nil
	}
}
