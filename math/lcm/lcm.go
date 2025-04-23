package lcm

import (
	"math/bits"

	"github.com/TheAlgorithms/Go/math/gcd"
)

// Lcm returns the lcm of two numbers using the fact that lcm(a,b) * gcd(a,b) = | a * b |
func Lcm(a, b int64) int64 {
	ua, ub := abs(a), abs(b)
	h, l := bits.Mul64(ua, ub)
	g := gcd.Iterative(a, b)
	r := bits.Div64(h, l, uint64(g))
	return int64(r)
}

func abs(x int64) uint64 {
	if x < 0 {
		return uint64(-(x+1))+1
	}
	return uint64(x)
}
