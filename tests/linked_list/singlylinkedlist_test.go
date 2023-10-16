package list_tests

import (
	"reflect"
	"testing"

	list "github.com/welschma/godsa/pkg/linkedlist"
)

func TestList_IsEmpty(t *testing.T) {
	l := list.NewSinglyLinkedList[int]()
	if !l.IsEmpty() {
		t.Errorf("Expected an empty list, but IsEmpty() returned false")
	}

	l.InsertAtHead(1)
	if l.IsEmpty() {
		t.Errorf("Expected a non-empty list, but IsEmpty() returned true")
	}
}

func TestList_InsertAtHead(t *testing.T) {
	l := list.NewSinglyLinkedList[int]()
	l.InsertAtHead(1)
	if size := l.Size(); size != 1 {
		t.Errorf("Expected list size to be 1 after InsertAtHead, got %d", size)
	}
}

func TestList_InsertAtTail(t *testing.T) {
	l := list.NewSinglyLinkedList[int]()
	l.InsertAtTail(1)
	if size := l.Size(); size != 1 {
		t.Errorf("Expected list size to be 1 after InsertAtTail, got %d", size)
	}
}

func TestList_DeleteFromHead(t *testing.T) {
	l := list.NewSinglyLinkedList[int]()
	l.InsertAtHead(1)

	item, err := l.DeleteFromHead()
	if err != nil {
		t.Errorf("DeleteFromHead returned an error: %v", err)
	}

	if size := l.Size(); size != 0 {
		t.Errorf("Expected list size to be 0 after DeleteFromHead, got %d", size)
	}

	if item != 1 {
		t.Errorf("Expected deleted item to be 1, got %v", item)
	}
}

func TestList_DeleteFromTail(t *testing.T) {
	l := list.NewSinglyLinkedList[int]()
	l.InsertAtHead(1)

	item, err := l.DeleteFromTail()
	if err != nil {
		t.Errorf("DeleteFromTail returned an error: %v", err)
	}

	if size := l.Size(); size != 0 {
		t.Errorf("Expected list size to be 0 after DeleteFromTail, got %d", size)
	}

	if item != 1 {
		t.Errorf("Expected deleted item to be 1, got %v", item)
	}
}

func TestList_InsertAndDeleteFromHead(t *testing.T) {
	l := list.NewSinglyLinkedList[int]()

	// Insert values at the head
	valuesToAdd := []int{1, 2, 3, 4, 5}
	for _, val := range valuesToAdd {
		l.InsertAtHead(val)
	}

	// Check if the size is correct after insertion
	if size := l.Size(); size != len(valuesToAdd) {
		t.Errorf("Expected list size to be %d after InsertAtHead, got %d", len(valuesToAdd), size)
	}

	// Check if the values are inserted in the correct order
	for i := len(valuesToAdd) - 1; i >= 0; i-- {
		if item, _ := l.DeleteFromHead(); item != valuesToAdd[i] {
			t.Errorf("Expected deleted item to be %d, got %d", valuesToAdd[i], item)
		}
	}

	// Check if the list is empty after deletion
	if size := l.Size(); size != 0 {
		t.Errorf("Expected list size to be 0 after all deletions, got %d", size)
	}
}

func TestList_InsertAndDeleteFromTail(t *testing.T) {
	l := list.NewSinglyLinkedList[int]()

	// Insert values at the tail
	valuesToAdd := []int{1, 2, 3, 4, 5}
	for _, val := range valuesToAdd {
		l.InsertAtTail(val)
	}

	// Check if the size is correct after insertion
	if size := l.Size(); size != len(valuesToAdd) {
		t.Errorf("Expected list size to be %d after InsertAtTail, got %d", len(valuesToAdd), size)
	}

	// Check if the values are inserted in the correct order
	for i := len(valuesToAdd) - 1; i >= 0; i-- {
		if item, _ := l.DeleteFromTail(); item != valuesToAdd[i] {
			t.Errorf("Expected deleted item to be %d, got %d", valuesToAdd[i], item)
		}
	}

	// Check if the list is empty after deletion
	if size := l.Size(); size != 0 {
		t.Errorf("Expected list size to be 0 after all deletions, got %d", size)
	}
}

