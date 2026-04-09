package paged

import (
	"math/bits"

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
		for w := uint64(mask); w != 0; {
			lo := bits.TrailingZeros64(w)
			// find end of contiguous run starting at lo
			inv := ^(w >> lo)
			runLen := bits.TrailingZeros64(inv) // 64 if all remaining bits are 1
			hi := lo + runLen                   // exclusive end within word

			j1 := lo
			if j1 == 0 {
				if carry == -1 {
					carry = k + j1
				}
			} else {
				if carry != -1 && !e(carry, k+j1) {
					return -1, true
				}
				carry = k + j1
			}

			if hi >= 64 {
				// run extends to end of word; carry continues into next word
				break
			}
			if !e(carry, k+hi-1) {
				return -1, true
			}
			carry = -1

			// clear bits 0..hi-1 and continue
			w &^= uint64(1)<<hi - 1
		}
	}
	return carry, false
}
