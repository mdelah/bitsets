package paged

import "github.com/mdelah/bitsets/bit64"

func IsNone(_ bool, _, _ int, mask bit64.Set) (bool, bool) {
	if mask.IsNone() {
		return true, false
	}
	return false, true
}

func IsAll(_ bool, _, _ int, mask bit64.Set) (bool, bool) {
	if mask.IsAll() {
		return true, false
	}
	return false, true
}

func Count(val, _, mult int, mask bit64.Set) (int, bool) {
	return val + mult*mask.Count(), false
}

func Max(_, i, _ int, mask bit64.Set) (int, bool) {
	val := mask.Max()
	if val == -1 {
		return -1, false
	}
	return 64*i + val, true
}

type LessCount struct{ Page, Offset int }

func (r LessCount) Walk(val, i, mult int, mask bit64.Set) (int, bool) {
	switch {
	case r.Page >= i+mult:
		return val + mask.Count()*mult, false
	case r.Page == i:
		return val + mask.LessCount(r.Offset), true
	default:
		return val + mask.Count()*(r.Page-i) + mask.LessCount(r.Offset), true
	}
}

func Min(_, i, _ int, mask bit64.Set) (int, bool) {
	val := mask.Min()
	if val == -1 {
		return -1, false
	}
	return 64*i + val, true
}

type MoreCount struct{ Page, Offset int }

func (r MoreCount) Walk(val, i, mult int, mask bit64.Set) (int, bool) {
	switch {
	case r.Page <= i-mult:
		return val + mask.Count()*mult, false
	case r.Page == i:
		return val + mask.MoreCount(r.Offset), true
	default:
		return val + mask.Count()*(i-r.Page) + mask.MoreCount(r.Offset), true
	}
}
