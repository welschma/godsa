package ds

type Queue[T comparable] interface {
    Collection[T]
    Offer(t T) bool
    Peek() (T, bool)
    Poll() (T, bool)
}
