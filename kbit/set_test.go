package kbit_test

import (
	"github.com/mdelah/bitsets/internal/expect"
	"github.com/mdelah/bitsets/kbit"
	"testing"
)

func TestCount(t *testing.T) {
	expect.Eq(t, 0, kbit.None().Count())
	expect.Eq(t, kbit.Cap, kbit.All().Count())
	expect.Eq(t, 1, kbit.Value(5).Count())
	expect.Eq(t, 2, kbit.Values(3, 5).Count())
	expect.Eq(t, 5, kbit.Less(5).Count())
	expect.Eq(t, kbit.Max-5, kbit.More(5).Count())
}

func TestString(t *testing.T) {
	expect.Eq(t, "{}", kbit.None().String())
	expect.Eq(t, "{0-}", kbit.All().String())
	expect.Eq(t, "{5}", kbit.Value(5).String())
	expect.Eq(t, "{3,5}", kbit.Values(3, 5).String())
	expect.Eq(t, "{0-4}", kbit.Less(5).String())
	expect.Eq(t, "{6-}", kbit.More(5).String())
}

func TestIsNone(t *testing.T) {
	expect.Eq(t, true, kbit.None().IsNone())
	expect.Eq(t, false, kbit.All().IsNone())
	expect.Eq(t, false, kbit.Value(5).IsNone())
	expect.Eq(t, false, kbit.Values(3, 5).IsNone())
	expect.Eq(t, false, kbit.Less(5).IsNone())
	expect.Eq(t, false, kbit.More(5).IsNone())
}

func TestIsAll(t *testing.T) {
	expect.Eq(t, false, kbit.None().IsAll())
	expect.Eq(t, true, kbit.All().IsAll())
	expect.Eq(t, false, kbit.Value(5).IsAll())
	expect.Eq(t, false, kbit.Values(3, 5).IsAll())
	expect.Eq(t, false, kbit.Less(5).IsAll())
	expect.Eq(t, false, kbit.More(5).IsAll())
}

func TestMin(t *testing.T) {
	expect.Eq(t, -1, kbit.None().Min())
	expect.Eq(t, 0, kbit.All().Min())
	expect.Eq(t, 5, kbit.Value(5).Min())
	expect.Eq(t, 3, kbit.Values(3, 5).Min())
	expect.Eq(t, 0, kbit.Less(5).Min())
	expect.Eq(t, 6, kbit.More(5).Min())
}

func TestMax(t *testing.T) {
	expect.Eq(t, -1, kbit.None().Max())
	expect.Eq(t, kbit.Max, kbit.All().Max())
	expect.Eq(t, 5, kbit.Value(5).Max())
	expect.Eq(t, 5, kbit.Values(3, 5).Max())
	expect.Eq(t, 4, kbit.Less(5).Max())
	expect.Eq(t, kbit.Max, kbit.More(5).Max())
}

func TestHas(t *testing.T) {
	expect.Eq(t, false, kbit.None().Has(3))
	expect.Eq(t, true, kbit.All().Has(3))
	expect.Eq(t, false, kbit.Value(5).Has(3))
	expect.Eq(t, true, kbit.Values(3, 5).Has(3))
	expect.Eq(t, true, kbit.Less(5).Has(3))
	expect.Eq(t, false, kbit.More(5).Has(3))
}

func TestLessCount(t *testing.T) {
	expect.Eq(t, 0, kbit.None().LessCount(5))
	expect.Eq(t, 5, kbit.All().LessCount(5))
	expect.Eq(t, 0, kbit.Value(5).LessCount(5))
	expect.Eq(t, 1, kbit.Values(3, 5).LessCount(5))
	expect.Eq(t, 5, kbit.Less(5).LessCount(5))
	expect.Eq(t, 0, kbit.More(5).LessCount(5))
}

func TestMoreCount(t *testing.T) {
	expect.Eq(t, 0, kbit.None().MoreCount(4))
	expect.Eq(t, kbit.Cap-5, kbit.All().MoreCount(4))
	expect.Eq(t, 1, kbit.Value(5).MoreCount(4))
	expect.Eq(t, 1, kbit.Values(3, 5).MoreCount(4))
	expect.Eq(t, 0, kbit.Less(5).MoreCount(4))
	expect.Eq(t, kbit.Cap-6, kbit.More(5).MoreCount(4))
}

