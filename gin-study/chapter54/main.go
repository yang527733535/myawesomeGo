package main

import "fmt"

//类型断言的最佳实践
//写一函数 循环判断传入参数的类型

func typeJudge(mytypes ...interface{}) {

	for index, value := range mytypes {
		//fmt.Println(index, value)
		switch value.(type) {
		case float64:
			fmt.Println("这是第", index, "个", "这个类型是float64")
		case float32:
			fmt.Println("这是第", index, "个", "这个类型是float32")
		}
	}
}
func main() {
	var n1 float64 = 1.1
	var n2 float32 = 1.1
	typeJudge(n1, n2)
}
