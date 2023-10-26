package slices_test

import (
	"testing"

	"github.com/robdasilva/stride/slices"
)

func TestMap(t *testing.T) {
	object := []string{"one", "two", "three"}

	m := map[string]int{
		"not found": -1,
		"one":       0,
		"two":       1,
		"three":     2,
	}

	subject := slices.Map(object, func(element string, index int) int {
		return m[element]
	})

	if len(subject) != len(object) {
		t.Errorf("expected %d entries, got %d", len(object), len(subject))
	}

	for key, value := range subject {
		if key != value {
			t.Errorf("expected %d on index %d, got %d", key, key, value)
		}
	}
}
