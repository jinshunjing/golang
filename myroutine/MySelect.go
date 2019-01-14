package myroutine

import (
	"fmt"
	"time"
)

func MySelect()  {
	var c1, c2 chan int8
	c1 = make(chan int8)
	c2 = make(chan int8)

	// 在另外一个goroutine里写入
	// 在这个goroutine里读出
	//go WriteChannel(c1)

	// 在另一个goroutine里读出
	// 在这个goroutine里写入
	go ReadChannel(c2)

	time.Sleep(1 * time.Second)

	var i1, i2 int8
	i2 = 78
	select {
		// 通道已经有数据
		case i1 = <-c1:
			fmt.Println("Read from c1: ", i1)
		// 通道可以写入数据
		case c2 <- i2:
			fmt.Println("Write to c2: ", i2)
		default:
			fmt.Println("Nothing")
	}
}

func WriteChannel(c chan int8) {
	// 数据写入通道
	c <- 77
}

func ReadChannel(c chan int8) {
	// 等待从通道读出数据
	i := <- c
	fmt.Println("Read from c: ", i)
}