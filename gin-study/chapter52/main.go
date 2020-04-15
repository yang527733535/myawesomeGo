package main

import "fmt"

type Ainterface interface {
	Say()
	//Run()
}

type myint int

func (myint myint) Say() {
	fmt.Println("myint myint()")
}

type Stu struct {
	Name string
}

func (stu Stu) Say() {
	fmt.Println("stu say()")
}

//接口本身不能创建实例 但是可以指向一个实现了该接口的自定义类型的变量
func main() {
	var stu Stu
	var a Ainterface = stu
	a.Say()
	var myint myint = 10
	var b Ainterface = myint
	b.Say()
}
