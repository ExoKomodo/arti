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

func Reduce[T any](collection []T, reducer func(T, T) T) T {
	length := len(collection)
	if length == 0 {
		var accumulator T
		return accumulator
	}
	if length == 1 {
		return collection[0]
	}
	return Fold(collection[:length-1], reducer, collection[length-1])
}
