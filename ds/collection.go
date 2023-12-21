package ds

type Collection[T comparable] interface {
    Add(t T)
    Contains(t T) bool
    Remove(t T) bool
    Size() int
    Clear()
    IsEmpty() bool
}
