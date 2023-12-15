package main

import (
	"reflect"
	"testing"
)

func TestUtopianTree(t *testing.T) {

	t.Run("Test Utopian Tree Sample", func(t *testing.T) {

		actual := []int32{utopianTree(0), utopianTree(1), utopianTree(4)}
		expected := []int32{int32(1), int32(2), int32(7)}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}

	})

}
