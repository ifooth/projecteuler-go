package problems

import (
	"fmt"
	"sync"

	"github.com/ifooth/projecteuler-go/euler/math/number"
)

// Problem64 : Odd period square roots
// 奇周期平方根
func Problem64() (result int64) {
	num := int64(23)
	a0 := number.SqrtFloorInt(num)
	r1 := num - a0*a0
	a1, b1 := number.Divmod(r1, a0)

	fmt.Println("lei", a0, a1, r1, b1)
	_r2 := num - b1*b1
	r2 := _r2 / r1
	a2, b2 := number.Divmod(r2, b1)
	fmt.Println("lei1", _r2, r2, a2, b2)
	return
}

func findDiophantine(d int64) (int64, int64) {
	for y := int64(1); ; y++ {
		x, ok := number.SqrtInt(d*y*y + 1)
		if ok {
			return x, y
		}
	}
}

// Problem66 : Diophantine equation
// 丢番图方程
// https://en.wikipedia.org/wiki/Pell%27s_equation
func Problem66() (result int64) {
	limit := int64(7)
	var (
		maxX int64
		maxY int64
		wg   sync.WaitGroup
		mtx  sync.Mutex
	)

	concurrency := 16

	diophantineChan := make(chan int64)

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for d := range diophantineChan {
				x, y := findDiophantine(d)
				fmt.Println("find diophantine done:", x, d, y)

				mtx.Lock()
				if x > maxX {
					maxX = x
					maxY = y
					result = d
				}
				mtx.Unlock()
			}
		}()
	}

	for diophantine := int64(2); diophantine <= limit; diophantine++ {
		// 丢番图数本身已经是平方时无解
		if _, ok := number.SqrtInt(diophantine); ok {
			continue
		}
		diophantineChan <- diophantine
		fmt.Println("try find diophantine:", diophantine)
	}

	close(diophantineChan)
	wg.Wait()

	fmt.Println("result", maxX, result, maxY)
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
