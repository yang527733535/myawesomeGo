package main

import "fmt"

//map相关的复习


type studeng struct{
	name string
	sex string
	age int8
}

func main()  {



	//map
	//var a map[string]string
	//在声明之后 要使用的话 一定要先make
	a :=make(map[string]string,10)  //也可以直接这样使用

	a["NO1"] = "yangtenghui"
	//注意  key是不能重复的 如果重复了 就以最后一个为准
	//value是可以相同的
	//map是无顺序的

	fmt.Println(a)

	//map 的第三种使用方式

	heros:= map[string]string{
		"hero1":"yangtenghui",
	}
	heros["hero2"] = "yantenghui2"
	fmt.Println(heros)

	//我们要存放3个学生信息 每个学生有name 和 sex 信息
	//思路 map[string]map[string]string

	var mymap map[string]map[string]string
	 mymap = make(map[string]map[string]string,3)
	 //fmt.Println(mymap)
	mymap["S1"] = make(map[string]string)
	fmt.Println(mymap["S1"])
	mymap["S1"]["name"] = "yangxiaohuo"
	mymap["S1"]["sex"] = "man"
	fmt.Println(mymap["S1"])
}