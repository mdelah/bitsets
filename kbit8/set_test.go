package kbit8_test

import (
	"github.com/mdelah/bitsets/internal/expect"
	"github.com/mdelah/bitsets/kbit8"
	"testing"
)

func TestCount(t *testing.T) {
	expect.Eq(t, 0, kbit8.None().Count())
	expect.Eq(t, kbit8.Cap, kbit8.All().Count())
	expect.Eq(t, 1, kbit8.Value(5).Count())
	expect.Eq(t, 2, kbit8.Values(3, 5).Count())
	expect.Eq(t, 5, kbit8.Less(5).Count())
	expect.Eq(t, kbit8.Max-5, kbit8.More(5).Count())
}

func TestString(t *testing.T) {
	expect.Eq(t, "{}", kbit8.None().String())
	expect.Eq(t, "{0-}", kbit8.All().String())
	expect.Eq(t, "{5}", kbit8.Value(5).String())
	expect.Eq(t, "{3,5}", kbit8.Values(3, 5).String())
	expect.Eq(t, "{0-4}", kbit8.Less(5).String())
	expect.Eq(t, "{6-}", kbit8.More(5).String())
}

func TestIsNone(t *testing.T) {
	expect.Eq(t, true, kbit8.None().IsNone())
	expect.Eq(t, false, kbit8.All().IsNone())
	expect.Eq(t, false, kbit8.Value(5).IsNone())
	expect.Eq(t, false, kbit8.Values(3, 5).IsNone())
	expect.Eq(t, false, kbit8.Less(5).IsNone())
	expect.Eq(t, false, kbit8.More(5).IsNone())
}

func TestIsAll(t *testing.T) {
	expect.Eq(t, false, kbit8.None().IsAll())
	expect.Eq(t, true, kbit8.All().IsAll())
	expect.Eq(t, false, kbit8.Value(5).IsAll())
	expect.Eq(t, false, kbit8.Values(3, 5).IsAll())
	expect.Eq(t, false, kbit8.Less(5).IsAll())
	expect.Eq(t, false, kbit8.More(5).IsAll())
}

func TestMin(t *testing.T) {
	expect.Eq(t, -1, kbit8.None().Min())
	expect.Eq(t, 0, kbit8.All().Min())
	expect.Eq(t, 5, kbit8.Value(5).Min())
	expect.Eq(t, 3, kbit8.Values(3, 5).Min())
	expect.Eq(t, 0, kbit8.Less(5).Min())
	expect.Eq(t, 6, kbit8.More(5).Min())
}

func TestMax(t *testing.T) {
	expect.Eq(t, -1, kbit8.None().Max())
	expect.Eq(t, kbit8.Max, kbit8.All().Max())
	expect.Eq(t, 5, kbit8.Value(5).Max())
	expect.Eq(t, 5, kbit8.Values(3, 5).Max())
	expect.Eq(t, 4, kbit8.Less(5).Max())
	expect.Eq(t, kbit8.Max, kbit8.More(5).Max())
}

func TestHas(t *testing.T) {
	expect.Eq(t, false, kbit8.None().Has(3))
	expect.Eq(t, true, kbit8.All().Has(3))
	expect.Eq(t, false, kbit8.Value(5).Has(3))
	expect.Eq(t, true, kbit8.Values(3, 5).Has(3))
	expect.Eq(t, true, kbit8.Less(5).Has(3))
	expect.Eq(t, false, kbit8.More(5).Has(3))
}

func TestLessCount(t *testing.T) {
	expect.Eq(t, 0, kbit8.None().LessCount(5))
	expect.Eq(t, 5, kbit8.All().LessCount(5))
	expect.Eq(t, 0, kbit8.Value(5).LessCount(5))
	expect.Eq(t, 1, kbit8.Values(3, 5).LessCount(5))
	expect.Eq(t, 5, kbit8.Less(5).LessCount(5))
	expect.Eq(t, 0, kbit8.More(5).LessCount(5))
}

