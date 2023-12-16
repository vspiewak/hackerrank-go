package main

import (
	"reflect"
	"testing"
)

func TestSaveThePrisoner(t *testing.T) {

	t.Run("Test Save The Prisoner Sample", func(t *testing.T) {

		actual := []int32{saveThePrisoner(4, 6, 2), saveThePrisoner(5, 2, 1), saveThePrisoner(5, 2, 2), saveThePrisoner(7, 19, 2), saveThePrisoner(3, 7, 3)}
		expected := []int32{3, 2, 3, 6, 3}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}

	})

}
