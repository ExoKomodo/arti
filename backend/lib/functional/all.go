package functional

func All[T any](collection []T, predicate func(T) bool) bool {
	for _, element := range collection {
		if !predicate(element) {
			return false
		}
	}
	return true
}
