package sets

import "github.com/robdasilva/stride/maps"

type Set[T comparable] map[T]struct{}

func (s *Set[T]) Add(value T) {
	if *s == nil {
		*s = make(Set[T])
	}

	(*s)[value] = struct{}{}
}

func (s *Set[T]) Contains(value T) bool {
	if _, ok := (*s)[value]; ok {
		return true
	}

	return false
}

func (s *Set[T]) Len() int {
	return len(*s)
}

func (s *Set[T]) Values() []T {
	return maps.Keys(*s)
}
