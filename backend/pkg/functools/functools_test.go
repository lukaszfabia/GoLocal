package functools_test

import (
	"backend/pkg/functools"
	"testing"
)

func TestIn(t *testing.T) {
	var lst = []int{1, 2, 3, 4, 11}

	var toFind = []int{3, 2}

	for _, e := range toFind {
		if !functools.In(e, lst) {
			t.Errorf("Failed")
		}
	}

	var notInList = []int{55, 222}

	for _, e := range notInList {
		if functools.In(e, lst) {
			t.Errorf("Failed")
		}
	}

}

func TestFilter(t *testing.T) {
	t.Run("Odd Numbers", func(t *testing.T) {
		var lst = []int{1, 2, 3, 4, 11}
		var pred = func(e int) bool { return e%2 == 1 }
		var filtered = functools.Filter(pred, lst)
		expected := []int{1, 3, 11}

		if len(expected) != len(filtered) {
			t.Errorf("Length mismatch: expected %d, got %d", len(expected), len(filtered))
		}

		for i, v := range filtered {
			if v != expected[i] {
				t.Errorf("Mismatch at index %d: expected %d, got %d", i, expected[i], v)
			}
		}
	})

	t.Run("Even Numbers", func(t *testing.T) {
		var lst = []int{1, 2, 3, 4, 11}
		var pred = func(e int) bool { return e%2 == 0 }
		var filtered = functools.Filter(pred, lst)
		expected := []int{2, 4}

		if len(expected) != len(filtered) {
			t.Errorf("Length mismatch: expected %d, got %d", len(expected), len(filtered))
		}

		for i, v := range filtered {
			if v != expected[i] {
				t.Errorf("Mismatch at index %d: expected %d, got %d", i, expected[i], v)
			}
		}
	})

	t.Run("Empty List", func(t *testing.T) {
		var lst = []int{}
		var pred = func(e int) bool { return e%2 == 1 }
		var filtered = functools.Filter(pred, lst)

		if len(filtered) != 0 {
			t.Errorf("Expected empty list, got %v", filtered)
		}
	})

	t.Run("No Matches", func(t *testing.T) {
		var lst = []int{2, 4, 6, 8}
		var pred = func(e int) bool { return e%2 == 1 }
		var filtered = functools.Filter(pred, lst)

		if len(filtered) != 0 {
			t.Errorf("Expected empty list, got %v", filtered)
		}
	})

	t.Run("All Match", func(t *testing.T) {
		var lst = []int{1, 3, 5, 7, 9}
		var pred = func(e int) bool { return e%2 == 1 }
		var filtered = functools.Filter(pred, lst)
		expected := []int{1, 3, 5, 7, 9}

		if len(expected) != len(filtered) {
			t.Errorf("Length mismatch: expected %d, got %d", len(expected), len(filtered))
		}

		for i, v := range filtered {
			if v != expected[i] {
				t.Errorf("Mismatch at index %d: expected %d, got %d", i, expected[i], v)
			}
		}
	})

	t.Run("Negative Numbers", func(t *testing.T) {
		var lst = []int{-3, -2, -1, 0, 1, 2, 3}
		var pred = func(e int) bool { return e < 0 }
		var filtered = functools.Filter(pred, lst)
		expected := []int{-3, -2, -1}

		if len(expected) != len(filtered) {
			t.Errorf("Length mismatch: expected %d, got %d", len(expected), len(filtered))
		}

		for i, v := range filtered {
			if v != expected[i] {
				t.Errorf("Mismatch at index %d: expected %d, got %d", i, expected[i], v)
			}
		}
	})
}
