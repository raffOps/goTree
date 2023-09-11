package main

import (
	"fmt"
	"gotree/cmd/tree"
)

type insertNodesTest[T tree.Node] struct {
	input []T
	want  *tree.BinaryTree[T]
}

func main() {
	valueListTest1 := []int{25, 22, 40, 30, 45, 27, 20, 21, 48}

	node21 := tree.NewBinaryTree(nil, nil, 21)
	node20 := tree.NewBinaryTree(nil, node21, 20)
	node22 := tree.NewBinaryTree(node20, nil, 22)
	node27 := tree.NewBinaryTree(nil, nil, 27)
	node30 := tree.NewBinaryTree(node27, nil, 30)
	node48 := tree.NewBinaryTree(nil, nil, 48)
	node45 := tree.NewBinaryTree(nil, node48, 45)
	node40 := tree.NewBinaryTree(node30, node45, 40)

	treeTest1 := tree.NewBinaryTree(node22, node40, 25)

	insertNodeTest1 := insertNodesTest[int]{input: valueListTest1, want: treeTest1}
	println(insertNodeTest1.want.GetValue())

	// string test
	stringTree := tree.NewBinaryTree(nil, nil, "foo")
	stringTree.InsertNewNode("bar")
	fmt.Printf("%v\n", stringTree.GetValue())

}
