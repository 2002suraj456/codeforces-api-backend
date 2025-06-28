package utils

func Filter[T any](slice []T, predicate func(T) bool) []T {
	result := make([]T, 0)

	for _, item := range slice {
		if predicate(item) {
			result = append(result, item)
		}
	}

	return result
}

func Find[T comparable](slice []T, item T) bool {
	for _, el := range slice {
		if el == item {
			return true
		}
	}

	return false
}
