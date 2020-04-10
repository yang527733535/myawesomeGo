package main

import "fmt"

//结构体初始化
type person struct {
	name string
	city string
	age int
}

// struct 结构体
func main()  {
 var p4 person
 fmt.Println(p4)


//使用键值对初始化
 p5 :=person{
 	name:"xiaowangzi",
 	city:"beijing",
 	age:12,
 }
 fmt.Println(p5)

 //也可以对结构体指针进行键值对初始化，例如
	p6 := &person{
		name: "小王子",
		city: "北京",
		//age:  18,
	}
	fmt.Println(p6)
}