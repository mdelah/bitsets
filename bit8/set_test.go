package bit8_test

import (
	"github.com/mdelah/bitsets/bit8"
	"github.com/mdelah/bitsets/internal/expect"
	"testing"
)

func TestCount(t *testing.T) {
	expect.Eq(t, 0, bit8.None.Count())
	expect.Eq(t, bit8.Cap, bit8.All.Count())
	expect.Eq(t, 1, bit8.Value(5).Count())
	expect.Eq(t, 2, bit8.Values(3, 5).Count())
	expect.Eq(t, 5, bit8.Less(5).Count())
	expect.Eq(t, bit8.Max-5, bit8.More(5).Count())
}

func TestString(t *testing.T) {
	expect.Eq(t, "{}", bit8.None.String())
	expect.Eq(t, "{0-7}", bit8.All.String())
	expect.Eq(t, "{5}", bit8.Value(5).String())
	expect.Eq(t, "{3,5}", bit8.Values(3, 5).String())
	expect.Eq(t, "{0-4}", bit8.Less(5).String())
	expect.Eq(t, "{6-7}", bit8.More(5).String())
}

func TestIsNone(t *testing.T) {
	expect.Eq(t, true, bit8.None.IsNone())
	expect.Eq(t, false, bit8.All.IsNone())
	expect.Eq(t, false, bit8.Value(5).IsNone())
	expect.Eq(t, false, bit8.Values(3, 5).IsNone())
	expect.Eq(t, false, bit8.Less(5).IsNone())
	expect.Eq(t, false, bit8.More(5).IsNone())
}

func TestIsAll(t *testing.T) {
	expect.Eq(t, false, bit8.None.IsAll())
	expect.Eq(t, true, bit8.All.IsAll())
	expect.Eq(t, false, bit8.Value(5).IsAll())
	expect.Eq(t, false, bit8.Values(3, 5).IsAll())
	expect.Eq(t, false, bit8.Less(5).IsAll())
	expect.Eq(t, false, bit8.More(5).IsAll())
}

func TestMin(t *testing.T) {
	expect.Eq(t, -1, bit8.None.Min())
	expect.Eq(t, 0, bit8.All.Min())
	expect.Eq(t, 5, bit8.Value(5).Min())
	expect.Eq(t, 3, bit8.Values(3, 5).Min())
	expect.Eq(t, 0, bit8.Less(5).Min())
	expect.Eq(t, 6, bit8.More(5).Min())
}

func TestMax(t *testing.T) {
	expect.Eq(t, -1, bit8.None.Max())
	expect.Eq(t, bit8.Max, bit8.All.Max())
	expect.Eq(t, 5, bit8.Value(5).Max())
	expect.Eq(t, 5, bit8.Values(3, 5).Max())
	expect.Eq(t, 4, bit8.Less(5).Max())
	expect.Eq(t, bit8.Max, bit8.More(5).Max())
}

func TestHas(t *testing.T) {
	expect.Eq(t, false, bit8.None.Has(3))
	expect.Eq(t, true, bit8.All.Has(3))
	expect.Eq(t, false, bit8.Value(5).Has(3))
	expect.Eq(t, true, bit8.Values(3, 5).Has(3))
	expect.Eq(t, true, bit8.Less(5).Has(3))
	expect.Eq(t, false, bit8.More(5).Has(3))
}

func TestLessCount(t *testing.T) {
	expect.Eq(t, 0, bit8.None.LessCount(5))
	expect.Eq(t, 5, bit8.All.LessCount(5))
	expect.Eq(t, 0, bit8.Value(5).LessCount(5))
	expect.Eq(t, 1, bit8.Values(3, 5).LessCount(5))
	expect.Eq(t, 5, bit8.Less(5).LessCount(5))
	expect.Eq(t, 0, bit8.More(5).LessCount(5))
}

func TestMoreCount(t *testing.T) {
	expect.Eq(t, 0, bit8.None.MoreCount(4))
	expect.Eq(t, bit8.Cap-5, bit8.All.MoreCount(4))
	expect.Eq(t, 1, bit8.Value(5).MoreCount(4))
	expect.Eq(t, 1, bit8.Values(3, 5).MoreCount(4))
	expect.Eq(t, 0, bit8.Less(5).MoreCount(4))
	expect.Eq(t, bit8.Cap-6, bit8.More(5).MoreCount(4))
}

