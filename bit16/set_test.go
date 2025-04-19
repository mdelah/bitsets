package bit16_test

import (
	"github.com/mdelah/bitsets/bit16"
	"github.com/mdelah/bitsets/internal/expect"
	"testing"
)

func TestCount(t *testing.T) {
	expect.Eq(t, 0, bit16.None.Count())
	expect.Eq(t, bit16.Cap, bit16.All.Count())
	expect.Eq(t, 1, bit16.Value(5).Count())
	expect.Eq(t, 2, bit16.Values(3, 5).Count())
	expect.Eq(t, 5, bit16.Less(5).Count())
	expect.Eq(t, bit16.Max-5, bit16.More(5).Count())
}

func TestString(t *testing.T) {
	expect.Eq(t, "{}", bit16.None.String())
	expect.Eq(t, "{0-15}", bit16.All.String())
	expect.Eq(t, "{5}", bit16.Value(5).String())
	expect.Eq(t, "{3,5}", bit16.Values(3, 5).String())
	expect.Eq(t, "{0-4}", bit16.Less(5).String())
	expect.Eq(t, "{6-15}", bit16.More(5).String())
}

func TestIsNone(t *testing.T) {
	expect.Eq(t, true, bit16.None.IsNone())
	expect.Eq(t, false, bit16.All.IsNone())
	expect.Eq(t, false, bit16.Value(5).IsNone())
	expect.Eq(t, false, bit16.Values(3, 5).IsNone())
	expect.Eq(t, false, bit16.Less(5).IsNone())
	expect.Eq(t, false, bit16.More(5).IsNone())
}

func TestIsAll(t *testing.T) {
	expect.Eq(t, false, bit16.None.IsAll())
	expect.Eq(t, true, bit16.All.IsAll())
	expect.Eq(t, false, bit16.Value(5).IsAll())
	expect.Eq(t, false, bit16.Values(3, 5).IsAll())
	expect.Eq(t, false, bit16.Less(5).IsAll())
	expect.Eq(t, false, bit16.More(5).IsAll())
}

func TestMin(t *testing.T) {
	expect.Eq(t, -1, bit16.None.Min())
	expect.Eq(t, 0, bit16.All.Min())
	expect.Eq(t, 5, bit16.Value(5).Min())
	expect.Eq(t, 3, bit16.Values(3, 5).Min())
	expect.Eq(t, 0, bit16.Less(5).Min())
	expect.Eq(t, 6, bit16.More(5).Min())
}

func TestMax(t *testing.T) {
	expect.Eq(t, -1, bit16.None.Max())
	expect.Eq(t, bit16.Max, bit16.All.Max())
	expect.Eq(t, 5, bit16.Value(5).Max())
	expect.Eq(t, 5, bit16.Values(3, 5).Max())
	expect.Eq(t, 4, bit16.Less(5).Max())
	expect.Eq(t, bit16.Max, bit16.More(5).Max())
}

func TestHas(t *testing.T) {
	expect.Eq(t, false, bit16.None.Has(3))
	expect.Eq(t, true, bit16.All.Has(3))
	expect.Eq(t, false, bit16.Value(5).Has(3))
	expect.Eq(t, true, bit16.Values(3, 5).Has(3))
	expect.Eq(t, true, bit16.Less(5).Has(3))
	expect.Eq(t, false, bit16.More(5).Has(3))
}

func TestLessCount(t *testing.T) {
	expect.Eq(t, 0, bit16.None.LessCount(5))
	expect.Eq(t, 5, bit16.All.LessCount(5))
	expect.Eq(t, 0, bit16.Value(5).LessCount(5))
	expect.Eq(t, 1, bit16.Values(3, 5).LessCount(5))
	expect.Eq(t, 5, bit16.Less(5).LessCount(5))
	expect.Eq(t, 0, bit16.More(5).LessCount(5))
}

func TestMoreCount(t *testing.T) {
	expect.Eq(t, 0, bit16.None.MoreCount(4))
	expect.Eq(t, bit16.Cap-5, bit16.All.MoreCount(4))
	expect.Eq(t, 1, bit16.Value(5).MoreCount(4))
	expect.Eq(t, 1, bit16.Values(3, 5).MoreCount(4))
	expect.Eq(t, 0, bit16.Less(5).MoreCount(4))
	expect.Eq(t, bit16.Cap-6, bit16.More(5).MoreCount(4))
}

