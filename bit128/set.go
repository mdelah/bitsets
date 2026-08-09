package bit128

import (
	"iter"
	"math/bits"

	"github.com/mdelah/bitsets/bit64"
	"github.com/mdelah/bitsets/internal/abstract"
	"github.com/mdelah/bitsets/internal/paged"
)

// Set is a fixed-width bitset able to store `int` values in the inclusive range [0, Max].
// Inputs outside this range are not supported and may panic or otherwise behave unexpectedly.
// The zero value is empty and ready to use.
type Set [2]uint64

const (
	// Cap is the maximum possible size for the set.
	Cap = 128

	// Max is the largest value that can be stored.
	Max = 127
)

// Value returns a set containing the one value.
// The value must be in [0, Max].
func Value(value int) Set {
	var s Set
	s.Add(value)
	return s
}

// All returns a set containing all possible values.
func All() Set {
	var s Set
	for i := range s {
		s[i] = ^uint64(0)
	}
	return s
}

// None returns an empty set.
func None() Set { return Set{} }

// Less returns a set containing all values smaller than that given.
// The value must be in [0, Max].
func Less(value int) Set {
	var s Set
	i := value / 64
	for k := range i {
		s[k] = ^uint64(0)
	}
	s[i] = 1<<(value%64) - 1
	return s
}

// More returns a set contains all values greater than that given.
// The value must be in [0, Max].
func More(value int) Set {
	var s Set
	i := value / 64
	for k := i + 1; k < 2; k++ {
		s[k] = ^uint64(0)
	}
	s[i] = ^uint64(1<<(value%64+1) - 1)
	return s
}

// Values returns a set containing the given values.
// Every value must be in [0, Max].
func Values(values ...int) Set {
	var s Set
	for _, value := range values {
		s.Add(value)
	}
	return s
}

// Count returns the number of values in the set.
func (s Set) Count() int { return paged.WalkSlice(s[:], 0, paged.Count) }

// String returns a human-readable form of the set.
func (s Set) String() string { return abstract.Format(s) }

// IsNone reports whether the set contains no values.
func (s Set) IsNone() bool { return paged.WalkSlice(s[:], true, paged.IsNone) }

// IsAll reports whether the set contains all possible values.
func (s Set) IsAll() bool { return paged.WalkSlice(s[:], true, paged.IsAll) }

// Min returns the smallest value in the set. It returns -1 if empty.
func (s Set) Min() int { return paged.WalkSlice(s[:], -1, paged.Min) }

// Max returns the largest value in the set. It returns -1 if empty.
func (s Set) Max() int { return paged.WalkSliceBack(s[:], -1, paged.Max) }

// Next returns the smallest member of the set that is >= from, and true; or (0, false) if there is
// none. Unlike Each/Ranges, this allocates nothing: Each returns an iter.Seq[int], which -- because it
// crosses this package's own API boundary as a bound method value or closure -- necessarily allocates
// once per call, even though the loop it drives never escapes the caller. Next is a plain method call
// with no closure of its own, so it is the preferred way to walk a Set's members in a hot loop:
//
//	for v, ok := s.Next(0); ok; v, ok = s.Next(v + 1) {
//		use(v)
//	}
//
// A negative from is treated as 0.
func (s Set) Next(from int) (int, bool) {
	if from < 0 {
		from = 0
	}
	w := from / 64
	if w >= len(s) {
		return 0, false
	}
	word := s[w] &^ (1<<uint(from%64) - 1)
	for word == 0 {
		w++
		if w >= len(s) {
			return 0, false
		}
		word = s[w]
	}
	return w*64 + bits.TrailingZeros64(word), true
}

// Has reports whether the set holds the value given.
// The value must be in [0, Max].
func (s Set) Has(value int) bool { return bit64.Set(s[value/64]).Has(value % 64) }

// Equal tests if the set is the same as another.
func (s Set) Equal(other Set) bool { return paged.WalkSlice2(s[:], other[:], true, paged.Equal) }

// HasAny reports whether the set has any values in common with another.
func (s Set) HasAny(other Set) bool { return !s.HasNone(other) }

// HasNone reports whether the set does not have any values in common with another.
func (s Set) HasNone(other Set) bool { return paged.WalkSlice2(s[:], other[:], true, paged.HasNone) }

// HasAll reports whether the set holds all values present in another.
func (s Set) HasAll(other Set) bool { return paged.WalkSlice2(s[:], other[:], true, paged.HasAll) }

// Compare returns 0 if the set are equal, else, for the smallest value whose presence is not the same between the two
// sets, a positive (negative) number if present in the left (right) hand side.
func (s Set) Compare(other Set) int { return paged.WalkSlice2(s[:], other[:], 0, paged.Compare) }