func TestEqual(t *testing.T) {
	expect.Eq(t, true, bit8.None.Equal(bit8.None))
	expect.Eq(t, false, bit8.None.Equal(bit8.All))
	expect.Eq(t, true, bit8.All.Equal(bit8.All))
	expect.Eq(t, false, bit8.All.Equal(bit8.Value(5)))
	expect.Eq(t, true, bit8.Value(5).Equal(bit8.Value(5)))
	expect.Eq(t, false, bit8.Value(5).Equal(bit8.Values(3, 5)))
	expect.Eq(t, true, bit8.Values(3, 5).Equal(bit8.Values(3, 5)))
	expect.Eq(t, false, bit8.Values(3, 5).Equal(bit8.Less(5)))
	expect.Eq(t, true, bit8.Less(5).Equal(bit8.Less(5)))
	expect.Eq(t, false, bit8.Less(5).Equal(bit8.More(5)))
	expect.Eq(t, true, bit8.More(5).Equal(bit8.More(5)))
	expect.Eq(t, false, bit8.More(5).Equal(bit8.None))
}

func TestHasNone(t *testing.T) {
	expect.Eq(t, true, bit8.None.HasNone(bit8.None))
	expect.Eq(t, true, bit8.None.HasNone(bit8.All))
	expect.Eq(t, false, bit8.All.HasNone(bit8.All))
	expect.Eq(t, false, bit8.All.HasNone(bit8.Value(5)))
	expect.Eq(t, false, bit8.Value(5).HasNone(bit8.Value(5)))
	expect.Eq(t, false, bit8.Value(5).HasNone(bit8.Values(3, 5)))
	expect.Eq(t, false, bit8.Values(3, 5).HasNone(bit8.Values(3, 5)))
	expect.Eq(t, false, bit8.Values(3, 5).HasNone(bit8.Less(5)))
	expect.Eq(t, false, bit8.Less(5).HasNone(bit8.Less(5)))
	expect.Eq(t, true, bit8.Less(5).HasNone(bit8.More(5)))
	expect.Eq(t, false, bit8.More(5).HasNone(bit8.More(5)))
	expect.Eq(t, true, bit8.More(5).HasNone(bit8.None))
}

func TestHasAll(t *testing.T) {
	expect.Eq(t, true, bit8.None.HasAll(bit8.None))
	expect.Eq(t, false, bit8.None.HasAll(bit8.All))
	expect.Eq(t, true, bit8.All.HasAll(bit8.All))
	expect.Eq(t, true, bit8.All.HasAll(bit8.Value(5)))
	expect.Eq(t, true, bit8.Value(5).HasAll(bit8.Value(5)))
	expect.Eq(t, false, bit8.Value(5).HasAll(bit8.Values(3, 5)))
	expect.Eq(t, true, bit8.Values(3, 5).HasAll(bit8.Values(3, 5)))
	expect.Eq(t, false, bit8.Values(3, 5).HasAll(bit8.Less(5)))
	expect.Eq(t, true, bit8.Less(5).HasAll(bit8.Less(5)))
	expect.Eq(t, false, bit8.Less(5).HasAll(bit8.More(5)))
	expect.Eq(t, true, bit8.More(5).HasAll(bit8.More(5)))
	expect.Eq(t, true, bit8.More(5).HasAll(bit8.None))
}

func TestCompare(t *testing.T) {
	expect.Eq(t, 0, bit8.None.Compare(bit8.None))
	expect.Eq(t, -1, bit8.None.Compare(bit8.All))
	expect.Eq(t, 0, bit8.All.Compare(bit8.All))
	expect.Eq(t, 1, bit8.All.Compare(bit8.Value(5)))
	expect.Eq(t, 0, bit8.Value(5).Compare(bit8.Value(5)))
	expect.Eq(t, -1, bit8.Value(5).Compare(bit8.Values(3, 5)))
	expect.Eq(t, 0, bit8.Values(3, 5).Compare(bit8.Values(3, 5)))
	expect.Eq(t, -1, bit8.Values(3, 5).Compare(bit8.Less(5)))
	expect.Eq(t, 0, bit8.Less(5).Compare(bit8.Less(5)))
	expect.Eq(t, 1, bit8.Less(5).Compare(bit8.More(5)))
	expect.Eq(t, 0, bit8.More(5).Compare(bit8.More(5)))
	expect.Eq(t, 1, bit8.More(5).Compare(bit8.None))
}