func TestEqual(t *testing.T) {
	expect.Eq(t, true, kbit.None().Equal(kbit.None()))
	expect.Eq(t, false, kbit.None().Equal(kbit.All()))
	expect.Eq(t, true, kbit.All().Equal(kbit.All()))
	expect.Eq(t, false, kbit.All().Equal(kbit.Value(5)))
	expect.Eq(t, true, kbit.Value(5).Equal(kbit.Value(5)))
	expect.Eq(t, false, kbit.Value(5).Equal(kbit.Values(3, 5)))
	expect.Eq(t, true, kbit.Values(3, 5).Equal(kbit.Values(3, 5)))
	expect.Eq(t, false, kbit.Values(3, 5).Equal(kbit.Less(5)))
	expect.Eq(t, true, kbit.Less(5).Equal(kbit.Less(5)))
	expect.Eq(t, false, kbit.Less(5).Equal(kbit.More(5)))
	expect.Eq(t, true, kbit.More(5).Equal(kbit.More(5)))
	expect.Eq(t, false, kbit.More(5).Equal(kbit.None()))
}

func TestHasNone(t *testing.T) {
	expect.Eq(t, true, kbit.None().HasNone(kbit.None()))
	expect.Eq(t, true, kbit.None().HasNone(kbit.All()))
	expect.Eq(t, false, kbit.All().HasNone(kbit.All()))
	expect.Eq(t, false, kbit.All().HasNone(kbit.Value(5)))
	expect.Eq(t, false, kbit.Value(5).HasNone(kbit.Value(5)))
	expect.Eq(t, false, kbit.Value(5).HasNone(kbit.Values(3, 5)))
	expect.Eq(t, false, kbit.Values(3, 5).HasNone(kbit.Values(3, 5)))
	expect.Eq(t, false, kbit.Values(3, 5).HasNone(kbit.Less(5)))
	expect.Eq(t, false, kbit.Less(5).HasNone(kbit.Less(5)))
	expect.Eq(t, true, kbit.Less(5).HasNone(kbit.More(5)))
	expect.Eq(t, false, kbit.More(5).HasNone(kbit.More(5)))
	expect.Eq(t, true, kbit.More(5).HasNone(kbit.None()))
}

func TestHasAll(t *testing.T) {
	expect.Eq(t, true, kbit.None().HasAll(kbit.None()))
	expect.Eq(t, false, kbit.None().HasAll(kbit.All()))
	expect.Eq(t, true, kbit.All().HasAll(kbit.All()))
	expect.Eq(t, true, kbit.All().HasAll(kbit.Value(5)))
	expect.Eq(t, true, kbit.Value(5).HasAll(kbit.Value(5)))
	expect.Eq(t, false, kbit.Value(5).HasAll(kbit.Values(3, 5)))
	expect.Eq(t, true, kbit.Values(3, 5).HasAll(kbit.Values(3, 5)))
	expect.Eq(t, false, kbit.Values(3, 5).HasAll(kbit.Less(5)))
	expect.Eq(t, true, kbit.Less(5).HasAll(kbit.Less(5)))
	expect.Eq(t, false, kbit.Less(5).HasAll(kbit.More(5)))
	expect.Eq(t, true, kbit.More(5).HasAll(kbit.More(5)))
	expect.Eq(t, true, kbit.More(5).HasAll(kbit.None()))
}

func TestCompare(t *testing.T) {
	expect.Eq(t, 0, kbit.None().Compare(kbit.None()))
	expect.Eq(t, -1, kbit.None().Compare(kbit.All()))
	expect.Eq(t, 0, kbit.All().Compare(kbit.All()))
	expect.Eq(t, 1, kbit.All().Compare(kbit.Value(5)))
	expect.Eq(t, 0, kbit.Value(5).Compare(kbit.Value(5)))
	expect.Eq(t, -1, kbit.Value(5).Compare(kbit.Values(3, 5)))
	expect.Eq(t, 0, kbit.Values(3, 5).Compare(kbit.Values(3, 5)))
	expect.Eq(t, -1, kbit.Values(3, 5).Compare(kbit.Less(5)))
	expect.Eq(t, 0, kbit.Less(5).Compare(kbit.Less(5)))
	expect.Eq(t, 1, kbit.Less(5).Compare(kbit.More(5)))
	expect.Eq(t, 0, kbit.More(5).Compare(kbit.More(5)))
	expect.Eq(t, 1, kbit.More(5).Compare(kbit.None()))
}

