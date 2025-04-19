package paged

import (
	"github.com/mdelah/bitsets/bit64"
	"math"
	"slices"
)

type Var struct {
	Begin            int
	Head, Tail, Body bit64.Set
	More             []bit64.Set
}

func (v *Var) Unalias() {
	if len(v.More) == 0 {
		v.More = nil
	} else {
		v.More = append([]bit64.Set(nil), v.More...)
	}
}

func (v *Var) AssignNone() {
	*v = Var{}
}

func (v *Var) AssignAll() {
	v.Head.AssignAll()
	v.Body.AssignAll()
	v.Tail.AssignAll()
	v.More = v.More[:0]
}

func (v *Var) AssignNot() {
	v.Head.AssignNot()
	v.Body.AssignNot()
	v.Tail.AssignNot()
	for i := range v.More {
		v.More[i].AssignNot()
	}
}

func (v *Var) Assign2(other *Var, fn func(*bit64.Set, bit64.Set)) {
	fn(&v.Head, other.Head)
	fn(&v.Tail, other.Tail)
	begin := min(v.Begin, other.Begin)
	end := max(v.End(), other.End())
	for i := begin; i != end; i++ {
		lhs := v.Get(i)
		rhs := other.Get(i)
		res := lhs
		fn(&res, rhs)
		if lhs != res {
			*v.Mut(i) = res
		}
	}
}

func (v *Var) End() int { return 1 + v.Begin + len(v.More) }

func (v *Var) Get(i int) bit64.Set {
	switch {
	case i < v.Begin:
		return v.Head
	case i == v.Begin:
		return v.Body
	case i >= v.End():
		return v.Tail
	default:
		return v.More[i-v.Begin-1]
	}
}

func (v *Var) Mut(i int) *bit64.Set {
	switch {
	case v.Head == v.Body && v.Body == v.Tail:
		v.Begin = i
		return &v.Body
	case i < v.Begin:
		n := v.Begin - i
		v.More = append(v.More, make([]bit64.Set, n)...)
		copy(v.More[n:], v.More)
		v.More[n-1] = v.Body
		v.Body = v.Head
		for k := range n - 1 {
			v.More[k] = v.Head
		}
		return &v.Body
	case i == v.Begin:
		return &v.Body
	case i >= v.End():
		n := 1 + v.End() - i
		v.More = slices.Grow(v.More, n)
		if v.Tail != 0 {
			for k := len(v.More) - 1 - n; k != len(v.More); k++ {
				v.More[k] = v.Tail
			}
		}
		return &v.More[len(v.More)-1]
	default:
		return &v.More[i-v.Begin-1]
	}
}

func (v Var) Each(yield func(int) bool) {
	WalkVar(v, struct{}{}, Each(yield).Walk)
}

func (v Var) EachRange(yield func(int, int) bool) {
	carry := WalkVar(v, -1, EachRange(yield).Walk)
	if carry != -1 {
		yield(carry, math.MaxInt)
	}
}
