package mytype

import (
	"fmt"
)

func MyArray()  {
	// 数组的初始化
	var ia1 [5]int32
	fmt.Println(ia1)

	ia2 := [5]int32{0, 1, 2, 3, 4}
	fmt.Println(ia2)

	// 数组的赋值: 复制整个数组
	ia11 := ia1
	fmt.Println(ia11)
	ia11[1] = 1
	fmt.Println(ia11)
	fmt.Println(ia1)

	// 子数组
	ia12 := ia2[1:4]
	fmt.Println(ia12)

	// 数组转换成切片
	var ia13 []int32
	ia13 = ia2[1:4]
	fmt.Println(ia13)

	ia14 := make([]int32, 10)
	ia14 = ia2[0:]
	fmt.Println(ia14)

	// 数组的遍历
	for idx, val := range ia2 {
		fmt.Println(idx, val)
	}
}
