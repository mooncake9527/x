package sliceUtil


func Contain[T comparable](collection []T, element T) bool {
	for i := range collection {
		if collection[i] == element {
			return true
		}
	}

	return false
}

func NotContain[T comparable](collection []T, element T) bool {
	return !Contain(collection, element)
}

func GroupBy[T any, U comparable](collection []T, iteratee func(item T) U) map[U][]T {
	result := make(map[U][]T)
	for i := range collection {
		key := iteratee(collection[i])
		result[key] = append(result[key], collection[i])
	}
	return result
}
