package main

import (
	"fmt"
	"github.com/rjribeiro/gotree/cmd/avltree"
	"math/rand"
)

func main() {
	tree := avltree.NewAvlTree[int]()
	r := rand.New(rand.NewSource(0))
	arr := r.Perm(10)
	for _, value := range arr {
		fmt.Printf("%d  ", value)
		tree = avltree.Insert(tree, value)
	}
	fmt.Printf("\n%v", tree)
	fmt.Printf("\n%v", avltree.ToArray(tree))
}
