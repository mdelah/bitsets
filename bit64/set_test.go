package bit64_test

import (
	"github.com/mdelah/bitsets/bit64"
	"github.com/mdelah/bitsets/internal/expect"
	"testing"
)

func TestCount(t *testing.T) {
	expect.Eq(t, 0, bit64.None.Count())
	expect.Eq(t, bit64.Cap, bit64.All.Count())
	expect.Eq(t, 1, bit64.Value(5).Count())
	expect.Eq(t, 2, bit64.Values(3, 5).Count())
	expect.Eq(t, 5, bit64.Less(5).Count())
	expect.Eq(t, bit64.Max-5, bit64.More(5).Count())
}

func TestString(t *testing.T) {
	expect.Eq(t, "{}", bit64.None.String())
	expect.Eq(t, "{0-63}", bit64.All.String())
	expect.Eq(t, "{5}", bit64.Value(5).String())
	expect.Eq(t, "{3,5}", bit64.Values(3, 5).String())
	expect.Eq(t, "{0-4}", bit64.Less(5).String())
	expect.Eq(t, "{6-63}", bit64.More(5).String())
}

func TestIsNone(t *testing.T) {
	expect.Eq(t, true, bit64.None.IsNone())
	expect.Eq(t, false, bit64.All.IsNone())
	expect.Eq(t, false, bit64.Value(5).IsNone())
	expect.Eq(t, false, bit64.Values(3, 5).IsNone())
	expect.Eq(t, false, bit64.Less(5).IsNone())
	expect.Eq(t, false, bit64.More(5).IsNone())
}

func TestIsAll(t *testing.T) {
	expect.Eq(t, false, bit64.None.IsAll())
	expect.Eq(t, true, bit64.All.IsAll())
	expect.Eq(t, false, bit64.Value(5).IsAll())
	expect.Eq(t, false, bit64.Values(3, 5).IsAll())
	expect.Eq(t, false, bit64.Less(5).IsAll())
	expect.Eq(t, false, bit64.More(5).IsAll())
}

func TestMin(t *testing.T) {
	expect.Eq(t, -1, bit64.None.Min())
	expect.Eq(t, 0, bit64.All.Min())
	expect.Eq(t, 5, bit64.Value(5).Min())
	expect.Eq(t, 3, bit64.Values(3, 5).Min())
	expect.Eq(t, 0, bit64.Less(5).Min())
	expect.Eq(t, 6, bit64.More(5).Min())
}

func TestMax(t *testing.T) {
	expect.Eq(t, -1, bit64.None.Max())
	expect.Eq(t, bit64.Max, bit64.All.Max())
	expect.Eq(t, 5, bit64.Value(5).Max())
	expect.Eq(t, 5, bit64.Values(3, 5).Max())
	expect.Eq(t, 4, bit64.Less(5).Max())
	expect.Eq(t, bit64.Max, bit64.More(5).Max())
}

func TestHas(t *testing.T) {
	expect.Eq(t, false, bit64.None.Has(3))
	expect.Eq(t, true, bit64.All.Has(3))
	expect.Eq(t, false, bit64.Value(5).Has(3))
	expect.Eq(t, true, bit64.Values(3, 5).Has(3))
	expect.Eq(t, true, bit64.Less(5).Has(3))
	expect.Eq(t, false, bit64.More(5).Has(3))
}

func TestLessCount(t *testing.T) {
	expect.Eq(t, 0, bit64.None.LessCount(5))
	expect.Eq(t, 5, bit64.All.LessCount(5))
	expect.Eq(t, 0, bit64.Value(5).LessCount(5))
	expect.Eq(t, 1, bit64.Values(3, 5).LessCount(5))
	expect.Eq(t, 5, bit64.Less(5).LessCount(5))
	expect.Eq(t, 0, bit64.More(5).LessCount(5))
}

func TestMoreCount(t *testing.T) {
	expect.Eq(t, 0, bit64.None.MoreCount(4))
	expect.Eq(t, bit64.Cap-5, bit64.All.MoreCount(4))
	expect.Eq(t, 1, bit64.Value(5).MoreCount(4))
	expect.Eq(t, 1, bit64.Values(3, 5).MoreCount(4))
	expect.Eq(t, 0, bit64.Less(5).MoreCount(4))
	expect.Eq(t, bit64.Cap-6, bit64.More(5).MoreCount(4))
}

