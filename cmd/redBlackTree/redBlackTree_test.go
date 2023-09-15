package redBlackTree

import (
	"cmp"
	"reflect"
	"testing"
)

func TestNewNodeTree(t *testing.T) {
	value := 3
	rootNode := NewNodeTree(nil, nil, value)

	got := rootNode.GetValue()
	if got != value {
		t.Errorf("got %d, wanted %d", got, value)
	}
}

func TestInsertNewNodeToTheLeft(t *testing.T) {
	rootNode := NewNodeTree(nil, nil, 3)

	newValue := 1
	rootNode.InsertNode(newValue)
	got := (*rootNode.leftNode).GetValue()
	if got != newValue {
		t.Errorf("got %d, wanted %d", got, newValue)
	}
}

func TestInsertNewNodeToTheRight(t *testing.T) {
	rootNode := NewNodeTree(nil, nil, 3)

	newValue := 5
	rootNode.InsertNode(newValue)
	got := (*rootNode.rightNode).GetValue()
	if got != newValue {
		t.Errorf("got %d, wanted %d", got, newValue)
	}
}

type BTreeTestCase[T cmp.Ordered] struct {
	input []T
	tree  *RedBlackTree[T]
}

func getIntegerNodesTestCases() ([]BTreeTestCase[int], error) {
	valueListTestCase1 := []int{25, 22, 40, 30, 45, 27, 20, 21, 48}

	case1Node21 := NewNodeTree(nil, nil, 21)
	case1Node20 := NewNodeTree(nil, case1Node21, 20)
	case1Node22 := NewNodeTree(case1Node20, nil, 22)
	case1Node27 := NewNodeTree(nil, nil, 27)
	case1Node30 := NewNodeTree(case1Node27, nil, 30)
	case1Node48 := NewNodeTree(nil, nil, 48)
	case1Node45 := NewNodeTree(nil, case1Node48, 45)
	case1node40 := NewNodeTree(case1Node30, case1Node45, 40)
	treeTestCase1 := NewNodeTree(case1Node22, case1node40, 25)

	testCase1 := BTreeTestCase[int]{input: valueListTestCase1, tree: treeTestCase1}

	// test case 2
	valueListTestCase2 := []int{40, 25, 20, 30, 45, 27, 22, 21, 48}
	case2Node21 := NewNodeTree(nil, nil, 21)
	case2Node22 := NewNodeTree(case2Node21, nil, 22)
	case2Node20 := NewNodeTree(nil, case2Node22, 20)
	case2Node27 := NewNodeTree(nil, nil, 27)
	case2Node30 := NewNodeTree(case2Node27, nil, 30)
	case2Node25 := NewNodeTree(case2Node20, case2Node30, 25)
	case2Node48 := NewNodeTree(nil, nil, 48)
	case2Node45 := NewNodeTree(nil, case2Node48, 45)
	treeTestCase2 := NewNodeTree(case2Node25, case2Node45, 40)

	testCase2 := BTreeTestCase[int]{input: valueListTestCase2, tree: treeTestCase2}

	// final list
	testCases := []BTreeTestCase[int]{
		testCase1,
		testCase2,
	}
	return testCases, nil
}

func getStringNodesTestCases() ([]BTreeTestCase[string], error) {
	valueListTestCase1 := []string{"Noah", "Oliver", "Eli", "James", "Sophie", "Will", "Omar", "Rachel"}
	caseNodeJames := NewNodeTree(nil, nil, "James")
	caseNodeEli := NewNodeTree(nil, caseNodeJames, "Eli")

	caseNodeRachel := NewNodeTree(nil, nil, "Rachel")
	caseNodeOmar := NewNodeTree(nil, caseNodeRachel, "Omar")
	caseNodeWill := NewNodeTree(nil, nil, "Will")
	caseNodeSophie := NewNodeTree(caseNodeOmar, caseNodeWill, "Sophie")
	caseNodeOliver := NewNodeTree(nil, caseNodeSophie, "Oliver")
	treeTestCase1 := NewNodeTree(caseNodeEli, caseNodeOliver, "Noah")

	testCase1 := BTreeTestCase[string]{
		input: valueListTestCase1,
		tree:  treeTestCase1,
	}

	testCase := []BTreeTestCase[string]{
		testCase1,
	}
	return testCase, nil
}

func TestInsertStringNodes(t *testing.T) {
	testCases, err := getStringNodesTestCases()
	if err != nil {
		t.Errorf("%s", err)
	}
	for _, testCase := range testCases {
		tree_ := NewNodeTree(nil, nil, "")
		for _, value := range testCase.input {
			tree_.InsertNode(value)
		}
		if reflect.DeepEqual(testCase.tree, tree_) == false {
			t.Fail()
		}
		tree_ = nil
	}
}

func TestInsertIntegerNodes(t *testing.T) {
	testCases, _ := getIntegerNodesTestCases()
	for _, testCase := range testCases {
		tree_ := NewNodeTree(nil, nil, 0)
		for _, value := range testCase.input {
			tree_.InsertNode(value)
		}
		if reflect.DeepEqual(testCase.tree, tree_) == false {
			t.Fail()
		}
		tree_ = nil
	}
}

func TestSearchStringValue(t *testing.T) {
	testCases, _ := getStringNodesTestCases()
	testCase := testCases[0]
	value := testCase.tree.rightNode.rightNode.rightNode.GetValue()
	got := testCase.tree.SearchValue(value)
	if got == false {
		t.Fail()
	}
}

func TestSearchStringValueNonExistent(t *testing.T) {
	testCases, _ := getStringNodesTestCases()
	testCase := testCases[0]
	value := "foo"
	got := testCase.tree.SearchValue(value)
	if got == true {
		t.Fail()
	}
}

func TestSearchIntegerValueNonExistent(t *testing.T) {
	testCases, _ := getIntegerNodesTestCases()
	testCase := testCases[0]
	value := testCase.tree.leftNode.leftNode.rightNode.GetValue()
	got := testCase.tree.SearchValue(value)
	if got == false {
		t.Fail()
	}
}

func TestSearchIntegerValue(t *testing.T) {
	testCases, _ := getIntegerNodesTestCases()
	testCase := testCases[0]
	value := 9999
	got := testCase.tree.SearchValue(value)
	if got == true {
		t.Fail()
	}
}