func TestAndCount(t *testing.T) {
	expect.Eq(t, 0, bit8.None.AndCount(bit8.None))
	expect.Eq(t, 0, bit8.None.AndCount(bit8.All))
	expect.Eq(t, bit8.Cap, bit8.All.AndCount(bit8.All))
	expect.Eq(t, 1, bit8.All.AndCount(bit8.Value(5)))
	expect.Eq(t, 1, bit8.Value(5).AndCount(bit8.Value(5)))
	expect.Eq(t, 1, bit8.Value(5).AndCount(bit8.Values(3, 5)))
	expect.Eq(t, 2, bit8.Values(3, 5).AndCount(bit8.Values(3, 5)))
	expect.Eq(t, 1, bit8.Values(3, 5).AndCount(bit8.Less(5)))
	expect.Eq(t, 5, bit8.Less(5).AndCount(bit8.Less(5)))
	expect.Eq(t, 0, bit8.Less(5).AndCount(bit8.More(5)))
	expect.Eq(t, bit8.Cap-6, bit8.More(5).AndCount(bit8.More(5)))
	expect.Eq(t, 0, bit8.More(5).AndCount(bit8.None))
}

func TestAdd(t *testing.T) {
	x := bit8.None
	x.Add(3)
	expect.Set(t, bit8.Value(3), x)
	x.Add(3)
	expect.Set(t, bit8.Value(3), x)
	x.Add(5)
	expect.Set(t, bit8.Values(3, 5), x)
}

func TestRemove(t *testing.T) {
	x := bit8.Values(3, 5)
	x.Remove(1)
	expect.Set(t, bit8.Values(3, 5), x)
	x.Remove(3)
	expect.Set(t, bit8.Value(5), x)
	x.Remove(5)
	expect.Set(t, bit8.None, x)
}

func TestAssign(t *testing.T) {
	x := bit8.None
	x.Assign(bit8.Value(5))
	expect.Set(t, bit8.Value(5), x)
}

func TestAssignNone(t *testing.T) {
	x := bit8.Values(3, 5)
	x.AssignNone()
	expect.Set(t, bit8.None, x)
}

func TestAssignAll(t *testing.T) {
	x := bit8.Values(3, 5)
	x.AssignAll()
	expect.Set(t, bit8.All, x)
}

func TestEach(t *testing.T) {
	expect.Ints(t, bit8.None.Each())
	expect.Ints(t, bit8.Value(5).Each(), 5)
	expect.Ints(t, bit8.Values(3, 5).Each(), 3, 5)
	expect.Ints(t, bit8.Less(5).Each(), 0, 1, 2, 3, 4)
}

func TestNot(t *testing.T) {
	expect.Set(t, bit8.All, bit8.None.Not())
	expect.Set(t, bit8.None, bit8.All.Not())
	expect.Set(t, bit8.Less(5).Or(bit8.More(5)), bit8.Value(5).Not())
	expect.Set(t, bit8.Less(3).Or(bit8.Value(4).Or(bit8.More(5))), bit8.Values(3, 5).Not())
	expect.Set(t, bit8.More(4), bit8.Less(5).Not())
	expect.Set(t, bit8.Less(6), bit8.More(5).Not())
}

func TestAssignNot(t *testing.T) {
	x := bit8.Values(3, 5)
	x.AssignNot()
	expect.Set(t, bit8.Less(3).Or(bit8.Value(4)).Or(bit8.More(5)), x)
}

func TestAssignSub(t *testing.T) {
	x := bit8.Values(3, 5)
	x.AssignSub(bit8.Values(3))
	expect.Set(t, bit8.Values(5), x)
}

func TestAssignAnd(t *testing.T) {
	x := bit8.Values(3, 5)
	x.AssignAnd(bit8.Values(3))
	expect.Set(t, bit8.Values(3), x)
}

func TestAssignOr(t *testing.T) {
	x := bit8.Values(3, 5)
	x.AssignOr(bit8.Values(4))
	expect.Set(t, bit8.Values(3, 4, 5), x)
}

func TestAssignXor(t *testing.T) {
	x := bit8.Values(3, 5)
	x.AssignXor(bit8.Values(3, 4))
	expect.Set(t, bit8.Values(4, 5), x)
}

func TestAssignNor(t *testing.T) {
	x := bit8.Values(3, 5)
	x.AssignNor(bit8.Values(3, 4))
	expect.Set(t, bit8.Less(3).Or(bit8.More(5)), x)
}

func TestAssignIff(t *testing.T) {
	x := bit8.Values(3, 5)
	x.AssignIff(bit8.Values(3, 4))
	expect.Set(t, bit8.Less(4).Or(bit8.More(5)), x)
}

