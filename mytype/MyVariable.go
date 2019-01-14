package mytype

import "fmt"

func MyVariables()  {
	// 声明变量，使用默认值
	var ua uint8
	fmt.Println(ua)
	// 给变量赋值
	ua = 65
	fmt.Println(ua)

	// 声明变量，赋值，但不给出类型
	var ub = 66
	fmt.Println(ub)

	// 简化声明变量
	uc := 67
	fmt.Println(uc)

	// 常量
	const ud uint8 = 88
	fmt.Println(ud)

	// iota
	const (
		ue = iota
		uf
		ug
	)
	fmt.Println(ue, uf, ug)

	const (
		uh = 3
		ui = iota
		uj
	)
	fmt.Println(uh, ui, uj)
}