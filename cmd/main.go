package main

import (
	"fmt"
	"github.com/rjribeiro/gotree/cmd/redblacktree"
	"math/rand"
)

func main() {
	//tree := avltree.NewAvlTree[int]()
	//r := rand.New(rand.NewSource(0))
	//arr := r.Perm(10)
	//for _, value := range arr {
	//	fmt.Printf("%d  ", value)
	//	tree = avltree.Insert(tree, value)
	//}
	//fmt.Printf("\n%v", tree)
	//fmt.Printf("\n%v", avltree.ToArray(tree))
	//fmt.Printf("\n%v", avltree.Search(tree, 5))
	//fmt.Printf("\n%v", avltree.Search(tree, 100))
	//
	//tree2 := redblacktree.NewRedBlackTree[int]()
	//fmt.Printf("%v", tree2)

	tree := redblacktree.NewRedBlackTree[int]()
	r := rand.New(rand.NewSource(1))
	arr := r.Perm(100)
	for _, value := range arr {
		fmt.Printf("%d  ", value)
		tree = redblacktree.Insert(tree, value)
	}
	fmt.Printf("\n%v", tree)

}
