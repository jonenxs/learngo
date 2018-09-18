package main

import "fmt"

func main() {
	m := map[string]string{
		"aaa": "AAA",
		"bbb": "BBB",
		"ccc": "CCC",
		"ddd": "DDD",
	}
	m2 := make(map[string]int) //m2 = empty map

	var m3 map[string]int //m3 = nil
	fmt.Println(m, m2, m3)
	// 遍历 map
	for k, v := range m {
		fmt.Println(k, v)
	}

	//取值
	aaa := m["aaa"]
	fmt.Println(aaa)
	aaaa, ok := m["aaaa"]
	fmt.Println(aaaa, ok)
	if bbbb, ok := m["bbbb"]; ok {
		fmt.Println(bbbb)
	} else {
		fmt.Println("key does not exist")
	}

	//删除
	delete(m, "ccc")
	ccc, ok := m["ccc"]
	fmt.Println(ccc, ok)
}