func TestAssignImply(t *testing.T) {
	x := bit8.Values(3, 5)
	x.AssignImply(bit8.Values(3, 4))
	expect.Set(t, bit8.Less(5).Or(bit8.More(5)), x)
}

func TestSub(t *testing.T) {
	expect.Set(t, bit8.None, bit8.None.Sub(bit8.None))
	expect.Set(t, bit8.None, bit8.None.Sub(bit8.All))
	expect.Set(t, bit8.None, bit8.All.Sub(bit8.All))
	expect.Set(t, bit8.Less(5).Or(bit8.More(5)), bit8.All.Sub(bit8.Value(5)))
	expect.Set(t, bit8.None, bit8.Value(5).Sub(bit8.Value(5)))
	expect.Set(t, bit8.None, bit8.Value(5).Sub(bit8.Values(3, 5)))
	expect.Set(t, bit8.None, bit8.Values(3, 5).Sub(bit8.Values(3, 5)))
	expect.Set(t, bit8.Value(5), bit8.Values(3, 5).Sub(bit8.Less(5)))
	expect.Set(t, bit8.None, bit8.Less(5).Sub(bit8.Less(5)))
	expect.Set(t, bit8.Less(5), bit8.Less(5).Sub(bit8.More(5)))
	expect.Set(t, bit8.None, bit8.More(5).Sub(bit8.More(5)))
	expect.Set(t, bit8.More(5), bit8.More(5).Sub(bit8.None))
}

func TestAnd(t *testing.T) {
	expect.Set(t, bit8.None, bit8.None.And(bit8.None))
	expect.Set(t, bit8.None, bit8.None.And(bit8.All))
	expect.Set(t, bit8.All, bit8.All.And(bit8.All))
	expect.Set(t, bit8.Value(5), bit8.All.And(bit8.Value(5)))
	expect.Set(t, bit8.Value(5), bit8.Value(5).And(bit8.Value(5)))
	expect.Set(t, bit8.Value(5), bit8.Value(5).And(bit8.Values(3, 5)))
	expect.Set(t, bit8.Values(3, 5), bit8.Values(3, 5).And(bit8.Values(3, 5)))
	expect.Set(t, bit8.Value(3), bit8.Values(3, 5).And(bit8.Less(5)))
	expect.Set(t, bit8.Less(5), bit8.Less(5).And(bit8.Less(5)))
	expect.Set(t, bit8.None, bit8.Less(5).And(bit8.More(5)))
	expect.Set(t, bit8.More(5), bit8.More(5).And(bit8.More(5)))
	expect.Set(t, bit8.None, bit8.More(5).And(bit8.None))
}

func TestOr(t *testing.T) {
	expect.Set(t, bit8.None, bit8.None.Or(bit8.None))
	expect.Set(t, bit8.All, bit8.None.Or(bit8.All))
	expect.Set(t, bit8.All, bit8.All.Or(bit8.All))
	expect.Set(t, bit8.All, bit8.All.Or(bit8.Value(5)))
	expect.Set(t, bit8.Value(5), bit8.Value(5).Or(bit8.Value(5)))
	expect.Set(t, bit8.Values(5, 3), bit8.Value(5).Or(bit8.Values(3, 5)))
	expect.Set(t, bit8.Values(3, 5), bit8.Values(3, 5).Or(bit8.Values(3, 5)))
	expect.Set(t, bit8.Less(6), bit8.Values(3, 5).Or(bit8.Less(5)))
	expect.Set(t, bit8.Less(5), bit8.Less(5).Or(bit8.Less(5)))
	expect.Set(t, bit8.Value(5).Not(), bit8.Less(5).Or(bit8.More(5)))
	expect.Set(t, bit8.More(5), bit8.More(5).Or(bit8.More(5)))
	expect.Set(t, bit8.More(5), bit8.More(5).Or(bit8.None))
}

