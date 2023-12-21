package ds

type List[T comparable] interface {
    Add(t T)
    Contains(t T) bool
    Remove(t T) bool
}