func TestList_MixedInsertAndDelete(t *testing.T) {
	l := list.NewSinglyLinkedList[int]()

	// Insert values at both head and tail
	valuesToAdd := []int{1, 2, 3, 4, 5}
	for i, val := range valuesToAdd {
		if i%2 == 0 {
			l.InsertAtHead(val)
		} else {
			l.InsertAtTail(val)
		}
	}

	// Check if the size is correct after insertion
	if size := l.Size(); size != len(valuesToAdd) {
		t.Errorf("Expected list size to be %d after MixedInsert, got %d", len(valuesToAdd), size)
	}

	// Check if the values are inserted in the correct order
	expectedOrder := []int{5, 3, 1, 2, 4} // Expected order after mixed insertion
	for _, val := range expectedOrder {
		if item, _ := l.DeleteFromHead(); item != val {
			t.Errorf("Expected deleted item from head to be %d, got %d", val, item)
		}
	}

	// Check if the list is empty after deletion from head
	if size := l.Size(); size != 0 {
		t.Errorf("Expected list size to be 0 after all deletions from head, got %d", size)
	}

	// Insert more values at both head and tail
	for i, val := range valuesToAdd {
		if i%2 == 0 {
			l.InsertAtTail(val)
		} else {
			l.InsertAtHead(val)
		}
	}

	// Check if the size is correct after additional insertion
	if size := l.Size(); size != len(valuesToAdd) {
		t.Errorf("Expected list size to be %d after additional MixedInsert, got %d", len(valuesToAdd), size)
	}

	expectedOrder = []int{4, 2, 1, 3, 5}
	// Check if the values are inserted in the correct order
	for i := len(expectedOrder) - 1; i >= 0; i-- {
		if item, _ := l.DeleteFromTail(); item != expectedOrder[i] {
			t.Errorf("Expected deleted item from tail to be %d, got %d", expectedOrder[i], item)
		}
	}

	// Check if the list is empty after deletion from tail
	if size := l.Size(); size != 0 {
		t.Errorf("Expected list size to be 0 after all deletions from tail, got %d", size)
	}
}

func TestList_Size(t *testing.T) {
	l := list.NewSinglyLinkedList[int]()
	l.InsertAtHead(1)
	l.InsertAtHead(2)
	l.InsertAtHead(3)

	if size := l.Size(); size != 3 {
		t.Errorf("Expected list size to be 3, got %d", size)
	}
}

func TestList_Search(t *testing.T) {
	l := list.NewSinglyLinkedList[int]()

	if l.Search(1) {
		t.Error("Expected value 1 to not be found in the list, but it was found")
	}

	l.InsertAtHead(1)
	if !l.Search(1) {
		t.Errorf("Expected value %d to be found in the list, but it wasn't found", 1)
	}

	// Insert values into the list
	valuesToAdd := []int{2, 3, 4, 5}
	for _, val := range valuesToAdd {
		l.InsertAtTail(val)
	}

	// Test searching for values that are present in the list
	for _, val := range []int{1, 2, 3, 4, 5} {
		if !l.Search(val) {
			t.Errorf("Expected value %d to be found in the list, but it wasn't found", val)
		}
	}

	// Test searching for a value that is not present in the list
	if l.Search(10) {
		t.Error("Expected value 10 to not be found in the list, but it was found")
	}
}

func TestLinkedList_Delete(t *testing.T) {
	// Create a new instance of your linked list implementation
	ll := list.NewSinglyLinkedList[int]()

	// Add some elements to the linked list
	ll.InsertAtTail(1)
	ll.InsertAtTail(2)
	ll.InsertAtTail(3)
	ll.InsertAtTail(4)
	ll.InsertAtTail(5)

	// Test deleting a node from the middle of the list
	ll.Delete(3)
	expectedResult := []int{1, 2, 4, 5}
	checkResult(t, "Delete(3)", &ll, expectedResult)

	// Test deleting the first node
	ll.Delete(1)
	expectedResult = []int{2, 4, 5}
	checkResult(t, "Delete(1)", &ll, expectedResult)

	// Test deleting the last node
	ll.Delete(5)
	expectedResult = []int{2, 4}
	checkResult(t, "Delete(5)", &ll, expectedResult)

	// Test deleting from an empty list
	emptyLL := list.NewSinglyLinkedList[int]()
	emptyLL.Delete(1) // should not panic or error

	// Test deleting from a list with a single node
	singleNodeLL := list.NewSinglyLinkedList[int]()
	singleNodeLL.InsertAtTail(42)
	singleNodeLL.Delete(42)
	checkResult(t, "Delete(42) from a single-node list", &singleNodeLL, []int{})
}

// checkResult is a helper function to check if the linked list has the expected values after a delete operation.
func checkResult[T comparable](t *testing.T, operation string, ll *list.SinglyLinkedList[T], expected []int) {
	result := ll.ToArray()
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("After %s, expected result %v, but got %v", operation, expected, result)
	}
}
