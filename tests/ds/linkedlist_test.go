package ds_test

import (
	"reflect"
	"testing"

	"github.com/welschma/godsa/ds"
)

func TestLinkedList_EmptyList(t *testing.T) {
	dll := ds.NewLinkedList[int]()

	t.Run("empty list size", func(t *testing.T) {
		if dll.Size() != 0 {
			t.Errorf("Expected size of an empty list to be zero, got %d", dll.Size())
		}
	})

	t.Run("empty list slice", func(t *testing.T) {
		emptySlice := []int{}
		resultSlice := dll.ToSlice()

		if !reflect.DeepEqual(emptySlice, resultSlice) {
			t.Errorf("ToSlice() on an empty list returned %v, expected %v", resultSlice, emptySlice)
		}
	})

	t.Run("empty list contains", func(t *testing.T) {
		// Test Contains on an empty list
		if dll.Contains(1) {
			t.Errorf("Contains(1) on an empty list returned true, expected false")
		}
	})

	t.Run("empty list remove", func(t *testing.T) {
		// Test Remove on an empty list
		if dll.Remove(1) {
			t.Errorf("Remove(1) on an empty list returned true, expected false")
		}

        v, err := dll.RemoveFirst()
        assertError(t, err)
        if v!= 0 {
            t.Errorf("calling RemoveFirst() on empty list, expected 0, got %d", v)
        }

        v, err = dll.RemoveLast()
        assertError(t, err)
        if v!= 0 {
            t.Errorf("calling RemoveLast() on empty list, expected 0, got %d", v)
        }
        
        v, err = dll.RemoveAt(0)
        assertError(t, err)
        if v!= 0 {
            t.Errorf("calling RemoveAt(0) on empty list, expected 0, got %d", v)
        }
	})


	t.Run("empty list get", func(t *testing.T) {
		// Test Remove on an empty list
        v, err := dll.Get(1)
        assertError(t, err)
        if v!= 0 {
			t.Errorf("calling Get(1) on an empty list , expected 0, got %d", v)
		}

        v, err = dll.GetFirst()
        assertError(t, err)
        if v!= 0 {
            t.Errorf("calling RemoveFirst() on empty list, expected 0, got %d", v)
        }

        v, err = dll.GetLast()
        assertError(t, err)
        if v!= 0 {
            t.Errorf("calling RemoveLast() on empty list, expected 0, got %d", v)
        }
	})
}

func TestLinkedList_IsEmpty(t *testing.T) {
	dll := ds.NewLinkedList[int]()

	t.Run("check is empty list is empty", func(t *testing.T) {
		if !dll.IsEmpty() {
			t.Errorf("Newly created list should be empty")
		}
	})

	dll.Add(1)

	t.Run("check is single element list is not empty", func(t *testing.T) {
		if dll.IsEmpty() {
			t.Errorf("After adding an element, list should not be empty")
		}
	})

}

func TestToLinkedList_SingleElement(t *testing.T) {
	dll := ds.NewLinkedList[int]()

	dll.Add(1)

	actualSlice := dll.ToSlice()

	expectedSlice := []int{1}

	if !reflect.DeepEqual(actualSlice, expectedSlice) {
		t.Errorf("ToSlice() on a list with a single element returned %v, expected %v", actualSlice, expectedSlice)
	}

	if dll.Size() != 1 {
		t.Errorf("Expected size of a single element list to be 1, got %d", dll.Size())
	}

	if !dll.Contains(1) {
		t.Errorf("Contains(1) on a list with a single element returned false, expected true")
	}

	if dll.Contains(42) {
		t.Errorf("Contains(42) on a list with a single element returned true, expected false")
	}

	if dll.Remove(42) {
		t.Errorf("Remove(42) on a list with a single element of 1 returned true, expected false")
	}

	if !dll.Remove(1) {
		t.Errorf("Remove(1) on a list with a single element of 1 returned false, expected true")
	}

	actualSlice = dll.ToSlice()
	expectedSlice = []int{}

	if !reflect.DeepEqual(actualSlice, expectedSlice) {
		t.Errorf("ToSlice() on a empty list returned %v, expected %v", actualSlice, expectedSlice)
	}

}

func TestLinkedList_MultipleElements(t *testing.T) {
	dll := ds.NewLinkedList[int]()

	dll.Add(1)
	dll.Add(2)
	dll.Add(3)
	dll.Add(4)

	actualSlice := dll.ToSlice()

	expectedSlice := []int{1, 2, 3, 4}

	if dll.Size() != 4 {
		t.Errorf("Expected size of four elements list to be 4, got %d", dll.Size())
	}

	if !reflect.DeepEqual(actualSlice, expectedSlice) {
		t.Errorf("expected %v, got %v", expectedSlice, actualSlice)
	}

}


func TestLinkedList_Add(t *testing.T) {
	dll := ds.NewLinkedList[int]()

    dll.Add(5)
    dll.AddLast(6)
    dll.AddFirst(1)
   
	actualSlice := dll.ToSlice()
	expectedSlice := []int{1, 5, 6}

	if dll.Size() != 3 {
		t.Errorf("Expected size of four elements list to be 4, got %d", dll.Size())
	}

	if !reflect.DeepEqual(actualSlice, expectedSlice) {
		t.Errorf("expected %v, got %v", expectedSlice, actualSlice)
	}

}

