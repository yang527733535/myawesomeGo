package main

import "fmt"

func main()  {
	//元素为map类型的切片
	//var sliceMap = make([]map[string]string,3)
	//for index,value:=range sliceMap{
	//	fmt.Println(index,value)
	//}
	//
	//sliceMap[0]= make(map[string]string,10)
	//sliceMap[0]["name"] = "xwz"
	//sliceMap[0]["password"] = "123456"
	//sliceMap[0]["address"] = "沙河"
	//for index,value :=range sliceMap{
	//	fmt.Println(index,value)
	//}

	//值为切片类型的map

	var sliceMap2 = make(map[string][]string,3)
	fmt.Println(sliceMap2)
	key :="中国"
	value,ok:=sliceMap2[key]
	if !ok{
		value = make([]string,0,2)
	}
	value = append(value,"北京","上海")
	sliceMap2[key]=value
	fmt.Println(sliceMap2)

}