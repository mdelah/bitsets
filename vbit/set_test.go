package vbit_test

import (
	"github.com/mdelah/bitsets/internal/expect"
	"github.com/mdelah/bitsets/vbit"
	"testing"
)

func TestCount(t *testing.T) {
	expect.Eq(t, 0, vbit.None().Count())
	expect.Eq(t, 1, vbit.Value(5).Count())
	expect.Eq(t, 2, vbit.Values(3, 5).Count())
	expect.Eq(t, 5, vbit.Less(5).Count())
}

func TestString(t *testing.T) {
	expect.Eq(t, "{}", vbit.None().String())
	expect.Eq(t, "{0-}", vbit.All().String())
	expect.Eq(t, "{5}", vbit.Value(5).String())
	expect.Eq(t, "{3,5}", vbit.Values(3, 5).String())
	expect.Eq(t, "{0-4}", vbit.Less(5).String())
	expect.Eq(t, "{6-}", vbit.More(5).String())
}

func TestIsNone(t *testing.T) {
	expect.Eq(t, true, vbit.None().IsNone())
	expect.Eq(t, false, vbit.All().IsNone())
	expect.Eq(t, false, vbit.Value(5).IsNone())
	expect.Eq(t, false, vbit.Values(3, 5).IsNone())
	expect.Eq(t, false, vbit.Less(5).IsNone())
	expect.Eq(t, false, vbit.More(5).IsNone())
}

func TestIsAll(t *testing.T) {
	expect.Eq(t, false, vbit.None().IsAll())
	expect.Eq(t, true, vbit.All().IsAll())
	expect.Eq(t, false, vbit.Value(5).IsAll())
	expect.Eq(t, false, vbit.Values(3, 5).IsAll())
	expect.Eq(t, false, vbit.Less(5).IsAll())
	expect.Eq(t, false, vbit.More(5).IsAll())
}

func TestMin(t *testing.T) {
	expect.Eq(t, -1, vbit.None().Min())
	expect.Eq(t, 0, vbit.All().Min())
	expect.Eq(t, 5, vbit.Value(5).Min())
	expect.Eq(t, 3, vbit.Values(3, 5).Min())
	expect.Eq(t, 0, vbit.Less(5).Min())
	expect.Eq(t, 6, vbit.More(5).Min())
}

func TestMax(t *testing.T) {
	expect.Eq(t, -1, vbit.None().Max())
	expect.Eq(t, 5, vbit.Value(5).Max())
	expect.Eq(t, 5, vbit.Values(3, 5).Max())
	expect.Eq(t, 4, vbit.Less(5).Max())
}

func TestHas(t *testing.T) {
	expect.Eq(t, false, vbit.None().Has(3))
	expect.Eq(t, true, vbit.All().Has(3))
	expect.Eq(t, false, vbit.Value(5).Has(3))
	expect.Eq(t, true, vbit.Values(3, 5).Has(3))
	expect.Eq(t, true, vbit.Less(5).Has(3))
	expect.Eq(t, false, vbit.More(5).Has(3))
}

func TestLessCount(t *testing.T) {
	expect.Eq(t, 0, vbit.None().LessCount(5))
	expect.Eq(t, 5, vbit.All().LessCount(5))
	expect.Eq(t, 0, vbit.Value(5).LessCount(5))
	expect.Eq(t, 1, vbit.Values(3, 5).LessCount(5))
	expect.Eq(t, 5, vbit.Less(5).LessCount(5))
	expect.Eq(t, 0, vbit.More(5).LessCount(5))
}

func TestMoreCount(t *testing.T) {
	expect.Eq(t, 0, vbit.None().MoreCount(4))
	expect.Eq(t, 1, vbit.Value(5).MoreCount(4))
	expect.Eq(t, 1, vbit.Values(3, 5).MoreCount(4))
	expect.Eq(t, 0, vbit.Less(5).MoreCount(4))
}

