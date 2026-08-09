package bit128_test

import (
	"github.com/mdelah/bitsets/bit128"
	"github.com/mdelah/bitsets/internal/expect"
	"reflect"
	"testing"
)

func TestCount(t *testing.T) {
	expect.Eq(t, 0, bit128.None().Count())
	expect.Eq(t, bit128.Cap, bit128.All().Count())
	expect.Eq(t, 1, bit128.Value(5).Count())
	expect.Eq(t, 2, bit128.Values(3, 5).Count())
	expect.Eq(t, 5, bit128.Less(5).Count())
	expect.Eq(t, bit128.Max-5, bit128.More(5).Count())
}

func TestString(t *testing.T) {
	expect.Eq(t, "{}", bit128.None().String())
	expect.Eq(t, "{0-127}", bit128.All().String())
	expect.Eq(t, "{5}", bit128.Value(5).String())
	expect.Eq(t, "{3,5}", bit128.Values(3, 5).String())
	expect.Eq(t, "{0-4}", bit128.Less(5).String())
	expect.Eq(t, "{6-127}", bit128.More(5).String())
}

func TestIsNone(t *testing.T) {
	expect.Eq(t, true, bit128.None().IsNone())
	expect.Eq(t, false, bit128.All().IsNone())
	expect.Eq(t, false, bit128.Value(5).IsNone())
	expect.Eq(t, false, bit128.Values(3, 5).IsNone())
	expect.Eq(t, false, bit128.Less(5).IsNone())
	expect.Eq(t, false, bit128.More(5).IsNone())
}

func TestIsAll(t *testing.T) {
	expect.Eq(t, false, bit128.None().IsAll())
	expect.Eq(t, true, bit128.All().IsAll())
	expect.Eq(t, false, bit128.Value(5).IsAll())
	expect.Eq(t, false, bit128.Values(3, 5).IsAll())
	expect.Eq(t, false, bit128.Less(5).IsAll())
	expect.Eq(t, false, bit128.More(5).IsAll())
}

func TestMin(t *testing.T) {
	expect.Eq(t, -1, bit128.None().Min())
	expect.Eq(t, 0, bit128.All().Min())
	expect.Eq(t, 5, bit128.Value(5).Min())
	expect.Eq(t, 3, bit128.Values(3, 5).Min())
	expect.Eq(t, 0, bit128.Less(5).Min())
	expect.Eq(t, 6, bit128.More(5).Min())
}

func TestMax(t *testing.T) {
	expect.Eq(t, -1, bit128.None().Max())
	expect.Eq(t, bit128.Max, bit128.All().Max())
	expect.Eq(t, 5, bit128.Value(5).Max())
	expect.Eq(t, 5, bit128.Values(3, 5).Max())
	expect.Eq(t, 4, bit128.Less(5).Max())
	expect.Eq(t, bit128.Max, bit128.More(5).Max())
}

func TestHas(t *testing.T) {
	expect.Eq(t, false, bit128.None().Has(3))
	expect.Eq(t, true, bit128.All().Has(3))
	expect.Eq(t, false, bit128.Value(5).Has(3))
	expect.Eq(t, true, bit128.Values(3, 5).Has(3))
	expect.Eq(t, true, bit128.Less(5).Has(3))
	expect.Eq(t, false, bit128.More(5).Has(3))
}

func TestLessCount(t *testing.T) {
	expect.Eq(t, 0, bit128.None().LessCount(5))
	expect.Eq(t, 5, bit128.All().LessCount(5))
	expect.Eq(t, 0, bit128.Value(5).LessCount(5))
	expect.Eq(t, 1, bit128.Values(3, 5).LessCount(5))
	expect.Eq(t, 5, bit128.Less(5).LessCount(5))
	expect.Eq(t, 0, bit128.More(5).LessCount(5))
}

func TestMoreCount(t *testing.T) {
	expect.Eq(t, 0, bit128.None().MoreCount(4))
	expect.Eq(t, bit128.Cap-5, bit128.All().MoreCount(4))
	expect.Eq(t, 1, bit128.Value(5).MoreCount(4))
	expect.Eq(t, 1, bit128.Values(3, 5).MoreCount(4))
	expect.Eq(t, 0, bit128.Less(5).MoreCount(4))
	expect.Eq(t, bit128.Cap-6, bit128.More(5).MoreCount(4))
}

