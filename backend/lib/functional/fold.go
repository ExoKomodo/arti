package functional

func Reduce[T, U any](collection []T, reducer func(T, U) U, accumulator U) U {
	return FoldRight(collection, reducer, accumulator)
}

func FoldRight[T, U any](collection []T, reducer func(T, U) U, accumulator U) U {
	for _, element := range collection {
		accumulator = reducer(element, accumulator)
	}

	return accumulator
}

func FoldLeft[T, U any](collection []T, reducer func(T, U) U, accumulator U) U {
	for i := len(collection) - 1; i >= 0; i-- {
		element := collection[i]
		accumulator = reducer(element, accumulator)
	}

	return accumulator
}
