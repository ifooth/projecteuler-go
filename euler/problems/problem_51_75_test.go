package problems

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDiophantine(t *testing.T) {
	d := int64(109)
	x, y := findDiophantine(d)
	fmt.Println(x, d, y)
	assert.Equal(t, x, int64(649))
	assert.Equal(t, y, int64(180))
}
