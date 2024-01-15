package ds_test

import (
	"reflect"
	"testing"

	"github.com/welschma/godsa/ds"
)

func TestResizeSlice(t *testing.T) {
	t.Run("resize to larger capacity", func(t *testing.T) {
		slice := []int{1, 2, 3}
		newSlice, err := ds.ResizeSlice(slice, 5)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if len(newSlice) != 5 {
			t.Errorf("Expected length 5, got %d", len(newSlice))
		}
		if !reflect.DeepEqual(newSlice[:3], slice) {
			t.Errorf("Expected %v, got %v", slice, newSlice[:3])
		}
	})

	t.Run("resize to smaller capacity", func(t *testing.T) {
		slice := []int{1, 2, 3}
		_, err := ds.ResizeSlice(slice, 2)
		if err == nil {
			t.Fatalf("Expected an error, got nil")
		}
	})

	t.Run("resize to negative capacity", func(t *testing.T) {
		slice := []int{1, 2, 3}
		_, err := ds.ResizeSlice(slice, -1)
		if err == nil {
			t.Fatalf("Expected an error, got nil")
		}
	})
}