func TestAndCount(t *testing.T) {
	expect.Eq(t, 0, kbit.None().AndCount(kbit.None()))
	expect.Eq(t, 0, kbit.None().AndCount(kbit.All()))
	expect.Eq(t, kbit.Cap, kbit.All().AndCount(kbit.All()))
	expect.Eq(t, 1, kbit.All().AndCount(kbit.Value(5)))
	expect.Eq(t, 1, kbit.Value(5).AndCount(kbit.Value(5)))
	expect.Eq(t, 1, kbit.Value(5).AndCount(kbit.Values(3, 5)))
	expect.Eq(t, 2, kbit.Values(3, 5).AndCount(kbit.Values(3, 5)))
	expect.Eq(t, 1, kbit.Values(3, 5).AndCount(kbit.Less(5)))
	expect.Eq(t, 5, kbit.Less(5).AndCount(kbit.Less(5)))
	expect.Eq(t, 0, kbit.Less(5).AndCount(kbit.More(5)))
	expect.Eq(t, kbit.Cap-6, kbit.More(5).AndCount(kbit.More(5)))
	expect.Eq(t, 0, kbit.More(5).AndCount(kbit.None()))
}

func TestAdd(t *testing.T) {
	x := kbit.None()
	x.Add(3)
	expect.Set(t, kbit.Value(3), x)
	x.Add(3)
	expect.Set(t, kbit.Value(3), x)
	x.Add(5)
	expect.Set(t, kbit.Values(3, 5), x)
}

func TestRemove(t *testing.T) {
	x := kbit.Values(3, 5)
	x.Remove(1)
	expect.Set(t, kbit.Values(3, 5), x)
	x.Remove(3)
	expect.Set(t, kbit.Value(5), x)
	x.Remove(5)
	expect.Set(t, kbit.None(), x)
}

func TestAssign(t *testing.T) {
	x := kbit.None()
	x.Assign(kbit.Value(5))
	expect.Set(t, kbit.Value(5), x)
}

func TestAssignNone(t *testing.T) {
	x := kbit.Values(3, 5)
	x.AssignNone()
	expect.Set(t, kbit.None(), x)
}

func TestAssignAll(t *testing.T) {
	x := kbit.Values(3, 5)
	x.AssignAll()
	expect.Set(t, kbit.All(), x)
}

func TestEach(t *testing.T) {
	expect.Ints(t, kbit.None().Each())
	expect.Ints(t, kbit.Value(5).Each(), 5)
	expect.Ints(t, kbit.Values(3, 5).Each(), 3, 5)
	expect.Ints(t, kbit.Less(5).Each(), 0, 1, 2, 3, 4)
}

func TestNot(t *testing.T) {
	expect.Set(t, kbit.All(), kbit.None().Not())
	expect.Set(t, kbit.None(), kbit.All().Not())
	expect.Set(t, kbit.Less(5).Or(kbit.More(5)), kbit.Value(5).Not())
	expect.Set(t, kbit.Less(3).Or(kbit.Value(4).Or(kbit.More(5))), kbit.Values(3, 5).Not())
	expect.Set(t, kbit.More(4), kbit.Less(5).Not())
	expect.Set(t, kbit.Less(6), kbit.More(5).Not())
}

func TestAssignNot(t *testing.T) {
	x := kbit.Values(3, 5)
	x.AssignNot()
	expect.Set(t, kbit.Less(3).Or(kbit.Value(4)).Or(kbit.More(5)), x)
}

func TestAssignSub(t *testing.T) {
	x := kbit.Values(3, 5)
	x.AssignSub(kbit.Values(3))
	expect.Set(t, kbit.Values(5), x)
}

func TestAssignAnd(t *testing.T) {
	x := kbit.Values(3, 5)
	x.AssignAnd(kbit.Values(3))
	expect.Set(t, kbit.Values(3), x)
}

func TestAssignOr(t *testing.T) {
	x := kbit.Values(3, 5)
	x.AssignOr(kbit.Values(4))
	expect.Set(t, kbit.Values(3, 4, 5), x)
}

func TestAssignXor(t *testing.T) {
	x := kbit.Values(3, 5)
	x.AssignXor(kbit.Values(3, 4))
	expect.Set(t, kbit.Values(4, 5), x)
}

func TestAssignNor(t *testing.T) {
	x := kbit.Values(3, 5)
	x.AssignNor(kbit.Values(3, 4))
	expect.Set(t, kbit.Less(3).Or(kbit.More(5)), x)
}

func TestAssignIff(t *testing.T) {
	x := kbit.Values(3, 5)
	x.AssignIff(kbit.Values(3, 4))
	expect.Set(t, kbit.Less(4).Or(kbit.More(5)), x)
}

