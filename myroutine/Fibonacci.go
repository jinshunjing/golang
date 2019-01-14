package myroutine

func Fibonacci(n int, c chan int) {
	x, y := 0, 1

	for i := 0; i < n; i++ {
		// 写入通道
		c <- x

		x, y = y, x+y
	}

	// 关闭通道
	close(c)
}
