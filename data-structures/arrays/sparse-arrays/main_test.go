package main

import (
	"reflect"
	"testing"
)

func TestMatchingStrings(t *testing.T) {

	t.Run("Matching Strings Sample 1", func(t *testing.T) {

		stringList := []string{"aba", "baba", "aba", "xzxb"}
		queries := []string{"aba", "xzxb", "ab"}
		actual := matchingStrings(stringList, queries)
		expected := []int32{2, 1, 0}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}

	})

	t.Run("Matching Strings Sample 2", func(t *testing.T) {

		stringList := []string{"def", "de", "fgh"}
		queries := []string{"de", "lmn", "fgh"}
		actual := matchingStrings(stringList, queries)
		expected := []int32{1, 0, 1}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}

	})

	t.Run("Matching Strings Sample 3", func(t *testing.T) {

		stringList := []string{"abcde", "sdaklfj", "asdjf", "na", "basdn", "sdaklfj", "asdjf", "na", "asdjf", "na", "basdn", "sdaklfj", "asdjf"}
		queries := []string{"abcde", "sdaklfj", "asdjf", "na", "basdn"}
		actual := matchingStrings(stringList, queries)
		expected := []int32{1, 3, 4, 3, 2}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}

	})

}