func TestAssignImply(t *testing.T) {
	x := kbit.Values(3, 5)
	x.AssignImply(kbit.Values(3, 4))
	expect.Set(t, kbit.Less(5).Or(kbit.More(5)), x)
}

func TestSub(t *testing.T) {
	expect.Set(t, kbit.None(), kbit.None().Sub(kbit.None()))
	expect.Set(t, kbit.None(), kbit.None().Sub(kbit.All()))
	expect.Set(t, kbit.None(), kbit.All().Sub(kbit.All()))
	expect.Set(t, kbit.Less(5).Or(kbit.More(5)), kbit.All().Sub(kbit.Value(5)))
	expect.Set(t, kbit.None(), kbit.Value(5).Sub(kbit.Value(5)))
	expect.Set(t, kbit.None(), kbit.Value(5).Sub(kbit.Values(3, 5)))
	expect.Set(t, kbit.None(), kbit.Values(3, 5).Sub(kbit.Values(3, 5)))
	expect.Set(t, kbit.Value(5), kbit.Values(3, 5).Sub(kbit.Less(5)))
	expect.Set(t, kbit.None(), kbit.Less(5).Sub(kbit.Less(5)))
	expect.Set(t, kbit.Less(5), kbit.Less(5).Sub(kbit.More(5)))
	expect.Set(t, kbit.None(), kbit.More(5).Sub(kbit.More(5)))
	expect.Set(t, kbit.More(5), kbit.More(5).Sub(kbit.None()))
}

func TestAnd(t *testing.T) {
	expect.Set(t, kbit.None(), kbit.None().And(kbit.None()))
	expect.Set(t, kbit.None(), kbit.None().And(kbit.All()))
	expect.Set(t, kbit.All(), kbit.All().And(kbit.All()))
	expect.Set(t, kbit.Value(5), kbit.All().And(kbit.Value(5)))
	expect.Set(t, kbit.Value(5), kbit.Value(5).And(kbit.Value(5)))
	expect.Set(t, kbit.Value(5), kbit.Value(5).And(kbit.Values(3, 5)))
	expect.Set(t, kbit.Values(3, 5), kbit.Values(3, 5).And(kbit.Values(3, 5)))
	expect.Set(t, kbit.Value(3), kbit.Values(3, 5).And(kbit.Less(5)))
	expect.Set(t, kbit.Less(5), kbit.Less(5).And(kbit.Less(5)))
	expect.Set(t, kbit.None(), kbit.Less(5).And(kbit.More(5)))
	expect.Set(t, kbit.More(5), kbit.More(5).And(kbit.More(5)))
	expect.Set(t, kbit.None(), kbit.More(5).And(kbit.None()))
}

func TestOr(t *testing.T) {
	expect.Set(t, kbit.None(), kbit.None().Or(kbit.None()))
	expect.Set(t, kbit.All(), kbit.None().Or(kbit.All()))
	expect.Set(t, kbit.All(), kbit.All().Or(kbit.All()))
	expect.Set(t, kbit.All(), kbit.All().Or(kbit.Value(5)))
	expect.Set(t, kbit.Value(5), kbit.Value(5).Or(kbit.Value(5)))
	expect.Set(t, kbit.Values(5, 3), kbit.Value(5).Or(kbit.Values(3, 5)))
	expect.Set(t, kbit.Values(3, 5), kbit.Values(3, 5).Or(kbit.Values(3, 5)))
	expect.Set(t, kbit.Less(6), kbit.Values(3, 5).Or(kbit.Less(5)))
	expect.Set(t, kbit.Less(5), kbit.Less(5).Or(kbit.Less(5)))
	expect.Set(t, kbit.Value(5).Not(), kbit.Less(5).Or(kbit.More(5)))
	expect.Set(t, kbit.More(5), kbit.More(5).Or(kbit.More(5)))
	expect.Set(t, kbit.More(5), kbit.More(5).Or(kbit.None()))
}

func TestXor(t *testing.T) {
	expect.Set(t, kbit.None(), kbit.None().Xor(kbit.None()))
	expect.Set(t, kbit.All(), kbit.None().Xor(kbit.All()))
	expect.Set(t, kbit.None(), kbit.All().Xor(kbit.All()))
	expect.Set(t, kbit.Value(5).Not(), kbit.All().Xor(kbit.Value(5)))
	expect.Set(t, kbit.None(), kbit.Value(5).Xor(kbit.Value(5)))
	expect.Set(t, kbit.Value(3), kbit.Value(5).Xor(kbit.Values(3, 5)))
	expect.Set(t, kbit.None(), kbit.Values(3, 5).Xor(kbit.Values(3, 5)))
	expect.Set(t, kbit.Values(0, 1, 2, 4, 5), kbit.Values(3, 5).Xor(kbit.Less(5)))
	expect.Set(t, kbit.None(), kbit.Less(5).Xor(kbit.Less(5)))
	expect.Set(t, kbit.Value(5).Not(), kbit.Less(5).Xor(kbit.More(5)))
	expect.Set(t, kbit.None(), kbit.More(5).Xor(kbit.More(5)))
	expect.Set(t, kbit.More(5), kbit.More(5).Xor(kbit.None()))
}