func TestEqual(t *testing.T) {
	expect.Eq(t, true, bit128.None().Equal(bit128.None()))
	expect.Eq(t, false, bit128.None().Equal(bit128.All()))
	expect.Eq(t, true, bit128.All().Equal(bit128.All()))
	expect.Eq(t, false, bit128.All().Equal(bit128.Value(5)))
	expect.Eq(t, true, bit128.Value(5).Equal(bit128.Value(5)))
	expect.Eq(t, false, bit128.Value(5).Equal(bit128.Values(3, 5)))
	expect.Eq(t, true, bit128.Values(3, 5).Equal(bit128.Values(3, 5)))
	expect.Eq(t, false, bit128.Values(3, 5).Equal(bit128.Less(5)))
	expect.Eq(t, true, bit128.Less(5).Equal(bit128.Less(5)))
	expect.Eq(t, false, bit128.Less(5).Equal(bit128.More(5)))
	expect.Eq(t, true, bit128.More(5).Equal(bit128.More(5)))
	expect.Eq(t, false, bit128.More(5).Equal(bit128.None()))
}

func TestHasNone(t *testing.T) {
	expect.Eq(t, true, bit128.None().HasNone(bit128.None()))
	expect.Eq(t, true, bit128.None().HasNone(bit128.All()))
	expect.Eq(t, false, bit128.All().HasNone(bit128.All()))
	expect.Eq(t, false, bit128.All().HasNone(bit128.Value(5)))
	expect.Eq(t, false, bit128.Value(5).HasNone(bit128.Value(5)))
	expect.Eq(t, false, bit128.Value(5).HasNone(bit128.Values(3, 5)))
	expect.Eq(t, false, bit128.Values(3, 5).HasNone(bit128.Values(3, 5)))
	expect.Eq(t, false, bit128.Values(3, 5).HasNone(bit128.Less(5)))
	expect.Eq(t, false, bit128.Less(5).HasNone(bit128.Less(5)))
	expect.Eq(t, true, bit128.Less(5).HasNone(bit128.More(5)))
	expect.Eq(t, false, bit128.More(5).HasNone(bit128.More(5)))
	expect.Eq(t, true, bit128.More(5).HasNone(bit128.None()))
}

func TestHasAll(t *testing.T) {
	expect.Eq(t, true, bit128.None().HasAll(bit128.None()))
	expect.Eq(t, false, bit128.None().HasAll(bit128.All()))
	expect.Eq(t, true, bit128.All().HasAll(bit128.All()))
	expect.Eq(t, true, bit128.All().HasAll(bit128.Value(5)))
	expect.Eq(t, true, bit128.Value(5).HasAll(bit128.Value(5)))
	expect.Eq(t, false, bit128.Value(5).HasAll(bit128.Values(3, 5)))
	expect.Eq(t, true, bit128.Values(3, 5).HasAll(bit128.Values(3, 5)))
	expect.Eq(t, false, bit128.Values(3, 5).HasAll(bit128.Less(5)))
	expect.Eq(t, true, bit128.Less(5).HasAll(bit128.Less(5)))
	expect.Eq(t, false, bit128.Less(5).HasAll(bit128.More(5)))
	expect.Eq(t, true, bit128.More(5).HasAll(bit128.More(5)))
	expect.Eq(t, true, bit128.More(5).HasAll(bit128.None()))
}

func TestCompare(t *testing.T) {
	expect.Eq(t, 0, bit128.None().Compare(bit128.None()))
	expect.Eq(t, -1, bit128.None().Compare(bit128.All()))
	expect.Eq(t, 0, bit128.All().Compare(bit128.All()))
	expect.Eq(t, 1, bit128.All().Compare(bit128.Value(5)))
	expect.Eq(t, 0, bit128.Value(5).Compare(bit128.Value(5)))
	expect.Eq(t, -1, bit128.Value(5).Compare(bit128.Values(3, 5)))
	expect.Eq(t, 0, bit128.Values(3, 5).Compare(bit128.Values(3, 5)))
	expect.Eq(t, -1, bit128.Values(3, 5).Compare(bit128.Less(5)))
	expect.Eq(t, 0, bit128.Less(5).Compare(bit128.Less(5)))
	expect.Eq(t, 1, bit128.Less(5).Compare(bit128.More(5)))
	expect.Eq(t, 0, bit128.More(5).Compare(bit128.More(5)))
	expect.Eq(t, 1, bit128.More(5).Compare(bit128.None()))
}