func TestEqual(t *testing.T) {
	expect.Eq(t, true, vbit.None().Equal(vbit.None()))
	expect.Eq(t, false, vbit.None().Equal(vbit.All()))
	expect.Eq(t, true, vbit.All().Equal(vbit.All()))
	expect.Eq(t, false, vbit.All().Equal(vbit.Value(5)))
	expect.Eq(t, true, vbit.Value(5).Equal(vbit.Value(5)))
	expect.Eq(t, false, vbit.Value(5).Equal(vbit.Values(3, 5)))
	expect.Eq(t, true, vbit.Values(3, 5).Equal(vbit.Values(3, 5)))
	expect.Eq(t, false, vbit.Values(3, 5).Equal(vbit.Less(5)))
	expect.Eq(t, true, vbit.Less(5).Equal(vbit.Less(5)))
	expect.Eq(t, false, vbit.Less(5).Equal(vbit.More(5)))
	expect.Eq(t, true, vbit.More(5).Equal(vbit.More(5)))
	expect.Eq(t, false, vbit.More(5).Equal(vbit.None()))
}

func TestHasNone(t *testing.T) {
	expect.Eq(t, true, vbit.None().HasNone(vbit.None()))
	expect.Eq(t, true, vbit.None().HasNone(vbit.All()))
	expect.Eq(t, false, vbit.All().HasNone(vbit.All()))
	expect.Eq(t, false, vbit.All().HasNone(vbit.Value(5)))
	expect.Eq(t, false, vbit.Value(5).HasNone(vbit.Value(5)))
	expect.Eq(t, false, vbit.Value(5).HasNone(vbit.Values(3, 5)))
	expect.Eq(t, false, vbit.Values(3, 5).HasNone(vbit.Values(3, 5)))
	expect.Eq(t, false, vbit.Values(3, 5).HasNone(vbit.Less(5)))
	expect.Eq(t, false, vbit.Less(5).HasNone(vbit.Less(5)))
	expect.Eq(t, true, vbit.Less(5).HasNone(vbit.More(5)))
	expect.Eq(t, false, vbit.More(5).HasNone(vbit.More(5)))
	expect.Eq(t, true, vbit.More(5).HasNone(vbit.None()))
}

func TestHasAll(t *testing.T) {
	expect.Eq(t, true, vbit.None().HasAll(vbit.None()))
	expect.Eq(t, false, vbit.None().HasAll(vbit.All()))
	expect.Eq(t, true, vbit.All().HasAll(vbit.All()))
	expect.Eq(t, true, vbit.All().HasAll(vbit.Value(5)))
	expect.Eq(t, true, vbit.Value(5).HasAll(vbit.Value(5)))
	expect.Eq(t, false, vbit.Value(5).HasAll(vbit.Values(3, 5)))
	expect.Eq(t, true, vbit.Values(3, 5).HasAll(vbit.Values(3, 5)))
	expect.Eq(t, false, vbit.Values(3, 5).HasAll(vbit.Less(5)))
	expect.Eq(t, true, vbit.Less(5).HasAll(vbit.Less(5)))
	expect.Eq(t, false, vbit.Less(5).HasAll(vbit.More(5)))
	expect.Eq(t, true, vbit.More(5).HasAll(vbit.More(5)))
	expect.Eq(t, true, vbit.More(5).HasAll(vbit.None()))
}

func TestCompare(t *testing.T) {
	expect.Eq(t, 0, vbit.None().Compare(vbit.None()))
	expect.Eq(t, -1, vbit.None().Compare(vbit.All()))
	expect.Eq(t, 0, vbit.All().Compare(vbit.All()))
	expect.Eq(t, 1, vbit.All().Compare(vbit.Value(5)))
	expect.Eq(t, 0, vbit.Value(5).Compare(vbit.Value(5)))
	expect.Eq(t, -1, vbit.Value(5).Compare(vbit.Values(3, 5)))
	expect.Eq(t, 0, vbit.Values(3, 5).Compare(vbit.Values(3, 5)))
	expect.Eq(t, -1, vbit.Values(3, 5).Compare(vbit.Less(5)))
	expect.Eq(t, 0, vbit.Less(5).Compare(vbit.Less(5)))
	expect.Eq(t, 1, vbit.Less(5).Compare(vbit.More(5)))
	expect.Eq(t, 0, vbit.More(5).Compare(vbit.More(5)))
	expect.Eq(t, 1, vbit.More(5).Compare(vbit.None()))
}

