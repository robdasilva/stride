package slices_test

import (
	"testing"

	"github.com/robdasilva/stride/slices"
)

func TestIndex(t *testing.T) {
	object := []string{"one", "two", "three"}

	m := map[string]int{
		"not found": -1,
		"one":       0,
		"two":       1,
		"three":     2,
	}

	for key, value := range m {
		subject := slices.Index(object, func(element string, index int) bool {
			return element == key
		})

		if subject != value {
			t.Errorf("expected \"%s\" to have index 0, got %d", key, subject)
		}
	}
}
