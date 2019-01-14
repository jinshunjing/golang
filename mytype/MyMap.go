package mytype

import "fmt"

func MyMap()  {
	var ctryMap map[string]string
	ctryMap = make(map[string]string)

	ctryMap["CN"] = "China"
	ctryMap["US"] = "United States"

	// 遍历主键
	for ctry := range ctryMap {
		fmt.Println(ctry, ctryMap[ctry])
	}

	// 查询
	name, ok := ctryMap["US"]
	if (ok) {
		fmt.Println(name)
	} else {
		fmt.Println("Unknown")
	}

	// 删除
	delete(ctryMap, "US")
	name, ok = ctryMap["US"]
	if (ok) {
		fmt.Println(name)
	} else {
		fmt.Println("Unknown")
	}
}
