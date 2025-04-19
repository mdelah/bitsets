package abstract

import (
	"iter"
)

type (
	// Set1 has the common value methods not requiring a type parameter.
	Set1 interface {
		Count() int
		String() string
		IsNone() bool
		IsAll() bool
		Has(int) bool
		LessCount(int) int
		MoreCount(int) int
		Min() int
		Max() int
		Each() iter.Seq[int]
		Ranges() iter.Seq2[int, int]
	}
	// Set2 has the common pointer methods not requiring a type parameter.
	Set2 interface {
		Set1
		Add(int)
		Remove(int)
		SetNot()
		SetNone()
		SetAll()
	}
	// Set3 has the common value methods.
	Set3[T Set1] interface {
		Set1
		Equal(T) bool
		Compare(T) int
		HasNone(T) bool
		HasAll(T) bool
		AndCount(T) int
		Not() T
		Sub(T) T
		And(T) T
		Or(T) T
		Xor(T) T
		Nor(T) T
		Iff(T) T
		Imply(T) T
	}
	// Set4 has all methods.
	Set4[T Set1] interface {
		Set2
		Set3[T]
		Assign(T)
		SetSub(T)
		SetAnd(T)
		SetOr(T)
		SetXor(T)
		SetNor(T)
		SetIff(T)
		SetImply(T)
	}
)
