package myroutine

func Sum(s []int, c chan int)  {
	sum := 0

	// 遍历切片
	for _, v := range s {
		sum += v
	}

	// 写入通道
	c <- sum
}