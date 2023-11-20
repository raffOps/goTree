package redblacktree

import (
	"cmp"
	"reflect"
	"testing"
)

type getRecoloredTestCase[T cmp.Ordered] struct {
	tree *RedBlackTree[T]
	want *RedBlackTree[T]
}

func TestGetRecoloredTree(t *testing.T) {
	testCases := getRecoloredTreeTestCases()

	for index, testCase := range testCases {
		getRecoloredTree(testCase.tree)
		if !reflect.DeepEqual(testCase.tree, testCase.want) {
			t.Errorf("test case %v: got \n%v \nwant \n%v", index, testCase.tree, testCase.want)
		}
	}
}

func getRecoloredTreeTestCases() []getRecoloredTestCase[int] {
	var testCases []getRecoloredTestCase[int]

	// tree
	node1 := &RedBlackTree[int]{
		value: &nodeValue[int]{1},
		isRed: true,
	}

	node3 := &RedBlackTree[int]{
		value: &nodeValue[int]{3},
		isRed: true,
	}

	node2 := &RedBlackTree[int]{
		value:     &nodeValue[int]{2},
		leftNode:  node1,
		rightNode: node3,
		isRed:     false,
	}

	// want

	wantNode1 := &RedBlackTree[int]{
		value: &nodeValue[int]{1},
		isRed: false,
	}

	wantNode3 := &RedBlackTree[int]{
		value: &nodeValue[int]{3},
		isRed: false,
	}

	wantNode2 := &RedBlackTree[int]{
		value:     &nodeValue[int]{2},
		leftNode:  wantNode1,
		rightNode: wantNode3,
		isRed:     true,
	}

	testCases = append(testCases, getRecoloredTestCase[int]{node2, wantNode2})
	return testCases
}

type insertTestCase[T cmp.Ordered] struct {
	tree  *RedBlackTree[T]
	value T
	want  *RedBlackTree[T]
}

func TestInsert(t *testing.T) {
	testCases := getInsertTestCases()

	for index, testCase := range testCases {
		got := Insert(testCase.tree, testCase.value)
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("test case %v: \ngot \n%v \nwant \n%v", index, testCase.tree, testCase.want)
		}
	}
}

