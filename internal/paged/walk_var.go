package paged

import "math"

func WalkVar[T any](s Var, val T, fn WalkFunc[T]) T {
	stop := false
	if s.Begin > 0 {
		val, stop = fn(val, 0, s.Begin, s.Head)
		if stop {
			return val
		}
	}
	val, stop = fn(val, s.Begin, 1, s.Body)
	if stop {
		return val
	}
	for i, mask := range s.More {
		val, stop = fn(val, 1+s.Begin+i, 1, mask)
		if stop {
			return val
		}
	}
	val, _ = fn(val, s.End(), math.MaxInt/64-s.End(), s.Tail)
	return val
}

func WalkVarBack[T any](s Var, val T, fn WalkFunc[T]) T {
	val, stop := fn(val, math.MaxInt/64, math.MaxInt/64-s.End(), s.Tail)
	if stop {
		return val
	}
	for i := s.End() - 1; i > s.Begin; i-- {
		val, stop = fn(val, i, 1, s.More[i])
		if stop {
			return val
		}
	}
	val, stop = fn(val, s.Begin, 1, s.Body)
	if stop {
		return val
	}
	if s.Begin > 0 {
		val, stop = fn(val, s.Begin-1, s.Begin, s.Head)
		if stop {
			return val
		}
	}
	return val
}

func WalkVar2[T any](lhs, rhs Var, val T, fn WalkFunc2[T]) T {
	stop := false
	begin := min(lhs.Begin, rhs.Begin)
	if begin > 0 {
		val, stop = fn(val, 0, begin, lhs.Head, rhs.Head)
		if stop {
			return val
		}
	}
	end := max(lhs.End(), rhs.End())
	for i := begin; i != end; i++ {
		val, stop = fn(val, i, 1, lhs.Get(i), rhs.Get(i))
		if stop {
			return val
		}
	}
	val, _ = fn(val, end, math.MaxInt/64-end, lhs.Tail, rhs.Tail)
	return val
}
