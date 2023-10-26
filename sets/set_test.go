package sets_test

import (
	"testing"

	"github.com/robdasilva/stride/sets"
	"github.com/robdasilva/stride/slices"
)

func TestSet(t *testing.T) {
	t.Run("Add", func(t *testing.T) {
		var subject sets.Set[string]

		subject.Add("one")
		subject.Add("two")
		subject.Add("three")

		if subject.Len() != 3 {
			t.Errorf("expected len 3, got %d", subject.Len())
		}

		if !subject.Contains("one") {
			t.Errorf("expected set to contain one")
		}

		if !subject.Contains("two") {
			t.Errorf("expected set to contain two")
		}

		if !subject.Contains("three") {
			t.Errorf("expected set to contain three")
		}
	})

	t.Run("Contains", func(t *testing.T) {
		var subject sets.Set[string]

		subject.Add("one")
		subject.Add("two")

		if !subject.Contains("one") {
			t.Errorf("expected set to contain one")
		}

		if !subject.Contains("two") {
			t.Errorf("expected set to contain two")
		}

		if subject.Contains("three") {
			t.Errorf("expected set to not contain three")
		}
	})

	t.Run("Len", func(t *testing.T) {
		var subject sets.Set[string]

		subject.Add("one")
		subject.Add("two")

		if subject.Len() != 2 {
			t.Errorf("expected len 2, got %d", subject.Len())
		}
	})

	t.Run("Values", func(t *testing.T) {
		var s sets.Set[string]

		object := []string{"one", "two", "three"}

		for _, value := range object {
			s.Add(value)
		}

		subject := s.Values()

		if len(subject) != 3 {
			t.Errorf("expected len 3, got %d", s.Len())
		}

		for _, value := range object {
			index := slices.Index(subject, func(v string, _ int) bool {
				return v == value
			})

			if index == -1 {
				t.Errorf("Expected values to contain %v", value)
			}
		}
	})
}
