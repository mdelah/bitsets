package expect

import (
	"github.com/mdelah/bitsets/internal/abstract"
	"iter"
	"testing"
)

func Eq[T comparable](t *testing.T, lhs, rhs T) {
	t.Helper()
	if lhs != rhs {
		t.Errorf("got %v; wanted %v", rhs, lhs)
	}
}

func Set[T abstract.Set3[T]](t *testing.T, lhs, rhs T) {
	t.Helper()
	if !lhs.Equal(rhs) {
		t.Errorf("got %v; wanted %v", rhs, lhs)
	}
}

func Ints(t *testing.T, lhs iter.Seq[int], rhs ...int) {
	t.Helper()
	i := 0
	for item := range lhs {
		if i >= len(rhs) {
			t.Errorf("got %v items; wanted %v", i+1, len(rhs))
		}
		if item != rhs[i] {
			t.Errorf("got %v index %d; wanted %v", item, i, rhs[i])
		}
		i++
	}
	if i != len(rhs) {
		t.Errorf("got %v items; wanted %v", i, len(rhs))
	}
}
