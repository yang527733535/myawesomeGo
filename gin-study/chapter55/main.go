package main

import "fmt"

//类型断言
//如何将一个接口变量 赋给自定义类型的变量 类型断言

type Point struct {
	x int
	y int
}

func main() {
	var a interface{}
	var point Point = Point{1, 2}
	a = point
	var b Point
	b = a.(Point) //这一步就是类型断言的关键了 表示判断a是否指向Point类型的变量
	fmt.Println(b)

	var x interface{}
	var b2 float64 = 2.3
	x = b2
	y := x.(float64) //要确保原来的空接口指向的就是断言的类型
	fmt.Println(y)
}