func TestAndCount(t *testing.T) {
	expect.Eq(t, 0, vbit.None().AndCount(vbit.None()))
	expect.Eq(t, 0, vbit.None().AndCount(vbit.All()))
	expect.Eq(t, 1, vbit.All().AndCount(vbit.Value(5)))
	expect.Eq(t, 1, vbit.Value(5).AndCount(vbit.Value(5)))
	expect.Eq(t, 1, vbit.Value(5).AndCount(vbit.Values(3, 5)))
	expect.Eq(t, 2, vbit.Values(3, 5).AndCount(vbit.Values(3, 5)))
	expect.Eq(t, 1, vbit.Values(3, 5).AndCount(vbit.Less(5)))
	expect.Eq(t, 5, vbit.Less(5).AndCount(vbit.Less(5)))
	expect.Eq(t, 0, vbit.Less(5).AndCount(vbit.More(5)))
	expect.Eq(t, 0, vbit.More(5).AndCount(vbit.None()))
}

func TestAdd(t *testing.T) {
	x := vbit.None()
	x.Add(3)
	expect.Set(t, vbit.Value(3), x)
	x.Add(3)
	expect.Set(t, vbit.Value(3), x)
	x.Add(5)
	expect.Set(t, vbit.Values(3, 5), x)
}

func TestRemove(t *testing.T) {
	x := vbit.Values(3, 5)
	x.Remove(1)
	expect.Set(t, vbit.Values(3, 5), x)
	x.Remove(3)
	expect.Set(t, vbit.Value(5), x)
	x.Remove(5)
	expect.Set(t, vbit.None(), x)
}

func TestAssign(t *testing.T) {
	x := vbit.None()
	x.Assign(vbit.Value(5))
	expect.Set(t, vbit.Value(5), x)
}

func TestAssignNone(t *testing.T) {
	x := vbit.Values(3, 5)
	x.AssignNone()
	expect.Set(t, vbit.None(), x)
}

func TestAssignAll(t *testing.T) {
	x := vbit.Values(3, 5)
	x.AssignAll()
	expect.Set(t, vbit.All(), x)
}

func TestEach(t *testing.T) {
	expect.Ints(t, vbit.None().Each())
	expect.Ints(t, vbit.Value(5).Each(), 5)
	expect.Ints(t, vbit.Values(3, 5).Each(), 3, 5)
	expect.Ints(t, vbit.Less(5).Each(), 0, 1, 2, 3, 4)
}

func TestNot(t *testing.T) {
	expect.Set(t, vbit.All(), vbit.None().Not())
	expect.Set(t, vbit.None(), vbit.All().Not())
	expect.Set(t, vbit.Less(5).Or(vbit.More(5)), vbit.Value(5).Not())
	expect.Set(t, vbit.Less(3).Or(vbit.Value(4).Or(vbit.More(5))), vbit.Values(3, 5).Not())
	expect.Set(t, vbit.More(4), vbit.Less(5).Not())
	expect.Set(t, vbit.Less(6), vbit.More(5).Not())
}

func TestAssignNot(t *testing.T) {
	x := vbit.Values(3, 5)
	x.AssignNot()
	expect.Set(t, vbit.Less(3).Or(vbit.Value(4)).Or(vbit.More(5)), x)
}

func TestAssignSub(t *testing.T) {
	x := vbit.Values(3, 5)
	x.AssignSub(vbit.Values(3))
	expect.Set(t, vbit.Values(5), x)
}

func TestAssignAnd(t *testing.T) {
	x := vbit.Values(3, 5)
	x.AssignAnd(vbit.Values(3))
	expect.Set(t, vbit.Values(3), x)
}

func TestAssignOr(t *testing.T) {
	x := vbit.Values(3, 5)
	x.AssignOr(vbit.Values(4))
	expect.Set(t, vbit.Values(3, 4, 5), x)
}

func TestAssignXor(t *testing.T) {
	x := vbit.Values(3, 5)
	x.AssignXor(vbit.Values(3, 4))
	expect.Set(t, vbit.Values(4, 5), x)
}

func TestAssignNor(t *testing.T) {
	x := vbit.Values(3, 5)
	x.AssignNor(vbit.Values(3, 4))
	expect.Set(t, vbit.Less(3).Or(vbit.More(5)), x)
}

func TestAssignIff(t *testing.T) {
	x := vbit.Values(3, 5)
	x.AssignIff(vbit.Values(3, 4))
	expect.Set(t, vbit.Less(4).Or(vbit.More(5)), x)
}

func TestAssignImply(t *testing.T) {
	x := vbit.Values(3, 5)
	x.AssignImply(vbit.Values(3, 4))
	expect.Set(t, vbit.Less(5).Or(vbit.More(5)), x)
}

