package slices_test

import (
	"testing"

	"github.com/robdasilva/stride/slices"
)

func TestReduce(t *testing.T) {
	object := []string{"one", "two", "three"}

	subject := slices.Reduce(object, map[string]int{}, func(accumulator map[string]int, element string, index int) map[string]int {
		accumulator[element] = index

		return accumulator
	})

	if len(subject) != len(object) {
		t.Errorf("expected %d entries, got %d", len(object), len(subject))
	}

	for index, value := range object {
		if subject[value] != index {
			t.Errorf("expected %d at key %s, got %d", index, value, subject[value])
		}
	}
}
