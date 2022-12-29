package problems

import (
	"context"
	"fmt"
	"math"
	"strconv"

	"github.com/ifooth/projecteuler-go/euler/assets"
	"github.com/ifooth/projecteuler-go/euler/math/number"
)

// Problem1 : Multiples of 3 and 5
// 3的倍数和5的倍数
func Problem1() (result int64) {
	var (
		num   int64 = 1000
		index int64 = 0
	)

	for ; index < num; index++ {
		if index%3 == 0 || index%5 == 0 {
			result += index
		}
	}
	return
}

// Problem2 : Even Fibonacci numbers
// 偶斐波那契数
func Problem2() (result int64) {
	result = 2
	var (
		n1 int64 = 1
		n2 int64 = 2
	)
	for {
		next := n1 + n2
		if next%2 == 0 {
			result += next
		}
		if next >= 4000000 {
			return
		}
		n1, n2 = n2, next
	}
}

// Problem3 : Largest prime factor
// 最大质因数
func Problem3() (result int64) {
	for factor := range number.FactorsGenerator(600851475143) {
		result = factor
	}
	return
}

// Problem4 : Largest palindrome product
// 最大回文乘积
func Problem4() (result int64) {
	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			num := int64(i * j)
			if number.IsPalindromic(num) && num > result {
				result = num
			}
		}
	}
	return
}

// Problem5 : Smallest multiple
// 最小倍数
func Problem5() (result int64) {
	maxFactorMap := map[int64]int64{}
	for i := 2; i <= 20; i++ {
		factorMap := number.Factors(int64(i))
		for k, v := range factorMap {
			if maxFactorMap[k] < v {
				maxFactorMap[k] = v
			}
		}
	}
	result = 1
	for k, v := range maxFactorMap {
		result *= int64(math.Pow(float64(k), float64(v)))
	}
	return
}

// Problem6 : Sum square difference
// 平方的和与和的平方之差
func Problem6() (result int64) {
	var (
		sumOfNatural int
		sumOfSquare  int
	)
	for i := 1; i <= 100; i++ {
		sumOfNatural += i
		sumOfSquare += i * i
	}

	result = int64(sumOfNatural*sumOfNatural - sumOfSquare)
	return
}

// Problem7 : 10001st prime
// 第10001个素数
func Problem7() (result int64) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	count := 1
	for prime := range number.PrimeGenerator(ctx) {
		if count == 10001 {
			result = prime
			break
		}
		count += 1
	}
	return
}

// Problem8 : Largest product in a series
// 连续数字最大乘积
func Problem8() (result int64) {
	dat, err := assets.GetTrimString("problem_8.txt")
	if err != nil {
		panic(err)
	}

	limit := len(dat)
	adjacentCount := 13

	for i := 0; i <= limit-adjacentCount; i++ {
		temp := int64(1)
		for _, v := range dat[i : i+adjacentCount] {
			k, err := strconv.Atoi(string(v))
			if err != nil {
				panic(err)
			}
			temp *= int64(k)
		}
		if temp > result {
			result = temp
		}
	}
	return
}

// Problem9 : Special Pythagorean triplet
// 特殊毕达哥拉斯三元组
func Problem9() (result int64) {
	for a := 1; a < 1000; a++ {
		for b := a + 1; b < 1000; b++ {
			c := (1000 - a - b)
			if c < b {
				continue
			}
			if a*a+b*b == c*c {
				result = int64(a * b * c)
				return
			}
		}
	}
	return
}

// Problem10 : Summation of primes
// 素数的和
func Problem10() (result int64) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for prime := range number.PrimeGenerator(ctx) {
		if prime >= 2000000 {
			break
		}
		result += prime
	}
	return
}

// Problem11 : Largest product in a grid
// 方阵中的最大乘积
func Problem11() (result int64) {
	dat, err := assets.GetMatrixInt("problem_11.txt")
	if err != nil {
		panic(err)
	}

	limit := len(dat)
	for i := 0; i < limit; i++ {
		for j := 0; j < limit; j++ {
			// 计算横向
			if j+3 < limit {
				t := dat[i][j] * dat[i][j+1] * dat[i][j+2] * dat[i][j+3]
				if t > result {
					result = t
				}
			}
			// 计算往下竖向
			if i+3 < limit {
				t := dat[i][j] * dat[i+1][j] * dat[i+2][j] * dat[i+3][j]
				if t > result {
					result = t
				}
			}
			// 计算右上 45 度
			if i-3 > 0 && j+3 < limit {
				t := dat[i][j] * dat[i-1][j+1] * dat[i-2][j+2] * dat[i-3][j+3]
				if t > result {
					result = t
				}
			}
			// 计算右下 45 度
			if i+3 < limit && j+3 < limit {
				t := dat[i][j] * dat[i+1][j+1] * dat[i+2][j+2] * dat[i+3][j+3]
				if t > result {
					result = t
				}
			}
		}
	}

	return
}

// Problem12 : Highly divisible triangular number
// 多约数的三角形数
func Problem12() (result int64) {
	fmt.Println(number.Factors(28))
	return
}
