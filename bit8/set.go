package bit8

import (
	"cmp"
	"iter"
	"math/bits"

	"github.com/mdelah/bitsets/internal/abstract"
)

// Set is fixed-width bitset able to store non-negative `int` values up to 7. The zero value is empty and ready to use.
type Set uint8

const (
	// Cap is the maximum possible size for the set.
	Cap = 8

	// Max is the largest value that can be stored.
	Max = 7

	// None is the empty set.
	None = Set(0)

	// All is the set containing all possible values.
	All = ^Set(0)
)

// Value returns a set containing the one value.
func Value(value int) Set { return Set(1 << value) }

// Less returns a set containing all values smaller than that given.
func Less(value int) Set { return Set((1 << value) - 1) }

// More returns a set contains all values greater than that given.
func More(value int) Set { return Set(^uint8(1<<(value+1) - 1)) }

// Values returns a set containing the given values.
func Values(values ...int) Set {
	var s Set
	for _, value := range values {
		s.Add(value)
	}
	return s
}

// Count returns the number of values in the set.
func (s Set) Count() int { return bits.OnesCount8(uint8(s)) }

// String returns a human-readable form of the set.
func (s Set) String() string { return abstract.Format(s) }

// IsNone reports whether the set contains no values.
func (s Set) IsNone() bool { return s == None }

// IsAll reports whether the set contains all possible values.
func (s Set) IsAll() bool { return s == All }

// Min returns the smallest value in the set. It returns -1 if empty.
func (s Set) Min() int {
	switch val := bits.TrailingZeros8(uint8(s)); val {
	case Cap:
		return -1
	default:
		return val
	}
}

// Max returns the largest value in the set. It returns -1 if empty.
func (s Set) Max() int {
	switch val := bits.LeadingZeros8(uint8(s)); val {
	case Cap:
		return -1
	default:
		return Max - val
	}
}

// Has reports whether the set holds the value given.
func (s Set) Has(value int) bool { return !s.HasNone(Value(value)) }

// Equal tests if the set is the same as another.
func (s Set) Equal(other Set) bool { return s == other }

// HasNone reports whether the set does not have any values in common with another.
func (s Set) HasNone(other Set) bool { return s&other == None }

// HasAll reports whether the set holds all values present in another.
func (s Set) HasAll(other Set) bool { return s|^other == All }

// Compare returns 0 if the set are equal, else, for the smallest value whose presence is not the same between the two
// sets, a positive (negative) number if present in the left (right) hand side.
func (s Set) Compare(other Set) int {
	return cmp.Compare(bits.Reverse8(uint8(s)), bits.Reverse8(uint8(other)))
}

// LessCount returns the number of values in the set less than the given value.
func (s Set) LessCount(value int) int { return s.And(Less(value)).Count() }

// MoreCount returns the number of values in the set greater than the given value.
func (s Set) MoreCount(value int) int { return s.And(More(value)).Count() }

// AndCount returns the number of values the set has in common the other.
func (s Set) AndCount(other Set) int { return s.And(other).Count() }

// Add puts a value into the set if not already present.
func (s *Set) Add(value int) { s.AssignOr(Value(value)) }

// Remove deletes a value from the set if present.
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
// Each iteration produces the first value of the range, and the smallest value greater than that absent from the set.
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