func TestAndCount(t *testing.T) {
	expect.Eq(t, 0, bit128.None().AndCount(bit128.None()))
	expect.Eq(t, 0, bit128.None().AndCount(bit128.All()))
	expect.Eq(t, bit128.Cap, bit128.All().AndCount(bit128.All()))
	expect.Eq(t, 1, bit128.All().AndCount(bit128.Value(5)))
	expect.Eq(t, 1, bit128.Value(5).AndCount(bit128.Value(5)))
	expect.Eq(t, 1, bit128.Value(5).AndCount(bit128.Values(3, 5)))
	expect.Eq(t, 2, bit128.Values(3, 5).AndCount(bit128.Values(3, 5)))
	expect.Eq(t, 1, bit128.Values(3, 5).AndCount(bit128.Less(5)))
	expect.Eq(t, 5, bit128.Less(5).AndCount(bit128.Less(5)))
	expect.Eq(t, 0, bit128.Less(5).AndCount(bit128.More(5)))
	expect.Eq(t, bit128.Cap-6, bit128.More(5).AndCount(bit128.More(5)))
	expect.Eq(t, 0, bit128.More(5).AndCount(bit128.None()))
}

func TestAdd(t *testing.T) {
	x := bit128.None()
	x.Add(3)
	expect.Set(t, bit128.Value(3), x)
	x.Add(3)
	expect.Set(t, bit128.Value(3), x)
	x.Add(5)
	expect.Set(t, bit128.Values(3, 5), x)
}

func TestRemove(t *testing.T) {
	x := bit128.Values(3, 5)
	x.Remove(1)
	expect.Set(t, bit128.Values(3, 5), x)
	x.Remove(3)
	expect.Set(t, bit128.Value(5), x)
	x.Remove(5)
	expect.Set(t, bit128.None(), x)
}

func TestAssign(t *testing.T) {
	x := bit128.None()
	x.Assign(bit128.Value(5))
	expect.Set(t, bit128.Value(5), x)
}

func TestAssignNone(t *testing.T) {
	x := bit128.Values(3, 5)
	x.AssignNone()
	expect.Set(t, bit128.None(), x)
}

func TestAssignAll(t *testing.T) {
	x := bit128.Values(3, 5)
	x.AssignAll()
	expect.Set(t, bit128.All(), x)
}

func TestEach(t *testing.T) {
	expect.Ints(t, bit128.None().Each())
	expect.Ints(t, bit128.Value(5).Each(), 5)
	expect.Ints(t, bit128.Values(3, 5).Each(), 3, 5)
	expect.Ints(t, bit128.Less(5).Each(), 0, 1, 2, 3, 4)
}