func TestMoreCount(t *testing.T) {
	expect.Eq(t, 0, kbit8.None().MoreCount(4))
	expect.Eq(t, kbit8.Cap-5, kbit8.All().MoreCount(4))
	expect.Eq(t, 1, kbit8.Value(5).MoreCount(4))
	expect.Eq(t, 1, kbit8.Values(3, 5).MoreCount(4))
	expect.Eq(t, 0, kbit8.Less(5).MoreCount(4))
	expect.Eq(t, kbit8.Cap-6, kbit8.More(5).MoreCount(4))
}

func TestEqual(t *testing.T) {
	expect.Eq(t, true, kbit8.None().Equal(kbit8.None()))
	expect.Eq(t, false, kbit8.None().Equal(kbit8.All()))
	expect.Eq(t, true, kbit8.All().Equal(kbit8.All()))
	expect.Eq(t, false, kbit8.All().Equal(kbit8.Value(5)))
	expect.Eq(t, true, kbit8.Value(5).Equal(kbit8.Value(5)))
	expect.Eq(t, false, kbit8.Value(5).Equal(kbit8.Values(3, 5)))
	expect.Eq(t, true, kbit8.Values(3, 5).Equal(kbit8.Values(3, 5)))
	expect.Eq(t, false, kbit8.Values(3, 5).Equal(kbit8.Less(5)))
	expect.Eq(t, true, kbit8.Less(5).Equal(kbit8.Less(5)))
	expect.Eq(t, false, kbit8.Less(5).Equal(kbit8.More(5)))
	expect.Eq(t, true, kbit8.More(5).Equal(kbit8.More(5)))
	expect.Eq(t, false, kbit8.More(5).Equal(kbit8.None()))
}

func TestHasNone(t *testing.T) {
	expect.Eq(t, true, kbit8.None().HasNone(kbit8.None()))
	expect.Eq(t, true, kbit8.None().HasNone(kbit8.All()))
	expect.Eq(t, false, kbit8.All().HasNone(kbit8.All()))
	expect.Eq(t, false, kbit8.All().HasNone(kbit8.Value(5)))
	expect.Eq(t, false, kbit8.Value(5).HasNone(kbit8.Value(5)))
	expect.Eq(t, false, kbit8.Value(5).HasNone(kbit8.Values(3, 5)))
	expect.Eq(t, false, kbit8.Values(3, 5).HasNone(kbit8.Values(3, 5)))
	expect.Eq(t, false, kbit8.Values(3, 5).HasNone(kbit8.Less(5)))
	expect.Eq(t, false, kbit8.Less(5).HasNone(kbit8.Less(5)))
	expect.Eq(t, true, kbit8.Less(5).HasNone(kbit8.More(5)))
	expect.Eq(t, false, kbit8.More(5).HasNone(kbit8.More(5)))
	expect.Eq(t, true, kbit8.More(5).HasNone(kbit8.None()))
}

func TestHasAll(t *testing.T) {
	expect.Eq(t, true, kbit8.None().HasAll(kbit8.None()))
	expect.Eq(t, false, kbit8.None().HasAll(kbit8.All()))
	expect.Eq(t, true, kbit8.All().HasAll(kbit8.All()))
	expect.Eq(t, true, kbit8.All().HasAll(kbit8.Value(5)))
	expect.Eq(t, true, kbit8.Value(5).HasAll(kbit8.Value(5)))
	expect.Eq(t, false, kbit8.Value(5).HasAll(kbit8.Values(3, 5)))
	expect.Eq(t, true, kbit8.Values(3, 5).HasAll(kbit8.Values(3, 5)))
	expect.Eq(t, false, kbit8.Values(3, 5).HasAll(kbit8.Less(5)))
	expect.Eq(t, true, kbit8.Less(5).HasAll(kbit8.Less(5)))
	expect.Eq(t, false, kbit8.Less(5).HasAll(kbit8.More(5)))
	expect.Eq(t, true, kbit8.More(5).HasAll(kbit8.More(5)))
	expect.Eq(t, true, kbit8.More(5).HasAll(kbit8.None()))
}

