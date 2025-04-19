package paged

import "github.com/mdelah/bitsets/bit64"

type (
	WalkFunc[T any]  func(T, int, int, bit64.Set) (T, bool)
	WalkFunc2[T any] func(T, int, int, bit64.Set, bit64.Set) (T, bool)
)
