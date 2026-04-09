package paged

import (
	"math/bits"

	"github.com/mdelah/bitsets/bit64"
)

type Each func(int) bool

func (e Each) Walk(_ struct{}, i, mult int, mask bit64.Set) (struct{}, bool) {
	k := 64 * i
	switch {
	case mask.IsAll():
		end := k + 64*mult
		for n := k; n != end; n++ {
			if !e(n) {
				return struct{}{}, true
			}
		}
	case mask.IsNone():
	default:
		for w := uint64(mask); w != 0; w &= w - 1 {
			if !e(k + bits.TrailingZeros64(w)) {
				return struct{}{}, true
			}
		}
	}
	return struct{}{}, false
}
