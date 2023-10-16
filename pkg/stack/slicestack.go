package stack

type SliceStack[T any] struct {
	values []T
}

func NewSliceStack[T any]() SliceStack[T] {
	return SliceStack[T]{values: make([]T, 0, 128)}
}

func (s *SliceStack[T]) Values() []T {
	return s.values
}

func (s *SliceStack[T]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *SliceStack[T]) Size() uint {
	return uint(len(s.values))
}

func (s *SliceStack[T]) Clear() {
	s.values = []T{}
}

func (s *SliceStack[T]) Push(newValue T) {
	s.values = append(s.values, newValue)
}

func (s *SliceStack[T]) Peek() (T, error) {

	if s.IsEmpty() {
		var zeroVal T
		return zeroVal, ErrEmptyStack
	}

	return s.values[s.Size()-1], nil
}

func (s *SliceStack[T]) Pop() (T, error) {

	popValue, err := s.Peek()
	if err != nil {
		return popValue, err
	}

	s.values = s.values[:s.Size()-1]

	return popValue, nil
}
