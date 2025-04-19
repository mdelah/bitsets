package paged

import (
	"github.com/mdelah/bitsets/bit64"
	"math"
)

type Slice []uint64

func (s Slice) Each(yield func(int) bool) {
	WalkSlice(s, struct{}{}, Each(yield).Walk)
}

func (s Slice) EachRange(yield func(int, int) bool) {
	carry := WalkSlice(s, -1, EachRange(yield).Walk)
	if carry != -1 {
		yield(carry, math.MaxInt)
	}
}

func (s Slice) Set2(other Slice, fn func(*bit64.Set, bit64.Set)) {
	for i, rhs := range other {
		fn((*bit64.Set)(&s[i]), bit64.Set(rhs))
	}
}
