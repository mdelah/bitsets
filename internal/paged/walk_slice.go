package paged

import "github.com/mdelah/bitsets/bit64"

func WalkSlice[T any](s []uint64, val T, fn WalkFunc[T]) T {
	stop := false
	for i, mask := range s {
		val, stop = fn(val, i, 1, bit64.Set(mask))
		if stop {
			return val
		}
	}
	return val
}

func WalkSliceBack[T any](s []uint64, val T, fn WalkFunc[T]) T {
	stop := false
	for i := len(s) - 1; i >= 0; i-- {
		val, stop = fn(val, i, 1, bit64.Set(s[i]))
		if stop {
			return val
		}
	}
	return val
}

func WalkSlice2[T any](lhs, rhs []uint64, val T, fn WalkFunc2[T]) T {
	stop := false
	for i, mask := range lhs {
		val, stop = fn(val, i, 1, bit64.Set(mask), bit64.Set(rhs[i]))
		if stop {
			return val
		}
	}
	return val
}
