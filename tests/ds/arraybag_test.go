package ds_test

import (
	"testing"

	"github.com/welschma/godsa/ds"
)

func TestArrayBag(t *testing.T) {
	bag := ds.NewArrayBag[int]()

	if !bag.IsEmpty() {
		t.Errorf("New bag should be empty")
	}

	bag.Add(1)
	if bag.Size() != 1 {
		t.Errorf("Bag size should be 1 after adding an item")
	}
	if bag.IsEmpty() {
		t.Errorf("Bag should not be empty after adding an item")
	}

	bag.Add(2)
	if bag.Size() != 2 {
		t.Errorf("Bag size should be 2 after adding a second item")
	}

	// Test iterator
	iterator := bag.CreateIterator()
	count := 0
	for iterator.HasNext() {
		_ = iterator.GetNext()
		count++
	}
	if count != 2 {
		t.Errorf("Iterator should iterate over 2 items")
	}
}
