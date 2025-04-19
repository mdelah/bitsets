package abstract

import (
	"fmt"
	"math"
	"strings"
)

func Format(s Set1) string {
	var b strings.Builder
	b.WriteString("{")
	for left, right := range s.Ranges() {
		if b.Len() > 1 {
			b.WriteString(",")
		}
		switch {
		case left == right:
			fmt.Fprintf(&b, "%d", left)
		case right == math.MaxInt:
			fmt.Fprintf(&b, "%d-", left)
		default:
			fmt.Fprintf(&b, "%d-%d", left, right)
		}
	}
	b.WriteString("}")
	return b.String()
}
