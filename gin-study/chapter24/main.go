package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

//map相关 映射关系容器为map
//map是一种无序的基于key-value的数据结构，Go语言中的map是引用类型，必须初始化才能使用。
func main()  {

	scoreMap  :=make(map[string]int,8)
	scoreMap ["张三"]=90
	scoreMap ["李四"]=100
	scoreMap["娜扎"] = 60
	fmt.Println(scoreMap )


	fmt.Println(scoreMap ["张三"])
	fmt.Printf("type of a:%T\n", scoreMap )

	//map也支持在声明的时候填充元素

	userInfo :=map[string]string {
		"username":"yth",
		"password":"123456",
	}
	fmt.Println(userInfo)

	//判断某个键是否存在

	value , isExit:=scoreMap["张三"]
	if isExit{
		fmt.Println(value)
	}else{
		fmt.Println("没有这个人")
	}

	//map的遍历  遍历map时的元素顺序与添加键值对的顺序无关

	for index,value:=range scoreMap{
		fmt.Println(index,value)
	}


	//使用delete()函数删除键值对
	delete(scoreMap,"娜扎")
	fmt.Println(scoreMap)

		//按照指定顺序遍历map
		rand.Seed(time.Now().UnixNano())

	 var scropeMap2 =make(map[string]int,200)

	 for i:=0;i<100;i++{
	 	key:=fmt.Sprintf("stu%02d",i) //生成stu开头的字符串
	 	value :=rand.Intn(100)
		 scropeMap2[key] = value
	 }
	//取出map中的所有key存入切片keys
	var keys = make([]string,0,200)
	for key :=range scropeMap2{
		keys = append(keys,key)
	}

	//对切片进行排序
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println(key, scropeMap2[key])
	}
}
