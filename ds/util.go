package ds


// ResizeSlice resizes the given slice to the given capacity. If the new capacity is less than the 
// length of the slice or less than 0, the function panics.
func ResizeSlice[T any](slice []T, newCapacity int) []T {

	if newCapacity < 0 {
		panic("newCapacity must be greater than or equal to 0")
	}	

	newSlice := make([]T, newCapacity)
	copy(newSlice, slice)
	return newSlice
}
