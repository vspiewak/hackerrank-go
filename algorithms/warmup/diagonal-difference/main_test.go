package main

import (
	"testing"
)

func TestDiagonalDifference(t *testing.T) {

	t.Run("Diagonal Difference Sample", func(t *testing.T) {

		initial := [][]int32{
			{11, 2, 4},
			{4, 5, 6},
			{10, 8, -12},
		}

		actual := diagonalDifference(initial)
		expected := int32(15)

		if actual != expected {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}

	})

}