func TestEqual(t *testing.T) {
	expect.Eq(t, true, bit64.None.Equal(bit64.None))
	expect.Eq(t, false, bit64.None.Equal(bit64.All))
	expect.Eq(t, true, bit64.All.Equal(bit64.All))
	expect.Eq(t, false, bit64.All.Equal(bit64.Value(5)))
	expect.Eq(t, true, bit64.Value(5).Equal(bit64.Value(5)))
	expect.Eq(t, false, bit64.Value(5).Equal(bit64.Values(3, 5)))
	expect.Eq(t, true, bit64.Values(3, 5).Equal(bit64.Values(3, 5)))
	expect.Eq(t, false, bit64.Values(3, 5).Equal(bit64.Less(5)))
	expect.Eq(t, true, bit64.Less(5).Equal(bit64.Less(5)))
	expect.Eq(t, false, bit64.Less(5).Equal(bit64.More(5)))
	expect.Eq(t, true, bit64.More(5).Equal(bit64.More(5)))
	expect.Eq(t, false, bit64.More(5).Equal(bit64.None))
}

func TestHasNone(t *testing.T) {
	expect.Eq(t, true, bit64.None.HasNone(bit64.None))
	expect.Eq(t, true, bit64.None.HasNone(bit64.All))
	expect.Eq(t, false, bit64.All.HasNone(bit64.All))
	expect.Eq(t, false, bit64.All.HasNone(bit64.Value(5)))
	expect.Eq(t, false, bit64.Value(5).HasNone(bit64.Value(5)))
	expect.Eq(t, false, bit64.Value(5).HasNone(bit64.Values(3, 5)))
	expect.Eq(t, false, bit64.Values(3, 5).HasNone(bit64.Values(3, 5)))
	expect.Eq(t, false, bit64.Values(3, 5).HasNone(bit64.Less(5)))
	expect.Eq(t, false, bit64.Less(5).HasNone(bit64.Less(5)))
	expect.Eq(t, true, bit64.Less(5).HasNone(bit64.More(5)))
	expect.Eq(t, false, bit64.More(5).HasNone(bit64.More(5)))
	expect.Eq(t, true, bit64.More(5).HasNone(bit64.None))
}

func TestHasAll(t *testing.T) {
	expect.Eq(t, true, bit64.None.HasAll(bit64.None))
	expect.Eq(t, false, bit64.None.HasAll(bit64.All))
	expect.Eq(t, true, bit64.All.HasAll(bit64.All))
	expect.Eq(t, true, bit64.All.HasAll(bit64.Value(5)))
	expect.Eq(t, true, bit64.Value(5).HasAll(bit64.Value(5)))
	expect.Eq(t, false, bit64.Value(5).HasAll(bit64.Values(3, 5)))
	expect.Eq(t, true, bit64.Values(3, 5).HasAll(bit64.Values(3, 5)))
	expect.Eq(t, false, bit64.Values(3, 5).HasAll(bit64.Less(5)))
	expect.Eq(t, true, bit64.Less(5).HasAll(bit64.Less(5)))
	expect.Eq(t, false, bit64.Less(5).HasAll(bit64.More(5)))
	expect.Eq(t, true, bit64.More(5).HasAll(bit64.More(5)))
	expect.Eq(t, true, bit64.More(5).HasAll(bit64.None))
}

func TestCompare(t *testing.T) {
	expect.Eq(t, 0, bit64.None.Compare(bit64.None))
	expect.Eq(t, -1, bit64.None.Compare(bit64.All))
	expect.Eq(t, 0, bit64.All.Compare(bit64.All))
	expect.Eq(t, 1, bit64.All.Compare(bit64.Value(5)))
	expect.Eq(t, 0, bit64.Value(5).Compare(bit64.Value(5)))
	expect.Eq(t, -1, bit64.Value(5).Compare(bit64.Values(3, 5)))
	expect.Eq(t, 0, bit64.Values(3, 5).Compare(bit64.Values(3, 5)))
	expect.Eq(t, -1, bit64.Values(3, 5).Compare(bit64.Less(5)))
	expect.Eq(t, 0, bit64.Less(5).Compare(bit64.Less(5)))
	expect.Eq(t, 1, bit64.Less(5).Compare(bit64.More(5)))
	expect.Eq(t, 0, bit64.More(5).Compare(bit64.More(5)))
	expect.Eq(t, 1, bit64.More(5).Compare(bit64.None))
}

