package main

import (
	"fmt"
	"github.com/jinshunjing/golang/myroutine"
	"github.com/jinshunjing/golang/mytype"
)

func main()  {
	testByte()
}

func testRoutine()  {
	go myroutine.Say("Hello")
	myroutine.Say("World")
}

func testChannel1() {
	s := []int{1, 2, 3, 4, 5, 6}

	// 没有缓冲的通道
	c := make(chan int)

	go myroutine.Sum(s[:len(s)/2], c)
	go myroutine.Sum(s[len(s)/2:], c)
	x, y := <-c, <-c

	fmt.Println(x, y, x + y)
}

func testChannel2() {
	// 有缓冲的通道
	c := make(chan int, 10)

	go myroutine.Fibonacci(cap(c), c)

	for i := range c {
		fmt.Println(i)
	}
}

func testChannel3()  {
	myroutine.MySelect()
}

func testVariable()  {
	mytype.MyVariables()
}

func testByte()  {
	mytype.MyByte()
}

func testArray() {
	mytype.MyArray()
}

func testSlice()  {
	mytype.MySlice()
}

func testPoint()  {
	mytype.MyPoint()
}

func testMap()  {
	mytype.MyMap()
}