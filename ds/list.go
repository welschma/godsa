package ds

type List[T comparable] interface {
    Collection[T]
    Queue[T]
    Deque[T]
    AddFirst(t T)
    AddLast(t T)
    GetFirst() (T, error)
    GetLast() (T, error)
    RemoveFirst() (T, error)
    RemoveLast() (T, error)
    IndexOf(t T) (int, bool)
    Get(index int) (T, error) 
    Set(index int) (T, error) 
}