func TestAndCount(t *testing.T) {
	expect.Eq(t, 0, bit64.None.AndCount(bit64.None))
	expect.Eq(t, 0, bit64.None.AndCount(bit64.All))
	expect.Eq(t, bit64.Cap, bit64.All.AndCount(bit64.All))
	expect.Eq(t, 1, bit64.All.AndCount(bit64.Value(5)))
	expect.Eq(t, 1, bit64.Value(5).AndCount(bit64.Value(5)))
	expect.Eq(t, 1, bit64.Value(5).AndCount(bit64.Values(3, 5)))
	expect.Eq(t, 2, bit64.Values(3, 5).AndCount(bit64.Values(3, 5)))
	expect.Eq(t, 1, bit64.Values(3, 5).AndCount(bit64.Less(5)))
	expect.Eq(t, 5, bit64.Less(5).AndCount(bit64.Less(5)))
	expect.Eq(t, 0, bit64.Less(5).AndCount(bit64.More(5)))
	expect.Eq(t, bit64.Cap-6, bit64.More(5).AndCount(bit64.More(5)))
	expect.Eq(t, 0, bit64.More(5).AndCount(bit64.None))
}

func TestAdd(t *testing.T) {
	x := bit64.None
	x.Add(3)
	expect.Set(t, bit64.Value(3), x)
	x.Add(3)
	expect.Set(t, bit64.Value(3), x)
	x.Add(5)
	expect.Set(t, bit64.Values(3, 5), x)
}

func TestRemove(t *testing.T) {
	x := bit64.Values(3, 5)
	x.Remove(1)
	expect.Set(t, bit64.Values(3, 5), x)
	x.Remove(3)
	expect.Set(t, bit64.Value(5), x)
	x.Remove(5)
	expect.Set(t, bit64.None, x)
}

func TestAssign(t *testing.T) {
	x := bit64.None
	x.Assign(bit64.Value(5))
	expect.Set(t, bit64.Value(5), x)
}

func TestAssignNone(t *testing.T) {
	x := bit64.Values(3, 5)
	x.AssignNone()
	expect.Set(t, bit64.None, x)
}

func TestAssignAll(t *testing.T) {
	x := bit64.Values(3, 5)
	x.AssignAll()
	expect.Set(t, bit64.All, x)
}

func TestEach(t *testing.T) {
	expect.Ints(t, bit64.None.Each())
	expect.Ints(t, bit64.Value(5).Each(), 5)
	expect.Ints(t, bit64.Values(3, 5).Each(), 3, 5)
	expect.Ints(t, bit64.Less(5).Each(), 0, 1, 2, 3, 4)
}

func TestNot(t *testing.T) {
	expect.Set(t, bit64.All, bit64.None.Not())
	expect.Set(t, bit64.None, bit64.All.Not())
	expect.Set(t, bit64.Less(5).Or(bit64.More(5)), bit64.Value(5).Not())
	expect.Set(t, bit64.Less(3).Or(bit64.Value(4).Or(bit64.More(5))), bit64.Values(3, 5).Not())
	expect.Set(t, bit64.More(4), bit64.Less(5).Not())
	expect.Set(t, bit64.Less(6), bit64.More(5).Not())
}

func TestAssignNot(t *testing.T) {
	x := bit64.Values(3, 5)
	x.AssignNot()
	expect.Set(t, bit64.Less(3).Or(bit64.Value(4)).Or(bit64.More(5)), x)
}

func TestAssignSub(t *testing.T) {
	x := bit64.Values(3, 5)
	x.AssignSub(bit64.Values(3))
	expect.Set(t, bit64.Values(5), x)
}

func TestAssignAnd(t *testing.T) {
	x := bit64.Values(3, 5)
	x.AssignAnd(bit64.Values(3))
	expect.Set(t, bit64.Values(3), x)
}

func TestAssignOr(t *testing.T) {
	x := bit64.Values(3, 5)
	x.AssignOr(bit64.Values(4))
	expect.Set(t, bit64.Values(3, 4, 5), x)
}

func TestAssignXor(t *testing.T) {
	x := bit64.Values(3, 5)
	x.AssignXor(bit64.Values(3, 4))
	expect.Set(t, bit64.Values(4, 5), x)
}

func TestAssignNor(t *testing.T) {
	x := bit64.Values(3, 5)
	x.AssignNor(bit64.Values(3, 4))
	expect.Set(t, bit64.Less(3).Or(bit64.More(5)), x)
}

func TestAssignIff(t *testing.T) {
	x := bit64.Values(3, 5)
	x.AssignIff(bit64.Values(3, 4))
	expect.Set(t, bit64.Less(4).Or(bit64.More(5)), x)
}

