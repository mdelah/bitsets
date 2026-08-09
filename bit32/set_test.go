package bit32_test

import (
	"github.com/mdelah/bitsets/bit32"
	"github.com/mdelah/bitsets/internal/expect"
	"reflect"
	"testing"
)

func TestCount(t *testing.T) {
	expect.Eq(t, 0, bit32.None.Count())
	expect.Eq(t, bit32.Cap, bit32.All.Count())
	expect.Eq(t, 1, bit32.Value(5).Count())
	expect.Eq(t, 2, bit32.Values(3, 5).Count())
	expect.Eq(t, 5, bit32.Less(5).Count())
	expect.Eq(t, bit32.Max-5, bit32.More(5).Count())
}

func TestString(t *testing.T) {
	expect.Eq(t, "{}", bit32.None.String())
	expect.Eq(t, "{0-31}", bit32.All.String())
	expect.Eq(t, "{5}", bit32.Value(5).String())
	expect.Eq(t, "{3,5}", bit32.Values(3, 5).String())
	expect.Eq(t, "{0-4}", bit32.Less(5).String())
	expect.Eq(t, "{6-31}", bit32.More(5).String())
}

func TestIsNone(t *testing.T) {
	expect.Eq(t, true, bit32.None.IsNone())
	expect.Eq(t, false, bit32.All.IsNone())
	expect.Eq(t, false, bit32.Value(5).IsNone())
	expect.Eq(t, false, bit32.Values(3, 5).IsNone())
	expect.Eq(t, false, bit32.Less(5).IsNone())
	expect.Eq(t, false, bit32.More(5).IsNone())
}

func TestIsAll(t *testing.T) {
	expect.Eq(t, false, bit32.None.IsAll())
	expect.Eq(t, true, bit32.All.IsAll())
	expect.Eq(t, false, bit32.Value(5).IsAll())
	expect.Eq(t, false, bit32.Values(3, 5).IsAll())
	expect.Eq(t, false, bit32.Less(5).IsAll())
	expect.Eq(t, false, bit32.More(5).IsAll())
}

func TestMin(t *testing.T) {
	expect.Eq(t, -1, bit32.None.Min())
	expect.Eq(t, 0, bit32.All.Min())
	expect.Eq(t, 5, bit32.Value(5).Min())
	expect.Eq(t, 3, bit32.Values(3, 5).Min())
	expect.Eq(t, 0, bit32.Less(5).Min())
	expect.Eq(t, 6, bit32.More(5).Min())
}

func TestMax(t *testing.T) {
	expect.Eq(t, -1, bit32.None.Max())
	expect.Eq(t, bit32.Max, bit32.All.Max())
	expect.Eq(t, 5, bit32.Value(5).Max())
	expect.Eq(t, 5, bit32.Values(3, 5).Max())
	expect.Eq(t, 4, bit32.Less(5).Max())
	expect.Eq(t, bit32.Max, bit32.More(5).Max())
}

func TestHas(t *testing.T) {
	expect.Eq(t, false, bit32.None.Has(3))
	expect.Eq(t, true, bit32.All.Has(3))
	expect.Eq(t, false, bit32.Value(5).Has(3))
	expect.Eq(t, true, bit32.Values(3, 5).Has(3))
	expect.Eq(t, true, bit32.Less(5).Has(3))
	expect.Eq(t, false, bit32.More(5).Has(3))
}

func TestLessCount(t *testing.T) {
	expect.Eq(t, 0, bit32.None.LessCount(5))
	expect.Eq(t, 5, bit32.All.LessCount(5))
	expect.Eq(t, 0, bit32.Value(5).LessCount(5))
	expect.Eq(t, 1, bit32.Values(3, 5).LessCount(5))
	expect.Eq(t, 5, bit32.Less(5).LessCount(5))
	expect.Eq(t, 0, bit32.More(5).LessCount(5))
}

func TestMoreCount(t *testing.T) {
	expect.Eq(t, 0, bit32.None.MoreCount(4))
	expect.Eq(t, bit32.Cap-5, bit32.All.MoreCount(4))
	expect.Eq(t, 1, bit32.Value(5).MoreCount(4))
	expect.Eq(t, 1, bit32.Values(3, 5).MoreCount(4))
	expect.Eq(t, 0, bit32.Less(5).MoreCount(4))
	expect.Eq(t, bit32.Cap-6, bit32.More(5).MoreCount(4))
}

