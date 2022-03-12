package problems

func Problem1() (result int64) {
	var (
		num   int64 = 1000
		index int64 = 0
	)

	for ; index < num; index++ {
		if index%3 == 0 || index%5 == 0 {
			result += index
		}
	}
	return
}

func Problem2() (result int64) {
	result = 2
	var (
		n1 int64 = 1
		n2 int64 = 2
	)
	for {
		next := n1 + n2
		if next%2 == 0 {
			result += next
		}
		if next >= 4000000 {
			return
		}
		n1, n2 = n2, next
	}
}
