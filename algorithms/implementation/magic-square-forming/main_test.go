package main

import (
	"testing"
)

func TestFormingMagicSquare(t *testing.T) {

	t.Run("Forming Magic Square Sample", func(t *testing.T) {

		initial := [][]int32{
			{5, 3, 4},
			{1, 5, 8},
			{6, 4, 2},
		}

		actual := formingMagicSquare(initial)

		expected := int32(7)

		if actual != expected {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}

	})

}
