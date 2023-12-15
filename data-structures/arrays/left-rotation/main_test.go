package main

import (
	"reflect"
	"testing"
)

func TestRotateLeft(t *testing.T) {

	t.Run("Rotate Left Sample", func(t *testing.T) {

		initial := []int32{1, 2, 3, 4, 5}

		actual := rotateLeft(2, initial)
		expected := []int32{3, 4, 5, 1, 2}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}

	})

}
