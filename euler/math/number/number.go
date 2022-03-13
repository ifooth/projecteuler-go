package number

import (
	"math"
	"strconv"
)

func Divmod(numerator, denominator int64) (quotient, remainder int64) {
	quotient = numerator / denominator // integer division, decimals are truncated
	remainder = numerator % denominator
	return
}

func SqrtInt(num int64) (int64, bool) {
	sqrt := SqrtCeilInt(num)
	if sqrt*sqrt == num {
		return sqrt, true
	}
	return 0, false
}

// SqrtInt 平方根, 往上取整
func SqrtCeilInt(num int64) int64 {
	return int64(math.Ceil(math.Sqrt(float64(num))))
}

// SqrtInt 平方根, 往下取整
func SqrtFloorInt(num int64) int64 {
	return int64(math.Floor(math.Sqrt(float64(num))))
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