func TestEqual(t *testing.T) {
	expect.Eq(t, true, bit16.None.Equal(bit16.None))
	expect.Eq(t, false, bit16.None.Equal(bit16.All))
	expect.Eq(t, true, bit16.All.Equal(bit16.All))
	expect.Eq(t, false, bit16.All.Equal(bit16.Value(5)))
	expect.Eq(t, true, bit16.Value(5).Equal(bit16.Value(5)))
	expect.Eq(t, false, bit16.Value(5).Equal(bit16.Values(3, 5)))
	expect.Eq(t, true, bit16.Values(3, 5).Equal(bit16.Values(3, 5)))
	expect.Eq(t, false, bit16.Values(3, 5).Equal(bit16.Less(5)))
	expect.Eq(t, true, bit16.Less(5).Equal(bit16.Less(5)))
	expect.Eq(t, false, bit16.Less(5).Equal(bit16.More(5)))
	expect.Eq(t, true, bit16.More(5).Equal(bit16.More(5)))
	expect.Eq(t, false, bit16.More(5).Equal(bit16.None))
}

func TestHasNone(t *testing.T) {
	expect.Eq(t, true, bit16.None.HasNone(bit16.None))
	expect.Eq(t, true, bit16.None.HasNone(bit16.All))
	expect.Eq(t, false, bit16.All.HasNone(bit16.All))
	expect.Eq(t, false, bit16.All.HasNone(bit16.Value(5)))
	expect.Eq(t, false, bit16.Value(5).HasNone(bit16.Value(5)))
	expect.Eq(t, false, bit16.Value(5).HasNone(bit16.Values(3, 5)))
	expect.Eq(t, false, bit16.Values(3, 5).HasNone(bit16.Values(3, 5)))
	expect.Eq(t, false, bit16.Values(3, 5).HasNone(bit16.Less(5)))
	expect.Eq(t, false, bit16.Less(5).HasNone(bit16.Less(5)))
	expect.Eq(t, true, bit16.Less(5).HasNone(bit16.More(5)))
	expect.Eq(t, false, bit16.More(5).HasNone(bit16.More(5)))
	expect.Eq(t, true, bit16.More(5).HasNone(bit16.None))
}

func TestHasAll(t *testing.T) {
	expect.Eq(t, true, bit16.None.HasAll(bit16.None))
	expect.Eq(t, false, bit16.None.HasAll(bit16.All))
	expect.Eq(t, true, bit16.All.HasAll(bit16.All))
	expect.Eq(t, true, bit16.All.HasAll(bit16.Value(5)))
	expect.Eq(t, true, bit16.Value(5).HasAll(bit16.Value(5)))
	expect.Eq(t, false, bit16.Value(5).HasAll(bit16.Values(3, 5)))
	expect.Eq(t, true, bit16.Values(3, 5).HasAll(bit16.Values(3, 5)))
	expect.Eq(t, false, bit16.Values(3, 5).HasAll(bit16.Less(5)))
	expect.Eq(t, true, bit16.Less(5).HasAll(bit16.Less(5)))
	expect.Eq(t, false, bit16.Less(5).HasAll(bit16.More(5)))
	expect.Eq(t, true, bit16.More(5).HasAll(bit16.More(5)))
	expect.Eq(t, true, bit16.More(5).HasAll(bit16.None))
}

func TestCompare(t *testing.T) {
	expect.Eq(t, 0, bit16.None.Compare(bit16.None))
	expect.Eq(t, -1, bit16.None.Compare(bit16.All))
	expect.Eq(t, 0, bit16.All.Compare(bit16.All))
	expect.Eq(t, 1, bit16.All.Compare(bit16.Value(5)))
	expect.Eq(t, 0, bit16.Value(5).Compare(bit16.Value(5)))
	expect.Eq(t, -1, bit16.Value(5).Compare(bit16.Values(3, 5)))
	expect.Eq(t, 0, bit16.Values(3, 5).Compare(bit16.Values(3, 5)))
	expect.Eq(t, -1, bit16.Values(3, 5).Compare(bit16.Less(5)))
	expect.Eq(t, 0, bit16.Less(5).Compare(bit16.Less(5)))
	expect.Eq(t, 1, bit16.Less(5).Compare(bit16.More(5)))
	expect.Eq(t, 0, bit16.More(5).Compare(bit16.More(5)))
	expect.Eq(t, 1, bit16.More(5).Compare(bit16.None))
}