func TestCompare(t *testing.T) {
	expect.Eq(t, 0, kbit8.None().Compare(kbit8.None()))
	expect.Eq(t, -1, kbit8.None().Compare(kbit8.All()))
	expect.Eq(t, 0, kbit8.All().Compare(kbit8.All()))
	expect.Eq(t, 1, kbit8.All().Compare(kbit8.Value(5)))
	expect.Eq(t, 0, kbit8.Value(5).Compare(kbit8.Value(5)))
	expect.Eq(t, -1, kbit8.Value(5).Compare(kbit8.Values(3, 5)))
	expect.Eq(t, 0, kbit8.Values(3, 5).Compare(kbit8.Values(3, 5)))
	expect.Eq(t, -1, kbit8.Values(3, 5).Compare(kbit8.Less(5)))
	expect.Eq(t, 0, kbit8.Less(5).Compare(kbit8.Less(5)))
	expect.Eq(t, 1, kbit8.Less(5).Compare(kbit8.More(5)))
	expect.Eq(t, 0, kbit8.More(5).Compare(kbit8.More(5)))
	expect.Eq(t, 1, kbit8.More(5).Compare(kbit8.None()))
}

func TestAndCount(t *testing.T) {
	expect.Eq(t, 0, kbit8.None().AndCount(kbit8.None()))
	expect.Eq(t, 0, kbit8.None().AndCount(kbit8.All()))
	expect.Eq(t, kbit8.Cap, kbit8.All().AndCount(kbit8.All()))
	expect.Eq(t, 1, kbit8.All().AndCount(kbit8.Value(5)))
	expect.Eq(t, 1, kbit8.Value(5).AndCount(kbit8.Value(5)))
	expect.Eq(t, 1, kbit8.Value(5).AndCount(kbit8.Values(3, 5)))
	expect.Eq(t, 2, kbit8.Values(3, 5).AndCount(kbit8.Values(3, 5)))
	expect.Eq(t, 1, kbit8.Values(3, 5).AndCount(kbit8.Less(5)))
	expect.Eq(t, 5, kbit8.Less(5).AndCount(kbit8.Less(5)))
	expect.Eq(t, 0, kbit8.Less(5).AndCount(kbit8.More(5)))
	expect.Eq(t, kbit8.Cap-6, kbit8.More(5).AndCount(kbit8.More(5)))
	expect.Eq(t, 0, kbit8.More(5).AndCount(kbit8.None()))
}

func TestAdd(t *testing.T) {
	x := kbit8.None()
	x.Add(3)
	expect.Set(t, kbit8.Value(3), x)
	x.Add(3)
	expect.Set(t, kbit8.Value(3), x)
	x.Add(5)
	expect.Set(t, kbit8.Values(3, 5), x)
}

func TestRemove(t *testing.T) {
	x := kbit8.Values(3, 5)
	x.Remove(1)
	expect.Set(t, kbit8.Values(3, 5), x)
	x.Remove(3)
	expect.Set(t, kbit8.Value(5), x)
	x.Remove(5)
	expect.Set(t, kbit8.None(), x)
}

func TestAssign(t *testing.T) {
	x := kbit8.None()
	x.Assign(kbit8.Value(5))
	expect.Set(t, kbit8.Value(5), x)
}

func TestAssignNone(t *testing.T) {
	x := kbit8.Values(3, 5)
	x.AssignNone()
	expect.Set(t, kbit8.None(), x)
}

func TestAssignAll(t *testing.T) {
	x := kbit8.Values(3, 5)
	x.AssignAll()
	expect.Set(t, kbit8.All(), x)
}

func TestEach(t *testing.T) {
	expect.Ints(t, kbit8.None().Each())
	expect.Ints(t, kbit8.Value(5).Each(), 5)
	expect.Ints(t, kbit8.Values(3, 5).Each(), 3, 5)
	expect.Ints(t, kbit8.Less(5).Each(), 0, 1, 2, 3, 4)
}

