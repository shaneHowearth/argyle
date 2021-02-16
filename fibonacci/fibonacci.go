// Package fibonacci -
package fibonacci

import "math/big"

// Compute -
func Compute(n int) big.Int {
	if n < 0 {
		return *big.NewInt(0)
	}
	memo := [2]*big.Int{big.NewInt(0), big.NewInt(1)}
	for count := 1; count <= n; count++ {
		memo[0].Add(memo[0], memo[1])
		memo[0], memo[1] = memo[1], memo[0]
	}
	return *memo[n%2]
}
