package ds

// CompareFunc is a function that compares two values of the same type and returns
// an integer representing the comparison. The return value should be:
//   -1 if a < b
type CompareFunc[T any] func(a, b T) int 


// Heap implements a heap data structure.
type Heap[T any] struct {
	data []T
	compareFunc CompareFunc[T]
	n int
}

// NewHeap creates a new heap with the given capacity and comparison function.
func NewHeap[T any](capacity int, compareFunc CompareFunc[T]) *Heap[T] {
	return &Heap[T]{data: make([]T, capacity), compareFunc: compareFunc, n: 0}
}

// Insert adds the given element to the heap.
func (h *Heap[T]) Insert(t T) {
}

// Max returns the maximum element in the heap.
func (h *Heap[T]) Max() T {
	var t T
	return t
}

// ExtractMax removes and returns the maximum element in the heap.
func (h *Heap[T]) ExtractMax() T {
	var t T
	return t
}

// Size returns the number of elements in the heap.
func (h *Heap[T]) Size() int {
	return h.n
}

// IsEmpty returns true if the heap is empty, false otherwise.
func (h *Heap[T]) IsEmpty() bool {
	return h.n == 0
}