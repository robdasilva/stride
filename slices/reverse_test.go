package slices_test

import (
	"testing"

	"github.com/robdasilva/stride/slices"
)

func TestReverse(t *testing.T) {
	object := []string{"one", "two", "three"}

	subject := slices.Reverse(object)

	if len(subject) != len(object) {
		t.Errorf("expected %d entries, got %d", len(object), len(subject))
	}

	for index, value := range subject {
		original := object[len(object)-index-1]
		if original != value {
			t.Errorf("expected %s at index %d, got %s", original, index, value)
		}
	}
}
