package myinterface

import (
	"fmt"
	"testing"
)

func TestSimpleAddress(t *testing.T)  {
	// 初始化变量
	var sa1 = SimpleAddress{
		network: "Test",
		address: "127.0.0.1",
	}
	fmt.Println(sa1)

	var sa2 = &sa1
	fmt.Println(sa2)

	// 只有指针可以识别成接口
	// 原因：值不能隐式的转换成指针
	Process(sa2)

}

func TestDefaultAddress(t *testing.T)  {
	// 初始化变量
	var sa1 = DefaultAddress{
		network: "Test",
		address: "127.0.0.1",
	}
	fmt.Println(sa1)

	var sa2 = &sa1
	fmt.Println(sa2)

	// 值和指针都可以识别成接口
	// 原因：指针可以隐式的转换成值
	Process(sa1)
	Process(sa2)
}
