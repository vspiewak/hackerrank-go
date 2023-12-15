package main

import (
	"reflect"
	"testing"
)

func TestDynamicArray(t *testing.T) {

	t.Run("Dynamic Array Sample", func(t *testing.T) {

		n := int32(2)
		queries := [][]int32{
			{1, 0, 5},
			{1, 1, 7},
			{1, 0, 3},
			{2, 1, 0},
			{2, 1, 1},
		}

		actual := dynamicArray(n, queries)

		expected := []int32{7, 3}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}

	})

}
