package main

import (
	"testing"
)

func TestArrayManipulation(t *testing.T) {

	t.Run("Array Manipulation Example", func(t *testing.T) {

		size := int32(10)
		queries := [][]int32{
			{1, 5, 3},
			{4, 8, 7},
			{6, 9, 1},
		}

		actual := arrayManipulation(size, queries)
		expected := int64(10)

		if actual != expected {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}

	})

	t.Run("Array Manipulation Sample", func(t *testing.T) {

		size := int32(5)
		queries := [][]int32{
			{1, 2, 100},
			{2, 5, 100},
			{3, 4, 100},
		}

		actual := arrayManipulation(size, queries)
		expected := int64(200)

		if actual != expected {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}

	})

}
