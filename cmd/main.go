package main

import (
	"fmt"
	"github.com/rjribeiro/gotree/cmd/avltree"
	"math/rand"
)

func main() {
	tree := avltree.NewAvlTree[int]()
	r := rand.New(rand.NewSource(99))
	arr := r.Perm(100)
	for _, value := range arr {
		fmt.Printf("%d  ", value)
		tree = avltree.Insert(tree, value)
	}
	fmt.Printf("\n%v", tree)
}
