package mapUtil

// 聚合
func Collect[K comparable, T any](m map[K][]T, choose func(x1, x2 T) T) map[K]T {
	ret := make(map[K]T)
	if len(m) == 0 {
		return ret
	}
	for k, collection := range m {
		// var initial T
		if len(collection) == 0 {
			// ret[k] = initial
			continue
		}
		if len(collection) == 1 {
			ret[k] = collection[0]
			continue
		}
		initial := collection[0]
		for i := range collection {
			if i == 0 {
				continue
			}
			initial = choose(initial, collection[i])
		}
		ret[k] = initial
	}
	return ret
}

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, len(m))

	var i int
	for k := range m {
		keys[i] = k
		i++
	}

	return keys
}

func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, len(m))

	var i int
	for _, v := range m {
		values[i] = v
		i++
	}

	return values
}
