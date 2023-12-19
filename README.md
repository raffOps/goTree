[![Go Coverage](https://github.com/raffops/goTree/wiki/coverage.svg)](.coverage.out)
[![Go Report Card](https://goreportcard.com/badge/github.com/raffops/gotree)](https://goreportcard.com/report/github.com/raffops/gotree)
# goTree

On this project I implement a tree data structure in Go for 2 types of self-balancing trees:

- AVL Tree: https://en.wikipedia.org/wiki/AVL_tree
- Red-Black Tree: https://en.wikipedia.org/wiki/Red%E2%80%93black_tree

For the AVL Tree I implemented the following methods:

- Insert
- Delete
- Search
- Preorder
- Height
- String

For the Red-Black Tree I implemented the following methods:

- Insert
- Search
- Preorder
- Height
- String_

Both trees are generic, so they can be used with any type of data, as long as the type implements the interface
Comparable.

# Setup

Just install Go 1.21 and clone the repository. Golang automatically downloads the dependencies.

# Run main with some examples

```bash
go run cmd/main.go
```

# Run tests

```bash
go test ./...
```

# Future work

- Implement the other methods for the AVL Tree and Red-Black Tree
- Create a interface shared both by the AVL Tree and Red-Black Tree
- Implement a dictionary of suffixes using the implemented trees, comparing the performance between them for different
  operations.