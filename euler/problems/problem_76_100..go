package problems

import (
	"fmt"

	"github.com/ifooth/projecteuler-go/euler/math/number"
)

// Problem76 : Counting summations
// 加和计数
// 算法：动态规划，背包问题(Knapsack problem),类似31题
func Problem76() (result int64) {
	return
}

// Problem77 : Prime summations
// 素数加和
func Problem77() (result int64) {
	primeList := []int64{}

	countMap := map[int64]int64{}
	for i := int64(2); ; i++ {
		if number.IsPrime(i) {
			primeList = append(primeList, i)
			// 素数个数算1
			countMap[i] = 1
			continue
		}

		var count int64
		visited := map[int64]struct{}{}
		for _, prime := range primeList {
			if _, ok := visited[prime]; ok {
				continue
			}

			if i-prime == 1 {
				continue
			}

			value := i - prime

			// 返回如果是合数， 加合数数量
			c := countMap[value]

			// 是素数的场景
			if c == 1 {
				visited[value] = struct{}{}
				count++
				fmt.Println(i, "=", value, prime, fmt.Sprintf("(%d)", 1))
				continue
			}

			// // 是合数的场景
			// if value%prime == 0 {
			// 	count += c - 1
			// 	fmt.Println(i, "=", value, prime, fmt.Sprintf("(%d)", c-1))
			// 	continue
			// }

			count += c
			fmt.Println(i, "=", value, prime, fmt.Sprintf("(%d)", c))
		}

		countMap[i] = count
		fmt.Println(i, "===", fmt.Sprintf("(%d)", count))
		if count >= 5000 {
			fmt.Println(countMap)
			result = i
			break
		}
	}
	return
}
