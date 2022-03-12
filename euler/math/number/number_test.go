package number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSqrtInt(t *testing.T) {
	num := SqrtInt(4)
	assert.Equal(t, num, int64(2))
	num = SqrtInt(5)
	assert.Equal(t, num, int64(3))
}

func TestIsPalindromic(t *testing.T) {
	ok := IsPalindromic(9009)
	assert.True(t, ok)
	ok = IsPalindromic(90092)
	assert.False(t, ok)
}
