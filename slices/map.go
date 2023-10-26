package slices

func Map[T any, V any](arr []T, fn func(element T, index int) V) []V {
	return Reduce(arr, make([]V, 0, len(arr)), func(accumulator []V, currentElement T, currentIndex int) []V {
		return append(accumulator, fn(currentElement, currentIndex))
	})
}
