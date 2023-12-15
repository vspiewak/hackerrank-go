package main

import (
	"testing"
)

func TestHurdleRace(t *testing.T) {

	t.Run("Test Hurdle Race Sample", func(t *testing.T) {

		k := int32(4)
		height := []int32{1, 6, 3, 5, 2}

		actual := hurdleRace(k, height)
		expected := int32(2)

		if actual != expected {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}

	})

}
