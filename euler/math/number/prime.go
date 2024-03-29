package number

import (
	"context"
	"fmt"
	"math"
	"math/big"

	"github.com/ifooth/projecteuler-go/euler/math/itertools"
)

const (
	// FirstPrime ..
	FirstPrime = int64(2)

	// SecondPrime ..
	SecondPrime = int64(3)
)

var (
	Zero = big.NewInt(0) // Zero ..
	One  = big.NewInt(1) // One ..
	Two  = big.NewInt(2) // Tow ..
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
	for i := int64(5); i <= SqrtCeilInt(num); i += 6 {
		if num%i == 0 || num%(i+2) == 0 {
			return false
		}
	}

	return true
}

// PrevPrime 上一个素数
func PrevPrime(num int64) int64 {
	start := PrevOddNumber(num)
	for {
		if IsPrime(start) {
			break
		}
		start -= 2
	}
	return start
}

// NextPrime 下一个素数
func NextPrime(num int64) int64 {
	start := NextOddNumber(num)
	for {
		if IsPrime(start) {
			break
		}
		start += 2
	}
	return start
}

// GeneratorOption : 素数迭代器配置
type GeneratorOption struct {
	Start   int64           // 开始
	Ctx     context.Context // 结束信号
	Stop    int64           // 结束
	Reverse bool            // 反向查找, 需要配合 Stop 参数
}

// Option sets values in Options
type Option func(o *GeneratorOption)

// StartOpt ..
func StartOpt(num int64) Option {
	return func(o *GeneratorOption) {
		o.Start = num
	}
}

// StopOpt ..
func StopOpt(num int64) Option {
	return func(o *GeneratorOption) {
		o.Stop = num
	}
}

// ReverseOpt ..
func ReverseOpt(reverse bool) Option {
	return func(o *GeneratorOption) {
		o.Reverse = reverse
	}
}

// PrimeGenerator : 素数迭代器
func PrimeGenerator(ctx context.Context, opts ...Option) <-chan int64 {
	result := make(chan int64)

	o := &GeneratorOption{}
	for _, opt := range opts {
		opt(o)
	}

	go func() {
		defer close(result)

		var num int64

		if o.Start <= FirstPrime {
			result <- FirstPrime
			num = SecondPrime
		} else {
			num = NextOddNumber(o.Start)
		}

		for {
			if IsPrime(num) {
				select {
				case <-ctx.Done():
					return
				case result <- num:
				}
			}

			if o.Reverse {
				num -= 2
			} else {
				num += 2
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
		fmt.Println("factor", factor, "limit", limit)
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

// FactorsGeneratorInt 因子生成器
// 12 = 1 * 2 * 2 * 3
func FactorsGeneratorInt(num *big.Int) <-chan *big.Int {
	result := make(chan *big.Int)

	go func() {
		defer close(result)

		result <- One
		factor, limit := Two, SqrtCeil(num)

		for factor.Cmp(limit) <= 0 {
			div := new(big.Int)
			mod := new(big.Int)
			div.DivMod(num, factor, mod)

			if mod.Cmp(Zero) == 0 {
				f := *factor
				result <- &f
				num.Set(div)
				limit = new(big.Int).Add(new(big.Int).Sqrt(num), One)
			} else {
				factor.Add(factor, One)
			}
		}
		if num.Cmp(One) > 0 {
			result <- num
		}
	}()
	return result
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

// Factors 大数分解
func FactorsInt(num *big.Int) map[*big.Int]*big.Int {
	factorMap := map[*big.Int]*big.Int{}
	for factor := range FactorsGeneratorInt(num) {
		if _, ok := factorMap[factor]; ok {
			factorMap[factor].Add(factorMap[factor], One)
		}
		factorMap[factor] = One
	}
	return factorMap
}

// ProperDivisors 真因子 小于n且整除n的正整数, 不包含自己
// 12 = [1, 2, 3, 4, 6]
func ProperDivisors(num int64) []int64 {
	factorMap := Factors(num)
	divisorSet := map[int64]struct{}{1: {}}

	fmt.Println(factorMap)

	limit := num / 2

	for prime, power := range factorMap {
		fmt.Println("prime, power", prime, power)
		for p := range itertools.IterInt(power + 1) {
			temp := map[int64]struct{}{}
			for d := range divisorSet {
				divisor := d * PowInt(prime, p)
				if divisor > limit {
					continue
				}
				temp[divisor] = struct{}{}
				fmt.Println("divisorSet", p, d, divisor, temp)
			}
			divisorSet = temp
		}
	}

	divisors := make([]int64, len(divisorSet))
	for d := range divisorSet {
		divisors = append(divisors, d)
	}
	return divisors
}

func Generator(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}