func TestAssignImply(t *testing.T) {
	x := bit64.Values(3, 5)
	x.AssignImply(bit64.Values(3, 4))
	expect.Set(t, bit64.Less(5).Or(bit64.More(5)), x)
}

func TestSub(t *testing.T) {
	expect.Set(t, bit64.None, bit64.None.Sub(bit64.None))
	expect.Set(t, bit64.None, bit64.None.Sub(bit64.All))
	expect.Set(t, bit64.None, bit64.All.Sub(bit64.All))
	expect.Set(t, bit64.Less(5).Or(bit64.More(5)), bit64.All.Sub(bit64.Value(5)))
	expect.Set(t, bit64.None, bit64.Value(5).Sub(bit64.Value(5)))
	expect.Set(t, bit64.None, bit64.Value(5).Sub(bit64.Values(3, 5)))
	expect.Set(t, bit64.None, bit64.Values(3, 5).Sub(bit64.Values(3, 5)))
	expect.Set(t, bit64.Value(5), bit64.Values(3, 5).Sub(bit64.Less(5)))
	expect.Set(t, bit64.None, bit64.Less(5).Sub(bit64.Less(5)))
	expect.Set(t, bit64.Less(5), bit64.Less(5).Sub(bit64.More(5)))
	expect.Set(t, bit64.None, bit64.More(5).Sub(bit64.More(5)))
	expect.Set(t, bit64.More(5), bit64.More(5).Sub(bit64.None))
}

func TestAnd(t *testing.T) {
	expect.Set(t, bit64.None, bit64.None.And(bit64.None))
	expect.Set(t, bit64.None, bit64.None.And(bit64.All))
	expect.Set(t, bit64.All, bit64.All.And(bit64.All))
	expect.Set(t, bit64.Value(5), bit64.All.And(bit64.Value(5)))
	expect.Set(t, bit64.Value(5), bit64.Value(5).And(bit64.Value(5)))
	expect.Set(t, bit64.Value(5), bit64.Value(5).And(bit64.Values(3, 5)))
	expect.Set(t, bit64.Values(3, 5), bit64.Values(3, 5).And(bit64.Values(3, 5)))
	expect.Set(t, bit64.Value(3), bit64.Values(3, 5).And(bit64.Less(5)))
	expect.Set(t, bit64.Less(5), bit64.Less(5).And(bit64.Less(5)))
	expect.Set(t, bit64.None, bit64.Less(5).And(bit64.More(5)))
	expect.Set(t, bit64.More(5), bit64.More(5).And(bit64.More(5)))
	expect.Set(t, bit64.None, bit64.More(5).And(bit64.None))
}

func TestOr(t *testing.T) {
	expect.Set(t, bit64.None, bit64.None.Or(bit64.None))
	expect.Set(t, bit64.All, bit64.None.Or(bit64.All))
	expect.Set(t, bit64.All, bit64.All.Or(bit64.All))
	expect.Set(t, bit64.All, bit64.All.Or(bit64.Value(5)))
	expect.Set(t, bit64.Value(5), bit64.Value(5).Or(bit64.Value(5)))
	expect.Set(t, bit64.Values(5, 3), bit64.Value(5).Or(bit64.Values(3, 5)))
	expect.Set(t, bit64.Values(3, 5), bit64.Values(3, 5).Or(bit64.Values(3, 5)))
	expect.Set(t, bit64.Less(6), bit64.Values(3, 5).Or(bit64.Less(5)))
	expect.Set(t, bit64.Less(5), bit64.Less(5).Or(bit64.Less(5)))
	expect.Set(t, bit64.Value(5).Not(), bit64.Less(5).Or(bit64.More(5)))
	expect.Set(t, bit64.More(5), bit64.More(5).Or(bit64.More(5)))
	expect.Set(t, bit64.More(5), bit64.More(5).Or(bit64.None))
}