// LessCount returns the number of values in the set less than the given value.
// The value must be in [0, Max].
func (s Set) LessCount(value int) int {
	return paged.WalkSlice(s[:], 0, paged.LessCount{Page: value / 64, Offset: value % 64}.Walk)
}

// MoreCount returns the number of values in the set greater than the given value.
// The value must be in [0, Max].
func (s Set) MoreCount(value int) int {
	return paged.WalkSliceBack(s[:], 0, paged.MoreCount{Page: value / 64, Offset: value % 64}.Walk)
}

// AndCount returns the number of values the set has in common the other.
func (s Set) AndCount(other Set) int { return paged.WalkSlice2(s[:], other[:], 0, paged.AndCount) }

// Add puts a value into the set if not already present.
// The value must be in [0, Max].
func (s *Set) Add(value int) { (*bit64.Set)(&s[value/64]).Add(value % 64) }

// Remove deletes a value from the set if present.
// The value must be in [0, Max].
func (s *Set) Remove(value int) { (*bit64.Set)(&s[value/64]).Remove(value % 64) }

// Assign replaces the values with those from another set.
func (s *Set) Assign(other Set) { copy(s[:], other[:]) }

// AssignNone removes all values from the set.
func (s *Set) AssignNone() { clear(s[:]) }

// AssignAll adds all possible values to the set.
func (s *Set) AssignAll() {
	for i := range s {
		(*bit64.Set)(&s[i]).AssignAll()
	}
}

// Each loops over the values of the set in ascending order.
func (s Set) Each() iter.Seq[int] { return paged.Slice(s[:]).Each }

// Ranges loops over contiguous sub-ranges of the set in ascending order.
// Each iteration yields inclusive range bounds [start, end].
func (s Set) Ranges() iter.Seq2[int, int] { return paged.Slice(s[:]).EachRange }

// Not returns the set of absent values.
func (s Set) Not() Set {
	var res Set
	res.Assign(s)
	res.AssignNot()
	return res
}

// AssignNot assigns the set of absent values.
func (s *Set) AssignNot() {
	for i := range s {
		(*bit64.Set)(&s[i]).AssignNot()
	}
}

// Sub returns the set of values present on the left but not right.
func (s Set) Sub(other Set) Set {
	var res Set
	res.Assign(s)
	res.AssignSub(other)
	return res
}

// And returns the set of values common to both sides.
func (s Set) And(other Set) Set {
	var res Set
	res.Assign(s)
	res.AssignAnd(other)
	return res
}

// Or returns the set of values on either side (or both).
func (s Set) Or(other Set) Set {
	var res Set
	res.Assign(s)
	res.AssignOr(other)
	return res
}

// Xor returns the set of values on exactly one side.
func (s Set) Xor(other Set) Set {
	var res Set
	res.Assign(s)
	res.AssignXor(other)
	return res
}

// Nor returns the set of values absent from both sides.
func (s Set) Nor(other Set) Set {
	var res Set
	res.Assign(s)
	res.AssignNor(other)
	return res
}

// Iff returns the set of values on both sides, or neither.
func (s Set) Iff(other Set) Set {
	var res Set
	res.Assign(s)
	res.AssignIff(other)
	return res
}

// Imply returns the set of values on the right side, or not the left.
func (s Set) Imply(other Set) Set {
	var res Set
	res.Assign(s)
	res.AssignImply(other)
	return res
}

// AssignSub assigns the set of values present on the left but not right.
func (s *Set) AssignSub(other Set) { paged.Slice(s[:]).Set2(other[:], (*bit64.Set).AssignSub) }

// AssignAnd assigns the set of values common to both sides.
func (s *Set) AssignAnd(other Set) { paged.Slice(s[:]).Set2(other[:], (*bit64.Set).AssignAnd) }

// AssignOr assigns the set of values on either side (or both).
func (s *Set) AssignOr(other Set) { paged.Slice(s[:]).Set2(other[:], (*bit64.Set).AssignOr) }

// AssignXor assigns the set of values on exactly one side.
func (s *Set) AssignXor(other Set) { paged.Slice(s[:]).Set2(other[:], (*bit64.Set).AssignXor) }

// AssignNor assigns the set of values absent from both sides.
func (s *Set) AssignNor(other Set) { paged.Slice(s[:]).Set2(other[:], (*bit64.Set).AssignNor) }

// AssignIff assigns the set of values on both sides, or neither.
func (s *Set) AssignIff(other Set) { paged.Slice(s[:]).Set2(other[:], (*bit64.Set).AssignIff) }

// AssignImply assigns the set of values on the right side, or not the left.
func (s *Set) AssignImply(other Set) { paged.Slice(s[:]).Set2(other[:], (*bit64.Set).AssignImply) }
