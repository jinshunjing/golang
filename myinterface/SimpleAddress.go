package myinterface

import "fmt"

// 定义结构
type SimpleAddress struct {
	network string
	address string
}

// 添加指针的方法
func (a *SimpleAddress) Network() string {
	return a.network
}
func (a *SimpleAddress) String() string {
	return a.address
}

// 定义结构
type DefaultAddress struct {
	network string
	address string
}

// 添加值的方法
func (a DefaultAddress) Network() string {
	return a.network
}
func (a DefaultAddress) String() string {
	return a.address
}

// 使用接口
// 接口直接传递值，不传递指针
func Process(a Address)  {
	fmt.Println(a.Network(), a.String())
}