func TestNor(t *testing.T) {
	expect.Set(t, kbit.All(), kbit.None().Nor(kbit.None()))
	expect.Set(t, kbit.None(), kbit.None().Nor(kbit.All()))
	expect.Set(t, kbit.None(), kbit.All().Nor(kbit.All()))
	expect.Set(t, kbit.None(), kbit.All().Nor(kbit.Value(5)))
	expect.Set(t, kbit.Less(5).Or(kbit.More(5)), kbit.Value(5).Nor(kbit.Value(5)))
	expect.Set(t, kbit.Values(0, 1, 2, 4).Or(kbit.More(5)), kbit.Value(5).Nor(kbit.Values(3, 5)))
	expect.Set(t, kbit.Values(0, 1, 2, 4).Or(kbit.More(5)), kbit.Values(3, 5).Nor(kbit.Values(3, 5)))
	expect.Set(t, kbit.More(5), kbit.Values(3, 5).Nor(kbit.Less(5)))
	expect.Set(t, kbit.More(4), kbit.Less(5).Nor(kbit.Less(5)))
	expect.Set(t, kbit.Value(5), kbit.Less(5).Nor(kbit.More(5)))
	expect.Set(t, kbit.Less(6), kbit.More(5).Nor(kbit.More(5)))
	expect.Set(t, kbit.Less(6), kbit.More(5).Nor(kbit.None()))
}

func TestIff(t *testing.T) {
	expect.Set(t, kbit.All(), kbit.None().Iff(kbit.None()))
	expect.Set(t, kbit.None(), kbit.None().Iff(kbit.All()))
	expect.Set(t, kbit.All(), kbit.All().Iff(kbit.All()))
	expect.Set(t, kbit.Value(5), kbit.All().Iff(kbit.Value(5)))
	expect.Set(t, kbit.All(), kbit.Value(5).Iff(kbit.Value(5)))
	expect.Set(t, kbit.Values(0, 1, 2).Or(kbit.More(3)), kbit.Value(5).Iff(kbit.Values(3, 5)))
	expect.Set(t, kbit.All(), kbit.Values(3, 5).Iff(kbit.Values(3, 5)))
	expect.Set(t, kbit.Value(3).Or(kbit.More(5)), kbit.Values(3, 5).Iff(kbit.Less(5)))
	expect.Set(t, kbit.All(), kbit.Less(5).Iff(kbit.Less(5)))
	expect.Set(t, kbit.Value(5), kbit.Less(5).Iff(kbit.More(5)))
	expect.Set(t, kbit.All(), kbit.More(5).Iff(kbit.More(5)))
	expect.Set(t, kbit.Less(6), kbit.More(5).Iff(kbit.None()))
}

func TestImply(t *testing.T) {
	expect.Set(t, kbit.All(), kbit.None().Imply(kbit.None()))
	expect.Set(t, kbit.All(), kbit.None().Imply(kbit.All()))
	expect.Set(t, kbit.All(), kbit.All().Imply(kbit.All()))
	expect.Set(t, kbit.Value(5), kbit.All().Imply(kbit.Value(5)))
	expect.Set(t, kbit.All(), kbit.Value(5).Imply(kbit.Value(5)))
	expect.Set(t, kbit.All(), kbit.Value(5).Imply(kbit.Values(3, 5)))
	expect.Set(t, kbit.All(), kbit.Values(3, 5).Imply(kbit.Values(3, 5)))
	expect.Set(t, kbit.Value(5).Not(), kbit.Values(3, 5).Imply(kbit.Less(5)))
	expect.Set(t, kbit.All(), kbit.Less(5).Imply(kbit.Less(5)))
	expect.Set(t, kbit.More(4), kbit.Less(5).Imply(kbit.More(5)))
	expect.Set(t, kbit.All(), kbit.More(5).Imply(kbit.More(5)))
	expect.Set(t, kbit.Less(6), kbit.More(5).Imply(kbit.None()))
}
