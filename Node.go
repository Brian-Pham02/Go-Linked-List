package golinkedlist

/*
Defines the node of the LinkedList structure
*/
type Node[T comparable] struct {
	data T
	prev *Node[T]
	next *Node[T]
}

/*
	Instantiates a Node object with a value and pointers
	to the previous Node object and the next Node object
*/
func NewNode[T comparable](prev *Node[T], elem T, next *Node[T]) *Node[T] {
	return &Node[T]{
		data: elem,
		prev: prev,
		next: next,
	}
}

// Get the node data
func (n *Node[T]) Data() T {
	return n.data
}

// Get the prev node
func (n *Node[T]) Prev() *Node[T] {
	return n.prev
}

// Get the next node
func (n *Node[T]) Next() *Node[T] {
	return n.next
}