func TestNot(t *testing.T) {
	expect.Set(t, kbit8.All(), kbit8.None().Not())
	expect.Set(t, kbit8.None(), kbit8.All().Not())
	expect.Set(t, kbit8.Less(5).Or(kbit8.More(5)), kbit8.Value(5).Not())
	expect.Set(t, kbit8.Less(3).Or(kbit8.Value(4).Or(kbit8.More(5))), kbit8.Values(3, 5).Not())
	expect.Set(t, kbit8.More(4), kbit8.Less(5).Not())
	expect.Set(t, kbit8.Less(6), kbit8.More(5).Not())
}

func TestAssignNot(t *testing.T) {
	x := kbit8.Values(3, 5)
	x.AssignNot()
	expect.Set(t, kbit8.Less(3).Or(kbit8.Value(4)).Or(kbit8.More(5)), x)
}

func TestAssignSub(t *testing.T) {
	x := kbit8.Values(3, 5)
	x.AssignSub(kbit8.Values(3))
	expect.Set(t, kbit8.Values(5), x)
}

func TestAssignAnd(t *testing.T) {
	x := kbit8.Values(3, 5)
	x.AssignAnd(kbit8.Values(3))
	expect.Set(t, kbit8.Values(3), x)
}

func TestAssignOr(t *testing.T) {
	x := kbit8.Values(3, 5)
	x.AssignOr(kbit8.Values(4))
	expect.Set(t, kbit8.Values(3, 4, 5), x)
}

func TestAssignXor(t *testing.T) {
	x := kbit8.Values(3, 5)
	x.AssignXor(kbit8.Values(3, 4))
	expect.Set(t, kbit8.Values(4, 5), x)
}

func TestAssignNor(t *testing.T) {
	x := kbit8.Values(3, 5)
	x.AssignNor(kbit8.Values(3, 4))
	expect.Set(t, kbit8.Less(3).Or(kbit8.More(5)), x)
}

func TestAssignIff(t *testing.T) {
	x := kbit8.Values(3, 5)
	x.AssignIff(kbit8.Values(3, 4))
	expect.Set(t, kbit8.Less(4).Or(kbit8.More(5)), x)
}

func TestAssignImply(t *testing.T) {
	x := kbit8.Values(3, 5)
	x.AssignImply(kbit8.Values(3, 4))
	expect.Set(t, kbit8.Less(5).Or(kbit8.More(5)), x)
}

func TestSub(t *testing.T) {
	expect.Set(t, kbit8.None(), kbit8.None().Sub(kbit8.None()))
	expect.Set(t, kbit8.None(), kbit8.None().Sub(kbit8.All()))
	expect.Set(t, kbit8.None(), kbit8.All().Sub(kbit8.All()))
	expect.Set(t, kbit8.Less(5).Or(kbit8.More(5)), kbit8.All().Sub(kbit8.Value(5)))
	expect.Set(t, kbit8.None(), kbit8.Value(5).Sub(kbit8.Value(5)))
	expect.Set(t, kbit8.None(), kbit8.Value(5).Sub(kbit8.Values(3, 5)))
	expect.Set(t, kbit8.None(), kbit8.Values(3, 5).Sub(kbit8.Values(3, 5)))
	expect.Set(t, kbit8.Value(5), kbit8.Values(3, 5).Sub(kbit8.Less(5)))
	expect.Set(t, kbit8.None(), kbit8.Less(5).Sub(kbit8.Less(5)))
	expect.Set(t, kbit8.Less(5), kbit8.Less(5).Sub(kbit8.More(5)))
	expect.Set(t, kbit8.None(), kbit8.More(5).Sub(kbit8.More(5)))
	expect.Set(t, kbit8.More(5), kbit8.More(5).Sub(kbit8.None()))
}

