package bit256_test

import (
	"github.com/mdelah/bitsets/bit256"
	"github.com/mdelah/bitsets/internal/expect"
	"reflect"
	"testing"
)

func TestCount(t *testing.T) {
	expect.Eq(t, 0, bit256.None().Count())
	expect.Eq(t, bit256.Cap, bit256.All().Count())
	expect.Eq(t, 1, bit256.Value(5).Count())
	expect.Eq(t, 2, bit256.Values(3, 5).Count())
	expect.Eq(t, 5, bit256.Less(5).Count())
	expect.Eq(t, bit256.Max-5, bit256.More(5).Count())
}

func TestString(t *testing.T) {
	expect.Eq(t, "{}", bit256.None().String())
	expect.Eq(t, "{0-255}", bit256.All().String())
	expect.Eq(t, "{5}", bit256.Value(5).String())
	expect.Eq(t, "{3,5}", bit256.Values(3, 5).String())
	expect.Eq(t, "{0-4}", bit256.Less(5).String())
	expect.Eq(t, "{6-255}", bit256.More(5).String())
}

func TestIsNone(t *testing.T) {
	expect.Eq(t, true, bit256.None().IsNone())
	expect.Eq(t, false, bit256.All().IsNone())
	expect.Eq(t, false, bit256.Value(5).IsNone())
	expect.Eq(t, false, bit256.Values(3, 5).IsNone())
	expect.Eq(t, false, bit256.Less(5).IsNone())
	expect.Eq(t, false, bit256.More(5).IsNone())
}

func TestIsAll(t *testing.T) {
	expect.Eq(t, false, bit256.None().IsAll())
	expect.Eq(t, true, bit256.All().IsAll())
	expect.Eq(t, false, bit256.Value(5).IsAll())
	expect.Eq(t, false, bit256.Values(3, 5).IsAll())
	expect.Eq(t, false, bit256.Less(5).IsAll())
	expect.Eq(t, false, bit256.More(5).IsAll())
}

func TestMin(t *testing.T) {
	expect.Eq(t, -1, bit256.None().Min())
	expect.Eq(t, 0, bit256.All().Min())
	expect.Eq(t, 5, bit256.Value(5).Min())
	expect.Eq(t, 3, bit256.Values(3, 5).Min())
	expect.Eq(t, 0, bit256.Less(5).Min())
	expect.Eq(t, 6, bit256.More(5).Min())
}

func TestMax(t *testing.T) {
	expect.Eq(t, -1, bit256.None().Max())
	expect.Eq(t, bit256.Max, bit256.All().Max())
	expect.Eq(t, 5, bit256.Value(5).Max())
	expect.Eq(t, 5, bit256.Values(3, 5).Max())
	expect.Eq(t, 4, bit256.Less(5).Max())
	expect.Eq(t, bit256.Max, bit256.More(5).Max())
}

func TestHas(t *testing.T) {
	expect.Eq(t, false, bit256.None().Has(3))
	expect.Eq(t, true, bit256.All().Has(3))
	expect.Eq(t, false, bit256.Value(5).Has(3))
	expect.Eq(t, true, bit256.Values(3, 5).Has(3))
	expect.Eq(t, true, bit256.Less(5).Has(3))
	expect.Eq(t, false, bit256.More(5).Has(3))
}

func TestLessCount(t *testing.T) {
	expect.Eq(t, 0, bit256.None().LessCount(5))
	expect.Eq(t, 5, bit256.All().LessCount(5))
	expect.Eq(t, 0, bit256.Value(5).LessCount(5))
	expect.Eq(t, 1, bit256.Values(3, 5).LessCount(5))
	expect.Eq(t, 5, bit256.Less(5).LessCount(5))
	expect.Eq(t, 0, bit256.More(5).LessCount(5))
}

func TestMoreCount(t *testing.T) {
	expect.Eq(t, 0, bit256.None().MoreCount(4))
	expect.Eq(t, bit256.Cap-5, bit256.All().MoreCount(4))
	expect.Eq(t, 1, bit256.Value(5).MoreCount(4))
	expect.Eq(t, 1, bit256.Values(3, 5).MoreCount(4))
	expect.Eq(t, 0, bit256.Less(5).MoreCount(4))
	expect.Eq(t, bit256.Cap-6, bit256.More(5).MoreCount(4))
}

