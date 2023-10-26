package slices

func Filter[T any](arr []T, fn func(element T, index int) bool) []T {
	return Reduce(arr, make([]T, 0), func(accumulator []T, currentElement T, currentIndex int) []T {
		if fn(currentElement, currentIndex) {
			return append(accumulator, currentElement)
		}

		return accumulator
	})
}