func TestAndCount(t *testing.T) {
	expect.Eq(t, 0, bit16.None.AndCount(bit16.None))
	expect.Eq(t, 0, bit16.None.AndCount(bit16.All))
	expect.Eq(t, bit16.Cap, bit16.All.AndCount(bit16.All))
	expect.Eq(t, 1, bit16.All.AndCount(bit16.Value(5)))
	expect.Eq(t, 1, bit16.Value(5).AndCount(bit16.Value(5)))
	expect.Eq(t, 1, bit16.Value(5).AndCount(bit16.Values(3, 5)))
	expect.Eq(t, 2, bit16.Values(3, 5).AndCount(bit16.Values(3, 5)))
	expect.Eq(t, 1, bit16.Values(3, 5).AndCount(bit16.Less(5)))
	expect.Eq(t, 5, bit16.Less(5).AndCount(bit16.Less(5)))
	expect.Eq(t, 0, bit16.Less(5).AndCount(bit16.More(5)))
	expect.Eq(t, bit16.Cap-6, bit16.More(5).AndCount(bit16.More(5)))
	expect.Eq(t, 0, bit16.More(5).AndCount(bit16.None))
}

func TestAdd(t *testing.T) {
	x := bit16.None
	x.Add(3)
	expect.Set(t, bit16.Value(3), x)
	x.Add(3)
	expect.Set(t, bit16.Value(3), x)
	x.Add(5)
	expect.Set(t, bit16.Values(3, 5), x)
}

func TestRemove(t *testing.T) {
	x := bit16.Values(3, 5)
	x.Remove(1)
	expect.Set(t, bit16.Values(3, 5), x)
	x.Remove(3)
	expect.Set(t, bit16.Value(5), x)
	x.Remove(5)
	expect.Set(t, bit16.None, x)
}

func TestAssign(t *testing.T) {
	x := bit16.None
	x.Assign(bit16.Value(5))
	expect.Set(t, bit16.Value(5), x)
}

func TestAssignNone(t *testing.T) {
	x := bit16.Values(3, 5)
	x.AssignNone()
	expect.Set(t, bit16.None, x)
}

func TestAssignAll(t *testing.T) {
	x := bit16.Values(3, 5)
	x.AssignAll()
	expect.Set(t, bit16.All, x)
}

func TestEach(t *testing.T) {
	expect.Ints(t, bit16.None.Each())
	expect.Ints(t, bit16.Value(5).Each(), 5)
	expect.Ints(t, bit16.Values(3, 5).Each(), 3, 5)
	expect.Ints(t, bit16.Less(5).Each(), 0, 1, 2, 3, 4)
}

func TestNot(t *testing.T) {
	expect.Set(t, bit16.All, bit16.None.Not())
	expect.Set(t, bit16.None, bit16.All.Not())
	expect.Set(t, bit16.Less(5).Or(bit16.More(5)), bit16.Value(5).Not())
	expect.Set(t, bit16.Less(3).Or(bit16.Value(4).Or(bit16.More(5))), bit16.Values(3, 5).Not())
	expect.Set(t, bit16.More(4), bit16.Less(5).Not())
	expect.Set(t, bit16.Less(6), bit16.More(5).Not())
}

func TestAssignNot(t *testing.T) {
	x := bit16.Values(3, 5)
	x.AssignNot()
	expect.Set(t, bit16.Less(3).Or(bit16.Value(4)).Or(bit16.More(5)), x)
}

func TestAssignSub(t *testing.T) {
	x := bit16.Values(3, 5)
	x.AssignSub(bit16.Values(3))
	expect.Set(t, bit16.Values(5), x)
}

func TestAssignAnd(t *testing.T) {
	x := bit16.Values(3, 5)
	x.AssignAnd(bit16.Values(3))
	expect.Set(t, bit16.Values(3), x)
}

func TestAssignOr(t *testing.T) {
	x := bit16.Values(3, 5)
	x.AssignOr(bit16.Values(4))
	expect.Set(t, bit16.Values(3, 4, 5), x)
}

func TestAssignXor(t *testing.T) {
	x := bit16.Values(3, 5)
	x.AssignXor(bit16.Values(3, 4))
	expect.Set(t, bit16.Values(4, 5), x)
}

func TestAssignNor(t *testing.T) {
	x := bit16.Values(3, 5)
	x.AssignNor(bit16.Values(3, 4))
	expect.Set(t, bit16.Less(3).Or(bit16.More(5)), x)
}

func TestAssignIff(t *testing.T) {
	x := bit16.Values(3, 5)
	x.AssignIff(bit16.Values(3, 4))
	expect.Set(t, bit16.Less(4).Or(bit16.More(5)), x)
}

