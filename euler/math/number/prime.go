package number

import (
	"math"
)

// 因子生成器
// 12 = 1 * 2 * 2 * 3
func FactorsGenerator(num int64) <-chan int64 {
	result := make(chan int64)

	go func() {
		defer close(result)

		result <- 1
		factor, limit := int64(2), math.Sqrt(float64(num))
		for float64(factor) < limit {
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
