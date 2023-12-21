package ds

import (
	"errors"
	"fmt"
)

// listNode represents a generic node for a doubly linked list. It holds the actual value,
// as well as pointers to the next and previous nodes in the list.
type listNode[T comparable] struct {
	val  T
	next *listNode[T]
	prev *listNode[T]
}

// LinkedList implements a doubly linked list structure.
type LinkedList[T comparable] struct {
	size int
	sentinel *listNode[T]
}

// Add adds the given element to the end of the list.
func (ll *LinkedList[T]) Add(t T) {
    ll.AddLast(t)
}

// AddFirst adds the given element to the beginning of the list.
func (ll *LinkedList[T]) AddFirst(t T) {
    ll.addBefore(t, ll.sentinel.next)
}

// AddLast adds the given element to the end of the list.
func (ll *LinkedList[T]) AddLast(t T) {
    ll.addBefore(t, ll.sentinel)
}


func (ll *LinkedList[T]) Clear() {
    ll.init()
    ll.size = 0
}

// Contains checks if the given element is present within the list.
func (ll *LinkedList[T]) Contains(t T) bool {
    _, ok := ll.GetIndexOf(t)
    return ok
}

// Get returns the value of the element stored at the given index. If the
// given index is negative or >= the list size, a zero value and an error
// is returned.
func (ll *LinkedList[T]) Get(index int) (T, error) {
    x, err := ll.getNodeAt(index)

    if err != nil {
        var t T
        return t, err
    }

    return x.val, nil
}

// GetIndexOf returns the index in this list of the first occurrence of the
// specified element. Succes is indicated by the second return value being
// true. If the element is not in the list the index is return as -1 and
// the second return value is false.
func (ll *LinkedList[T]) GetIndexOf(t T) (int, bool) {
    x := ll.sentinel.next
    var index int

    for x != ll.sentinel {
        if x.val == t {
            return index, true
        }

        x = x.next
        index += 1
    }

    return -1, false
}


// GetFirst returns the first element from the linked list. If the list is empty
// a non-nil error is returned
func (ll *LinkedList[T]) GetFirst() (T, error) {

    if ll.IsEmpty() {
        var t T
        return t, errors.New("cannot return element from empty list")
    }

    return ll.sentinel.next.val, nil
}

// GetLast returns the last element from the linked list. If the list is empty
// a non-nil error is returned
func (ll *LinkedList[T]) GetLast() (T, error) {

    if ll.IsEmpty() {
        var t T
        return t, errors.New("cannot return element from empty list")
    }

    return ll.sentinel.prev.val, nil
}

// IsEmpty checks if the list contains at least one element.
func (ll *LinkedList[T]) IsEmpty() bool {
	return ll.sentinel.next == ll.sentinel && ll.sentinel.prev == ll.sentinel
}

// Remove removes the first ListNode storing the given element from the list and returns true.
// If the element is not present in the list, false is returned.
func (ll *LinkedList[T]) Remove(t T) bool {

    node := ll.findNode(t)
    
    if node == nil {
        return false
    }

    t, err := ll.removeNode(node)

    return err == nil
}

// RemoveFirst removes the first node of the list and returns the stored element.
// If the list is empty, an non-nil error is returned.
func (ll *LinkedList[T]) RemoveFirst() (T, error) {
    return ll.removeNode(ll.sentinel.next)
}

// RemoveLast removes the last node of the list and returns the stored element.
// If the list is empty, an non-nil error is returned.
func (ll *LinkedList[T]) RemoveLast() (T, error) {
    return ll.removeNode(ll.sentinel.prev)
}

// RemoveAt removes the ListNode present at the given index and returns the stored element.
// If the index is out of bounds, a non-nil error is returned.
func (ll *LinkedList[T]) RemoveAt(index int) (T, error) {
    
    node, err := ll.getNodeAt(index)

    var t T

    if err != nil {
        return t, err
    }

    return ll.removeNode(node)
}

// Size returns the current number of elements in the list.
func (ll *LinkedList[T]) Size() int {
	return ll.size
}

// ToSlice returns the slice representation of the linked list.
func (ll *LinkedList[T]) ToSlice() []T {
	t := []T{}

	x := ll.sentinel.next

	for x != ll.sentinel {
		t = append(t, x.val)
		x = x.next
	}

	return t
}

// CreateIterator returns a LinkedListIterator implementing the Iterator interface.
func (ll *LinkedList[T]) CreateIterator() LinkedListIterator[T] {
    return LinkedListIterator[T]{ll.sentinel, ll.sentinel.next}
}

// addBefore inserts a new ListNode holding the given value before the specified ListNode.
func (ll *LinkedList[T]) addBefore(t T, node *listNode[T]) {
    newNode := &listNode[T]{val: t, next: node, prev: node.prev}

    newNode.prev.next = newNode
    newNode.next.prev = newNode

    ll.size += 1
}

// init initializes/resets the linked list by setting the sentinel node to an empty listNode
// with the previous and next pointer pointing to itself
func (ll *LinkedList[T]) init() {
    ll.sentinel = &listNode[T]{}
    ll.sentinel.next = ll.sentinel
    ll.sentinel.prev = ll.sentinel
}

// removeNode removes the given node from the list and returns the stored element.
func (ll *LinkedList[T]) removeNode(node *listNode[T]) (T, error) {
    var t T

    if node == ll.sentinel {
        return t, errors.New("cannot remove header node in linked list")
    }

    t = node.val
    node.prev.next = node.next
    node.next.prev = node.prev

    ll.size -= 1

    return t, nil
}

// getNodeAt returns the node present at the given index. In canse the index is out of bounds,
// an error is returned.
func (ll *LinkedList[T]) getNodeAt(index int) (*listNode[T], error) {
    if index < 0 || index >= ll.size {
        return nil, fmt.Errorf("Index %d is out of bounds for a list of size %d", index, ll.size) 
    }

    x := ll.sentinel.next

    for currentIndex := 0; currentIndex != index; currentIndex++ {
        x = x.next
    }

    return x, nil
}

// findNode returns a pointer to the first listNode storing the given element.
// If the element is not contained within the list, a nil pointer is returned.
func (ll *LinkedList[T]) findNode(t T) *listNode[T] {
	x := ll.sentinel.next

	for x != ll.sentinel {

		if x.val == t {
			return x
		}

		x = x.next
	}

	return nil
}

// NewLinkedList creates a new and empty LinkedList.
func NewLinkedList[T comparable]() LinkedList[T] {
	ll := LinkedList[T]{}
    ll.init()
	return ll
}
