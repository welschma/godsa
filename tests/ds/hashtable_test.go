package ds_test

import (
	"fmt"
	"testing"

	"github.com/welschma/godsa/ds"
)

func TestHashTable(t *testing.T) {
	ht := ds.NewHashTable[string, int]()

	// Test Put and Get on an empty hashtable
	err := ht.Put("testKey", 123)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	value, err := ht.Get("testKey")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if value != 123 {
		t.Errorf("Expected value to be 123, but got %d", value)
	}

	// Test Get on a non-existing key
	_, err = ht.Get("nonExistingKey")
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
}

func TestHashTableIncrease(t *testing.T) {
	ht := ds.NewHashTable[string, int]()

	// Insert more items than the initial capacity
	for i := 0; i < ds.INITIAL_CAPACITY+1; i++ {
		err := ht.Put(fmt.Sprintf("key%d", i), i)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	}

	// Check that all items are still in the table after the resize
	for i := 0; i < ds.INITIAL_CAPACITY+1; i++ {
		value, err := ht.Get(fmt.Sprintf("key%d", i))
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if value != i {
			t.Errorf("Expected value to be %d, but got %d", i, value)
		}
	}

	// Check that the capacity of the table has increased
	if ht.Capacity() <= ds.INITIAL_CAPACITY {
		t.Errorf("Expected capacity to be greater than %d, but got %d", ds.INITIAL_CAPACITY, ht.Capacity())
	}
}

func TestHashTableDecrease(t *testing.T) {
    ht := ds.NewHashTable[string, int]()

    // Insert items into the hash table
    for i := 0; i < ds.INITIAL_CAPACITY*2; i++ {
        err := ht.Put(fmt.Sprintf("key%d", i), i)
        if err != nil {
            t.Errorf("Unexpected error: %v", err)
        }
    }

    // Delete some items to trigger a decrease in table size
    for i := 0; i < ds.INITIAL_CAPACITY*3/2; i++ {
        deleted := ht.Delete(fmt.Sprintf("key%d", i))
        if !deleted {
            t.Errorf("Expected key%d to be deleted, but it was not", i)
        }
    }

    // Check that the remaining items are still in the table
    for i := ds.INITIAL_CAPACITY*3/2; i < ds.INITIAL_CAPACITY*2; i++ {
        value, err := ht.Get(fmt.Sprintf("key%d", i))
        if err != nil {
            t.Errorf("Unexpected error: %v", err)
        }

        if value != i {
            t.Errorf("Expected value to be %d, but got %d", i, value)
        }
    }

    // Check that the capacity of the table has decreased
    if ht.Capacity() >= ds.INITIAL_CAPACITY*2 {
        t.Errorf("Expected capacity to be less than %d, but got %d", ds.INITIAL_CAPACITY*2, ht.Capacity())
    }
}


