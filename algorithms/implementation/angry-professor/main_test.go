package main

import (
	"reflect"
	"testing"
)

func TestAngryProfessor(t *testing.T) {

	t.Run("Test Angry Professor Sample", func(t *testing.T) {

		actual := []string{angryProfessor(3, []int32{-1, -3, 4, 2}), angryProfessor(2, []int32{0, -1, 1, 2})}
		expected := []string{"YES", "NO"}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}

	})

}