func TestSub(t *testing.T) {
	expect.Set(t, vbit.None(), vbit.None().Sub(vbit.None()))
	expect.Set(t, vbit.None(), vbit.None().Sub(vbit.All()))
	expect.Set(t, vbit.None(), vbit.All().Sub(vbit.All()))
	expect.Set(t, vbit.Less(5).Or(vbit.More(5)), vbit.All().Sub(vbit.Value(5)))
	expect.Set(t, vbit.None(), vbit.Value(5).Sub(vbit.Value(5)))
	expect.Set(t, vbit.None(), vbit.Value(5).Sub(vbit.Values(3, 5)))
	expect.Set(t, vbit.None(), vbit.Values(3, 5).Sub(vbit.Values(3, 5)))
	expect.Set(t, vbit.Value(5), vbit.Values(3, 5).Sub(vbit.Less(5)))
	expect.Set(t, vbit.None(), vbit.Less(5).Sub(vbit.Less(5)))
	expect.Set(t, vbit.Less(5), vbit.Less(5).Sub(vbit.More(5)))
	expect.Set(t, vbit.None(), vbit.More(5).Sub(vbit.More(5)))
	expect.Set(t, vbit.More(5), vbit.More(5).Sub(vbit.None()))
}

func TestAnd(t *testing.T) {
	expect.Set(t, vbit.None(), vbit.None().And(vbit.None()))
	expect.Set(t, vbit.None(), vbit.None().And(vbit.All()))
	expect.Set(t, vbit.All(), vbit.All().And(vbit.All()))
	expect.Set(t, vbit.Value(5), vbit.All().And(vbit.Value(5)))
	expect.Set(t, vbit.Value(5), vbit.Value(5).And(vbit.Value(5)))
	expect.Set(t, vbit.Value(5), vbit.Value(5).And(vbit.Values(3, 5)))
	expect.Set(t, vbit.Values(3, 5), vbit.Values(3, 5).And(vbit.Values(3, 5)))
	expect.Set(t, vbit.Value(3), vbit.Values(3, 5).And(vbit.Less(5)))
	expect.Set(t, vbit.Less(5), vbit.Less(5).And(vbit.Less(5)))
	expect.Set(t, vbit.None(), vbit.Less(5).And(vbit.More(5)))
	expect.Set(t, vbit.More(5), vbit.More(5).And(vbit.More(5)))
	expect.Set(t, vbit.None(), vbit.More(5).And(vbit.None()))
}

func TestOr(t *testing.T) {
	expect.Set(t, vbit.None(), vbit.None().Or(vbit.None()))
	expect.Set(t, vbit.All(), vbit.None().Or(vbit.All()))
	expect.Set(t, vbit.All(), vbit.All().Or(vbit.All()))
	expect.Set(t, vbit.All(), vbit.All().Or(vbit.Value(5)))
	expect.Set(t, vbit.Value(5), vbit.Value(5).Or(vbit.Value(5)))
	expect.Set(t, vbit.Values(5, 3), vbit.Value(5).Or(vbit.Values(3, 5)))
	expect.Set(t, vbit.Values(3, 5), vbit.Values(3, 5).Or(vbit.Values(3, 5)))
	expect.Set(t, vbit.Less(6), vbit.Values(3, 5).Or(vbit.Less(5)))
	expect.Set(t, vbit.Less(5), vbit.Less(5).Or(vbit.Less(5)))
	expect.Set(t, vbit.Value(5).Not(), vbit.Less(5).Or(vbit.More(5)))
	expect.Set(t, vbit.More(5), vbit.More(5).Or(vbit.More(5)))
	expect.Set(t, vbit.More(5), vbit.More(5).Or(vbit.None()))
}

func TestXor(t *testing.T) {
	expect.Set(t, vbit.None(), vbit.None().Xor(vbit.None()))
	expect.Set(t, vbit.All(), vbit.None().Xor(vbit.All()))
	expect.Set(t, vbit.None(), vbit.All().Xor(vbit.All()))
	expect.Set(t, vbit.Value(5).Not(), vbit.All().Xor(vbit.Value(5)))
	expect.Set(t, vbit.None(), vbit.Value(5).Xor(vbit.Value(5)))
	expect.Set(t, vbit.Value(3), vbit.Value(5).Xor(vbit.Values(3, 5)))
	expect.Set(t, vbit.None(), vbit.Values(3, 5).Xor(vbit.Values(3, 5)))
	expect.Set(t, vbit.Values(0, 1, 2, 4, 5), vbit.Values(3, 5).Xor(vbit.Less(5)))
	expect.Set(t, vbit.None(), vbit.Less(5).Xor(vbit.Less(5)))
	expect.Set(t, vbit.Value(5).Not(), vbit.Less(5).Xor(vbit.More(5)))
	expect.Set(t, vbit.None(), vbit.More(5).Xor(vbit.More(5)))
	expect.Set(t, vbit.More(5), vbit.More(5).Xor(vbit.None()))
}

