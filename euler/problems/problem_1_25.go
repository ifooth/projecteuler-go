package problems

import "github.com/ifooth/projecteuler-go/euler/math/number"

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
