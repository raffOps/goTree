package main

import (
	"fmt"
	"github.com/rjribeiro/gotree/cmd/avltree"
)

func main() {
	tree := avltree.NewAvlTree[int]()
	tree = avltree.Insert(tree, 1)

	fmt.Printf("%v", tree)
}
