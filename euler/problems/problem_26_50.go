package problems

import "fmt"

// Problem31 : Coin sums
// 硬币求和
// 动态规划，背包问题(Knapsack problem)
func Problem31() (result int64) {
	coins := []int64{1, 2, 5, 10, 20, 50, 100, 200}
	wayMap := map[int64]int64{}
	limit := int64(20)
	for i := int64(1); i <= limit; i++ {
		var count int64
		for _, coin := range coins {
			if i < coin {
				break
			}
			if i == coin {
				wayMap[i] = 1
				count++
				continue
			}

			value := i - coin
			count += wayMap[value]
		}
		wayMap[i] = count
		fmt.Println(i, count)
	}
	result = wayMap[200]
	return
}
