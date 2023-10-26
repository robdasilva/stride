package slices

func Reverse[T any](arr []T) []T {
	return Map(arr, func(element T, index int) T {
		return arr[len(arr)-(index+1)]
	})
}
