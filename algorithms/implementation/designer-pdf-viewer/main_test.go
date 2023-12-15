package main

import (
	"testing"
)

func TestDesignerPdfViewer(t *testing.T) {

	t.Run("Designer Pdf Viewer Sample", func(t *testing.T) {

		h := []int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5}
		word := "abc"

		actual := designerPdfViewer(h, word)

		expected := int32(9)

		if actual != expected {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}

	})

	t.Run("Designer Pdf Viewer Sample 2", func(t *testing.T) {

		h := []int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 7}
		word := "zaba"

		actual := designerPdfViewer(h, word)

		expected := int32(28)

		if actual != expected {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}

	})
}
