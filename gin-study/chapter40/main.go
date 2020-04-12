package main

import "fmt"

type Person struct {
	Name string
}

//给Person类型绑定一个方法

func (p Person) test()  {
	fmt.Println("test() name=",p.Name)
}

//方法
func main()  {

	var p Person
	p.Name = "yangtenghui"
	p.test()
	//test 方法只能通过Person类型的变量来调用 而不能直接调用
}