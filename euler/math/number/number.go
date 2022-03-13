package number

import (
	"math"
	"strconv"
)

// SqrtInt 平方根, 往上取整
func SqrtInt(num int64) int64 {
	return int64(math.Ceil(math.Sqrt(float64(num))))
}

// 指数计算, 整数
func PowInt(x, y int64) int64 {
	return int64(math.Pow(float64(x), float64(y)))
}

// NextOddNumber 下一个奇数
func NextOddNumber(num int64) int64 {
	if num%2 == 0 {
		return num + 1
	}
	return num
}

// PrevOddNumber 上一个奇数
func PrevOddNumber(num int64) int64 {
	if num%2 == 0 {
		return num - 1
	}
	return num
}

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
