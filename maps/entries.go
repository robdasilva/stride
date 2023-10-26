package maps

type Entry[K comparable, V any] struct {
	Key   K
	Value V
}

func Entries[K comparable, V any](m map[K]V) []Entry[K, V] {
	values := make([]Entry[K, V], 0, len(m))

	for key, value := range m {
		values = append(values, Entry[K, V]{Key: key, Value: value})
	}

	return values
}

func FromEntries[K comparable, V any](entries []Entry[K, V]) map[K]V {
	m := make(map[K]V, len(entries))

	for _, entry := range entries {
		m[entry.Key] = entry.Value
	}

	return m
}