func TestEqual(t *testing.T) {
	expect.Eq(t, true, bit32.None.Equal(bit32.None))
	expect.Eq(t, false, bit32.None.Equal(bit32.All))
	expect.Eq(t, true, bit32.All.Equal(bit32.All))
	expect.Eq(t, false, bit32.All.Equal(bit32.Value(5)))
	expect.Eq(t, true, bit32.Value(5).Equal(bit32.Value(5)))
	expect.Eq(t, false, bit32.Value(5).Equal(bit32.Values(3, 5)))
	expect.Eq(t, true, bit32.Values(3, 5).Equal(bit32.Values(3, 5)))
	expect.Eq(t, false, bit32.Values(3, 5).Equal(bit32.Less(5)))
	expect.Eq(t, true, bit32.Less(5).Equal(bit32.Less(5)))
	expect.Eq(t, false, bit32.Less(5).Equal(bit32.More(5)))
	expect.Eq(t, true, bit32.More(5).Equal(bit32.More(5)))
	expect.Eq(t, false, bit32.More(5).Equal(bit32.None))
}

func TestHasNone(t *testing.T) {
	expect.Eq(t, true, bit32.None.HasNone(bit32.None))
	expect.Eq(t, true, bit32.None.HasNone(bit32.All))
	expect.Eq(t, false, bit32.All.HasNone(bit32.All))
	expect.Eq(t, false, bit32.All.HasNone(bit32.Value(5)))
	expect.Eq(t, false, bit32.Value(5).HasNone(bit32.Value(5)))
	expect.Eq(t, false, bit32.Value(5).HasNone(bit32.Values(3, 5)))
	expect.Eq(t, false, bit32.Values(3, 5).HasNone(bit32.Values(3, 5)))
	expect.Eq(t, false, bit32.Values(3, 5).HasNone(bit32.Less(5)))
	expect.Eq(t, false, bit32.Less(5).HasNone(bit32.Less(5)))
	expect.Eq(t, true, bit32.Less(5).HasNone(bit32.More(5)))
	expect.Eq(t, false, bit32.More(5).HasNone(bit32.More(5)))
	expect.Eq(t, true, bit32.More(5).HasNone(bit32.None))
}

func TestHasAll(t *testing.T) {
	expect.Eq(t, true, bit32.None.HasAll(bit32.None))
	expect.Eq(t, false, bit32.None.HasAll(bit32.All))
	expect.Eq(t, true, bit32.All.HasAll(bit32.All))
	expect.Eq(t, true, bit32.All.HasAll(bit32.Value(5)))
	expect.Eq(t, true, bit32.Value(5).HasAll(bit32.Value(5)))
	expect.Eq(t, false, bit32.Value(5).HasAll(bit32.Values(3, 5)))
	expect.Eq(t, true, bit32.Values(3, 5).HasAll(bit32.Values(3, 5)))
	expect.Eq(t, false, bit32.Values(3, 5).HasAll(bit32.Less(5)))
	expect.Eq(t, true, bit32.Less(5).HasAll(bit32.Less(5)))
	expect.Eq(t, false, bit32.Less(5).HasAll(bit32.More(5)))
	expect.Eq(t, true, bit32.More(5).HasAll(bit32.More(5)))
	expect.Eq(t, true, bit32.More(5).HasAll(bit32.None))
}

func TestCompare(t *testing.T) {
	expect.Eq(t, 0, bit32.None.Compare(bit32.None))
	expect.Eq(t, -1, bit32.None.Compare(bit32.All))
	expect.Eq(t, 0, bit32.All.Compare(bit32.All))
	expect.Eq(t, 1, bit32.All.Compare(bit32.Value(5)))
	expect.Eq(t, 0, bit32.Value(5).Compare(bit32.Value(5)))
	expect.Eq(t, -1, bit32.Value(5).Compare(bit32.Values(3, 5)))
	expect.Eq(t, 0, bit32.Values(3, 5).Compare(bit32.Values(3, 5)))
	expect.Eq(t, -1, bit32.Values(3, 5).Compare(bit32.Less(5)))
	expect.Eq(t, 0, bit32.Less(5).Compare(bit32.Less(5)))
	expect.Eq(t, 1, bit32.Less(5).Compare(bit32.More(5)))
	expect.Eq(t, 0, bit32.More(5).Compare(bit32.More(5)))
	expect.Eq(t, 1, bit32.More(5).Compare(bit32.None))
}

