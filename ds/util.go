package ds

import "fmt"

// ResizeSlice resizes the given slice to the given capacity. If the new capacity is less than the 
// length of the slice or less than 0, the function panics.
func ResizeSlice[T any](slice []T, newCapacity int) []T {

	if newCapacity < 0 {
		panic("newCapacity must be greater than or equal to 0")
	}	

	if newCapacity < len(slice) {
		panic(fmt.Sprintf("newCapacity must be greater than or equal to the length of the slice (%d)", len(slice)))
	}

	newSlice := make([]T, newCapacity)
	copy(newSlice, slice)
	return newSlice
}
