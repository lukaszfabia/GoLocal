package recommendation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCosineSimilarity(t *testing.T) {
	s := &recommendationServiceImpl{}

	tests := []struct {
		vec1     []int
		vec2     []int
		expected float64
	}{
		{[]int{1, 0, -1}, []int{-1, 0, 1}, -1},
		{[]int{1, 1, 1}, []int{1, 1, 1}, 1},
		{[]int{1, 2, 3}, []int{4, 5, 6}, 0.9746318461970762},
		{[]int{0, 0, 0}, []int{0, 0, 0}, 0},
	}

	for _, test := range tests {
		result := s.cosineSimilarity(test.vec1, test.vec2)
		assert.InDelta(t, test.expected, result, 1e-9, "cosineSimilarity(%v, %v) = %v; want %v", test.vec1, test.vec2, result, test.expected)
	}
}

func TestSortIndicesByCosineSim(t *testing.T) {
	tests := []struct {
		cosineSim []float64
		expected  []int
	}{
		{[]float64{0.1, 0.3, 0.2}, []int{1, 2, 0}},
		{[]float64{0.5, 0.5, 0.5}, []int{0, 1, 2}},
		{[]float64{0.9, 0.1, 0.5}, []int{0, 2, 1}},
	}

	for _, test := range tests {
		result := sortIndicesByCosineSim(test.cosineSim)
		assert.Equal(t, test.expected, result, "sortIndicesByCosineSim(%v) = %v; want %v", test.cosineSim, result, test.expected)
	}
}
