package number

import (
	"strconv"
)

func ReverseSlice(values []int64) []int64 {
	for i, j := 0, len(values)-1; i < j; i, j = i+1, j-1 {
		values[i], values[j] = values[j], values[i]
	}
	return values
}

func ReverseSliceChar(values []rune) []rune {
	for i, j := 0, len(values)-1; i < j; i, j = i+1, j-1 {
		values[i], values[j] = values[j], values[i]
	}
	return values
}

// IsPalindromic : 回文数 9009
func IsPalindromic(num int64) bool {
	s := strconv.FormatInt(num, 10)
	numSlice := []rune{}
	for _, v := range s {
		numSlice = append(numSlice, v)
	}
	reverse := string(ReverseSliceChar(numSlice))
	return s == reverse
}
