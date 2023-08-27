package itertools

// IterSlice ..
func IterSlice(values []int64) <-chan int64 {
	result := make(chan int64)
	go func() {
		defer close(result)

		for _, v := range values {
			result <- v
		}
	}()
	return result
}

// IterInt ..
func IterInt(max int64) <-chan int64 {
	result := make(chan int64)
	go func() {
		defer close(result)

		for i := int64(0); i < max; i++ {
			result <- i
		}
	}()
	return result
}