func TestNor(t *testing.T) {
	expect.Set(t, vbit.All(), vbit.None().Nor(vbit.None()))
	expect.Set(t, vbit.None(), vbit.None().Nor(vbit.All()))
	expect.Set(t, vbit.None(), vbit.All().Nor(vbit.All()))
	expect.Set(t, vbit.None(), vbit.All().Nor(vbit.Value(5)))
	expect.Set(t, vbit.Less(5).Or(vbit.More(5)), vbit.Value(5).Nor(vbit.Value(5)))
	expect.Set(t, vbit.Values(0, 1, 2, 4).Or(vbit.More(5)), vbit.Value(5).Nor(vbit.Values(3, 5)))
	expect.Set(t, vbit.Values(0, 1, 2, 4).Or(vbit.More(5)), vbit.Values(3, 5).Nor(vbit.Values(3, 5)))
	expect.Set(t, vbit.More(5), vbit.Values(3, 5).Nor(vbit.Less(5)))
	expect.Set(t, vbit.More(4), vbit.Less(5).Nor(vbit.Less(5)))
	expect.Set(t, vbit.Value(5), vbit.Less(5).Nor(vbit.More(5)))
	expect.Set(t, vbit.Less(6), vbit.More(5).Nor(vbit.More(5)))
	expect.Set(t, vbit.Less(6), vbit.More(5).Nor(vbit.None()))
}

func TestIff(t *testing.T) {
	expect.Set(t, vbit.All(), vbit.None().Iff(vbit.None()))
	expect.Set(t, vbit.None(), vbit.None().Iff(vbit.All()))
	expect.Set(t, vbit.All(), vbit.All().Iff(vbit.All()))
	expect.Set(t, vbit.Value(5), vbit.All().Iff(vbit.Value(5)))
	expect.Set(t, vbit.All(), vbit.Value(5).Iff(vbit.Value(5)))
	expect.Set(t, vbit.Values(0, 1, 2).Or(vbit.More(3)), vbit.Value(5).Iff(vbit.Values(3, 5)))
	expect.Set(t, vbit.All(), vbit.Values(3, 5).Iff(vbit.Values(3, 5)))
	expect.Set(t, vbit.Value(3).Or(vbit.More(5)), vbit.Values(3, 5).Iff(vbit.Less(5)))
	expect.Set(t, vbit.All(), vbit.Less(5).Iff(vbit.Less(5)))
	expect.Set(t, vbit.Value(5), vbit.Less(5).Iff(vbit.More(5)))
	expect.Set(t, vbit.All(), vbit.More(5).Iff(vbit.More(5)))
	expect.Set(t, vbit.Less(6), vbit.More(5).Iff(vbit.None()))
}

func TestImply(t *testing.T) {
	expect.Set(t, vbit.All(), vbit.None().Imply(vbit.None()))
	expect.Set(t, vbit.All(), vbit.None().Imply(vbit.All()))
	expect.Set(t, vbit.All(), vbit.All().Imply(vbit.All()))
	expect.Set(t, vbit.Value(5), vbit.All().Imply(vbit.Value(5)))
	expect.Set(t, vbit.All(), vbit.Value(5).Imply(vbit.Value(5)))
	expect.Set(t, vbit.All(), vbit.Value(5).Imply(vbit.Values(3, 5)))
	expect.Set(t, vbit.All(), vbit.Values(3, 5).Imply(vbit.Values(3, 5)))
	expect.Set(t, vbit.Value(5).Not(), vbit.Values(3, 5).Imply(vbit.Less(5)))
	expect.Set(t, vbit.All(), vbit.Less(5).Imply(vbit.Less(5)))
	expect.Set(t, vbit.More(4), vbit.Less(5).Imply(vbit.More(5)))
	expect.Set(t, vbit.All(), vbit.More(5).Imply(vbit.More(5)))
	expect.Set(t, vbit.Less(6), vbit.More(5).Imply(vbit.None()))
}
