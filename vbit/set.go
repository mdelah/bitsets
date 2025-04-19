package vbit

import (
	"iter"

	"github.com/mdelah/bitsets/bit64"
	"github.com/mdelah/bitsets/internal/abstract"
	"github.com/mdelah/bitsets/internal/paged"
)

// Set is variable-width bitset able to store non-negative `int` values. The zero value is empty and ready to use.
type Set struct{ v paged.Var }

// Value returns a set containing the one value.
func Value(value int) Set { return Set{paged.Var{Begin: value / 64, Body: bit64.Value(value % 64)}} }

// All returns a set containing all possible values.
func All() Set { return Set{paged.Var{Head: bit64.All, Body: bit64.All, Tail: bit64.All}} }

// None returns an empty set.
func None() Set { return Set{} }

// Less returns a set containing all values smaller than that given.
func Less(value int) Set {
	return Set{paged.Var{Head: bit64.All, Begin: value / 64, Body: bit64.Less(value % 64)}}
}

// More returns a set contains all values greater than that given.
func More(value int) Set {
	return Set{paged.Var{Begin: value / 64, Body: bit64.More(value % 64), Tail: bit64.All}}
}

// Values returns a set containing the given values.
func Values(values ...int) Set {
	var s Set
	for _, value := range values {
		s.Add(value)
	}
	return s
}

// Count returns the number of values in the set.
func (s Set) Count() int { return paged.WalkVar(s.v, 0, paged.Count) }

// String returns a human-readable form of the set.
func (s Set) String() string { return abstract.Format(s) }

// IsNone reports whether the set contains no values.
func (s Set) IsNone() bool { return paged.WalkVar(s.v, true, paged.IsNone) }

// IsAll reports whether the set contains all possible values.
func (s Set) IsAll() bool { return paged.WalkVar(s.v, true, paged.IsAll) }

// Min returns the smallest value in the set. It returns -1 if empty.
func (s Set) Min() int { return paged.WalkVar(s.v, -1, paged.Min) }

// Max returns the largest value in the set. It returns -1 if empty.
func (s Set) Max() int { return paged.WalkVarBack(s.v, -1, paged.Max) }

// Has reports whether the set holds the value given.
func (s Set) Has(value int) bool { return s.v.Get(value / 64).Has(value % 64) }

// Equal tests if the set is the same as another.
func (s Set) Equal(other Set) bool { return paged.WalkVar2(s.v, other.v, true, paged.Equal) }

// HasNone reports whether the set does not have any values in common with another.
func (s Set) HasNone(other Set) bool { return paged.WalkVar2(s.v, other.v, true, paged.HasNone) }

// HasAll reports whether the set holds all values present in another.
func (s Set) HasAll(other Set) bool { return paged.WalkVar2(s.v, other.v, true, paged.HasAll) }

// Compare returns 0 if the set are equal, else, for the smallest value whose presence is not the same between the two
// sets, a positive (negative) number if present in the left (right) hand side.
func (s Set) Compare(other Set) int { return paged.WalkVar2(s.v, other.v, 0, paged.Compare) }

// LessCount returns the number of values in the set less than the given value.
func (s Set) LessCount(value int) int {
	return paged.WalkVar(s.v, 0, paged.LessCount{value / 64, value % 64}.Walk)
}

// MoreCount returns the number of values in the set greater than the given value.
func (s Set) MoreCount(value int) int {
	return paged.WalkVarBack(s.v, 0, paged.MoreCount{value / 64, value % 64}.Walk)
}

// AndCount returns the number of values the set has in common the other.
func (s Set) AndCount(other Set) int { return paged.WalkVar2(s.v, other.v, 0, paged.AndCount) }

// Add puts a value into the set if not already present.
func (s *Set) Add(value int) { s.v.Mut(value / 64).Add(value % 64) }

// Remove deletes a value from the set if present.
func (s *Set) Remove(value int) { s.v.Mut(value / 64).Remove(value % 64) }

// Assign replaces the values with those from another set.
func (s *Set) Assign(other Set) {
	s.v = other.v
	s.v.Unalias()
}

// AssignNone removes all values from the set.
func (s *Set) AssignNone() { s.v.AssignNone() }

// AssignAll adds all possible values to the set.
func (s *Set) AssignAll() { s.v.AssignAll() }

// Each loops over the values of the set in ascending order.
func (s Set) Each() iter.Seq[int] { return s.v.Each }

// Ranges loops over contiguous sub-ranges of the set in ascending order.
// Each iteration produces the first value of the range, and the smallest value greater than that absent from the set.
func (s Set) Ranges() iter.Seq2[int, int] { return s.v.EachRange }

// Not returns the set of absent values.
func (s Set) Not() Set {
	s.v.Unalias()
	s.AssignNot()
	return s
}

// AssignNot assigns the set of absent values.
func (s *Set) AssignNot() { s.v.AssignNot() }

// Sub returns the set of values present on the left but not right.
func (s Set) Sub(other Set) Set {
	s.v.Unalias()
	s.AssignSub(other)
	return s
}

// And returns the set of values common to both sides.
func (s Set) And(other Set) Set {
	s.v.Unalias()
	s.AssignAnd(other)
	return s
}

// Or returns the set of values on either side (or both).
func (s Set) Or(other Set) Set {
	s.v.Unalias()
	s.AssignOr(other)
	return s
}

// Xor returns the set of values on exactly one side.
func (s Set) Xor(other Set) Set {
	s.v.Unalias()
	s.AssignXor(other)
	return s
}

// Nor returns the set of values absent from both sides.
func (s Set) Nor(other Set) Set {
	s.v.Unalias()
	s.AssignNor(other)
	return s
}

// Iff returns the set of values on both sides, or neither.
func (s Set) Iff(other Set) Set {
	s.v.Unalias()
	s.AssignIff(other)
	return s
}

// Imply returns the set of values on the right side, or not the left.
func (s Set) Imply(other Set) Set {
	s.v.Unalias()
	s.AssignImply(other)
	return s
}

// AssignSub assigns the set of values present on the left but not right.
func (s *Set) AssignSub(other Set) { s.v.Assign2(&other.v, (*bit64.Set).AssignSub) }

// AssignAnd assigns the set of values common to both sides.
func (s *Set) AssignAnd(other Set) { s.v.Assign2(&other.v, (*bit64.Set).AssignAnd) }

// AssignOr assigns the set of values on either side (or both).
func (s *Set) AssignOr(other Set) { s.v.Assign2(&other.v, (*bit64.Set).AssignOr) }

// AssignXor assigns the set of values on exactly one side.
func (s *Set) AssignXor(other Set) { s.v.Assign2(&other.v, (*bit64.Set).AssignXor) }

// AssignNor assigns the set of values absent from both sides.
func (s *Set) AssignNor(other Set) { s.v.Assign2(&other.v, (*bit64.Set).AssignNor) }

// AssignIff assigns the set of values on both sides, or neither.
func (s *Set) AssignIff(other Set) { s.v.Assign2(&other.v, (*bit64.Set).AssignIff) }

// AssignImply assigns the set of values on the right side, or not the left.
func (s *Set) AssignImply(other Set) { s.v.Assign2(&other.v, (*bit64.Set).AssignImply) }
