package number

import (
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

func TestFactorsGenerator(t *testing.T) {
	var result []int64
	for factor := range FactorsGenerator(12) {
		result = append(result, factor)
	}
	assert.Equal(t, len(result), 4)
	assert.Equal(t, result, []int64{1, 2, 2, 3})
}
