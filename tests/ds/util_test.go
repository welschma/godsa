package ds_test

import (
	"reflect"
	"testing"

	"github.com/welschma/godsa/ds"
)

func TestResizeSlice(t *testing.T) {
	t.Run("resize to larger capacity", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Fatalf("Expected no panic, got %v", r)
			}
		}()

		slice := []int{1, 2, 3}
		newSlice := ds.ResizeSlice(slice, 5)
		if len(newSlice) != 5 {
			t.Errorf("Expected length 5, got %d", len(newSlice))
		}
		if !reflect.DeepEqual(newSlice[:3], slice) {
			t.Errorf("Expected %v, got %v", slice, newSlice[:3])
		}
	})

	t.Run("resize to negative capacity", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Fatalf("Expected a panic, got nil")
			}
		}()

		slice := []int{1, 2, 3}
		ds.ResizeSlice(slice, -1)
	})
}