func TestAssignImply(t *testing.T) {
	x := bit16.Values(3, 5)
	x.AssignImply(bit16.Values(3, 4))
	expect.Set(t, bit16.Less(5).Or(bit16.More(5)), x)
}

func TestSub(t *testing.T) {
	expect.Set(t, bit16.None, bit16.None.Sub(bit16.None))
	expect.Set(t, bit16.None, bit16.None.Sub(bit16.All))
	expect.Set(t, bit16.None, bit16.All.Sub(bit16.All))
	expect.Set(t, bit16.Less(5).Or(bit16.More(5)), bit16.All.Sub(bit16.Value(5)))
	expect.Set(t, bit16.None, bit16.Value(5).Sub(bit16.Value(5)))
	expect.Set(t, bit16.None, bit16.Value(5).Sub(bit16.Values(3, 5)))
	expect.Set(t, bit16.None, bit16.Values(3, 5).Sub(bit16.Values(3, 5)))
	expect.Set(t, bit16.Value(5), bit16.Values(3, 5).Sub(bit16.Less(5)))
	expect.Set(t, bit16.None, bit16.Less(5).Sub(bit16.Less(5)))
	expect.Set(t, bit16.Less(5), bit16.Less(5).Sub(bit16.More(5)))
	expect.Set(t, bit16.None, bit16.More(5).Sub(bit16.More(5)))
	expect.Set(t, bit16.More(5), bit16.More(5).Sub(bit16.None))
}

func TestAnd(t *testing.T) {
	expect.Set(t, bit16.None, bit16.None.And(bit16.None))
	expect.Set(t, bit16.None, bit16.None.And(bit16.All))
	expect.Set(t, bit16.All, bit16.All.And(bit16.All))
	expect.Set(t, bit16.Value(5), bit16.All.And(bit16.Value(5)))
	expect.Set(t, bit16.Value(5), bit16.Value(5).And(bit16.Value(5)))
	expect.Set(t, bit16.Value(5), bit16.Value(5).And(bit16.Values(3, 5)))
	expect.Set(t, bit16.Values(3, 5), bit16.Values(3, 5).And(bit16.Values(3, 5)))
	expect.Set(t, bit16.Value(3), bit16.Values(3, 5).And(bit16.Less(5)))
	expect.Set(t, bit16.Less(5), bit16.Less(5).And(bit16.Less(5)))
	expect.Set(t, bit16.None, bit16.Less(5).And(bit16.More(5)))
	expect.Set(t, bit16.More(5), bit16.More(5).And(bit16.More(5)))
	expect.Set(t, bit16.None, bit16.More(5).And(bit16.None))
}

func TestOr(t *testing.T) {
	expect.Set(t, bit16.None, bit16.None.Or(bit16.None))
	expect.Set(t, bit16.All, bit16.None.Or(bit16.All))
	expect.Set(t, bit16.All, bit16.All.Or(bit16.All))
	expect.Set(t, bit16.All, bit16.All.Or(bit16.Value(5)))
	expect.Set(t, bit16.Value(5), bit16.Value(5).Or(bit16.Value(5)))
	expect.Set(t, bit16.Values(5, 3), bit16.Value(5).Or(bit16.Values(3, 5)))
	expect.Set(t, bit16.Values(3, 5), bit16.Values(3, 5).Or(bit16.Values(3, 5)))
	expect.Set(t, bit16.Less(6), bit16.Values(3, 5).Or(bit16.Less(5)))
	expect.Set(t, bit16.Less(5), bit16.Less(5).Or(bit16.Less(5)))
	expect.Set(t, bit16.Value(5).Not(), bit16.Less(5).Or(bit16.More(5)))
	expect.Set(t, bit16.More(5), bit16.More(5).Or(bit16.More(5)))
	expect.Set(t, bit16.More(5), bit16.More(5).Or(bit16.None))
}

