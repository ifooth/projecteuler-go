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

func Problem2() int64 {
	var (
		n1  int64 = 1
		n2  int64 = 2
		sum int64 = 2
	)
	for {
		next := n1 + n2
		if next%2 == 0 {
			sum += next
		}
		if next >= 4000000 {
			return sum
		}
		n1, n2 = n2, next
	}
}
