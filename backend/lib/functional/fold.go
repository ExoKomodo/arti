package functional

func Fold[T any](collection []T, reducer func(T, T) T, accumulator T) T {
	for i := len(collection) - 1; i >= 0; i-- {
		element := collection[i]
		accumulator = reducer(element, accumulator)
	}

	return accumulator
}

func FoldRight[T any](collection []T, reducer func(T, T) T, accumulator T) T {
	for _, element := range collection {
		accumulator = reducer(element, accumulator)
	}

	return accumulator
}
