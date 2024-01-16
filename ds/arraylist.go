package ds

import (
	"fmt"
)

// ArrayList implements a list using a dynamically resizing array.
type ArrayList[T comparable] struct {
	a   []T
	n   int
}

func NewArrayList[T comparable](capacity int) *ArrayList[T] {
	return &ArrayList[T]{a: make([]T, capacity), n: 0}
}

// Add adds the given element to the end of this list.
func (l *ArrayList[T]) Add(t T) {
	l.AddLast(t)
}

// AddAt adds the given element to the given index in this list.
func (l *ArrayList[T]) AddAt(i int, t T) error {

	if err := l.checkBounds(i); err != nil {
		return err
	}

	if l.n == l.Capacity() {
		l.a = ResizeSlice[T](l.a, l.Capacity()*2)
	}
	copy(l.a[i+1:], l.a[i:])
	l.a[i] = t
	l.n++
	return nil
}

// AddLast adds the given element to the end of the list.
func (l *ArrayList[T]) AddLast(t T) {
	if l.n == l.Capacity() {
		l.a = ResizeSlice[T](l.a, l.Capacity()*2)
	}

	l.a[l.n] = t
	l.n++
}

// Get returns the element at the given index in the list.
func (l *ArrayList[T]) Get(i int) (T, error) {
	if err := l.checkBounds(i); err != nil {
		var t T
		return t, err
	}
	return l.a[i], nil
}

// Set sets the element at the given index in the list to the given value.
func (l *ArrayList[T]) Set(i int, t T) (T, error) {
	var x T

	if err := l.checkBounds(i); err != nil {
		return x, err
	}

	x = l.a[i]
	l.a[i] = t

	return x, nil
}



// RemoveAt removes the element at the given index in the list.
func (l *ArrayList[T]) RemoveAt(i int) (T, error) {
	if err := l.checkBounds(i); err != nil {
		var t T
		return t, err
	}

	x := l.a[i]
	copy(l.a[i:], l.a[i+1:])

	l.n--

	if l.n <= l.Capacity()/4 {
		l.a = ResizeSlice[T](l.a, l.Capacity()/2)
	}

	return x, nil
}

// RemoveLast removes and returns the last element of this list.
func (l *ArrayList[T]) RemoveLast() (T, error) {
	if l.IsEmpty() {
		var t T
		return t, fmt.Errorf("list is empty")
	}
	return l.RemoveAt(l.n - 1)
}


// Push adds the given element to the end of this list.
func (l *ArrayList[T]) Push(t T) {
	l.Add(t)
}

// Pop removes and returns the last element of this list.
func (l *ArrayList[T]) Pop() (T, error) {
	return l.RemoveLast()
}


// IndexOf returns the index of the first occurrence of the given element in the list.
// If the element is not in the list, -1 is returned.
func (l *ArrayList[T]) IndexOf(t T) int {
	for i := 0; i < l.n; i++ {
		if l.a[i] == t {
			return i
		}
	}
	return -1
}

// IsEmpty returns true if the list is empty.
func (l *ArrayList[T]) IsEmpty() bool {
	return l.n == 0
}

// Size returns the number of elements in the list.
func (l *ArrayList[T]) Size() int {
	return l.n
}

// Capacity returns the current capacity of the list.
func (l *ArrayList[T]) Capacity() int {
	return cap(l.a)
}

// checkBounds checks if the given index is within the bounds of the list.
func (l *ArrayList[T]) checkBounds(i int) error {
	if i < 0 || i >= l.n {
		return fmt.Errorf("index out of bounds")
	}
	return nil
}
