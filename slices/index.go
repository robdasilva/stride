package slices

func Index[T any](arr []T, fn func(element T, index int) bool) int {
	for currentIndex, currentElement := range arr {
		if fn(currentElement, currentIndex) {
			return currentIndex
		}
	}

	return -1
}
