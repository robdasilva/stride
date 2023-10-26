package slices_test

import (
	"testing"

	"github.com/robdasilva/stride/slices"
)

func TestFilter(t *testing.T) {
	object := []string{"one", "two", "three"}

	subject := slices.Filter(object, func(element string, index int) bool {
		return element != "two"
	})

	if len(subject) != 2 {
		t.Errorf("expected len 2, got %d", len(subject))
	}

	for _, value := range subject {
		if value == "two" {
			t.Errorf("expected subject to not contain two")
		}
	}
}
