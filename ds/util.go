package ds

import "fmt"

// ResizeSlice resizes the given slice to the given capacity. If the new capacity is less than the length of the slice,	
// an error is returned.
func ResizeSlice[T any](slice []T, newCapacity int) ([]T, error) {

	if newCapacity < 0 {
		return nil, fmt.Errorf("newCapacity must be non-negative")
	}	

	if newCapacity < len(slice) {
		return nil, fmt.Errorf("newCapacity must be greater than or equal to the length of the slice")
	}

	newSlice := make([]T, newCapacity)
	copy(newSlice, slice)
	return newSlice, nil
}
