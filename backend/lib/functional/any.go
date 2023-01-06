package functional

func Any[T any](collection []T, predicate func(T) bool) bool {
	for _, element := range collection {
		if predicate(element) {
			return true
		}
	}
	return false
}
