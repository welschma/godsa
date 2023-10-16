package binarysearchtree

import (
	"errors"

	"golang.org/x/exp/constraints"
)

var ErrEmptyBinaryTree = errors.New("emtpy binary search tree")

// BST represents a binary search tree interface.
type BST[T constraints.Ordered] interface {
	// Insert adds a value to the binary search tree.
	Insert(value T)

	// Search searches for a value in the binary search tree and returns true if found, false otherwise.
	Search(value T) bool

	// Delete removes a value from the binary search tree.
	Delete(value T)

	// InOrderTraversal performs an in-order traversal of the binary search tree.
	InOrderTraversal() []T

	// Min returns the minimum value in the binary search tree.
	Min() (T, error)

	// Max returns the maximum value in the binary search tree.
	Max() (T, error)

	// IsEmpty returns true if the root is nil.
	IsEmpty() bool
}

// Node represents a node in the binary search tree.
type Node[T constraints.Ordered] struct {
	value T
	left  *Node[T]
	right *Node[T]
	parent *Node[T]
}


// BinarySearchTree implements the BST interface.
type BinarySearchTree[T constraints.Ordered] struct {
	root *Node[T]
}

func NewBinarySearchTree[T constraints.Ordered]() BinarySearchTree[T] {
	return BinarySearchTree[T]{}
}

func (bst *BinarySearchTree[T]) IsEmpty() bool {
	return bst.root == nil
}

func (bst *BinarySearchTree[T]) InOrderTraversal() []T {
	values := []T{}
	values = bst.inOrder(bst.root, values)
	return values
}

func (bst *BinarySearchTree[T]) inOrder(node *Node[T], result []T) []T {

	if node != nil {
		result = bst.inOrder(node.left, result)
		result = append(result, node.value)
		result = bst.inOrder(node.right, result)
	}

	return result
}

func (bst *BinarySearchTree[T]) Insert(value T) {

	newNode := &Node[T]{value: value}

	if bst.IsEmpty() {
		bst.root = newNode
		return
	}

	var previousNode *Node[T] = nil
	currentNode := bst.root

	for currentNode != nil {
		previousNode = currentNode

		if value < currentNode.value {
			currentNode = currentNode.left
		} else {
			currentNode = currentNode.right
		}
	}

	newNode.parent = previousNode

	if value < previousNode.value {
		previousNode.left = newNode
	} else {
		previousNode.right = newNode
	}
}

func (bst *BinarySearchTree[T]) findNode(value T) (*Node[T], error) {

	if bst.IsEmpty() {
		return nil, ErrEmptyBinaryTree
	}

	currentNode := bst.root

	for currentNode != nil {
		if value == currentNode.value {
			return currentNode, nil
		} else if value < currentNode.value {
			currentNode = currentNode.left
		} else {
			currentNode = currentNode.right
		}
	}

	return nil, errors.New("value not found")
}


func (bst *BinarySearchTree[T]) Search(value T) bool {

	_, err := bst.findNode(value)

	return err == nil
}

func (bst *BinarySearchTree[T]) transplant(u *Node[T], v*Node[T]) {
	
	if u.parent == nil {
		bst.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}

	if v != nil {
		v.parent = u. parent
	}

}

func (bst *BinarySearchTree[T]) Delete(value T) {
	
	deleteNode, err := bst.findNode(value)

	if err != nil {
		return
	}

	if deleteNode.left == nil {
		bst.transplant(deleteNode, deleteNode.right)
	} else if deleteNode.right == nil {
		bst.transplant(deleteNode, deleteNode.left)
	} else {
		y := bst.MinNode(deleteNode.right)

		if y.parent != deleteNode {
			bst.transplant(y, y.right)
			y.right = deleteNode.right
			y.right.parent = y
		}

		bst.transplant(deleteNode, y)
		y.left = deleteNode.left
		y.left.parent = y
	} 
}

func (bst *BinarySearchTree[T]) MinNode(n *Node[T]) *Node[T] {
	
	for n.left != nil {
		n = n.left
	}

	return n
}

func (bst *BinarySearchTree[T]) MaxNode(n *Node[T]) *Node[T] {

	for n.right != nil {
		n = n.right
	}

	return n
}

// Min returns the minimum value in the binary search tree.
func (bst *BinarySearchTree[T]) Min() (T, error) {

	if bst.IsEmpty() {
		var emptyValue T
		return emptyValue, ErrEmptyBinaryTree
	}

	minNode := bst.MinNode(bst.root)
	return minNode.value, nil
}

// Max returns the maximum value in the binary search tree.
func (bst *BinarySearchTree[T]) Max() (T, error) {

	if bst.IsEmpty() {
		var emptyValue T
		return emptyValue, ErrEmptyBinaryTree
	}

	maxNode := bst.MaxNode(bst.root)
	return maxNode.value, nil
}


