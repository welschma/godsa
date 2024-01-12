package ds_test

import (
	"reflect"
	"testing"

	"github.com/welschma/godsa/ds"
)

func compareInt(new int, current int) int {
    return new - current
}

func TestBinarySearchTreeEmpty(t *testing.T) {

	bst := ds.NewBinarySearchTree[int, string](compareInt)

	if bst.Size() != 0 {
		t.Errorf("empty tree size: expected 0, got %d", bst.Size())
	}

	if !bst.IsEmpty() {
		t.Error("new tree is not empty")
	}

	want := []int{}
	got := bst.Keys()

	if !reflect.DeepEqual(want, got) {
		t.Errorf("key slice: want %v, got %v", want, got)
	}

    val, ok := bst.Get(5)
    
    if ok {
        t.Error("tree should return a false flag indicating key is not in tree")
    }

    if val != "" {
        t.Error("tree should return a zero value for the value of a missing key")
    }
}

func TestBinarySearchTreeOneElement(t *testing.T) {

	bst := ds.NewBinarySearchTree[int, string](compareInt)

	bst.Put(5, "5")

	if bst.IsEmpty() {
		t.Error("tree with one element should not be empty")
	}

	if bst.Size() != 1 {
		t.Errorf("tree size with one node is wrong, expected 1, got %d", bst.Size())
	}

	want := []int{5}
	got := bst.Keys()

	if !reflect.DeepEqual(want, got) {
		t.Errorf("key slice: want %v, got %v", want, got)
	}

    valueGot, ok := bst.Get(5)
	valueExp := "5"

	if !ok {
		t.Error("expected key 5 to be present in tree")
	}

	if valueGot != valueExp{
		t.Errorf("getting wrong value from key 5, expected %s, got %s", valueExp, valueGot)
	}
}

func TestBinarySearchTreeUpdateOneElement(t *testing.T) {

	bst := ds.NewBinarySearchTree[int, string](compareInt)

	bst.Put(5, "5")
	bst.Put(5, "five")

	if bst.Size() != 1 {
		t.Errorf("tree size with one node is wrong, expected 1, got %d", bst.Size())
	}

	want := []int{5}
	got := bst.Keys()

	if !reflect.DeepEqual(want, got) {
		t.Errorf("key slice: want %v, got %v", want, got)
	}

    valueGot, ok := bst.Get(5)
	valueExpNew := "five"

	if !ok {
		t.Error("expected key 5 to be present in tree")
	}

	if valueGot != valueExpNew{
		t.Errorf("getting wrong value from key 5, expected %s, got %s", valueExpNew, valueGot)
	}
}


func TestBinarySearchTreeSeveralElements(t *testing.T) {

	bst := ds.NewBinarySearchTree[int, string](compareInt)

    testCases := map[int]string{3: "3", 4: "4", 5: "5", 7: "7", 10: "10"}

    for k, v := range testCases {
		bst.Put(k, v)
	}

	if bst.Size() != len(testCases) {
		t.Errorf("tree size with one node is wrong, expected 3, got %d", bst.Size())
	}

	keysExp := []int{3, 4, 5, 7, 10}
	keysGot := bst.Keys()

	if !reflect.DeepEqual(keysExp, keysGot) {
		t.Errorf("key slice: want %v, got %v", keysExp, keysGot)
	}


    
    for k, valueExp := range testCases {
        valueGot, ok := bst.Get(k)

        if !ok {
            t.Errorf("expected key %d to be present in tree", k)
        }

        if valueGot != valueExp{
            t.Errorf("getting wrong value from key 5, expected %s, got %s", valueExp, valueGot)
        }

    }
}


func TestBinarySearchTreeUpdateSeveralElements(t *testing.T) {

	bst := ds.NewBinarySearchTree[int, string](compareInt)

	bst.Put(5, "5")
	bst.Put(4, "4")
	bst.Put(6, "6")

	bst.Put(5, "five")
	bst.Put(4, "four")
	bst.Put(6, "six")

	if bst.Size() != 3 {
		t.Errorf("tree size with one node is wrong, expected 3, got %d", bst.Size())
	}

	keysExp := []int{4, 5, 6}
	keysGot := bst.Keys()

	if !reflect.DeepEqual(keysExp, keysGot) {
		t.Errorf("key slice: want %v, got %v", keysExp, keysGot)
	}


    testCases := map[int]string{4: "four", 5: "five", 6: "six"}
    
    for k, valueExp := range testCases {
        valueGot, ok := bst.Get(k)

        if !ok {
            t.Errorf("expected key %d to be present in tree", k)
        }

        if valueGot != valueExp{
            t.Errorf("getting wrong value from key 5, expected %s, got %s", valueExp, valueGot)
        }
    }

}

