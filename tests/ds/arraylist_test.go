package ds_test

import (
	"testing"

	"github.com/welschma/godsa/ds"
)

func TestArrayList(t *testing.T) {
	list := ds.NewArrayList[int](4) // assuming NewArrayList function exists

	if !list.IsEmpty() {
		t.Errorf("New list should be empty")
	}

	if list.Size() != 0 {
		t.Errorf("New list should have size 0")
	}

	// Test adding elements
	for i := 0; i < 10; i++ {
		list.Add(i) // assuming Add function exists
	}

	if list.IsEmpty() {
		t.Errorf("List should not be empty")
	}

	if list.Size() != 10 {
		t.Errorf("List should have size 10")
	}

	// Test removing elements
	for i := 0; i < 5; i++ {
		_, err := list.RemoveAt(0)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	}

	if list.Size() != 5 {
		t.Errorf("List should have size 5")
	}

	// Test out of bounds
	_, err := list.RemoveAt(10)
	if err == nil {
		t.Errorf("Expected error for out of bounds, got nil")
	}
}

func TestArrayListReallocation(t *testing.T) {
	initialCapacity := 10
	list := ds.NewArrayList[int](initialCapacity) // assuming NewArrayList function exists

	if list.Capacity() != initialCapacity {
		t.Errorf("Expected initial capacity %d, got %d", initialCapacity, list.Capacity())
	}

	// Test growing reallocation
	for i := 0; i < 1000; i++ {
		list.Add(i)
		if list.Size() != i+1 {
			t.Errorf("Expected size %d, got %d", i+1, list.Size())
		}
		if list.Capacity() < list.Size() {
			t.Errorf("Expected capacity >= size, got capacity %d, size %d", list.Capacity(), list.Size())
		}
	}

	// Test shrinking reallocation
	for i := 0; i < 750; i++ {
		_, err := list.RemoveAt(0)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if list.Size() != 1000-i-1 {
			t.Errorf("Expected size %d, got %d", 1000-i-1, list.Size())
		}
		if list.Capacity() < list.Size() {
			t.Errorf("Expected capacity >= size, got capacity %d, size %d", list.Capacity(), list.Size())
		}
	}
	list.Capacity()
}