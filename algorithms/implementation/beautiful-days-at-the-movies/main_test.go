package main

import (
	"reflect"
	"testing"
)

func TestBeautifulDays(t *testing.T) {

	t.Run("Reverse Number", func(t *testing.T) {

		actual := []int32{reverseNumber(12), reverseNumber(20), reverseNumber(120)}

		expected := []int32{21, 2, 21}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}

	})

	t.Run("Beautiful Days Sample", func(t *testing.T) {

		actual := beautifulDays(20, 23, 6)

		expected := int32(2)

		if actual != expected {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}

	})

}
