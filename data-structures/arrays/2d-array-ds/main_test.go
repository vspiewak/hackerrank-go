package main

import (
	"testing"
)

func TestHourglassSumAt(t *testing.T) {

	t.Run("Hourglass Sum At", func(t *testing.T) {

		initial := [][]int32{
			{1, 1, 1, 0, 0, 0},
			{0, 1, 0, 0, 0, 0},
			{1, 1, 1, 0, 0, 0},
			{0, 0, 2, 4, 4, 0},
			{0, 0, 0, 2, 0, 0},
			{0, 0, 1, 2, 4, 0},
		}

		// 0, 0
		actual := hourglassSumAt(0, 0, initial)
		expected := int32(7)

		if actual != expected {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}

		// 3, 3
		actual = hourglassSumAt(3, 3, initial)
		expected = int32(14)

		if actual != expected {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}

	})

}

func TestHourglassSum(t *testing.T) {

	t.Run("Hourglass Sample", func(t *testing.T) {

		initial := [][]int32{
			{1, 1, 1, 0, 0, 0},
			{0, 1, 0, 0, 0, 0},
			{1, 1, 1, 0, 0, 0},
			{0, 0, 2, 4, 4, 0},
			{0, 0, 0, 2, 0, 0},
			{0, 0, 1, 2, 4, 0},
		}

		actual := hourglassSum(initial)

		expected := int32(19)

		if actual != expected {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}

	})

}