func TestXor(t *testing.T) {
	expect.Set(t, bit8.None, bit8.None.Xor(bit8.None))
	expect.Set(t, bit8.All, bit8.None.Xor(bit8.All))
	expect.Set(t, bit8.None, bit8.All.Xor(bit8.All))
	expect.Set(t, bit8.Value(5).Not(), bit8.All.Xor(bit8.Value(5)))
	expect.Set(t, bit8.None, bit8.Value(5).Xor(bit8.Value(5)))
	expect.Set(t, bit8.Value(3), bit8.Value(5).Xor(bit8.Values(3, 5)))
	expect.Set(t, bit8.None, bit8.Values(3, 5).Xor(bit8.Values(3, 5)))
	expect.Set(t, bit8.Values(0, 1, 2, 4, 5), bit8.Values(3, 5).Xor(bit8.Less(5)))
	expect.Set(t, bit8.None, bit8.Less(5).Xor(bit8.Less(5)))
	expect.Set(t, bit8.Value(5).Not(), bit8.Less(5).Xor(bit8.More(5)))
	expect.Set(t, bit8.None, bit8.More(5).Xor(bit8.More(5)))
	expect.Set(t, bit8.More(5), bit8.More(5).Xor(bit8.None))
}

func TestNor(t *testing.T) {
	expect.Set(t, bit8.All, bit8.None.Nor(bit8.None))
	expect.Set(t, bit8.None, bit8.None.Nor(bit8.All))
	expect.Set(t, bit8.None, bit8.All.Nor(bit8.All))
	expect.Set(t, bit8.None, bit8.All.Nor(bit8.Value(5)))
	expect.Set(t, bit8.Less(5).Or(bit8.More(5)), bit8.Value(5).Nor(bit8.Value(5)))
	expect.Set(t, bit8.Values(0, 1, 2, 4).Or(bit8.More(5)), bit8.Value(5).Nor(bit8.Values(3, 5)))
	expect.Set(t, bit8.Values(0, 1, 2, 4).Or(bit8.More(5)), bit8.Values(3, 5).Nor(bit8.Values(3, 5)))
	expect.Set(t, bit8.More(5), bit8.Values(3, 5).Nor(bit8.Less(5)))
	expect.Set(t, bit8.More(4), bit8.Less(5).Nor(bit8.Less(5)))
	expect.Set(t, bit8.Value(5), bit8.Less(5).Nor(bit8.More(5)))
	expect.Set(t, bit8.Less(6), bit8.More(5).Nor(bit8.More(5)))
	expect.Set(t, bit8.Less(6), bit8.More(5).Nor(bit8.None))
}

func TestIff(t *testing.T) {
	expect.Set(t, bit8.All, bit8.None.Iff(bit8.None))
	expect.Set(t, bit8.None, bit8.None.Iff(bit8.All))
	expect.Set(t, bit8.All, bit8.All.Iff(bit8.All))
	expect.Set(t, bit8.Value(5), bit8.All.Iff(bit8.Value(5)))
	expect.Set(t, bit8.All, bit8.Value(5).Iff(bit8.Value(5)))
	expect.Set(t, bit8.Values(0, 1, 2).Or(bit8.More(3)), bit8.Value(5).Iff(bit8.Values(3, 5)))
	expect.Set(t, bit8.All, bit8.Values(3, 5).Iff(bit8.Values(3, 5)))
	expect.Set(t, bit8.Value(3).Or(bit8.More(5)), bit8.Values(3, 5).Iff(bit8.Less(5)))
	expect.Set(t, bit8.All, bit8.Less(5).Iff(bit8.Less(5)))
	expect.Set(t, bit8.Value(5), bit8.Less(5).Iff(bit8.More(5)))
	expect.Set(t, bit8.All, bit8.More(5).Iff(bit8.More(5)))
	expect.Set(t, bit8.Less(6), bit8.More(5).Iff(bit8.None))
}

func TestImply(t *testing.T) {
	expect.Set(t, bit8.All, bit8.None.Imply(bit8.None))
	expect.Set(t, bit8.All, bit8.None.Imply(bit8.All))
	expect.Set(t, bit8.All, bit8.All.Imply(bit8.All))
	expect.Set(t, bit8.Value(5), bit8.All.Imply(bit8.Value(5)))
	expect.Set(t, bit8.All, bit8.Value(5).Imply(bit8.Value(5)))
	expect.Set(t, bit8.All, bit8.Value(5).Imply(bit8.Values(3, 5)))
	expect.Set(t, bit8.All, bit8.Values(3, 5).Imply(bit8.Values(3, 5)))
	expect.Set(t, bit8.Value(5).Not(), bit8.Values(3, 5).Imply(bit8.Less(5)))
	expect.Set(t, bit8.All, bit8.Less(5).Imply(bit8.Less(5)))
	expect.Set(t, bit8.More(4), bit8.Less(5).Imply(bit8.More(5)))
	expect.Set(t, bit8.All, bit8.More(5).Imply(bit8.More(5)))
	expect.Set(t, bit8.Less(6), bit8.More(5).Imply(bit8.None))
}
