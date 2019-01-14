package mytype

import "fmt"

func MySlice()  {
	// 切片
	s1 := []int {1, 2, 3, 4, 5, 6}

	// 截取切片
	s2 := s1[:3]
	s3 := s1[3:]

	fmt.Println(len(s1), len(s2), len(s3))

	// 创建切片
	s4 := make([]int, 6)
	fmt.Println(len(s4))

	// 声明一个空切片
	var s5 []int
	fmt.Println(len(s5))

	// 增加
	s1 = append(s1, 7)
	fmt.Println(len(s1))

	// 拷贝
	s5 = make([]int, len(s1))
	copy(s5, s1)
	fmt.Println(s5)

}