func TestLinkedList_Iterator(t *testing.T) {
	dll := ds.NewLinkedList[int]()

    t.Run("empty iterator", func(t *testing.T) {

        dllIterator := dll.CreateIterator()

        if dllIterator.HasNext() {
            t.Error("calling HasNext() on iterator created from empty list returns true, expect false")
        }

        v := dllIterator.GetNext()
        if v != 0 {
            t.Errorf("calling GetNext() on iterator created from empty list, got %d, expected %v", v, 0)
        }
    })

    t.Run("non-empty iterator", func(t *testing.T) {
        dll.Add(1)
        dll.Add(2)
        dll.Add(3)
        dll.Add(4)

        actualSlice := []int{}
        expectedSlice := []int{1, 2, 3, 4}

        dllIterator := dll.CreateIterator()

        for dllIterator.HasNext() {
            actualSlice = append(actualSlice, dllIterator.GetNext())
        }

        assertSliceEqual(t, actualSlice, expectedSlice)
    })
}

func TestLinkedList_Get(t *testing.T) {
	dll := ds.NewLinkedList[int]()

	v, err := dll.Get(0)
	assertError(t, err)
	if v != 0 {
		t.Errorf("Called Get() with index 0 on an empty list, expected return value 0, got %d", v)
	}

	v, err = dll.Get(-1)
	assertError(t, err)
	if v != 0 {
		t.Errorf("Called Get() with an negative index, expected return value 0, got %d", v)
	}

	dll.Add(1)
	v, err = dll.Get(0)
	assertNoError(t, err)
	if v != 1 {
		t.Errorf("expect value 1, got value %d and expected no error, got %s", v, err)
	}

	dll.Add(2)
	v, err = dll.Get(1)
	assertNoError(t, err)
	if v != 2 {
		t.Errorf("expect value 2, got value %d", v)
	}

	v, err = dll.Get(2)
	assertError(t, err)
	if v != 0 {
		t.Errorf("Called Get() with an out of bound index, expected return value 0, got %d", v)
	}
}


func TestLinkedList_Remove(t *testing.T) {
	dll := ds.NewLinkedList[int]()
    
    t.Run("remove from empty list", func(t *testing.T) {
        if dll.Remove(1) {
            t.Error("Remove(1) called on empty list returns true, expect false")
        }

        _, err := dll.RemoveAt(0)
        assertError(t, err)

        _, err = dll.RemoveFirst()
        assertError(t, err)
        
        _, err = dll.RemoveLast()
        assertError(t, err)
    })


	dll.Add(1)
	dll.Add(2)
	dll.Add(3)
	dll.Add(4)
	dll.Add(3)

    t.Run("remove element by value", func(t *testing.T) {
        if !dll.Remove(3) {
            t.Error("Remove(3) called ist containing 3 returns false, expected true")
        }

        actualSlice := dll.ToSlice()
        expectedSlice := []int{1, 2, 4, 3}
        assertSliceEqual(t, actualSlice, expectedSlice)
    })

    t.Run("remove element by index", func(t *testing.T) {

        actual, err := dll.Get(0)
        assertNoError(t, err)

        got, err := dll.RemoveAt(0)
        assertNoError(t, err)

        if got != actual {
            t.Errorf("Removed element at index 0, got %d, expected %d", got, actual)
        }

        actualSlice := dll.ToSlice()
        expectedSlice := []int{2,4, 3}
        assertSliceEqual(t, actualSlice, expectedSlice)
    })

    t.Run("remove first element", func(t *testing.T) {

        actual, err := dll.GetFirst()
        assertNoError(t, err)

        got, err := dll.RemoveFirst()
        assertNoError(t, err)

        if got != actual {
            t.Errorf("Removed element at index 0, got %d, expected %d", got, actual)
        }

        actualSlice := dll.ToSlice()
        expectedSlice := []int{4, 3}
        assertSliceEqual(t, actualSlice, expectedSlice)
        
    })

    t.Run("remove last element", func(t *testing.T) {

        actual, err := dll.GetLast()
        assertNoError(t, err)

        got, err := dll.RemoveLast()
        assertNoError(t, err)

        if got != actual {
            t.Errorf("Removed element at index 0, got %d, expected %d", got, actual)
        }

        actualSlice := dll.ToSlice()
        expectedSlice := []int{4}
        assertSliceEqual(t, actualSlice, expectedSlice)
    })
}

func assertSliceEqual[T any](t testing.TB, actual []T, expected []T) {
        if !reflect.DeepEqual(actual, expected) {
            t.Errorf("expected %v, got %v", expected, actual)
        }
}

func assertError(t testing.TB, err error) {
	t.Helper()
	if err == nil {
		t.Error("wanted an error but didn't get one")
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("wanted no error but did get %v", err)
	}
}
