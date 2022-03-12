package number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactorsGenerator(t *testing.T) {
	var result []int64
	for factor := range FactorsGenerator(12) {
		result = append(result, factor)
	}
	assert.Equal(t, len(result), 4)
	assert.Equal(t, result, []int64{1, 2, 2, 3})
}
