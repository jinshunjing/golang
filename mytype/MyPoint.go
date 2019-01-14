package mytype

import "fmt"

func MyPoint()  {
	var va int8 = 65
	// 变量的地址
	fmt.Println(&va)

	// 声明指针
	var pa *int8
	pa = &va
	// 指针的值是地址
	fmt.Println(pa)
	// 指针的内容
	fmt.Println(*pa)


}