func TestAnd(t *testing.T) {
	expect.Set(t, kbit8.None(), kbit8.None().And(kbit8.None()))
	expect.Set(t, kbit8.None(), kbit8.None().And(kbit8.All()))
	expect.Set(t, kbit8.All(), kbit8.All().And(kbit8.All()))
	expect.Set(t, kbit8.Value(5), kbit8.All().And(kbit8.Value(5)))
	expect.Set(t, kbit8.Value(5), kbit8.Value(5).And(kbit8.Value(5)))
	expect.Set(t, kbit8.Value(5), kbit8.Value(5).And(kbit8.Values(3, 5)))
	expect.Set(t, kbit8.Values(3, 5), kbit8.Values(3, 5).And(kbit8.Values(3, 5)))
	expect.Set(t, kbit8.Value(3), kbit8.Values(3, 5).And(kbit8.Less(5)))
	expect.Set(t, kbit8.Less(5), kbit8.Less(5).And(kbit8.Less(5)))
	expect.Set(t, kbit8.None(), kbit8.Less(5).And(kbit8.More(5)))
	expect.Set(t, kbit8.More(5), kbit8.More(5).And(kbit8.More(5)))
	expect.Set(t, kbit8.None(), kbit8.More(5).And(kbit8.None()))
}

func TestOr(t *testing.T) {
	expect.Set(t, kbit8.None(), kbit8.None().Or(kbit8.None()))
	expect.Set(t, kbit8.All(), kbit8.None().Or(kbit8.All()))
	expect.Set(t, kbit8.All(), kbit8.All().Or(kbit8.All()))
	expect.Set(t, kbit8.All(), kbit8.All().Or(kbit8.Value(5)))
	expect.Set(t, kbit8.Value(5), kbit8.Value(5).Or(kbit8.Value(5)))
	expect.Set(t, kbit8.Values(5, 3), kbit8.Value(5).Or(kbit8.Values(3, 5)))
	expect.Set(t, kbit8.Values(3, 5), kbit8.Values(3, 5).Or(kbit8.Values(3, 5)))
	expect.Set(t, kbit8.Less(6), kbit8.Values(3, 5).Or(kbit8.Less(5)))
	expect.Set(t, kbit8.Less(5), kbit8.Less(5).Or(kbit8.Less(5)))
	expect.Set(t, kbit8.Value(5).Not(), kbit8.Less(5).Or(kbit8.More(5)))
	expect.Set(t, kbit8.More(5), kbit8.More(5).Or(kbit8.More(5)))
	expect.Set(t, kbit8.More(5), kbit8.More(5).Or(kbit8.None()))
}

func TestXor(t *testing.T) {
	expect.Set(t, kbit8.None(), kbit8.None().Xor(kbit8.None()))
	expect.Set(t, kbit8.All(), kbit8.None().Xor(kbit8.All()))
	expect.Set(t, kbit8.None(), kbit8.All().Xor(kbit8.All()))
	expect.Set(t, kbit8.Value(5).Not(), kbit8.All().Xor(kbit8.Value(5)))
	expect.Set(t, kbit8.None(), kbit8.Value(5).Xor(kbit8.Value(5)))
	expect.Set(t, kbit8.Value(3), kbit8.Value(5).Xor(kbit8.Values(3, 5)))
	expect.Set(t, kbit8.None(), kbit8.Values(3, 5).Xor(kbit8.Values(3, 5)))
	expect.Set(t, kbit8.Values(0, 1, 2, 4, 5), kbit8.Values(3, 5).Xor(kbit8.Less(5)))
	expect.Set(t, kbit8.None(), kbit8.Less(5).Xor(kbit8.Less(5)))
	expect.Set(t, kbit8.Value(5).Not(), kbit8.Less(5).Xor(kbit8.More(5)))
	expect.Set(t, kbit8.None(), kbit8.More(5).Xor(kbit8.More(5)))
	expect.Set(t, kbit8.More(5), kbit8.More(5).Xor(kbit8.None()))
}

