package ds_test

import (
	"testing"

	"github.com/welschma/godsa/ds"
)

func TestArrayBag(t *testing.T) {
    bag := ds.NewArrayBag[int]() // assuming NewArrayBag takes an initial capacity

    for i := 0; i < 10; i++ {
        bag.Add(i)
    }

    if bag.Size() != 10 {
        t.Errorf("Expected size 10, got %d", bag.Size())
    }

    if bag.IsEmpty() {
        t.Errorf("Expected bag not to be empty")
    }

    if bag.Capacity() < 10 {
        t.Errorf("Expected capacity at least 10, got %d", bag.Capacity())
    }

    // Test iterator
    iterator := bag.CreateIterator() // assuming ArrayBag has an Iterator method
    for i := 0; i < 10; i++ {
        if !iterator.HasNext() {
            t.Errorf("Expected iterator to have next at index %d", i)
        }

        next, err := iterator.GetNext()
        if err != nil {
            t.Fatalf("Expected no error, got %v", err)
        }

        if next != i {
            t.Errorf("Expected %d, got %d", i, next)
        }
    }

    if iterator.HasNext() {
        t.Errorf("Expected iterator to have no next after 10 items")
    }

    _, err := iterator.GetNext()
    if err == nil {
        t.Fatalf("Expected error, got nil")
    }
}