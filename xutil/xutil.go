package xutil

// InArray 将在数组中搜索任意类型的元素，返回匹配元素的布尔值。
func InArray[T comparable](element T, collection []T) bool {
	for _, item := range collection {
		if item == element {
			return true
		}
	}
	return false
}

func Contains[T comparable](element T, collection []T) bool {
	for _, item := range collection {
		if item == element {
			return true
		}
	}
	return false
}