func TestNor(t *testing.T) {
	expect.Set(t, kbit8.All(), kbit8.None().Nor(kbit8.None()))
	expect.Set(t, kbit8.None(), kbit8.None().Nor(kbit8.All()))
	expect.Set(t, kbit8.None(), kbit8.All().Nor(kbit8.All()))
	expect.Set(t, kbit8.None(), kbit8.All().Nor(kbit8.Value(5)))
	expect.Set(t, kbit8.Less(5).Or(kbit8.More(5)), kbit8.Value(5).Nor(kbit8.Value(5)))
	expect.Set(t, kbit8.Values(0, 1, 2, 4).Or(kbit8.More(5)), kbit8.Value(5).Nor(kbit8.Values(3, 5)))
	expect.Set(t, kbit8.Values(0, 1, 2, 4).Or(kbit8.More(5)), kbit8.Values(3, 5).Nor(kbit8.Values(3, 5)))
	expect.Set(t, kbit8.More(5), kbit8.Values(3, 5).Nor(kbit8.Less(5)))
	expect.Set(t, kbit8.More(4), kbit8.Less(5).Nor(kbit8.Less(5)))
	expect.Set(t, kbit8.Value(5), kbit8.Less(5).Nor(kbit8.More(5)))
	expect.Set(t, kbit8.Less(6), kbit8.More(5).Nor(kbit8.More(5)))
	expect.Set(t, kbit8.Less(6), kbit8.More(5).Nor(kbit8.None()))
}

func TestIff(t *testing.T) {
	expect.Set(t, kbit8.All(), kbit8.None().Iff(kbit8.None()))
	expect.Set(t, kbit8.None(), kbit8.None().Iff(kbit8.All()))
	expect.Set(t, kbit8.All(), kbit8.All().Iff(kbit8.All()))
	expect.Set(t, kbit8.Value(5), kbit8.All().Iff(kbit8.Value(5)))
	expect.Set(t, kbit8.All(), kbit8.Value(5).Iff(kbit8.Value(5)))
	expect.Set(t, kbit8.Values(0, 1, 2).Or(kbit8.More(3)), kbit8.Value(5).Iff(kbit8.Values(3, 5)))
	expect.Set(t, kbit8.All(), kbit8.Values(3, 5).Iff(kbit8.Values(3, 5)))
	expect.Set(t, kbit8.Value(3).Or(kbit8.More(5)), kbit8.Values(3, 5).Iff(kbit8.Less(5)))
	expect.Set(t, kbit8.All(), kbit8.Less(5).Iff(kbit8.Less(5)))
	expect.Set(t, kbit8.Value(5), kbit8.Less(5).Iff(kbit8.More(5)))
	expect.Set(t, kbit8.All(), kbit8.More(5).Iff(kbit8.More(5)))
	expect.Set(t, kbit8.Less(6), kbit8.More(5).Iff(kbit8.None()))
}

func TestImply(t *testing.T) {
	expect.Set(t, kbit8.All(), kbit8.None().Imply(kbit8.None()))
	expect.Set(t, kbit8.All(), kbit8.None().Imply(kbit8.All()))
	expect.Set(t, kbit8.All(), kbit8.All().Imply(kbit8.All()))
	expect.Set(t, kbit8.Value(5), kbit8.All().Imply(kbit8.Value(5)))
	expect.Set(t, kbit8.All(), kbit8.Value(5).Imply(kbit8.Value(5)))
	expect.Set(t, kbit8.All(), kbit8.Value(5).Imply(kbit8.Values(3, 5)))
	expect.Set(t, kbit8.All(), kbit8.Values(3, 5).Imply(kbit8.Values(3, 5)))
	expect.Set(t, kbit8.Value(5).Not(), kbit8.Values(3, 5).Imply(kbit8.Less(5)))
	expect.Set(t, kbit8.All(), kbit8.Less(5).Imply(kbit8.Less(5)))
	expect.Set(t, kbit8.More(4), kbit8.Less(5).Imply(kbit8.More(5)))
	expect.Set(t, kbit8.All(), kbit8.More(5).Imply(kbit8.More(5)))
	expect.Set(t, kbit8.Less(6), kbit8.More(5).Imply(kbit8.None()))
}
