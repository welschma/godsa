package ds

import (
	"fmt"
)

const (
	INITIAL_CAPACITY int = 16 
)


// Entry is a key-value pair
type Entry[K string, V any] struct {
	key   string
	value V
}

// HashTable is a data structure that implements a hash table using open addressing and linear probing
type HashTable[K string, V any] struct {
	table    []*Entry[K, V]
	capacity int
	size int
}

// NewHashTable returns a new hash table
func NewHashTable[K string, V any]() *HashTable[K, V] {
	return &HashTable[K, V]{
		table:    make([]*Entry[K, V], INITIAL_CAPACITY),
		capacity: INITIAL_CAPACITY,
	}
}

// Size returns the number of items in the hash table
func (ht *HashTable[K, V]) Size() int {
	return ht.size
}

// Capacity returns the capacity of the hash table
func (ht *HashTable[K, V]) Capacity() int {
	return ht.capacity
}

// loadFactor returns the load factor of the hash table
func (ht *HashTable[K, V]) loadFactor() float64 {
	return float64(ht.size) / float64(ht.capacity)
}

// resizeTable resizes the hash table
func (ht *HashTable[K, V]) resizeTable(newCapacity int) {

	oldTable := ht.table
	oldCapacity := ht.capacity

	ht.table = make([]*Entry[K, V], newCapacity)
	ht.capacity = newCapacity
	ht.size = 0

	for i := 0; i < oldCapacity; i++ {
		if oldTable[i] != nil {
			ht.Put(oldTable[i].key, oldTable[i].value)
		}
	}
}

// hashString hashes a string 
func (ht *HashTable[K, V]) hashString(key string) int {
	hash := 0
	for i := 0; i < len(key); i++ {
		hash = (31*hash + int(key[i])) % ht.capacity
	}
	return hash
}

// hash hashes a key using linear probing
func (ht *HashTable[K, V]) hash(key string, i int) int {
	hash := ht.hashString(key)
	return (hash + i) % ht.capacity
}

// Put puts a key-value pair into the hash table using linear probing
func (ht *HashTable[K, V]) Put(key string, value V) error {
	entry := &Entry[K, V]{key: key, value: value}

	for i := 0; i < ht.capacity; i++ {
		index := ht.hash(key, i)
		if ht.table[index] == nil {
			ht.table[index] = entry
			ht.size++

			if ht.loadFactor() >= 0.75 {
				ht.resizeTable(ht.capacity * 2)
			}
			
			return nil
		}
	}
	return fmt.Errorf("hashtable is full")
}

// Get gets a value from the hash table using linear probing
func (ht *HashTable[K, V]) Get(key string) (V, error) {

	for i := 0; i < ht.capacity; i++ {
		index := ht.hash(key, i)

		if ht.table[index] == nil {
			var v V
			return v, fmt.Errorf("key '%s' not found", key)
		}

		if ht.table[index].key == key {
			return ht.table[index].value, nil
		}
	}

	var v V
	return v, fmt.Errorf("key not found")
}

// Delete deletes a key-value pair from the hash table using linear probing
func (ht *HashTable[K, V]) Delete(key string) bool {
	for i := 0; i < ht.capacity; i++ {
		index := ht.hash(key, i)

		if ht.table[index] == nil {
			return false
		}

		if ht.table[index].key == key {
			 // Set the slot to nil
			 ht.table[index] = nil
			 ht.size--
	 
			 // Rehash all the keys that come after the deleted key in the probe sequence
			 index = (index + 1) % ht.capacity
			 for ht.table[index] != nil {
				 entry := ht.table[index]
				 ht.table[index] = nil
				 ht.size--
				 ht.Put(entry.key, entry.value)
				 index = (index + 1) % ht.capacity
			 }
	 
			 if ht.loadFactor() <= 0.25 {
				 ht.resizeTable(ht.capacity / 2)
			 }
			 return true
		}
	}

	return false
}
