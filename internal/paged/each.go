package paged

import "github.com/mdelah/bitsets/bit64"

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
		for j := range mask.Each() {
			if j == -1 {
				return struct{}{}, false
			}
			if !e(k + j) {
				return struct{}{}, true
			}
		}
	}
	return struct{}{}, false
}