func TestEqual(t *testing.T) {
	expect.Eq(t, true, bit256.None().Equal(bit256.None()))
	expect.Eq(t, false, bit256.None().Equal(bit256.All()))
	expect.Eq(t, true, bit256.All().Equal(bit256.All()))
	expect.Eq(t, false, bit256.All().Equal(bit256.Value(5)))
	expect.Eq(t, true, bit256.Value(5).Equal(bit256.Value(5)))
	expect.Eq(t, false, bit256.Value(5).Equal(bit256.Values(3, 5)))
	expect.Eq(t, true, bit256.Values(3, 5).Equal(bit256.Values(3, 5)))
	expect.Eq(t, false, bit256.Values(3, 5).Equal(bit256.Less(5)))
	expect.Eq(t, true, bit256.Less(5).Equal(bit256.Less(5)))
	expect.Eq(t, false, bit256.Less(5).Equal(bit256.More(5)))
	expect.Eq(t, true, bit256.More(5).Equal(bit256.More(5)))
	expect.Eq(t, false, bit256.More(5).Equal(bit256.None()))
}

func TestHasNone(t *testing.T) {
	expect.Eq(t, true, bit256.None().HasNone(bit256.None()))
	expect.Eq(t, true, bit256.None().HasNone(bit256.All()))
	expect.Eq(t, false, bit256.All().HasNone(bit256.All()))
	expect.Eq(t, false, bit256.All().HasNone(bit256.Value(5)))
	expect.Eq(t, false, bit256.Value(5).HasNone(bit256.Value(5)))
	expect.Eq(t, false, bit256.Value(5).HasNone(bit256.Values(3, 5)))
	expect.Eq(t, false, bit256.Values(3, 5).HasNone(bit256.Values(3, 5)))
	expect.Eq(t, false, bit256.Values(3, 5).HasNone(bit256.Less(5)))
	expect.Eq(t, false, bit256.Less(5).HasNone(bit256.Less(5)))
	expect.Eq(t, true, bit256.Less(5).HasNone(bit256.More(5)))
	expect.Eq(t, false, bit256.More(5).HasNone(bit256.More(5)))
	expect.Eq(t, true, bit256.More(5).HasNone(bit256.None()))
}

func TestHasAll(t *testing.T) {
	expect.Eq(t, true, bit256.None().HasAll(bit256.None()))
	expect.Eq(t, false, bit256.None().HasAll(bit256.All()))
	expect.Eq(t, true, bit256.All().HasAll(bit256.All()))
	expect.Eq(t, true, bit256.All().HasAll(bit256.Value(5)))
	expect.Eq(t, true, bit256.Value(5).HasAll(bit256.Value(5)))
	expect.Eq(t, false, bit256.Value(5).HasAll(bit256.Values(3, 5)))
	expect.Eq(t, true, bit256.Values(3, 5).HasAll(bit256.Values(3, 5)))
	expect.Eq(t, false, bit256.Values(3, 5).HasAll(bit256.Less(5)))
	expect.Eq(t, true, bit256.Less(5).HasAll(bit256.Less(5)))
	expect.Eq(t, false, bit256.Less(5).HasAll(bit256.More(5)))
	expect.Eq(t, true, bit256.More(5).HasAll(bit256.More(5)))
	expect.Eq(t, true, bit256.More(5).HasAll(bit256.None()))
}

func TestCompare(t *testing.T) {
	expect.Eq(t, 0, bit256.None().Compare(bit256.None()))
	expect.Eq(t, -1, bit256.None().Compare(bit256.All()))
	expect.Eq(t, 0, bit256.All().Compare(bit256.All()))
	expect.Eq(t, 1, bit256.All().Compare(bit256.Value(5)))
	expect.Eq(t, 0, bit256.Value(5).Compare(bit256.Value(5)))
	expect.Eq(t, -1, bit256.Value(5).Compare(bit256.Values(3, 5)))
	expect.Eq(t, 0, bit256.Values(3, 5).Compare(bit256.Values(3, 5)))
	expect.Eq(t, -1, bit256.Values(3, 5).Compare(bit256.Less(5)))
	expect.Eq(t, 0, bit256.Less(5).Compare(bit256.Less(5)))
	expect.Eq(t, 1, bit256.Less(5).Compare(bit256.More(5)))
	expect.Eq(t, 0, bit256.More(5).Compare(bit256.More(5)))
	expect.Eq(t, 1, bit256.More(5).Compare(bit256.None()))
}