func TestAndCount(t *testing.T) {
	expect.Eq(t, 0, bit32.None.AndCount(bit32.None))
	expect.Eq(t, 0, bit32.None.AndCount(bit32.All))
	expect.Eq(t, bit32.Cap, bit32.All.AndCount(bit32.All))
	expect.Eq(t, 1, bit32.All.AndCount(bit32.Value(5)))
	expect.Eq(t, 1, bit32.Value(5).AndCount(bit32.Value(5)))
	expect.Eq(t, 1, bit32.Value(5).AndCount(bit32.Values(3, 5)))
	expect.Eq(t, 2, bit32.Values(3, 5).AndCount(bit32.Values(3, 5)))
	expect.Eq(t, 1, bit32.Values(3, 5).AndCount(bit32.Less(5)))
	expect.Eq(t, 5, bit32.Less(5).AndCount(bit32.Less(5)))
	expect.Eq(t, 0, bit32.Less(5).AndCount(bit32.More(5)))
	expect.Eq(t, bit32.Cap-6, bit32.More(5).AndCount(bit32.More(5)))
	expect.Eq(t, 0, bit32.More(5).AndCount(bit32.None))
}

func TestAdd(t *testing.T) {
	x := bit32.None
	x.Add(3)
	expect.Set(t, bit32.Value(3), x)
	x.Add(3)
	expect.Set(t, bit32.Value(3), x)
	x.Add(5)
	expect.Set(t, bit32.Values(3, 5), x)
}

func TestRemove(t *testing.T) {
	x := bit32.Values(3, 5)
	x.Remove(1)
	expect.Set(t, bit32.Values(3, 5), x)
	x.Remove(3)
	expect.Set(t, bit32.Value(5), x)
	x.Remove(5)
	expect.Set(t, bit32.None, x)
}

func TestAssign(t *testing.T) {
	x := bit32.None
	x.Assign(bit32.Value(5))
	expect.Set(t, bit32.Value(5), x)
}

func TestAssignNone(t *testing.T) {
	x := bit32.Values(3, 5)
	x.AssignNone()
	expect.Set(t, bit32.None, x)
}

func TestAssignAll(t *testing.T) {
	x := bit32.Values(3, 5)
	x.AssignAll()
	expect.Set(t, bit32.All, x)
}

func TestEach(t *testing.T) {
	expect.Ints(t, bit32.None.Each())
	expect.Ints(t, bit32.Value(5).Each(), 5)
	expect.Ints(t, bit32.Values(3, 5).Each(), 3, 5)
	expect.Ints(t, bit32.Less(5).Each(), 0, 1, 2, 3, 4)
}

