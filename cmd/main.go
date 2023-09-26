package main

import (
	"fmt"
	"github.com/rjribeiro/gotree/cmd/avltree"
)

func main() {
	tree := avltree.NewAVLTree(1)
	fmt.Printf("%v", tree)
}
