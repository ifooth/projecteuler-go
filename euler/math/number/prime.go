package number

import (
	"context"
	"math"
)

const (
	FirstPrime  = int64(2)
	SecondPrime = int64(3)
)

// IsPrime : 素数检测法
// Simple methods https://en.wikipedia.org/wiki/Primality_test
func IsPrime(num int64) bool {
	if num < 2 {
		return false
	}
	if num%2 == 0 && num != 2 {
		return false
	}
	if num%3 == 0 && num != 3 {
		return false
	}

	// +2序列必须加1, int向上浮动1
	limit := int64(math.Floor(math.Sqrt(float64(num))))
	for i := int64(5); i <= limit; i += 6 {
		if num%i == 0 || num%(i+2) == 0 {
			return false
		}
	}

	return true
}

// Factors 因数分解 12 = 1^1 * 2^2 * 3^1
// return {1: 1, 2: 2, : 3: 1}
func Factors(num int64) map[int64]int64 {
	factorMap := map[int64]int64{}
	for factor := range FactorsGenerator(num) {
		factorMap[factor] += 1
	}
	return factorMap
}

func PrimeGenerator(ctx context.Context) <-chan int64 {
	result := make(chan int64)
	go func() {
		defer close(result)
		result <- FirstPrime
		result <- SecondPrime

		num := SecondPrime
		for {
			num += 2
			if IsPrime(num) {
				select {
				case <-ctx.Done():
					return
				case result <- num:
				}
			}

		}
	}()
	return result
}

// 因子生成器
// 12 = 1 * 2 * 2 * 3
func FactorsGenerator(num int64) <-chan int64 {
	result := make(chan int64)

	go func() {
		defer close(result)

		result <- 1
		factor, limit := FirstPrime, math.Sqrt(float64(num))
		for float64(factor) <= limit {
			if num%factor == 0 {
				result <- factor

				num /= factor
				limit = math.Sqrt(float64(num))
			} else {
				factor += 1
			}
		}
		if num > 1 {
			result <- num
		}
	}()
	return result
}