func TestXor(t *testing.T) {
	expect.Set(t, bit16.None, bit16.None.Xor(bit16.None))
	expect.Set(t, bit16.All, bit16.None.Xor(bit16.All))
	expect.Set(t, bit16.None, bit16.All.Xor(bit16.All))
	expect.Set(t, bit16.Value(5).Not(), bit16.All.Xor(bit16.Value(5)))
	expect.Set(t, bit16.None, bit16.Value(5).Xor(bit16.Value(5)))
	expect.Set(t, bit16.Value(3), bit16.Value(5).Xor(bit16.Values(3, 5)))
	expect.Set(t, bit16.None, bit16.Values(3, 5).Xor(bit16.Values(3, 5)))
	expect.Set(t, bit16.Values(0, 1, 2, 4, 5), bit16.Values(3, 5).Xor(bit16.Less(5)))
	expect.Set(t, bit16.None, bit16.Less(5).Xor(bit16.Less(5)))
	expect.Set(t, bit16.Value(5).Not(), bit16.Less(5).Xor(bit16.More(5)))
	expect.Set(t, bit16.None, bit16.More(5).Xor(bit16.More(5)))
	expect.Set(t, bit16.More(5), bit16.More(5).Xor(bit16.None))
}

func TestNor(t *testing.T) {
	expect.Set(t, bit16.All, bit16.None.Nor(bit16.None))
	expect.Set(t, bit16.None, bit16.None.Nor(bit16.All))
	expect.Set(t, bit16.None, bit16.All.Nor(bit16.All))
	expect.Set(t, bit16.None, bit16.All.Nor(bit16.Value(5)))
	expect.Set(t, bit16.Less(5).Or(bit16.More(5)), bit16.Value(5).Nor(bit16.Value(5)))
	expect.Set(t, bit16.Values(0, 1, 2, 4).Or(bit16.More(5)), bit16.Value(5).Nor(bit16.Values(3, 5)))
	expect.Set(t, bit16.Values(0, 1, 2, 4).Or(bit16.More(5)), bit16.Values(3, 5).Nor(bit16.Values(3, 5)))
	expect.Set(t, bit16.More(5), bit16.Values(3, 5).Nor(bit16.Less(5)))
	expect.Set(t, bit16.More(4), bit16.Less(5).Nor(bit16.Less(5)))
	expect.Set(t, bit16.Value(5), bit16.Less(5).Nor(bit16.More(5)))
	expect.Set(t, bit16.Less(6), bit16.More(5).Nor(bit16.More(5)))
	expect.Set(t, bit16.Less(6), bit16.More(5).Nor(bit16.None))
}

func TestIff(t *testing.T) {
	expect.Set(t, bit16.All, bit16.None.Iff(bit16.None))
	expect.Set(t, bit16.None, bit16.None.Iff(bit16.All))
	expect.Set(t, bit16.All, bit16.All.Iff(bit16.All))
	expect.Set(t, bit16.Value(5), bit16.All.Iff(bit16.Value(5)))
	expect.Set(t, bit16.All, bit16.Value(5).Iff(bit16.Value(5)))
	expect.Set(t, bit16.Values(0, 1, 2).Or(bit16.More(3)), bit16.Value(5).Iff(bit16.Values(3, 5)))
	expect.Set(t, bit16.All, bit16.Values(3, 5).Iff(bit16.Values(3, 5)))
	expect.Set(t, bit16.Value(3).Or(bit16.More(5)), bit16.Values(3, 5).Iff(bit16.Less(5)))
	expect.Set(t, bit16.All, bit16.Less(5).Iff(bit16.Less(5)))
	expect.Set(t, bit16.Value(5), bit16.Less(5).Iff(bit16.More(5)))
	expect.Set(t, bit16.All, bit16.More(5).Iff(bit16.More(5)))
	expect.Set(t, bit16.Less(6), bit16.More(5).Iff(bit16.None))
}

func TestImply(t *testing.T) {
	expect.Set(t, bit16.All, bit16.None.Imply(bit16.None))
	expect.Set(t, bit16.All, bit16.None.Imply(bit16.All))
	expect.Set(t, bit16.All, bit16.All.Imply(bit16.All))
	expect.Set(t, bit16.Value(5), bit16.All.Imply(bit16.Value(5)))
	expect.Set(t, bit16.All, bit16.Value(5).Imply(bit16.Value(5)))
	expect.Set(t, bit16.All, bit16.Value(5).Imply(bit16.Values(3, 5)))
	expect.Set(t, bit16.All, bit16.Values(3, 5).Imply(bit16.Values(3, 5)))
	expect.Set(t, bit16.Value(5).Not(), bit16.Values(3, 5).Imply(bit16.Less(5)))
	expect.Set(t, bit16.All, bit16.Less(5).Imply(bit16.Less(5)))
	expect.Set(t, bit16.More(4), bit16.Less(5).Imply(bit16.More(5)))
	expect.Set(t, bit16.All, bit16.More(5).Imply(bit16.More(5)))
	expect.Set(t, bit16.Less(6), bit16.More(5).Imply(bit16.None))
}
