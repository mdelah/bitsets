package paged

import (
	"github.com/mdelah/bitsets/bit64"
)

type EachRange func(int, int) bool

func (e EachRange) Walk(carry, i, _ int, mask bit64.Set) (int, bool) {
	k := 64 * i
	switch {
	case mask.IsAll():
		if carry == -1 {
			carry = k
		}
	case mask.IsNone():
		if carry != -1 {
			if !e(carry, k) {
				return -1, true
			}
			carry = -1
		}
	default:
		for j1, j2 := range mask.Ranges() {
			if j1 == 0 {
				if carry == -1 {
					carry = k + j1
				}
			} else {
				if carry != -1 && !e(carry, k) {
					return -1, true
				}
				carry = k + j1
			}
			if j2 == bit64.Max {
				continue
			}
			if !e(carry, k+j2) {
				return -1, true
			}
			carry = -1
		}
	}
	return carry, false
}
