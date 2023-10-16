package stack

import "errors"

var ErrEmptyStack = errors.New("stack is empty")

type Stack[T any] interface{
	Push(newValue T)
	Pop() (T, error)
	Peek() (T, error)
	IsEmpty() bool
	Size() uint
	Clear()
}

