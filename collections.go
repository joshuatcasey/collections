package collections

// FilterFunc returns an array containing only those elements of a for which f returns true
func FilterFunc[T any](a []T, f func(T) bool) []T {
	result := make([]T, 0)

	for _, t := range a {
		if f(t) {
			result = append(result, t)
		}
	}

	return result
}

// TransformFunc returns an array containing the result of f for all elements of a
func TransformFunc[T, U any](a []T, f func(T) U) []U {
	result := make([]U, 0)

	for _, t := range a {
		result = append(result, f(t))
	}

	return result
}

// TransformFuncWithError returns an array containing the result of f for all elements of a
// and allows f to return an error
func TransformFuncWithError[T, U any](a []T, f func(T) (U, error)) ([]U, error) {
	result := make([]U, 0)

	for _, t := range a {
		if u, err := f(t); err != nil {
			return []U{}, err
		} else {
			result = append(result, u)
		}
	}

	return result, nil
}
