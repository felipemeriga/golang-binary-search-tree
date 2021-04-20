package golang_binary_search_tree

import (
	"github.com/cheekybits/genny/generic"
)

// Item the type of the binary search tree
type Item generic.Type

// Node a single node that composes the tree
type Node struct {
	key   int
	value Item
	left  *Node //left
	right *Node //right
}