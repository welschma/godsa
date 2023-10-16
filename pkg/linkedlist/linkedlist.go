package list

import "errors"

var ErrEmptyList error = errors.New("list is empty")

// LinkedList is an interface for a basic singly linked list.
type LinkedList[T comparable] interface {
    // InsertAtHead inserts a new node with the given data at the beginning of the list.
    InsertAtHead(data T)

    // InsertAtTail inserts a new node with the given data at the end of the list.
    InsertAtTail(data T)

    // DeleteFromHead deletes the node at the beginning of the list and returns its data.
    // It returns an error if the list is empty.
    DeleteFromHead() (T, error)

    // DeleteFromTail deletes the node at the end of the list and returns its data.
    // It returns an error if the list is empty.
    DeleteFromTail() (T, error)

    // IsEmpty returns true if the list is empty, false otherwise.
    IsEmpty() bool

    // Size returns the number of nodes in the list.
    Size() uint

    // Search returns true if the given data is present in the list, false otherwise.
    Search(data T) bool

    // Delete deletes the first occurrence of the given data from the list.
    // It returns true if the data was found and deleted, false otherwise.
    Delete(data T) bool
}

