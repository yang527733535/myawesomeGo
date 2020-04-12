package main

import "fmt"

type Monster struct {
	Name string
	Age int8
}

//面向对象编程
func main()  {


	//结构体有一个要注意的就是 如果这个字段的类型是引用类型的话 一定要make才可以使用
	//还有一个要注意的一点就是 结构体是值类型
	var monster1 Monster
	monster1.Name= "niumowang"
	monster1.Age = 100

	 monster2 := monster1
	monster2.Name = "XIXI"
	monster2.Age = 21
	fmt.Println(monster1,monster2)
}