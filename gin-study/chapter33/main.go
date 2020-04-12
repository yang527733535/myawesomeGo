package main

import (
	"fmt"
	"sort"
)

//map切片
func main()  {


	//使用一个map来记录一个Monster信息

	var monsters []map[string]string
	monsters = make([]map[string]string,2) //准备放入两个妖怪
	if monsters[0]==nil{
		monsters[0] = make(map[string]string,2)
		monsters[0]["name"] = "niumowang"
		monsters[0]["age"] = "500"
	}


	if monsters[1]==nil{
		monsters[1] = make(map[string]string,2)
		monsters[1]["name"] = "BAIGU"
		monsters[1]["age"] = "23"
	}
	//容量已经满了 不能够再添加了 但是可以append
	fmt.Println(monsters)

	for index,value := range monsters{
		fmt.Println(index,value)
	}

	newMonster:= map[string]string{
		"name":"新的妖怪",
		"age":"200",
	}
	monsters  = append(monsters,newMonster)
	fmt.Println(monsters)

	//map的排序

	map1 :=make(map[int]int,10)
	map1[10] = 100
	map1[1] = 13
	map1[4] = 56
	map1[8] = 80

	fmt.Println(map1)


	for key,value := range map1{
		fmt.Println(key,value)
	}
		//如果要按照key的顺序来排列的话
		//先将map的key放入到切片当中
		//对切片进行排序
		//遍历切片 然后按照key的值来输出map

	var keys []int
	for k,_:=range map1{
		keys = append(keys,k)
	}

	fmt.Println(keys)
	sort.Ints(keys)
	fmt.Println(keys)

	for _,k:=range keys{
		fmt.Println(map1[k])
	}
	//MAP 是引用类型



}