func TestBinaryDeleteMinMax(t *testing.T) {

	bst := ds.NewBinarySearchTree[int, string](compareInt)

	bst.Put(100, "hundred")
	bst.DeleteMax()

	if !bst.IsEmpty() {
		t.Error("tree should be empty after deleting max element of one element tree")
	}

	bst.Put(99, "ninety-nine")
	bst.Put(100, "hundred")
	bst.DeleteMin()

	keysExp := []int{100}
	keysGot := bst.Keys()

	if !reflect.DeepEqual(keysExp, keysGot) {
		t.Errorf("key slice: want %v, got %v", keysExp, keysGot)
	}

	bst.Put(99, "ninety-nine")
	bst.DeleteMax()

	keysExp = []int{99}
	keysGot = bst.Keys()

	if !reflect.DeepEqual(keysExp, keysGot) {
		t.Errorf("key slice: want %v, got %v", keysExp, keysGot)
	}


	bst.Clear()

    testCases := map[int]string{3: "3", 4: "4", 5: "5", 7: "7", 10: "10"}

    for k, v := range testCases {
		bst.Put(k, v)
	}

    bst.DeleteMin()

	keysExp = []int{4, 5, 7, 10}
	keysGot = bst.Keys()

	if !reflect.DeepEqual(keysExp, keysGot) {
		t.Errorf("key slice: want %v, got %v", keysExp, keysGot)
	}

    bst.DeleteMax()

	keysExp = []int{4, 5, 7}
	keysGot = bst.Keys()

	if !reflect.DeepEqual(keysExp, keysGot) {
		t.Errorf("key slice: want %v, got %v", keysExp, keysGot)
	}
}

func TestBinarySearchTreeValues(t *testing.T) {
    bst := ds.NewBinarySearchTree[int, string](compareInt)

    bst.Put(1, "one")
    bst.Put(2, "two")
    bst.Put(3, "three")

    values := bst.Values()

    expectedValues := []string{"one", "two", "three"}

    if !reflect.DeepEqual(values, expectedValues) {
        t.Errorf("Expected %v, but got %v", expectedValues, values)
    }
}

func TestBinaryDelete(t *testing.T) {

	bst := ds.NewBinarySearchTree[int, string](compareInt)

	bst.Put(100, "hundred")
	bst.Delete(100)

	if !bst.IsEmpty() {
		t.Error("tree should be empty after deleting max element of one element tree")
	}

	bst = ds.NewBinarySearchTree[int, string](compareInt)

	bst.Put(100, "hundred")
	bst.Put(110, "hundred")
	bst.Delete(110)

	keysExp := []int{100}
	keysGot := bst.Keys()

	if !reflect.DeepEqual(keysExp, keysGot) {
		t.Errorf("key slice: want %v, got %v", keysExp, keysGot)
	}

	bst = ds.NewBinarySearchTree[int, string](compareInt)

	bst.Put(100, "hundred")
	bst.Put(110, "hundred-ten")
	bst.Put(120, "hundred-twenty")
	bst.Delete(110)

	keysExp = []int{100,120}
	keysGot = bst.Keys()

	if !reflect.DeepEqual(keysExp, keysGot) {
		t.Errorf("key slice: want %v, got %v", keysExp, keysGot)
	}

	bst = ds.NewBinarySearchTree[int, string](compareInt)

	bst.Put(100, "hundred")
	bst.Put(110, "hundred-ten")
	bst.Put(105, "hundred-five")
	bst.Put(120, "hundred-twenty")
	bst.Delete(110)

	keysExp = []int{100,120}
	keysGot = bst.Keys()

	if !reflect.DeepEqual(keysExp, keysGot) {
		t.Errorf("key slice: want %v, got %v", keysExp, keysGot)
	}

	bst = ds.NewBinarySearchTree[int, string](compareInt)

	bst.Put(6, "six")
	bst.Put(10, "ten")
	bst.Put(7, "seven")
	bst.Put(8, "eight")
	bst.Put(2, "two")
	bst.Put(1, "one")
	bst.Put(-1, "minus one")
	bst.Put(4, "four")
	bst.Put(5, "five")
	bst.Put(9, "nine")
	bst.Put(100, "one-hundred")

	keysExp = []int{-1, 1, 2, 4, 5, 6, 7, 8, 9, 10, 100}
	keysGot = bst.Keys()

	if !reflect.DeepEqual(keysExp, keysGot) {
		t.Errorf("key slice: want %v, got %v", keysExp, keysGot)
	}

	bst.Delete(-1)

	keysExp = []int{1, 2, 4, 5, 6, 7, 8, 9, 10, 100}
	keysGot = bst.Keys()

	if !reflect.DeepEqual(keysExp, keysGot) {
		t.Errorf("key slice: want %v, got %v", keysExp, keysGot)
	}
	
	bst.Delete(10)

	keysExp = []int{1, 2, 4, 5, 6, 7, 8, 9, 100}
	keysGot = bst.Keys()

	if !reflect.DeepEqual(keysExp, keysGot) {
		t.Errorf("key slice: want %v, got %v", keysExp, keysGot)
	}

	bst.Delete(2)

	keysExp = []int{1, 4, 5, 6, 7, 8, 9, 100}
	keysGot = bst.Keys()

	if !reflect.DeepEqual(keysExp, keysGot) {
		t.Errorf("key slice: want %v, got %v", keysExp, keysGot)
	}
}
