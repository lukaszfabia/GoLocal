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
