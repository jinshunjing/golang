package mystring

import (
	"fmt"
	"net"
	"strconv"
	"testing"
)

func TestNetString(t *testing.T)  {
	// 字符串与数字互相转换
	port := strconv.FormatUint(uint64(18333), 10)
	fmt.Println(port)

	portnum, _ := strconv.ParseUint(port, 10, 0)
	fmt.Println(portnum)

	full := net.JoinHostPort("127.0.0.1", port)
	fmt.Println(full)
}
