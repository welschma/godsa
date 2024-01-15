package ds 

type Iterator[T any] interface {
    GetNext() (T, error)
    HasNext() bool
}

type Iterable[T any] interface {
    CreateIterator() Iterator[T]
}