func TestAndCount(t *testing.T) {
	expect.Eq(t, 0, bit256.None().AndCount(bit256.None()))
	expect.Eq(t, 0, bit256.None().AndCount(bit256.All()))
	expect.Eq(t, bit256.Cap, bit256.All().AndCount(bit256.All()))
	expect.Eq(t, 1, bit256.All().AndCount(bit256.Value(5)))
	expect.Eq(t, 1, bit256.Value(5).AndCount(bit256.Value(5)))
	expect.Eq(t, 1, bit256.Value(5).AndCount(bit256.Values(3, 5)))
	expect.Eq(t, 2, bit256.Values(3, 5).AndCount(bit256.Values(3, 5)))
	expect.Eq(t, 1, bit256.Values(3, 5).AndCount(bit256.Less(5)))
	expect.Eq(t, 5, bit256.Less(5).AndCount(bit256.Less(5)))
	expect.Eq(t, 0, bit256.Less(5).AndCount(bit256.More(5)))
	expect.Eq(t, bit256.Cap-6, bit256.More(5).AndCount(bit256.More(5)))
	expect.Eq(t, 0, bit256.More(5).AndCount(bit256.None()))
}

func TestAdd(t *testing.T) {
	x := bit256.None()
	x.Add(3)
	expect.Set(t, bit256.Value(3), x)
	x.Add(3)
	expect.Set(t, bit256.Value(3), x)
	x.Add(5)
	expect.Set(t, bit256.Values(3, 5), x)
}

func TestRemove(t *testing.T) {
	x := bit256.Values(3, 5)
	x.Remove(1)
	expect.Set(t, bit256.Values(3, 5), x)
	x.Remove(3)
	expect.Set(t, bit256.Value(5), x)
	x.Remove(5)
	expect.Set(t, bit256.None(), x)
}

func TestAssign(t *testing.T) {
	x := bit256.None()
	x.Assign(bit256.Value(5))
	expect.Set(t, bit256.Value(5), x)
}

func TestAssignNone(t *testing.T) {
	x := bit256.Values(3, 5)
	x.AssignNone()
	expect.Set(t, bit256.None(), x)
}

func TestAssignAll(t *testing.T) {
	x := bit256.Values(3, 5)
	x.AssignAll()
	expect.Set(t, bit256.All(), x)
}

func TestEach(t *testing.T) {
	expect.Ints(t, bit256.None().Each())
	expect.Ints(t, bit256.Value(5).Each(), 5)
	expect.Ints(t, bit256.Values(3, 5).Each(), 3, 5)
	expect.Ints(t, bit256.Less(5).Each(), 0, 1, 2, 3, 4)
}

