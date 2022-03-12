package number

import (
	"math"
)

// Factors 因数分解 12 = 1^1 * 2^2 * 3^1
// return {1: 1, 2: 2, : 3: 1}
func Factors(num int64) map[int64]int64 {
	factorMap := map[int64]int64{}
	for factor := range FactorsGenerator(num) {
		factorMap[factor] += 1
	}
	return factorMap
}

// 因子生成器
// 12 = 1 * 2 * 2 * 3
func FactorsGenerator(num int64) <-chan int64 {
	result := make(chan int64)

	go func() {
		defer close(result)

		result <- 1
		factor, limit := int64(2), math.Sqrt(float64(num))
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
