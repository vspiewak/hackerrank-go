package main

import (
	"reflect"
	"testing"
)

func TestViralAdvertising(t *testing.T) {

	t.Run("Test Viral Advertising Sample", func(t *testing.T) {

		actual := []int32{viralAdvertising(1), viralAdvertising(2), viralAdvertising(3), viralAdvertising(4), viralAdvertising(5)}
		expected := []int32{int32(2), int32(5), int32(9), int32(15), int32(24)}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}

	})

}
