package slices

func Reduce[T any, V any](arr []T, acc V, fn func(accumulator V, element T, index int) V) V {
	for currentIndex, currentElement := range arr {
		acc = fn(acc, currentElement, currentIndex)
	}

	return acc
}
