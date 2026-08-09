package bit32

import (
	"cmp"
	"iter"
	"math/bits"

	"github.com/mdelah/bitsets/internal/abstract"
)

// Set is a fixed-width bitset able to store `int` values in the inclusive range [0, Max].
// Inputs outside this range are not supported and may panic or otherwise behave unexpectedly.
// The zero value is empty and ready to use.
type Set uint32

const (
	// Cap is the maximum possible size for the set.
	Cap = 32

	// Max is the largest value that can be stored.
	Max = 31

	// None is the empty set.
	None = Set(0)

	// All is the set containing all possible values.
	All = ^Set(0)
)

// Value returns a set containing the one value.
// The value must be in [0, Max].
func Value(value int) Set { return Set(1 << value) }

// Less returns a set containing all values smaller than that given.
// The value must be in [0, Max].
func Less(value int) Set { return Set((1 << value) - 1) }

// More returns a set contains all values greater than that given.
// The value must be in [0, Max].
func More(value int) Set { return Set(^uint32(1<<(value+1) - 1)) }

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
func (s Set) Count() int { return bits.OnesCount32(uint32(s)) }

// String returns a human-readable form of the set.
func (s Set) String() string { return abstract.Format(s) }

// IsNone reports whether the set contains no values.
func (s Set) IsNone() bool { return s == None }

// IsAll reports whether the set contains all possible values.
func (s Set) IsAll() bool { return s == All }

// Min returns the smallest value in the set. It returns -1 if empty.
func (s Set) Min() int {
	switch val := bits.TrailingZeros32(uint32(s)); val {
	case Cap:
		return -1
	default:
		return val
	}
}

// Max returns the largest value in the set. It returns -1 if empty.
func (s Set) Max() int {
	switch val := bits.LeadingZeros32(uint32(s)); val {
	case Cap:
		return -1
	default:
		return Max - val
	}
}

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
	if from > Max {
		return 0, false
	}
	word := uint32(s) &^ (1<<uint(from) - 1)
	if word == 0 {
		return 0, false
	}
	return bits.TrailingZeros32(word), true
}

// Has reports whether the set holds the value given.
// The value must be in [0, Max].
func (s Set) Has(value int) bool { return !s.HasNone(Value(value)) }

// Equal tests if the set is the same as another.
func (s Set) Equal(other Set) bool { return s == other }

// HasAny reports whether the set has any values in common with another.
func (s Set) HasAny(other Set) bool { return !s.HasNone(other) }

// HasNone reports whether the set does not have any values in common with another.
func (s Set) HasNone(other Set) bool { return s&other == None }

// HasAll reports whether the set holds all values present in another.
func (s Set) HasAll(other Set) bool { return s|^other == All }

// Compare returns 0 if the set are equal, else, for the smallest value whose presence is not the same between the two
// sets, a positive (negative) number if present in the left (right) hand side.
func (s Set) Compare(other Set) int {
	return cmp.Compare(bits.Reverse32(uint32(s)), bits.Reverse32(uint32(other)))
}

// LessCount returns the number of values in the set less than the given value.
// The value must be in [0, Max].
func (s Set) LessCount(value int) int { return s.And(Less(value)).Count() }

// MoreCount returns the number of values in the set greater than the given value.
// The value must be in [0, Max].
func (s Set) MoreCount(value int) int { return s.And(More(value)).Count() }

// AndCount returns the number of values the set has in common the other.
func (s Set) AndCount(other Set) int { return s.And(other).Count() }

// Add puts a value into the set if not already present.
// The value must be in [0, Max].
func (s *Set) Add(value int) { s.AssignOr(Value(value)) }

// Remove deletes a value from the set if present.
// The value must be in [0, Max].
func (s *Set) Remove(value int) { s.AssignSub(Value(value)) }

// Assign replaces the values with those from another set.
func (s *Set) Assign(other Set) { *s = other }

// AssignNone removes all values from the set.
func (s *Set) AssignNone() { *s = None }

// AssignAll adds all possible values to the set.
func (s *Set) AssignAll() { *s = All }

// Each loops over the values of the set in ascending order.
func (s Set) Each() iter.Seq[int] { return s.each }

// Ranges loops over contiguous sub-ranges of the set in ascending order.
// Each iteration yields inclusive range bounds [start, end].
func (s Set) Ranges() iter.Seq2[int, int] { return s.eachRange }

// Not returns the set of absent values.
func (s Set) Not() Set { return ^s }

// AssignNot assigns the set of absent values.
func (s *Set) AssignNot() { *s = ^*s }

// Sub returns the set of values present on the left but not right.
func (s Set) Sub(other Set) Set { return s &^ other }

// And returns the set of values common to both sides.
func (s Set) And(other Set) Set { return s & other }

// Or returns the set of values on either side (or both).
func (s Set) Or(other Set) Set { return s | other }

// Xor returns the set of values on exactly one side.
func (s Set) Xor(other Set) Set { return s ^ other }

// Nor returns the set of values absent from both sides.
func (s Set) Nor(other Set) Set { return ^s &^ other }

// Iff returns the set of values on both sides, or neither.
func (s Set) Iff(other Set) Set { return s ^ ^other }

// Imply returns the set of values on the right side, or not the left.
func (s Set) Imply(other Set) Set { return ^s | other }

// AssignSub assigns the set of values present on the left but not right.
func (s *Set) AssignSub(other Set) { *s &= ^other }

// AssignAnd assigns the set of values common to both sides.
func (s *Set) AssignAnd(other Set) { *s &= other }

// AssignOr assigns the set of values on either side (or both).
func (s *Set) AssignOr(other Set) { *s |= other }

// AssignXor assigns the set of values on exactly one side.
func (s *Set) AssignXor(other Set) { *s ^= other }

// AssignNor assigns the set of values absent from both sides.
func (s *Set) AssignNor(other Set) { *s = ^*s &^ other }

// AssignIff assigns the set of values on both sides, or neither.
func (s *Set) AssignIff(other Set) { *s ^= ^other }

// AssignImply assigns the set of values on the right side, or not the left.
func (s *Set) AssignImply(other Set) { *s = ^*s | other }

func (s Set) each(yield func(int) bool) {
	for {
		value := s.Min()
		if value == -1 || !yield(value) {
			return
		}
		s.Remove(value)
	}
}

func (s Set) eachRange(yield func(int, int) bool) {
	for {
		left := s.Min()
		if left == -1 {
			return
		}
		s.AssignNor(Less(left))
		right := s.Min()
		if right == -1 {
			yield(left, Max)
			return
		}
		if !yield(left, right-1) {
			return
		}
		s.AssignNor(Less(right))
	}
}
