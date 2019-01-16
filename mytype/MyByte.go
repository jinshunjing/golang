package mytype

import "fmt"

func MyByte()  {
	// byte等价于uint8，可以隐式转换
	var u1 uint8 = 65
	var b1 byte = u1
	fmt.Println(b1)

	// byte不等于int8,需要显示转换
	var i1 int8 = 65
	b1 = byte(i1)
	fmt.Println(b1)
}
