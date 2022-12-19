package functional

func Filter[T any](collection []T, predicate func(T) bool) []T {
	// Set to 0 capacity and max possible length, resizing the slice at the end
	filtered := make([]T, 0, len(collection))

	newSize := 0
	for _, element := range collection {
		if predicate(element) {
			newSize++
			filtered = append(filtered, element)
		}
	}

	// Three-index slice sets the capacity to a smaller subsection of the slice,
	// allowing GC to clean this up more quickly
	return append([]T(nil), filtered[:newSize:newSize]...)
}
