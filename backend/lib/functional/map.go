package functional

func GoMap[T, U any](collection []T, mapper func(T) U) []U {
	success := make(chan struct {
		index  int
		mapped U
	})
	for i, element := range collection {
		go func(
			i int,
			element T,
			success chan struct {
				index  int
				mapped U
			},
		) {
			success <- struct {
				index  int
				mapped U
			}{i, mapper(element)}
		}(i, element, success)
	}

	results := make([]U, len(collection))
	for range collection {
		result := <-success
		results[result.index] = result.mapped
	}
	close(success)
	return results
}

func Map[T, U any](collection []T, mapper func(T) U) []U {
	mapped := make([]U, len(collection))

	for i, element := range collection {
		mapped[i] = mapper(element)
	}

	return mapped
}
