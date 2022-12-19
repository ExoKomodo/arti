package functional

func Reduce[T, U any](collection []T, reducer func(T, U) U, accumulator U) U {
	for _, element := range collection {
		accumulator = reducer(element, accumulator)
	}

	return accumulator
}
