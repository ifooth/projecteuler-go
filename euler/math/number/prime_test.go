package number

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactors(t *testing.T) {
	factors := Factors(12)
	assert.Equal(t, len(factors), 3)
	assert.Equal(t, factors, map[int64]int64{
		1: 1,
		2: 2,
		3: 1,
	})
}

func TestPrimeGenerator(t *testing.T) {
	in := 10
	c := 0

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for prime := range PrimeGenerator(ctx) {
		if c > in {
			break
		}
		fmt.Println(prime)
		c += 1
	}
}

func BenchmarkPrimeGenerator(b *testing.B) {
	in := 10000
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i := 0; i < b.N; i++ {
		c := 1
		for range PrimeGenerator(ctx, StartOpt(10)) {
			if c > in {
				break
			}
			c += 1
		}
	}
}

func TestFactorsGenerator(t *testing.T) {
	var result []int64
	for factor := range FactorsGenerator(12) {
		result = append(result, factor)
	}
	assert.Equal(t, len(result), 4)
	assert.Equal(t, result, []int64{1, 2, 2, 3})
}

func TestProperDivisors(t *testing.T) {
	result := ProperDivisors(12)
	fmt.Println(result)
	assert.Equal(t, len(result), 5)
	assert.Equal(t, result, []int64{1, 2, 3, 4, 6})
}
