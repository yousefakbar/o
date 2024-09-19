package utils

// Contains: Checks if a `slice` contains a specific `item`, regardless of type
func Contains[T comparable](slice []T, item T) bool {
	for _, val := range slice {
		if val == item {
			return true
		}
	}
	return false
}
