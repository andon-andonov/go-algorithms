package sorting_test

import (
	"reflect"
	"testing"

	"github.com/andon-andonov/go-algorithms/sorting"
)

func TestInsertionSort(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2, 3, 4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{-1, -2, -3, -4, -5, -6, 0, 1, 2, 3, 4, 5, 6}, []int{-6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6}},
		{[]int{5, 2, 4, 6, 1, 3}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{1, 1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1, 1}},
	}
	for _, test := range tests {
		input := make([]int, len(test.input))
		copy(input, test.input)
		sorting.InsertionSort(test.input)
		if !reflect.DeepEqual(test.input, test.want) {
			t.Errorf("InsertionSort(%v) got %v, want %v", input, test.input, test.want)
		}
	}
}
