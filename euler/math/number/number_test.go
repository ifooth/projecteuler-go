package number

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSqrtInt(t *testing.T) {
	num, ok := SqrtInt(4)
	assert.Equal(t, num, int64(2))
	assert.True(t, ok)

	num, ok = SqrtInt(5)
	assert.Equal(t, num, int64(0))
	assert.False(t, ok)

	for i := 0; i < 1000; i++ {
		num, ok = SqrtInt(int64(i))
		if ok {
			fmt.Println(i, num)
		}
	}
}

func TestIsPalindromic(t *testing.T) {
	ok := IsPalindromic(9009)
	assert.True(t, ok)
	ok = IsPalindromic(90092)
	assert.False(t, ok)
}
