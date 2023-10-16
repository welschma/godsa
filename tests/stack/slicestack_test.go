package stack_tests

import (
	"testing"

	"github.com/welschma/godsa/pkg/stack"
)

func TestEmptyStack(t *testing.T) {
	s := stack.NewSliceStack[int]()

	if !s.IsEmpty() {
		t.Error("Newly created stack is not empty")
	}

	s.Push(1)
	if s.IsEmpty() {
		t.Error("Stack with stored items is empty")
	}
}


func TestClear(t *testing.T) {
	s := stack.NewSliceStack[int]()
	s.Push(1)
	s.Push(3)
	s.Push(42)
	
	if s.IsEmpty() {
		t.Errorf("Stack should not be empty")
	}

	s.Clear()
	if !s.IsEmpty() {
		t.Errorf("Stack should be empty")
	}
}

func TestPush(t *testing.T) {
	s := stack.NewSliceStack[int]()
	s.Push(1)

	want := uint(1)
	got := s.Size()

	if got != want {
		t.Errorf("Testing size: got %d, want %d", got, want)
	}

	s.Push(3)
	s.Push(42)
	want = uint(3)
	got = s.Size()
	if got != want {
		t.Errorf("Testing size: got %d, want %d", got, want)
	}
}

func TestPeek(t *testing.T) {
	s := stack.NewSliceStack[int]()

	test_val := 1
	s.Push(test_val)

	want := test_val
	got, _ := s.Peek()
	if got != want {
		t.Errorf("Test pop with single element: got %d, want %d", got, want)
	}
	got, _ = s.Peek()
	if got != want {
		t.Errorf("Test pop with single element: got %d, want %d", got, want)
	}

	test_val = 42
	s.Push(test_val)
	want = test_val
	got, _ = s.Peek()
	if got != want {
		t.Errorf("Test pop with single element: got %d, want %d", got, want)
	}

}

func TestPop(t *testing.T) {
	s := stack.NewSliceStack[int]()

	_, err := s.Pop()

	if err != stack.ErrEmptyStack {
		t.Error("Popping from empty stack should give EmptyStackError")
	}

	test_val := 1
	s.Push(test_val)

	want := test_val
	got, err := s.Pop()

	if got != want {
		t.Errorf("Test pop with single element: got %d, want %d", got, want)
	}
	if err != nil {
		t.Error("Pop from a non-empty stack should yield no error")
	}
	if !s.IsEmpty() {
		t.Error("Stack is not empty after popping its only element")
	}

	test_val2 := 3
	test_val3 := 42
	s.Push(test_val2)
	s.Push(test_val3)

	got = int(s.Size())
	want = 2
	if got != want {
		t.Errorf("Stack should be of size %d, got %d", want, got)
	}

	want = test_val3
	got, _ = s.Pop()
	if got != want {
		t.Errorf("Popping from stack: got %d, want %d", got, want)
	}

	got = int(s.Size())
	want = 1
	if got != want {
		t.Errorf("Stack should be of size %d, got %d", want, got)
	}

	want = test_val2
	got, _ = s.Pop()
	if got != want {
		t.Errorf("Popping from stack: got %d, want %d", got, want)
	}

	_, err = s.Pop()
	if err != stack.ErrEmptyStack {
		t.Error("Popping from empty stack should give EmptyStackError")
	}

}
