package maps_test

import (
	"testing"

	"github.com/robdasilva/stride/maps"
	"github.com/robdasilva/stride/slices"
)

func TestValues(t *testing.T) {
	object := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	subject := maps.Values(object)

	if len(subject) != len(object) {
		t.Errorf("Expected %d entries, got %d", len(object), len(subject))
	}

	for _, value := range object {
		index := slices.Index(subject, func(v int, _ int) bool {
			return v == value
		})

		if index == -1 {
			t.Errorf("Expected values to contain %v", value)
		}
	}
}
