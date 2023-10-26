package maps_test

import (
	"testing"

	"github.com/robdasilva/stride/maps"
	"github.com/robdasilva/stride/slices"
)

func TestKeys(t *testing.T) {
	object := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	subject := maps.Keys(object)

	if len(subject) != len(object) {
		t.Errorf("Expected %d entries, got %d", len(object), len(subject))
	}

	for key := range object {
		index := slices.Index(subject, func(k string, _ int) bool {
			return k == key
		})

		if index == -1 {
			t.Errorf("Expected keys to contain %v", key)
		}
	}
}