func TestNext(t *testing.T) {
	collect := func(s bit256.Set) []int {
		var got []int
		for v, ok := s.Next(0); ok; v, ok = s.Next(v + 1) {
			got = append(got, v)
		}
		return got
	}
	check := func(t *testing.T, s bit256.Set, want []int) {
		t.Helper()
		if got := collect(s); !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	}
	t.Run("none", func(t *testing.T) { check(t, bit256.None(), nil) })
	t.Run("single", func(t *testing.T) { check(t, bit256.Value(5), []int{5}) })
	t.Run("multiple", func(t *testing.T) { check(t, bit256.Values(3, 5), []int{3, 5}) })
	t.Run("contiguous", func(t *testing.T) { check(t, bit256.Less(5), []int{0, 1, 2, 3, 4}) })
	t.Run("last value", func(t *testing.T) { check(t, bit256.Value(255), []int{255}) })
	t.Run("spans a page boundary", func(t *testing.T) { check(t, bit256.Values(63, 64, 65), []int{63, 64, 65}) })
	t.Run("matches Each", func(t *testing.T) {
		s := bit256.Values(1, 2, 5, 255)
		var want []int
		for v := range s.Each() {
			want = append(want, v)
		}
		check(t, s, want)
	})
	t.Run("negative from clamps to 0", func(t *testing.T) {
		s := bit256.Value(5)
		v, ok := s.Next(-3)
		if !ok || v != 5 {
			t.Fatalf("Next(-3) = (%v, %v); want (5, true)", v, ok)
		}
	})
}
func TestRanges(t *testing.T) {
	t.Run("none", func(t *testing.T) {
		var got [][2]int
		for left, right := range bit256.None().Ranges() {
			got = append(got, [2]int{left, right})
		}
		if !reflect.DeepEqual(got, [][2]int(nil)) {
			t.Fatalf("got %v; want %v", got, [][2]int(nil))
		}
	})

	t.Run("single", func(t *testing.T) {
		var got [][2]int
		for left, right := range bit256.Value(5).Ranges() {
			got = append(got, [2]int{left, right})
		}
		want := [][2]int{{5, 5}}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})

	t.Run("contiguous", func(t *testing.T) {
		var got [][2]int
		for left, right := range bit256.Less(5).Ranges() {
			got = append(got, [2]int{left, right})
		}
		want := [][2]int{{0, 4}}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})

	t.Run("disjoint", func(t *testing.T) {
		var got [][2]int
		for left, right := range bit256.Values(0, 2, 3, 6).Ranges() {
			got = append(got, [2]int{left, right})
		}
		want := [][2]int{{0, 0}, {2, 3}, {6, 6}}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("all", func(t *testing.T) {
		var got [][2]int
		for left, right := range bit256.All().Ranges() {
			got = append(got, [2]int{left, right})
		}
		want := [][2]int{{0, bit256.Max}}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
}

func TestRangesEarlyStop(t *testing.T) {
	set := bit256.Values(0, 2, 3, 6)
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
	expect.Set(t, bit256.All(), bit256.None().Not())
	expect.Set(t, bit256.None(), bit256.All().Not())
	expect.Set(t, bit256.Less(5).Or(bit256.More(5)), bit256.Value(5).Not())
	expect.Set(t, bit256.Less(3).Or(bit256.Value(4).Or(bit256.More(5))), bit256.Values(3, 5).Not())
	expect.Set(t, bit256.More(4), bit256.Less(5).Not())
	expect.Set(t, bit256.Less(6), bit256.More(5).Not())
}

func TestAssignNot(t *testing.T) {
	x := bit256.Values(3, 5)
	x.AssignNot()
	expect.Set(t, bit256.Less(3).Or(bit256.Value(4)).Or(bit256.More(5)), x)
}

func TestAssignSub(t *testing.T) {
	x := bit256.Values(3, 5)
	x.AssignSub(bit256.Values(3))
	expect.Set(t, bit256.Values(5), x)
}

func TestAssignAnd(t *testing.T) {
	x := bit256.Values(3, 5)
	x.AssignAnd(bit256.Values(3))
	expect.Set(t, bit256.Values(3), x)
}

func TestAssignOr(t *testing.T) {
	x := bit256.Values(3, 5)
	x.AssignOr(bit256.Values(4))
	expect.Set(t, bit256.Values(3, 4, 5), x)
}

func TestAssignXor(t *testing.T) {
	x := bit256.Values(3, 5)
	x.AssignXor(bit256.Values(3, 4))
	expect.Set(t, bit256.Values(4, 5), x)
}

func TestAssignNor(t *testing.T) {
	x := bit256.Values(3, 5)
	x.AssignNor(bit256.Values(3, 4))
	expect.Set(t, bit256.Less(3).Or(bit256.More(5)), x)
}

func TestAssignIff(t *testing.T) {
	x := bit256.Values(3, 5)
	x.AssignIff(bit256.Values(3, 4))
	expect.Set(t, bit256.Less(4).Or(bit256.More(5)), x)
}

func TestAssignImply(t *testing.T) {
	x := bit256.Values(3, 5)
	x.AssignImply(bit256.Values(3, 4))
	expect.Set(t, bit256.Less(5).Or(bit256.More(5)), x)
}

func TestSub(t *testing.T) {
	expect.Set(t, bit256.None(), bit256.None().Sub(bit256.None()))
	expect.Set(t, bit256.None(), bit256.None().Sub(bit256.All()))
	expect.Set(t, bit256.None(), bit256.All().Sub(bit256.All()))
	expect.Set(t, bit256.Less(5).Or(bit256.More(5)), bit256.All().Sub(bit256.Value(5)))
	expect.Set(t, bit256.None(), bit256.Value(5).Sub(bit256.Value(5)))
	expect.Set(t, bit256.None(), bit256.Value(5).Sub(bit256.Values(3, 5)))
	expect.Set(t, bit256.None(), bit256.Values(3, 5).Sub(bit256.Values(3, 5)))
	expect.Set(t, bit256.Value(5), bit256.Values(3, 5).Sub(bit256.Less(5)))
	expect.Set(t, bit256.None(), bit256.Less(5).Sub(bit256.Less(5)))
	expect.Set(t, bit256.Less(5), bit256.Less(5).Sub(bit256.More(5)))
	expect.Set(t, bit256.None(), bit256.More(5).Sub(bit256.More(5)))
	expect.Set(t, bit256.More(5), bit256.More(5).Sub(bit256.None()))
}

func TestAnd(t *testing.T) {
	expect.Set(t, bit256.None(), bit256.None().And(bit256.None()))
	expect.Set(t, bit256.None(), bit256.None().And(bit256.All()))
	expect.Set(t, bit256.All(), bit256.All().And(bit256.All()))
	expect.Set(t, bit256.Value(5), bit256.All().And(bit256.Value(5)))
	expect.Set(t, bit256.Value(5), bit256.Value(5).And(bit256.Value(5)))
	expect.Set(t, bit256.Value(5), bit256.Value(5).And(bit256.Values(3, 5)))
	expect.Set(t, bit256.Values(3, 5), bit256.Values(3, 5).And(bit256.Values(3, 5)))
	expect.Set(t, bit256.Value(3), bit256.Values(3, 5).And(bit256.Less(5)))
	expect.Set(t, bit256.Less(5), bit256.Less(5).And(bit256.Less(5)))
	expect.Set(t, bit256.None(), bit256.Less(5).And(bit256.More(5)))
	expect.Set(t, bit256.More(5), bit256.More(5).And(bit256.More(5)))
	expect.Set(t, bit256.None(), bit256.More(5).And(bit256.None()))
}

func TestOr(t *testing.T) {
	expect.Set(t, bit256.None(), bit256.None().Or(bit256.None()))
	expect.Set(t, bit256.All(), bit256.None().Or(bit256.All()))
	expect.Set(t, bit256.All(), bit256.All().Or(bit256.All()))
	expect.Set(t, bit256.All(), bit256.All().Or(bit256.Value(5)))
	expect.Set(t, bit256.Value(5), bit256.Value(5).Or(bit256.Value(5)))
	expect.Set(t, bit256.Values(5, 3), bit256.Value(5).Or(bit256.Values(3, 5)))
	expect.Set(t, bit256.Values(3, 5), bit256.Values(3, 5).Or(bit256.Values(3, 5)))
	expect.Set(t, bit256.Less(6), bit256.Values(3, 5).Or(bit256.Less(5)))
	expect.Set(t, bit256.Less(5), bit256.Less(5).Or(bit256.Less(5)))
	expect.Set(t, bit256.Value(5).Not(), bit256.Less(5).Or(bit256.More(5)))
	expect.Set(t, bit256.More(5), bit256.More(5).Or(bit256.More(5)))
	expect.Set(t, bit256.More(5), bit256.More(5).Or(bit256.None()))
}

func TestXor(t *testing.T) {
	expect.Set(t, bit256.None(), bit256.None().Xor(bit256.None()))
	expect.Set(t, bit256.All(), bit256.None().Xor(bit256.All()))
	expect.Set(t, bit256.None(), bit256.All().Xor(bit256.All()))
	expect.Set(t, bit256.Value(5).Not(), bit256.All().Xor(bit256.Value(5)))
	expect.Set(t, bit256.None(), bit256.Value(5).Xor(bit256.Value(5)))
	expect.Set(t, bit256.Value(3), bit256.Value(5).Xor(bit256.Values(3, 5)))
	expect.Set(t, bit256.None(), bit256.Values(3, 5).Xor(bit256.Values(3, 5)))
	expect.Set(t, bit256.Values(0, 1, 2, 4, 5), bit256.Values(3, 5).Xor(bit256.Less(5)))
	expect.Set(t, bit256.None(), bit256.Less(5).Xor(bit256.Less(5)))
	expect.Set(t, bit256.Value(5).Not(), bit256.Less(5).Xor(bit256.More(5)))
	expect.Set(t, bit256.None(), bit256.More(5).Xor(bit256.More(5)))
	expect.Set(t, bit256.More(5), bit256.More(5).Xor(bit256.None()))
}

func TestNor(t *testing.T) {
	expect.Set(t, bit256.All(), bit256.None().Nor(bit256.None()))
	expect.Set(t, bit256.None(), bit256.None().Nor(bit256.All()))
	expect.Set(t, bit256.None(), bit256.All().Nor(bit256.All()))
	expect.Set(t, bit256.None(), bit256.All().Nor(bit256.Value(5)))
	expect.Set(t, bit256.Less(5).Or(bit256.More(5)), bit256.Value(5).Nor(bit256.Value(5)))
	expect.Set(t, bit256.Values(0, 1, 2, 4).Or(bit256.More(5)), bit256.Value(5).Nor(bit256.Values(3, 5)))
	expect.Set(t, bit256.Values(0, 1, 2, 4).Or(bit256.More(5)), bit256.Values(3, 5).Nor(bit256.Values(3, 5)))
	expect.Set(t, bit256.More(5), bit256.Values(3, 5).Nor(bit256.Less(5)))
	expect.Set(t, bit256.More(4), bit256.Less(5).Nor(bit256.Less(5)))
	expect.Set(t, bit256.Value(5), bit256.Less(5).Nor(bit256.More(5)))
	expect.Set(t, bit256.Less(6), bit256.More(5).Nor(bit256.More(5)))
	expect.Set(t, bit256.Less(6), bit256.More(5).Nor(bit256.None()))
}

func TestIff(t *testing.T) {
	expect.Set(t, bit256.All(), bit256.None().Iff(bit256.None()))
	expect.Set(t, bit256.None(), bit256.None().Iff(bit256.All()))
	expect.Set(t, bit256.All(), bit256.All().Iff(bit256.All()))
	expect.Set(t, bit256.Value(5), bit256.All().Iff(bit256.Value(5)))
	expect.Set(t, bit256.All(), bit256.Value(5).Iff(bit256.Value(5)))
	expect.Set(t, bit256.Values(0, 1, 2).Or(bit256.More(3)), bit256.Value(5).Iff(bit256.Values(3, 5)))
	expect.Set(t, bit256.All(), bit256.Values(3, 5).Iff(bit256.Values(3, 5)))
	expect.Set(t, bit256.Value(3).Or(bit256.More(5)), bit256.Values(3, 5).Iff(bit256.Less(5)))
	expect.Set(t, bit256.All(), bit256.Less(5).Iff(bit256.Less(5)))
	expect.Set(t, bit256.Value(5), bit256.Less(5).Iff(bit256.More(5)))
	expect.Set(t, bit256.All(), bit256.More(5).Iff(bit256.More(5)))
	expect.Set(t, bit256.Less(6), bit256.More(5).Iff(bit256.None()))
}

func TestImply(t *testing.T) {
	expect.Set(t, bit256.All(), bit256.None().Imply(bit256.None()))
	expect.Set(t, bit256.All(), bit256.None().Imply(bit256.All()))
	expect.Set(t, bit256.All(), bit256.All().Imply(bit256.All()))
	expect.Set(t, bit256.Value(5), bit256.All().Imply(bit256.Value(5)))
	expect.Set(t, bit256.All(), bit256.Value(5).Imply(bit256.Value(5)))
	expect.Set(t, bit256.All(), bit256.Value(5).Imply(bit256.Values(3, 5)))
	expect.Set(t, bit256.All(), bit256.Values(3, 5).Imply(bit256.Values(3, 5)))
	expect.Set(t, bit256.Value(5).Not(), bit256.Values(3, 5).Imply(bit256.Less(5)))
	expect.Set(t, bit256.All(), bit256.Less(5).Imply(bit256.Less(5)))
	expect.Set(t, bit256.More(4), bit256.Less(5).Imply(bit256.More(5)))
	expect.Set(t, bit256.All(), bit256.More(5).Imply(bit256.More(5)))
	expect.Set(t, bit256.Less(6), bit256.More(5).Imply(bit256.None()))
}