func TestNext(t *testing.T) {
	collect := func(s bit32.Set) []int {
		var got []int
		for v, ok := s.Next(0); ok; v, ok = s.Next(v + 1) {
			got = append(got, v)
		}
		return got
	}
	check := func(t *testing.T, s bit32.Set, want []int) {
		t.Helper()
		if got := collect(s); !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	}
	t.Run("none", func(t *testing.T) { check(t, bit32.None, nil) })
	t.Run("single", func(t *testing.T) { check(t, bit32.Value(5), []int{5}) })
	t.Run("multiple", func(t *testing.T) { check(t, bit32.Values(3, 5), []int{3, 5}) })
	t.Run("contiguous", func(t *testing.T) { check(t, bit32.Less(5), []int{0, 1, 2, 3, 4}) })
	t.Run("last value", func(t *testing.T) { check(t, bit32.Value(31), []int{31}) })
	t.Run("matches Each", func(t *testing.T) {
		s := bit32.Values(1, 2, 5, 31)
		var want []int
		for v := range s.Each() {
			want = append(want, v)
		}
		check(t, s, want)
	})
	t.Run("negative from clamps to 0", func(t *testing.T) {
		s := bit32.Value(5)
		v, ok := s.Next(-3)
		if !ok || v != 5 {
			t.Fatalf("Next(-3) = (%v, %v); want (5, true)", v, ok)
		}
	})
}
func TestRanges(t *testing.T) {
	t.Run("none", func(t *testing.T) {
		var got [][2]int
		for left, right := range bit32.None.Ranges() {
			got = append(got, [2]int{left, right})
		}
		if !reflect.DeepEqual(got, [][2]int(nil)) {
			t.Fatalf("got %v; want %v", got, [][2]int(nil))
		}
	})

	t.Run("single", func(t *testing.T) {
		var got [][2]int
		for left, right := range bit32.Value(5).Ranges() {
			got = append(got, [2]int{left, right})
		}
		want := [][2]int{{5, 5}}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})

	t.Run("contiguous", func(t *testing.T) {
		var got [][2]int
		for left, right := range bit32.Less(5).Ranges() {
			got = append(got, [2]int{left, right})
		}
		want := [][2]int{{0, 4}}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})

	t.Run("disjoint", func(t *testing.T) {
		var got [][2]int
		for left, right := range bit32.Values(0, 2, 3, 6).Ranges() {
			got = append(got, [2]int{left, right})
		}
		want := [][2]int{{0, 0}, {2, 3}, {6, 6}}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("all", func(t *testing.T) {
		var got [][2]int
		for left, right := range bit32.All.Ranges() {
			got = append(got, [2]int{left, right})
		}
		want := [][2]int{{0, bit32.Max}}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
}

func TestRangesEarlyStop(t *testing.T) {
	set := bit32.Values(0, 2, 3, 6)
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
	expect.Set(t, bit32.All, bit32.None.Not())
	expect.Set(t, bit32.None, bit32.All.Not())
	expect.Set(t, bit32.Less(5).Or(bit32.More(5)), bit32.Value(5).Not())
	expect.Set(t, bit32.Less(3).Or(bit32.Value(4).Or(bit32.More(5))), bit32.Values(3, 5).Not())
	expect.Set(t, bit32.More(4), bit32.Less(5).Not())
	expect.Set(t, bit32.Less(6), bit32.More(5).Not())
}

func TestAssignNot(t *testing.T) {
	x := bit32.Values(3, 5)
	x.AssignNot()
	expect.Set(t, bit32.Less(3).Or(bit32.Value(4)).Or(bit32.More(5)), x)
}

func TestAssignSub(t *testing.T) {
	x := bit32.Values(3, 5)
	x.AssignSub(bit32.Values(3))
	expect.Set(t, bit32.Values(5), x)
}

func TestAssignAnd(t *testing.T) {
	x := bit32.Values(3, 5)
	x.AssignAnd(bit32.Values(3))
	expect.Set(t, bit32.Values(3), x)
}

func TestAssignOr(t *testing.T) {
	x := bit32.Values(3, 5)
	x.AssignOr(bit32.Values(4))
	expect.Set(t, bit32.Values(3, 4, 5), x)
}

func TestAssignXor(t *testing.T) {
	x := bit32.Values(3, 5)
	x.AssignXor(bit32.Values(3, 4))
	expect.Set(t, bit32.Values(4, 5), x)
}

func TestAssignNor(t *testing.T) {
	x := bit32.Values(3, 5)
	x.AssignNor(bit32.Values(3, 4))
	expect.Set(t, bit32.Less(3).Or(bit32.More(5)), x)
}

func TestAssignIff(t *testing.T) {
	x := bit32.Values(3, 5)
	x.AssignIff(bit32.Values(3, 4))
	expect.Set(t, bit32.Less(4).Or(bit32.More(5)), x)
}

func TestAssignImply(t *testing.T) {
	x := bit32.Values(3, 5)
	x.AssignImply(bit32.Values(3, 4))
	expect.Set(t, bit32.Less(5).Or(bit32.More(5)), x)
}

func TestSub(t *testing.T) {
	expect.Set(t, bit32.None, bit32.None.Sub(bit32.None))
	expect.Set(t, bit32.None, bit32.None.Sub(bit32.All))
	expect.Set(t, bit32.None, bit32.All.Sub(bit32.All))
	expect.Set(t, bit32.Less(5).Or(bit32.More(5)), bit32.All.Sub(bit32.Value(5)))
	expect.Set(t, bit32.None, bit32.Value(5).Sub(bit32.Value(5)))
	expect.Set(t, bit32.None, bit32.Value(5).Sub(bit32.Values(3, 5)))
	expect.Set(t, bit32.None, bit32.Values(3, 5).Sub(bit32.Values(3, 5)))
	expect.Set(t, bit32.Value(5), bit32.Values(3, 5).Sub(bit32.Less(5)))
	expect.Set(t, bit32.None, bit32.Less(5).Sub(bit32.Less(5)))
	expect.Set(t, bit32.Less(5), bit32.Less(5).Sub(bit32.More(5)))
	expect.Set(t, bit32.None, bit32.More(5).Sub(bit32.More(5)))
	expect.Set(t, bit32.More(5), bit32.More(5).Sub(bit32.None))
}

func TestAnd(t *testing.T) {
	expect.Set(t, bit32.None, bit32.None.And(bit32.None))
	expect.Set(t, bit32.None, bit32.None.And(bit32.All))
	expect.Set(t, bit32.All, bit32.All.And(bit32.All))
	expect.Set(t, bit32.Value(5), bit32.All.And(bit32.Value(5)))
	expect.Set(t, bit32.Value(5), bit32.Value(5).And(bit32.Value(5)))
	expect.Set(t, bit32.Value(5), bit32.Value(5).And(bit32.Values(3, 5)))
	expect.Set(t, bit32.Values(3, 5), bit32.Values(3, 5).And(bit32.Values(3, 5)))
	expect.Set(t, bit32.Value(3), bit32.Values(3, 5).And(bit32.Less(5)))
	expect.Set(t, bit32.Less(5), bit32.Less(5).And(bit32.Less(5)))
	expect.Set(t, bit32.None, bit32.Less(5).And(bit32.More(5)))
	expect.Set(t, bit32.More(5), bit32.More(5).And(bit32.More(5)))
	expect.Set(t, bit32.None, bit32.More(5).And(bit32.None))
}

func TestOr(t *testing.T) {
	expect.Set(t, bit32.None, bit32.None.Or(bit32.None))
	expect.Set(t, bit32.All, bit32.None.Or(bit32.All))
	expect.Set(t, bit32.All, bit32.All.Or(bit32.All))
	expect.Set(t, bit32.All, bit32.All.Or(bit32.Value(5)))
	expect.Set(t, bit32.Value(5), bit32.Value(5).Or(bit32.Value(5)))
	expect.Set(t, bit32.Values(5, 3), bit32.Value(5).Or(bit32.Values(3, 5)))
	expect.Set(t, bit32.Values(3, 5), bit32.Values(3, 5).Or(bit32.Values(3, 5)))
	expect.Set(t, bit32.Less(6), bit32.Values(3, 5).Or(bit32.Less(5)))
	expect.Set(t, bit32.Less(5), bit32.Less(5).Or(bit32.Less(5)))
	expect.Set(t, bit32.Value(5).Not(), bit32.Less(5).Or(bit32.More(5)))
	expect.Set(t, bit32.More(5), bit32.More(5).Or(bit32.More(5)))
	expect.Set(t, bit32.More(5), bit32.More(5).Or(bit32.None))
}

func TestXor(t *testing.T) {
	expect.Set(t, bit32.None, bit32.None.Xor(bit32.None))
	expect.Set(t, bit32.All, bit32.None.Xor(bit32.All))
	expect.Set(t, bit32.None, bit32.All.Xor(bit32.All))
	expect.Set(t, bit32.Value(5).Not(), bit32.All.Xor(bit32.Value(5)))
	expect.Set(t, bit32.None, bit32.Value(5).Xor(bit32.Value(5)))
	expect.Set(t, bit32.Value(3), bit32.Value(5).Xor(bit32.Values(3, 5)))
	expect.Set(t, bit32.None, bit32.Values(3, 5).Xor(bit32.Values(3, 5)))
	expect.Set(t, bit32.Values(0, 1, 2, 4, 5), bit32.Values(3, 5).Xor(bit32.Less(5)))
	expect.Set(t, bit32.None, bit32.Less(5).Xor(bit32.Less(5)))
	expect.Set(t, bit32.Value(5).Not(), bit32.Less(5).Xor(bit32.More(5)))
	expect.Set(t, bit32.None, bit32.More(5).Xor(bit32.More(5)))
	expect.Set(t, bit32.More(5), bit32.More(5).Xor(bit32.None))
}

func TestNor(t *testing.T) {
	expect.Set(t, bit32.All, bit32.None.Nor(bit32.None))
	expect.Set(t, bit32.None, bit32.None.Nor(bit32.All))
	expect.Set(t, bit32.None, bit32.All.Nor(bit32.All))
	expect.Set(t, bit32.None, bit32.All.Nor(bit32.Value(5)))
	expect.Set(t, bit32.Less(5).Or(bit32.More(5)), bit32.Value(5).Nor(bit32.Value(5)))
	expect.Set(t, bit32.Values(0, 1, 2, 4).Or(bit32.More(5)), bit32.Value(5).Nor(bit32.Values(3, 5)))
	expect.Set(t, bit32.Values(0, 1, 2, 4).Or(bit32.More(5)), bit32.Values(3, 5).Nor(bit32.Values(3, 5)))
	expect.Set(t, bit32.More(5), bit32.Values(3, 5).Nor(bit32.Less(5)))
	expect.Set(t, bit32.More(4), bit32.Less(5).Nor(bit32.Less(5)))
	expect.Set(t, bit32.Value(5), bit32.Less(5).Nor(bit32.More(5)))
	expect.Set(t, bit32.Less(6), bit32.More(5).Nor(bit32.More(5)))
	expect.Set(t, bit32.Less(6), bit32.More(5).Nor(bit32.None))
}

func TestIff(t *testing.T) {
	expect.Set(t, bit32.All, bit32.None.Iff(bit32.None))
	expect.Set(t, bit32.None, bit32.None.Iff(bit32.All))
	expect.Set(t, bit32.All, bit32.All.Iff(bit32.All))
	expect.Set(t, bit32.Value(5), bit32.All.Iff(bit32.Value(5)))
	expect.Set(t, bit32.All, bit32.Value(5).Iff(bit32.Value(5)))
	expect.Set(t, bit32.Values(0, 1, 2).Or(bit32.More(3)), bit32.Value(5).Iff(bit32.Values(3, 5)))
	expect.Set(t, bit32.All, bit32.Values(3, 5).Iff(bit32.Values(3, 5)))
	expect.Set(t, bit32.Value(3).Or(bit32.More(5)), bit32.Values(3, 5).Iff(bit32.Less(5)))
	expect.Set(t, bit32.All, bit32.Less(5).Iff(bit32.Less(5)))
	expect.Set(t, bit32.Value(5), bit32.Less(5).Iff(bit32.More(5)))
	expect.Set(t, bit32.All, bit32.More(5).Iff(bit32.More(5)))
	expect.Set(t, bit32.Less(6), bit32.More(5).Iff(bit32.None))
}

func TestImply(t *testing.T) {
	expect.Set(t, bit32.All, bit32.None.Imply(bit32.None))
	expect.Set(t, bit32.All, bit32.None.Imply(bit32.All))
	expect.Set(t, bit32.All, bit32.All.Imply(bit32.All))
	expect.Set(t, bit32.Value(5), bit32.All.Imply(bit32.Value(5)))
	expect.Set(t, bit32.All, bit32.Value(5).Imply(bit32.Value(5)))
	expect.Set(t, bit32.All, bit32.Value(5).Imply(bit32.Values(3, 5)))
	expect.Set(t, bit32.All, bit32.Values(3, 5).Imply(bit32.Values(3, 5)))
	expect.Set(t, bit32.Value(5).Not(), bit32.Values(3, 5).Imply(bit32.Less(5)))
	expect.Set(t, bit32.All, bit32.Less(5).Imply(bit32.Less(5)))
	expect.Set(t, bit32.More(4), bit32.Less(5).Imply(bit32.More(5)))
	expect.Set(t, bit32.All, bit32.More(5).Imply(bit32.More(5)))
	expect.Set(t, bit32.Less(6), bit32.More(5).Imply(bit32.None))
}
