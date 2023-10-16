package binarysearchtree

import (
	"reflect"
	"testing"

	"golang.org/x/exp/constraints"

	"github.com/welschma/godsa/pkg/binarysearchtree"
)

func TestBinarySearchTree_Insert(t *testing.T) {
	bst := binarysearchtree.NewBinarySearchTree[int]()

	valuesToAdd := []int{5, 3, 7, 2, 4, 6, 8}

	for _, val := range valuesToAdd {
		bst.Insert(val)
	}

	// Check if the tree is not empty after insertion
	if bst.IsEmpty() {
		t.Error("Expected the tree to be non-empty after insertion, but it is empty")
	}

	// Check if the values are present in the tree after insertion
	for _, val := range valuesToAdd {
		if !bst.Search(val) {
			t.Errorf("Expected value %d to be present in the tree after insertion, but it is not found", val)
		}
	}

	if bst.Search(10) {
		t.Errorf("Expected value %d not to be present in the tree after insertion, but it is found", 10)
	}

}

func TestBinarySearchTree_Search(t *testing.T) {
	bst := binarysearchtree.NewBinarySearchTree[int]()

	// Insert values into the tree
	valuesToAdd := []int{5, 3, 7, 2, 4, 6, 8}
	for _, val := range valuesToAdd {
		bst.Insert(val)
	}

	// Test searching for values that are present in the tree
	for _, val := range valuesToAdd {
		if !bst.Search(val) {
			t.Errorf("Expected value %d to be found in the tree, but it wasn't found", val)
		}
	}

	// Test searching for a value that is not present in the tree
	if bst.Search(10) {
		t.Error("Expected value 10 to not be found in the tree, but it was found")
	}
}

func TestBST_MinMax(t *testing.T) {
	// Create a new instance of your BST implementation
	bst := binarysearchtree.NewBinarySearchTree[int]()

	// Test empty tree
	_, err := bst.Min()
	if err == nil {
		t.Error("Expected error for Min() on an empty tree, but got none.")
	}

	_, err = bst.Max()
	if err == nil {
		t.Error("Expected error for Max() on an empty tree, but got none.")
	}

	// Insert values into the tree
	values := []int{5, 3, 7, 2, 4, 6, 8}
	for _, val := range values {
		bst.Insert(val)
	}

	// Test Min
	minExpected := 2
	minVal, err := bst.Min()
	if err != nil {
		t.Errorf("Unexpected error for Min(): %v", err)
	}
	if minVal != minExpected {
		t.Errorf("Expected Min() to return %d, but got %d", minExpected, minVal)
	}

	// Test Max
	maxExpected := 8
	maxVal, err := bst.Max()
	if err != nil {
		t.Errorf("Unexpected error for Max(): %v", err)
	}
	if maxVal != maxExpected {
		t.Errorf("Expected Max() to return %d, but got %d", maxExpected, maxVal)
	}
}

func TestBST_InOrderTraversal(t *testing.T) {
	// Create a new instance of your BST implementation
	bst := binarysearchtree.NewBinarySearchTree[int]()

	// Insert values into the tree
	values := []int{5, 3, 7, 2, 4, 6, 8}
	for _, val := range values {
		bst.Insert(val)
	}

	// Perform in-order traversal
	traversalResult := bst.InOrderTraversal()

	// Define the expected result
	expectedResult := []int{2, 3, 4, 5, 6, 7, 8}

	// Compare the results
	if !reflect.DeepEqual(traversalResult, expectedResult) {
		t.Errorf("Expected in-order traversal result %v, but got %v", expectedResult, traversalResult)
	}
}

func TestBST_Delete(t *testing.T) {
	// Create a new instance of your BST implementation
	bst := binarysearchtree.NewBinarySearchTree[int]()

	// Insert values into the tree
	values := []int{5, 3, 7, 2, 4, 6, 8}
	for _, val := range values {
		bst.Insert(val)
	}

	// Test deleting a node with no children
	bst.Delete(2)
	expectedResult := []int{3, 4, 5, 6, 7, 8}
	checkResult[int](t, "Delete(2)", bst, expectedResult)

	//Test deleting a node with one child
	bst.Delete(7)
	expectedResult = []int{3, 4, 5, 6, 8}
	checkResult[int](t, "Delete(7)", bst, expectedResult)

	// Test deleting a node with two children
	bst.Delete(5)
	expectedResult = []int{3, 4, 6, 8}
	checkResult[int](t, "Delete(5)", bst, expectedResult)
}

// checkResult is a helper function to check if the BST has the expected values after a delete operation.
func checkResult[T constraints.Ordered](t *testing.T, operation string, bst binarysearchtree.BinarySearchTree[T], expected []int) {
	result := bst.InOrderTraversal()
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("After %s, expected result %v, but got %v", operation, expected, result)
	}
}
