package collections

import "golang.org/x/exp/constraints"

// DeriveFunc will scan the array to gain an insight (a derivation) about the array.
// For example, given an array of ints it can calculate the max or min.
func DeriveFunc[T any](a []T, f func(current, derived T) T) T {
	var derived T

	if len(a) > 1 {
		derived = a[0]
		for i := 1; i < len(a); i++ {
			derived = f(a[i], derived)
		}
	}

	return derived
}

// Max will return the max of the ordered elements
func Max[T constraints.Ordered](a []T) T {
	return DeriveFunc(a, func(current, derived T) T {
		if current > derived {
			return current
		}
		return derived
	})
}

// MaxParams will return the max of the ordered elements
func MaxParams[T constraints.Ordered](a ...T) T {
	return Max(a)
}

// Min will return the min of the ordered elements
func Min[T constraints.Ordered](a []T) T {
	return DeriveFunc(a, func(current, derived T) T {
		if current < derived {
			return current
		}
		return derived
	})
}

// MinParams will return the min of the ordered elements
func MinParams[T constraints.Ordered](a ...T) T {
	return Min(a)
}

// Sum will return the sum of the ordered elements
func Sum[T constraints.Ordered](a []T) T {
	return DeriveFunc(a, func(current, derived T) T {
		return derived + current
	})
}

// SumParams will return the sum of the ordered elements
func SumParams[T constraints.Ordered](a ...T) T {
	return Sum(a)
}
