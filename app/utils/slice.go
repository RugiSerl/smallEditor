package utils

// Slow function to remove element if the order in the slice matters.
func Remove[T any](slice []T, i int) []T {
	a := append(slice[:i], slice[i+1:]...)
	return a
}
