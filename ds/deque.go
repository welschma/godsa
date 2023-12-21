package ds

type Deque[T comparable] interface {
    Collection[T]
    Queue[T]
    OfferFirst(t T) bool
    OfferLast(t T) bool
    PeekFirst() (T, bool)
    PeekLast() (T, bool)
    PollFirst() (T, bool)
    PollLast() (T, bool)
    Push(t T)
    Pop() (T, error)
}

