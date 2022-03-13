package problems

import (
	"fmt"

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
	for x := int64(2); ; x++ {
		quotient, remainder := number.Divmod(x*x-1, d)
		if remainder != 0 {
			continue
		}
		sqrt, ok := number.SqrtInt(quotient)
		if ok {
			return x, sqrt
		}
	}
}

// Problem66 : Diophantine equation
// 丢番图方程
func Problem66() (result int64) {
	limit := int64(1000)
	var (
		maxX int64
		maxY int64
	)

	for diophantine := int64(2); diophantine <= limit; diophantine++ {
		// 丢番图数本身已经是平方时无解
		if _, ok := number.SqrtInt(diophantine); ok {
			continue
		}

		x, y := findDiophantine(diophantine)
		fmt.Println(x, diophantine, y)

		if x > maxX {
			maxX = x
			maxY = y
			result = diophantine
		}

	}
	fmt.Println("result", maxX, result, maxY)
	return
}
