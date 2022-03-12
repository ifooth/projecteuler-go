package number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindromic(t *testing.T) {
	ok := IsPalindromic(9009)
	assert.True(t, ok)
	ok = IsPalindromic(90092)
	assert.False(t, ok)
}