func TestXor(t *testing.T) {
	expect.Set(t, bit64.None, bit64.None.Xor(bit64.None))
	expect.Set(t, bit64.All, bit64.None.Xor(bit64.All))
	expect.Set(t, bit64.None, bit64.All.Xor(bit64.All))
	expect.Set(t, bit64.Value(5).Not(), bit64.All.Xor(bit64.Value(5)))
	expect.Set(t, bit64.None, bit64.Value(5).Xor(bit64.Value(5)))
	expect.Set(t, bit64.Value(3), bit64.Value(5).Xor(bit64.Values(3, 5)))
	expect.Set(t, bit64.None, bit64.Values(3, 5).Xor(bit64.Values(3, 5)))
	expect.Set(t, bit64.Values(0, 1, 2, 4, 5), bit64.Values(3, 5).Xor(bit64.Less(5)))
	expect.Set(t, bit64.None, bit64.Less(5).Xor(bit64.Less(5)))
	expect.Set(t, bit64.Value(5).Not(), bit64.Less(5).Xor(bit64.More(5)))
	expect.Set(t, bit64.None, bit64.More(5).Xor(bit64.More(5)))
	expect.Set(t, bit64.More(5), bit64.More(5).Xor(bit64.None))
}

func TestNor(t *testing.T) {
	expect.Set(t, bit64.All, bit64.None.Nor(bit64.None))
	expect.Set(t, bit64.None, bit64.None.Nor(bit64.All))
	expect.Set(t, bit64.None, bit64.All.Nor(bit64.All))
	expect.Set(t, bit64.None, bit64.All.Nor(bit64.Value(5)))
	expect.Set(t, bit64.Less(5).Or(bit64.More(5)), bit64.Value(5).Nor(bit64.Value(5)))
	expect.Set(t, bit64.Values(0, 1, 2, 4).Or(bit64.More(5)), bit64.Value(5).Nor(bit64.Values(3, 5)))
	expect.Set(t, bit64.Values(0, 1, 2, 4).Or(bit64.More(5)), bit64.Values(3, 5).Nor(bit64.Values(3, 5)))
	expect.Set(t, bit64.More(5), bit64.Values(3, 5).Nor(bit64.Less(5)))
	expect.Set(t, bit64.More(4), bit64.Less(5).Nor(bit64.Less(5)))
	expect.Set(t, bit64.Value(5), bit64.Less(5).Nor(bit64.More(5)))
	expect.Set(t, bit64.Less(6), bit64.More(5).Nor(bit64.More(5)))
	expect.Set(t, bit64.Less(6), bit64.More(5).Nor(bit64.None))
}

func TestIff(t *testing.T) {
	expect.Set(t, bit64.All, bit64.None.Iff(bit64.None))
	expect.Set(t, bit64.None, bit64.None.Iff(bit64.All))
	expect.Set(t, bit64.All, bit64.All.Iff(bit64.All))
	expect.Set(t, bit64.Value(5), bit64.All.Iff(bit64.Value(5)))
	expect.Set(t, bit64.All, bit64.Value(5).Iff(bit64.Value(5)))
	expect.Set(t, bit64.Values(0, 1, 2).Or(bit64.More(3)), bit64.Value(5).Iff(bit64.Values(3, 5)))
	expect.Set(t, bit64.All, bit64.Values(3, 5).Iff(bit64.Values(3, 5)))
	expect.Set(t, bit64.Value(3).Or(bit64.More(5)), bit64.Values(3, 5).Iff(bit64.Less(5)))
	expect.Set(t, bit64.All, bit64.Less(5).Iff(bit64.Less(5)))
	expect.Set(t, bit64.Value(5), bit64.Less(5).Iff(bit64.More(5)))
	expect.Set(t, bit64.All, bit64.More(5).Iff(bit64.More(5)))
	expect.Set(t, bit64.Less(6), bit64.More(5).Iff(bit64.None))
}

func TestImply(t *testing.T) {
	expect.Set(t, bit64.All, bit64.None.Imply(bit64.None))
	expect.Set(t, bit64.All, bit64.None.Imply(bit64.All))
	expect.Set(t, bit64.All, bit64.All.Imply(bit64.All))
	expect.Set(t, bit64.Value(5), bit64.All.Imply(bit64.Value(5)))
	expect.Set(t, bit64.All, bit64.Value(5).Imply(bit64.Value(5)))
	expect.Set(t, bit64.All, bit64.Value(5).Imply(bit64.Values(3, 5)))
	expect.Set(t, bit64.All, bit64.Values(3, 5).Imply(bit64.Values(3, 5)))
	expect.Set(t, bit64.Value(5).Not(), bit64.Values(3, 5).Imply(bit64.Less(5)))
	expect.Set(t, bit64.All, bit64.Less(5).Imply(bit64.Less(5)))
	expect.Set(t, bit64.More(4), bit64.Less(5).Imply(bit64.More(5)))
	expect.Set(t, bit64.All, bit64.More(5).Imply(bit64.More(5)))
	expect.Set(t, bit64.Less(6), bit64.More(5).Imply(bit64.None))
}