func getInsertTestCases() (testCases []insertTestCase[int]) {

	var tree, want *RedBlackTree[int]
	var value int

	// test case 1: empty tree
	tree = &RedBlackTree[int]{}
	value = 3
	want = &RedBlackTree[int]{
		value:     &nodeValue[int]{value: 3},
		leftNode:  NewRedBlackTree[int](),
		rightNode: NewRedBlackTree[int](),
		isRed:     false,
	}
	testCases = append(testCases, insertTestCase[int]{tree, value, want})

	// test case 2: value lesser than root value.
	tree = &RedBlackTree[int]{
		value:     &nodeValue[int]{value: 3},
		leftNode:  NewRedBlackTree[int](),
		rightNode: NewRedBlackTree[int](),
	}
	value = 2
	want = &RedBlackTree[int]{
		value: &nodeValue[int]{value: 3},
		leftNode: &RedBlackTree[int]{
			value:     &nodeValue[int]{value: 2},
			isRed:     true,
			leftNode:  NewRedBlackTree[int](),
			rightNode: NewRedBlackTree[int](),
		},
		rightNode: NewRedBlackTree[int](),
	}

	testCases = append(testCases, insertTestCase[int]{tree, value, want})

	// test case 3: value greater than root value.
	tree = &RedBlackTree[int]{
		value:     &nodeValue[int]{value: 3},
		leftNode:  NewRedBlackTree[int](),
		rightNode: NewRedBlackTree[int](),
	}
	value = 4
	want = &RedBlackTree[int]{
		value: &nodeValue[int]{value: 4},
		leftNode: &RedBlackTree[int]{
			value:     &nodeValue[int]{value: 3},
			isRed:     true,
			leftNode:  NewRedBlackTree[int](),
			rightNode: NewRedBlackTree[int](),
		},
		rightNode: NewRedBlackTree[int](),
	}

	testCases = append(testCases, insertTestCase[int]{tree, value, want})

	// test case 4: value greater than root. Root with left child.

	tree = &RedBlackTree[int]{
		value: &nodeValue[int]{value: 3},
		leftNode: &RedBlackTree[int]{
			value:     &nodeValue[int]{value: 2},
			isRed:     true,
			leftNode:  NewRedBlackTree[int](),
			rightNode: NewRedBlackTree[int](),
		},
		rightNode: NewRedBlackTree[int](),
	}

	value = 4
	want = &RedBlackTree[int]{
		value: &nodeValue[int]{value: 3},
		isRed: false,
		leftNode: &RedBlackTree[int]{
			value:     &nodeValue[int]{value: 2},
			isRed:     false,
			leftNode:  NewRedBlackTree[int](),
			rightNode: NewRedBlackTree[int](),
		},
		rightNode: &RedBlackTree[int]{
			value:     &nodeValue[int]{value: 4},
			isRed:     false,
			leftNode:  NewRedBlackTree[int](),
			rightNode: NewRedBlackTree[int](),
		},
	}

	testCases = append(testCases, insertTestCase[int]{tree, value, want})

	// test case 5: value lesser than left son.

	tree = &RedBlackTree[int]{
		value: &nodeValue[int]{value: 3},
		leftNode: &RedBlackTree[int]{
			value:     &nodeValue[int]{value: 2},
			isRed:     true,
			leftNode:  NewRedBlackTree[int](),
			rightNode: NewRedBlackTree[int](),
		},
		rightNode: NewRedBlackTree[int](),
	}

	value = 1
	want = &RedBlackTree[int]{
		value: &nodeValue[int]{value: 2},
		isRed: false,
		leftNode: &RedBlackTree[int]{
			value:     &nodeValue[int]{value: 1},
			isRed:     false,
			leftNode:  NewRedBlackTree[int](),
			rightNode: NewRedBlackTree[int](),
		},
		rightNode: &RedBlackTree[int]{
			value:     &nodeValue[int]{value: 3},
			isRed:     false,
			leftNode:  NewRedBlackTree[int](),
			rightNode: NewRedBlackTree[int](),
		},
	}

	testCases = append(testCases, insertTestCase[int]{tree, value, want})

	// test case 6: value greater than left son, lesser than root.

	tree = &RedBlackTree[int]{
		value: &nodeValue[int]{value: 3},
		leftNode: &RedBlackTree[int]{
			value:     &nodeValue[int]{value: 1},
			isRed:     true,
			leftNode:  NewRedBlackTree[int](),
			rightNode: NewRedBlackTree[int](),
		},
		rightNode: NewRedBlackTree[int](),
	}
	value = 2
	want = &RedBlackTree[int]{
		value: &nodeValue[int]{value: 2},
		isRed: false,
		leftNode: &RedBlackTree[int]{
			value:     &nodeValue[int]{value: 1},
			isRed:     false,
			leftNode:  NewRedBlackTree[int](),
			rightNode: NewRedBlackTree[int](),
		},
		rightNode: &RedBlackTree[int]{
			value:     &nodeValue[int]{value: 3},
			isRed:     false,
			leftNode:  NewRedBlackTree[int](),
			rightNode: NewRedBlackTree[int](),
		},
	}

	testCases = append(testCases, insertTestCase[int]{tree, value, want})

	// test case 7: value already in tree.
	// tree
	tree = &RedBlackTree[int]{
		value: &nodeValue[int]{value: 3},
		leftNode: &RedBlackTree[int]{
			value:     &nodeValue[int]{value: 1},
			isRed:     true,
			leftNode:  NewRedBlackTree[int](),
			rightNode: NewRedBlackTree[int](),
		},
		rightNode: NewRedBlackTree[int](),
	}
	value = 1
	want = &RedBlackTree[int]{
		value: &nodeValue[int]{value: 3},
		leftNode: &RedBlackTree[int]{
			value:     &nodeValue[int]{value: 1},
			isRed:     true,
			leftNode:  NewRedBlackTree[int](),
			rightNode: NewRedBlackTree[int](),
		},
		rightNode: NewRedBlackTree[int](),
	}

	testCases = append(testCases, insertTestCase[int]{tree, value, want})

	return testCases
}
