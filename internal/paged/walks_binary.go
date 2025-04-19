package paged

import "github.com/mdelah/bitsets/bit64"

func Compare(_, _, _ int, lhs, rhs bit64.Set) (int, bool) {
	d := lhs.Compare(rhs)
	if d != 0 {
		return d, true
	}
	return 0, false
}

func AndCount(val, _, _ int, lhs, rhs bit64.Set) (int, bool) {
	return val + lhs.And(rhs).Count(), false
}

func Equal(_ bool, _, _ int, lhs, rhs bit64.Set) (bool, bool) {
	if lhs.Equal(rhs) {
		return true, false
	}
	return false, true
}

func HasNone(_ bool, _, _ int, lhs, rhs bit64.Set) (bool, bool) {
	if lhs.HasNone(rhs) {
		return true, false
	}
	return false, true
}

func HasAll(_ bool, _, _ int, lhs, rhs bit64.Set) (bool, bool) {
	if lhs.HasAll(rhs) {
		return true, false
	}
	return false, true
}
