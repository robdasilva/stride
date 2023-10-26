package maps_test

import (
	"testing"

	"github.com/robdasilva/stride/maps"
	"github.com/robdasilva/stride/slices"
)

func TestEntries(t *testing.T) {
	object := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	subject := maps.Entries(object)

	if len(subject) != len(object) {
		t.Errorf("Expected %d entries, got %d", len(object), len(subject))
	}

	for key, value := range object {
		index := slices.Index(subject, func(entry maps.Entry[string, int], _ int) bool {
			return entry.Key == key && entry.Value == value
		})

		if index == -1 {
			t.Errorf("Expected entries to contain [%v, %v]", key, value)
		}
	}
}

func TestFromEntries(t *testing.T) {
	object := []maps.Entry[string, int]{{Key: "one", Value: 1}, {Key: "two", Value: 2}, {Key: "three", Value: 3}}

	subject := maps.FromEntries(object)

	if len(subject) != len(object) {
		t.Errorf("Expected %d entries, got %d", len(object), len(subject))
	}

	for _, entry := range object {
		value, ok := subject[entry.Key]

		if !ok {
			t.Errorf("Expected map to contain key %v", entry.Key)
		}

		if value != entry.Value {
			t.Errorf("Expected map to contain value %v for key %v", entry.Value, entry.Key)
		}
	}
}
