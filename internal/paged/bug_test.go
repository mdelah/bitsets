package paged

import (
	"fmt"
	"testing"
)

func TestMutBug(t *testing.T) {
	var v Var
	// First add: Head==Body==Tail==0, so we go into the first case: v.Begin=1, return &v.Body
	p := v.Mut(1)
	fmt.Printf("After Mut(1): Begin=%d, End=%d, len(More)=%d\n", v.Begin, v.End(), len(v.More))
	*p = 1

	// Second add: i=2, v.Begin=1, v.End()=2, so i >= v.End() case
	// n = 1 + 2 - 2 = 1... wait let me recalculate
	// n = 1 + v.End() - i = 1 + 2 - 2 = 1
	p2 := v.Mut(2)
	fmt.Printf("After Mut(2): Begin=%d, End=%d, len(More)=%d\n", v.Begin, v.End(), len(v.More))
	*p2 = 2

	fmt.Printf("Body=%v, More=%v\n", v.Body, v.More)

	// Now start fresh and add value at page 2 first (simulating "freshly created set")
	var v2 Var
	// First add: Head==Body==Tail==0, so first case: v.Begin=0, return &v.Body
	p3 := v2.Mut(0)
	*p3 = 1
	fmt.Printf("After Mut(0): Begin=%d, End=%d, len(More)=%d\n", v2.Begin, v2.End(), len(v2.More))

	// Second add: i=2, Begin=0, End=1, i >= End
	// n = 1 + 1 - 2 = 0!! negative or zero
	fmt.Printf("n would be: %d\n", 1 + v2.End() - 2)
	p4 := v2.Mut(2)
	fmt.Printf("After Mut(2): Begin=%d, End=%d, len(More)=%d\n", v2.Begin, v2.End(), len(v2.More))
	*p4 = 2
	fmt.Printf("Body=%v, More=%v\n", v2.Body, v2.More)
}
