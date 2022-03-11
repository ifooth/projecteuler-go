package problems

func Problem1() int64 {
	var (
		sum   int64 = 0
		num   int64 = 1000
		index int64 = 0
	)

	for ; index < num; index++ {
		if index%3 == 0 || index%5 == 0 {
			sum += index
		}
	}
	return sum
}
