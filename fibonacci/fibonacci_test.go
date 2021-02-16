package fibonacci_test

import (
	"fmt"
	"math/big"
	"math/rand"
	"testing"

	"github.com/shanehowearth/argyle/fibonacci"
)

func TestCompute(t *testing.T) {
	testcases := map[string]struct {
		input  int
		output big.Int
	}{
		"example 1": {
			input:  1,
			output: *big.NewInt(1),
		},
		"example 2": {
			input:  10,
			output: *big.NewInt(55),
		},
		"negative input": {
			input:  -15,
			output: *big.NewInt(0),
		},
		"0th possible value": {
			input:  0,
			output: *big.NewInt(0),
		},
		"29th possible value": {
			input:  28,
			output: *big.NewInt(317811),
		},
		"51st possible value": {
			input:  50,
			output: *big.NewInt(12586269025),
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			output := fibonacci.Compute(tc.input)
			if output.Cmp(&tc.output) != 0 {
				t.Errorf(fmt.Sprint("Compute gave a different value expected: ", tc.output, ", got ", output, " for ", tc.input))
			}
		})
	}
}

func BenchmarkCompute(b *testing.B) {
	input := rand.Perm(b.N)
	for i := 0; i < b.N; i++ {
		fibonacci.Compute(input[i])
	}
}
