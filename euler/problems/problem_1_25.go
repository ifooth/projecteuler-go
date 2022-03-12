package problems

import (
	"context"
	"math"

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

// Problem6 : 10001st prime
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
