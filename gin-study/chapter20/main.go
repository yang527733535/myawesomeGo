package main

import "fmt"

//切片相关
//切片是一个引用类型，它的内部结构包含地址、长度和容量。切片一般用于快速地操作一块数据集合
func main()  {
	//声明切片类型
	var a []string //声明一个字符串切片
	var b  = []int{}  //声明一个整型切片并初始化
	var c = []bool{false, true} //声明一个布尔切片并初始化
	var d = []bool{false, true} //声明一个布尔切片并初始化
	fmt.Println(a)              //[]
	fmt.Println(b)              //[]
	fmt.Println(c)              //[false true]
	fmt.Println(d)              //[false true]
	fmt.Println(a == nil)       //true
	fmt.Println(b == nil)       //false  因为是初始化了 所以是false
	fmt.Println(c == nil)       //false
	//切片是引用类型，不支持直接比较，只能和nil比较

}