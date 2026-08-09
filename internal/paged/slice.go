package paged

import (
	"github.com/mdelah/bitsets/bit64"
)

type Slice []uint64

func (s Slice) Each(yield func(int) bool) {
	WalkSlice(s, struct{}{}, Each(yield).Walk)
}

// EachRange loops over contiguous sub-ranges. Unlike Var (an unbounded set with a genuine infinite
// Tail), a Slice is a fixed-width bitset -- so a trailing open run (the walk finished with carry still
// set, meaning the last word's top bit was part of a run) ends at the slice's own last bit, not at
// math.MaxInt. Using math.MaxInt here (copied from Var's EachRange, which is correct for ITS unbounded
// case) was the bug: Ranges() on a full Large-variant set (bit128/kbit/kbit8) reported its single range
// as [0, math.MaxInt] instead of [0, Max].
func (s Slice) EachRange(yield func(int, int) bool) {
	carry := WalkSlice(s, -1, EachRange(yield).Walk)
	if carry != -1 {
		yield(carry, 64*len(s)-1)
	}
}

func (s Slice) Set2(other Slice, fn func(*bit64.Set, bit64.Set)) {
	for i, rhs := range other {
		fn((*bit64.Set)(&s[i]), bit64.Set(rhs))
	}
}