func TestNext(t *testing.T) {
	collect := func(s bit128.Set) []int {
		var got []int
		for v, ok := s.Next(0); ok; v, ok = s.Next(v + 1) {
			got = append(got, v)
		}
		return got
	}
	check := func(t *testing.T, s bit128.Set, want []int) {
		t.Helper()
		if got := collect(s); !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	}
	t.Run("none", func(t *testing.T) { check(t, bit128.None(), nil) })
	t.Run("single", func(t *testing.T) { check(t, bit128.Value(5), []int{5}) })
	t.Run("multiple", func(t *testing.T) { check(t, bit128.Values(3, 5), []int{3, 5}) })
	t.Run("contiguous", func(t *testing.T) { check(t, bit128.Less(5), []int{0, 1, 2, 3, 4}) })
	t.Run("last value", func(t *testing.T) { check(t, bit128.Value(127), []int{127}) })
	t.Run("spans a page boundary", func(t *testing.T) { check(t, bit128.Values(63, 64, 65), []int{63, 64, 65}) })
	t.Run("matches Each", func(t *testing.T) {
		s := bit128.Values(1, 2, 5, 127)
		var want []int
		for v := range s.Each() {
			want = append(want, v)
		}
		check(t, s, want)
	})
	t.Run("negative from clamps to 0", func(t *testing.T) {
		s := bit128.Value(5)
		v, ok := s.Next(-3)
		if !ok || v != 5 {
			t.Fatalf("Next(-3) = (%v, %v); want (5, true)", v, ok)
		}
	})
}
func TestRanges(t *testing.T) {
	t.Run("none", func(t *testing.T) {
		var got [][2]int
		for left, right := range bit128.None().Ranges() {
			got = append(got, [2]int{left, right})
		}
		if !reflect.DeepEqual(got, [][2]int(nil)) {
			t.Fatalf("got %v; want %v", got, [][2]int(nil))
		}
	})

	t.Run("single", func(t *testing.T) {
		var got [][2]int
		for left, right := range bit128.Value(5).Ranges() {
			got = append(got, [2]int{left, right})
		}
		want := [][2]int{{5, 5}}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})

	t.Run("contiguous", func(t *testing.T) {
		var got [][2]int
		for left, right := range bit128.Less(5).Ranges() {
			got = append(got, [2]int{left, right})
		}
		want := [][2]int{{0, 4}}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})

	t.Run("disjoint", func(t *testing.T) {
		var got [][2]int
		for left, right := range bit128.Values(0, 2, 3, 6).Ranges() {
			got = append(got, [2]int{left, right})
		}
		want := [][2]int{{0, 0}, {2, 3}, {6, 6}}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("all", func(t *testing.T) {
		var got [][2]int
		for left, right := range bit128.All().Ranges() {
			got = append(got, [2]int{left, right})
		}
		want := [][2]int{{0, bit128.Max}}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
}

func TestRangesEarlyStop(t *testing.T) {
	set := bit128.Values(0, 2, 3, 6)
	var got [][2]int
	for left, right := range set.Ranges() {
		got = append(got, [2]int{left, right})
		break
	}
	want := [][2]int{{0, 0}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v; want %v", got, want)
	}
}

func TestNot(t *testing.T) {
	expect.Set(t, bit128.All(), bit128.None().Not())
	expect.Set(t, bit128.None(), bit128.All().Not())
	expect.Set(t, bit128.Less(5).Or(bit128.More(5)), bit128.Value(5).Not())
	expect.Set(t, bit128.Less(3).Or(bit128.Value(4).Or(bit128.More(5))), bit128.Values(3, 5).Not())
	expect.Set(t, bit128.More(4), bit128.Less(5).Not())
	expect.Set(t, bit128.Less(6), bit128.More(5).Not())
}

func TestAssignNot(t *testing.T) {
	x := bit128.Values(3, 5)
	x.AssignNot()
	expect.Set(t, bit128.Less(3).Or(bit128.Value(4)).Or(bit128.More(5)), x)
}

func TestAssignSub(t *testing.T) {
	x := bit128.Values(3, 5)
	x.AssignSub(bit128.Values(3))
	expect.Set(t, bit128.Values(5), x)
}

func TestAssignAnd(t *testing.T) {
	x := bit128.Values(3, 5)
	x.AssignAnd(bit128.Values(3))
	expect.Set(t, bit128.Values(3), x)
}

func TestAssignOr(t *testing.T) {
	x := bit128.Values(3, 5)
	x.AssignOr(bit128.Values(4))
	expect.Set(t, bit128.Values(3, 4, 5), x)
}

func TestAssignXor(t *testing.T) {
	x := bit128.Values(3, 5)
	x.AssignXor(bit128.Values(3, 4))
	expect.Set(t, bit128.Values(4, 5), x)
}

func TestAssignNor(t *testing.T) {
	x := bit128.Values(3, 5)
	x.AssignNor(bit128.Values(3, 4))
	expect.Set(t, bit128.Less(3).Or(bit128.More(5)), x)
}

func TestAssignIff(t *testing.T) {
	x := bit128.Values(3, 5)
	x.AssignIff(bit128.Values(3, 4))
	expect.Set(t, bit128.Less(4).Or(bit128.More(5)), x)
}

func TestAssignImply(t *testing.T) {
	x := bit128.Values(3, 5)
	x.AssignImply(bit128.Values(3, 4))
	expect.Set(t, bit128.Less(5).Or(bit128.More(5)), x)
}

func TestSub(t *testing.T) {
	expect.Set(t, bit128.None(), bit128.None().Sub(bit128.None()))
	expect.Set(t, bit128.None(), bit128.None().Sub(bit128.All()))
	expect.Set(t, bit128.None(), bit128.All().Sub(bit128.All()))
	expect.Set(t, bit128.Less(5).Or(bit128.More(5)), bit128.All().Sub(bit128.Value(5)))
	expect.Set(t, bit128.None(), bit128.Value(5).Sub(bit128.Value(5)))
	expect.Set(t, bit128.None(), bit128.Value(5).Sub(bit128.Values(3, 5)))
	expect.Set(t, bit128.None(), bit128.Values(3, 5).Sub(bit128.Values(3, 5)))
	expect.Set(t, bit128.Value(5), bit128.Values(3, 5).Sub(bit128.Less(5)))
	expect.Set(t, bit128.None(), bit128.Less(5).Sub(bit128.Less(5)))
	expect.Set(t, bit128.Less(5), bit128.Less(5).Sub(bit128.More(5)))
	expect.Set(t, bit128.None(), bit128.More(5).Sub(bit128.More(5)))
	expect.Set(t, bit128.More(5), bit128.More(5).Sub(bit128.None()))
}

func TestAnd(t *testing.T) {
	expect.Set(t, bit128.None(), bit128.None().And(bit128.None()))
	expect.Set(t, bit128.None(), bit128.None().And(bit128.All()))
	expect.Set(t, bit128.All(), bit128.All().And(bit128.All()))
	expect.Set(t, bit128.Value(5), bit128.All().And(bit128.Value(5)))
	expect.Set(t, bit128.Value(5), bit128.Value(5).And(bit128.Value(5)))
	expect.Set(t, bit128.Value(5), bit128.Value(5).And(bit128.Values(3, 5)))
	expect.Set(t, bit128.Values(3, 5), bit128.Values(3, 5).And(bit128.Values(3, 5)))
	expect.Set(t, bit128.Value(3), bit128.Values(3, 5).And(bit128.Less(5)))
	expect.Set(t, bit128.Less(5), bit128.Less(5).And(bit128.Less(5)))
	expect.Set(t, bit128.None(), bit128.Less(5).And(bit128.More(5)))
	expect.Set(t, bit128.More(5), bit128.More(5).And(bit128.More(5)))
	expect.Set(t, bit128.None(), bit128.More(5).And(bit128.None()))
}

func TestOr(t *testing.T) {
	expect.Set(t, bit128.None(), bit128.None().Or(bit128.None()))
	expect.Set(t, bit128.All(), bit128.None().Or(bit128.All()))
	expect.Set(t, bit128.All(), bit128.All().Or(bit128.All()))
	expect.Set(t, bit128.All(), bit128.All().Or(bit128.Value(5)))
	expect.Set(t, bit128.Value(5), bit128.Value(5).Or(bit128.Value(5)))
	expect.Set(t, bit128.Values(5, 3), bit128.Value(5).Or(bit128.Values(3, 5)))
	expect.Set(t, bit128.Values(3, 5), bit128.Values(3, 5).Or(bit128.Values(3, 5)))
	expect.Set(t, bit128.Less(6), bit128.Values(3, 5).Or(bit128.Less(5)))
	expect.Set(t, bit128.Less(5), bit128.Less(5).Or(bit128.Less(5)))
	expect.Set(t, bit128.Value(5).Not(), bit128.Less(5).Or(bit128.More(5)))
	expect.Set(t, bit128.More(5), bit128.More(5).Or(bit128.More(5)))
	expect.Set(t, bit128.More(5), bit128.More(5).Or(bit128.None()))
}

func TestXor(t *testing.T) {
	expect.Set(t, bit128.None(), bit128.None().Xor(bit128.None()))
	expect.Set(t, bit128.All(), bit128.None().Xor(bit128.All()))
	expect.Set(t, bit128.None(), bit128.All().Xor(bit128.All()))
	expect.Set(t, bit128.Value(5).Not(), bit128.All().Xor(bit128.Value(5)))
	expect.Set(t, bit128.None(), bit128.Value(5).Xor(bit128.Value(5)))
	expect.Set(t, bit128.Value(3), bit128.Value(5).Xor(bit128.Values(3, 5)))
	expect.Set(t, bit128.None(), bit128.Values(3, 5).Xor(bit128.Values(3, 5)))
	expect.Set(t, bit128.Values(0, 1, 2, 4, 5), bit128.Values(3, 5).Xor(bit128.Less(5)))
	expect.Set(t, bit128.None(), bit128.Less(5).Xor(bit128.Less(5)))
	expect.Set(t, bit128.Value(5).Not(), bit128.Less(5).Xor(bit128.More(5)))
	expect.Set(t, bit128.None(), bit128.More(5).Xor(bit128.More(5)))
	expect.Set(t, bit128.More(5), bit128.More(5).Xor(bit128.None()))
}

func TestNor(t *testing.T) {
	expect.Set(t, bit128.All(), bit128.None().Nor(bit128.None()))
	expect.Set(t, bit128.None(), bit128.None().Nor(bit128.All()))
	expect.Set(t, bit128.None(), bit128.All().Nor(bit128.All()))
	expect.Set(t, bit128.None(), bit128.All().Nor(bit128.Value(5)))
	expect.Set(t, bit128.Less(5).Or(bit128.More(5)), bit128.Value(5).Nor(bit128.Value(5)))
	expect.Set(t, bit128.Values(0, 1, 2, 4).Or(bit128.More(5)), bit128.Value(5).Nor(bit128.Values(3, 5)))
	expect.Set(t, bit128.Values(0, 1, 2, 4).Or(bit128.More(5)), bit128.Values(3, 5).Nor(bit128.Values(3, 5)))
	expect.Set(t, bit128.More(5), bit128.Values(3, 5).Nor(bit128.Less(5)))
	expect.Set(t, bit128.More(4), bit128.Less(5).Nor(bit128.Less(5)))
	expect.Set(t, bit128.Value(5), bit128.Less(5).Nor(bit128.More(5)))
	expect.Set(t, bit128.Less(6), bit128.More(5).Nor(bit128.More(5)))
	expect.Set(t, bit128.Less(6), bit128.More(5).Nor(bit128.None()))
}

func TestIff(t *testing.T) {
	expect.Set(t, bit128.All(), bit128.None().Iff(bit128.None()))
	expect.Set(t, bit128.None(), bit128.None().Iff(bit128.All()))
	expect.Set(t, bit128.All(), bit128.All().Iff(bit128.All()))
	expect.Set(t, bit128.Value(5), bit128.All().Iff(bit128.Value(5)))
	expect.Set(t, bit128.All(), bit128.Value(5).Iff(bit128.Value(5)))
	expect.Set(t, bit128.Values(0, 1, 2).Or(bit128.More(3)), bit128.Value(5).Iff(bit128.Values(3, 5)))
	expect.Set(t, bit128.All(), bit128.Values(3, 5).Iff(bit128.Values(3, 5)))
	expect.Set(t, bit128.Value(3).Or(bit128.More(5)), bit128.Values(3, 5).Iff(bit128.Less(5)))
	expect.Set(t, bit128.All(), bit128.Less(5).Iff(bit128.Less(5)))
	expect.Set(t, bit128.Value(5), bit128.Less(5).Iff(bit128.More(5)))
	expect.Set(t, bit128.All(), bit128.More(5).Iff(bit128.More(5)))
	expect.Set(t, bit128.Less(6), bit128.More(5).Iff(bit128.None()))
}

func TestImply(t *testing.T) {
	expect.Set(t, bit128.All(), bit128.None().Imply(bit128.None()))
	expect.Set(t, bit128.All(), bit128.None().Imply(bit128.All()))
	expect.Set(t, bit128.All(), bit128.All().Imply(bit128.All()))
	expect.Set(t, bit128.Value(5), bit128.All().Imply(bit128.Value(5)))
	expect.Set(t, bit128.All(), bit128.Value(5).Imply(bit128.Value(5)))
	expect.Set(t, bit128.All(), bit128.Value(5).Imply(bit128.Values(3, 5)))
	expect.Set(t, bit128.All(), bit128.Values(3, 5).Imply(bit128.Values(3, 5)))
	expect.Set(t, bit128.Value(5).Not(), bit128.Values(3, 5).Imply(bit128.Less(5)))
	expect.Set(t, bit128.All(), bit128.Less(5).Imply(bit128.Less(5)))
	expect.Set(t, bit128.More(4), bit128.Less(5).Imply(bit128.More(5)))
	expect.Set(t, bit128.All(), bit128.More(5).Imply(bit128.More(5)))
	expect.Set(t, bit128.Less(6), bit128.More(5).Imply(bit128.None()))
}
