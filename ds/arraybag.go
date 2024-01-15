package ds

const (
	INITIAL_BAG_CAPACITY int = 32
)

// Bag represents a generic bag. It provides clients with the ability to collect items
// and iterate over them.
type ArrayBag[T any] struct {
	items []T
	n     int
}

// Add adds the given item to the bag.
func (b *ArrayBag[T]) Add(item T) {

	if b.n == cap(b.items) {
		ResizeSlice[T](b.items, 2*cap(b.items))
	}

	b.items[b.n] = item
	b.n += 1

}

// Size returns the number of items in the bag.
func (b *ArrayBag[T]) Size() int {
	return b.n
}

// IsEmpty returns true if the bag is empty, false otherwise.
func (b *ArrayBag[T]) IsEmpty() bool {
	return b.n == 0
}

// Capacity returns the capacity of the bag.
func (b *ArrayBag[T]) Capacity() int {
	return cap(b.items)
}

// CreateIterator returns an iterator for the bag.
func (b *ArrayBag[T]) CreateIterator() Iterator[T] {
	return &ArrayBagIterator[T]{
		items:        b.items,
		currentIndex: 0,
		maxIndex: b.n,
	}
}

// NewArrayBag returns a new bag.
func NewArrayBag[T any]() ArrayBag[T] {
	return ArrayBag[T]{
		items: make([]T, INITIAL_BAG_CAPACITY),
	}
}

// BagIterator represents an iterator for a bag.
type ArrayBagIterator[T any] struct {
	items        []T
	currentIndex int
	maxIndex	 int
}

// HasNext returns true if there are more items to iterate over, false otherwise.
func (it *ArrayBagIterator[T]) HasNext() bool {
	return it.currentIndex < it.maxIndex
}

func (it *ArrayBagIterator[T]) GetNext() T {
	var t T

	if it.HasNext() {
		t = it.items[it.currentIndex]
		it.currentIndex += 1
		return t
	}

	return t
}
